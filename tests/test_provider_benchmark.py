#!/usr/bin/env python3
"""Execute benchmark evidence boundaries without compiling the provider."""

import contextlib
import importlib.util
import io
import json
import os
import shutil
import signal
import subprocess
import tempfile
import time
import unittest
from pathlib import Path
from unittest.mock import patch

ROOT = Path(__file__).resolve().parents[1]


class BenchmarkEvidenceTest(unittest.TestCase):
    def setUp(self):
        self.stack = contextlib.ExitStack()
        self.addCleanup(self.stack.close)
        self.work = Path(self.stack.enter_context(tempfile.TemporaryDirectory()))
        self.repo = self.work / "repo"
        self.repo.mkdir()
        shutil.copytree(ROOT / "scripts", self.repo / "scripts")
        self.bin = self.work / "bin"
        self.bin.mkdir()
        self.write(
            self.bin / "go",
            """#!/usr/bin/env bash
set -euo pipefail
if [ "$1" = env ]; then
  case "$2" in
    GOVERSION) echo go1.25.13 ;;
    GOCACHE) echo "$GOCACHE" ;;
    GOMODCACHE) echo "$GOMODCACHE" ;;
  esac
elif [ "$1" = list ]; then
  echo 'example.test/provider example.test/provider@v1.0.0'
elif [ "$1" = run ] && [ "${FAIL_FIRST_GENERATOR:-false}" = true ]; then
  exit 7
elif [ "$1" = mod ] && [ "${FAIL_FIRST_GENERATOR:-false}" = true ]; then
  printf '%s\n' after-failed-generator >"$FOLLOW_ON_LOG"
elif [ "$1" = build ]; then
  [ "$GOMAXPROCS" = "$EXPECTED_CONCURRENCY" ]
  [ "$GOFLAGS" = "-p=$EXPECTED_CONCURRENCY" ]
  exit "${WORKLOAD_EXIT:-0}"
fi
""",
            executable=True,
        )
        self.write(
            self.repo / "scripts/generate-provider-docs.sh",
            """#!/usr/bin/env bash
set -euo pipefail
[ "$GOMAXPROCS" = "$EXPECTED_CONCURRENCY" ]
[ "$GOFLAGS" = "-p=$EXPECTED_CONCURRENCY" ]
exit "${WORKLOAD_EXIT:-0}"
""",
            executable=True,
        )
        self.write(
            self.repo / "scripts/go-retry.sh",
            '#!/usr/bin/env bash\nshift\nexec "$@"\n',
            executable=True,
        )
        self.write(
            self.repo / ".runner-harness/scripts/runner-profile.py",
            """import json, pathlib, subprocess, sys
args=sys.argv[1:]
output=pathlib.Path(args[args.index('--output')+1])
status=subprocess.run(args[args.index('--')+1:]).returncode
output.write_text(json.dumps({'exit':{'code':status}}))
raise SystemExit(status)
""",
        )
        self.env = os.environ | {
            "PATH": str(self.bin) + os.pathsep + os.environ["PATH"],
            "RUNNER_TEMP": str(self.work / "runner-temp"),
            "GOCACHE": str(self.work / "build-cache"),
            "GOMODCACHE": str(self.work / "module-cache"),
            "GITHUB_REPOSITORY": "example/provider",
            "RUNNER_IMAGE_DIGEST": "github-hosted",
        }
        for command in (
            ["git", "init", "-q"],
            ["git", "config", "user.name", "Benchmark Test"],
            ["git", "config", "user.email", "benchmark@example.com"],
            ["git", "add", "."],
            ["git", "commit", "-qm", "fixture"],
        ):
            # These are fixed isolated Git fixture operations without a shell.
            subprocess.run(  # noqa: S603 - fixed isolated Git fixture operation
                command, cwd=self.repo, env=self.env, check=True, capture_output=True
            )
        self.source = subprocess.check_output(  # noqa: S603 - fixed Git query
            [self.executable("git"), "rev-parse", "HEAD"], cwd=self.repo, text=True
        ).strip()

    @staticmethod
    def executable(name):
        path = shutil.which(name)
        assert path is not None, f"required test executable {name} is unavailable"
        return path

    @staticmethod
    def write(path, text, executable=False):
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(text)
        if executable:
            path.chmod(0o755)

    def benchmark(self, cache, concurrency, phase, status=0):
        evidence = self.work / f"evidence-{cache}-{concurrency}-{phase}-{status}"
        env = self.env | {
            "EXPECTED_CONCURRENCY": str(concurrency),
            "WORKLOAD_EXIT": str(status),
            "GOFLAGS": "-p=99",
        }
        result = subprocess.run(  # noqa: S603 - fixed local fixture and test arguments
            [
                self.executable("bash"),
                str(self.repo / "scripts/run-provider-benchmark.sh"),
                self.source,
                "ghcr.io/f5-sales-demo/self-hosted-runner@sha256:" + "a" * 64,
                "hosted",
                cache,
                "1",
                str(concurrency),
                phase,
                str(evidence),
            ],
            cwd=self.repo,
            env=env,
            text=True,
            capture_output=True,
            check=False,
        )
        return result, evidence

    def test_effective_concurrency_and_cache_disk_evidence(self):
        for cache in ("cold", "warm"):
            for concurrency in (1, 4):
                for phase in ("build", "documentation-generation"):
                    with self.subTest(
                        cache=cache, concurrency=concurrency, phase=phase
                    ):
                        result, evidence = self.benchmark(cache, concurrency, phase)
                        assert result.returncode == 0, result.stdout + result.stderr
                        runtime = json.loads(
                            (evidence / "runtime-environment.json").read_text()
                        )
                        assert runtime["cache_state"] == cache
                        assert runtime["source_sha"] == self.source
                        assert runtime["go_environment"]["GOMAXPROCS"] == str(
                            concurrency
                        )
                        assert (
                            runtime["go_environment"]["GOFLAGS"] == f"-p={concurrency}"
                        )
                        for key in ("GOCACHE", "GOMODCACHE"):
                            assert Path(runtime["go_environment"][key]).is_absolute()
                            if cache == "cold":
                                assert "cold-1-" in runtime["go_environment"][key]
                        assert runtime["exit_code"] == 0
                        assert runtime["sample_count"] >= 2
                        assert runtime["filesystems"]
                        for disk in runtime["filesystems"]:
                            assert disk["total_bytes"] > 0
                            assert disk["peak_used_bytes"] >= disk["start_used_bytes"]
                            assert disk["peak_used_bytes"] >= disk["end_used_bytes"]

    def test_workload_failure_is_retained(self):
        result, evidence = self.benchmark(
            "cold", 4, "documentation-generation", status=7
        )
        assert result.returncode == 7, result.stdout + result.stderr
        assert (
            json.loads((evidence / "runtime-environment.json").read_text())["exit_code"]
            == 7
        )
        assert (
            json.loads((evidence / "output-manifest.json").read_text())["exit_code"]
            == 7
        )

    def test_sampling_failure_rejects_successful_workload(self):
        spec = importlib.util.spec_from_file_location(
            "runtime_profile", ROOT / "scripts/profile_provider_runtime.py"
        )
        assert spec is not None
        assert spec.loader is not None
        profiler = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(profiler)
        calls = 0

        def sample(_):
            nonlocal calls
            calls += 1
            if calls == 2:
                message = "fixture filesystem unavailable"
                raise OSError(message)
            return {"total_bytes": 1000, "used_bytes": 200, "available_bytes": 700}

        evidence = self.work / "sampling-error"
        evidence.mkdir()
        # All paths share one filesystem fixture so call 2 comes from the monitor.
        env = self.env | {
            key: str(self.repo) for key in ("RUNNER_TEMP", "GOCACHE", "GOMODCACHE")
        }
        env |= {
            "GOMAXPROCS": "4",
            "GOFLAGS": "-p=4",
            "GOMEMLIMIT": "4GiB",
            "GOGC": "20",
        }
        args = [
            "profile",
            str(evidence),
            self.source,
            "github-hosted",
            "cold",
            "4",
            "--",
            "python3",
            "-c",
            "import time; time.sleep(0.4)",
        ]
        previous = Path.cwd()
        try:
            os.chdir(self.repo)
            with (
                patch.dict(os.environ, env),
                patch.object(profiler.sys, "argv", args),
                patch.object(profiler, "disk_sample", sample),
                contextlib.redirect_stderr(io.StringIO()) as stderr,
            ):
                assert profiler.main() == 1
                assert "fixture filesystem unavailable" in stderr.getvalue()
        finally:
            os.chdir(previous)
        runtime = json.loads((evidence / "runtime-environment.json").read_text())
        assert runtime["exit_code"] == 0
        assert runtime["sampling_errors"] == ["fixture filesystem unavailable"]

    def test_signal_is_forwarded_and_retained(self):
        evidence = self.work / "signal"
        evidence.mkdir()
        marker = self.work / "child-ready"
        child_script = self.work / "child.py"
        self.write(
            child_script,
            "import pathlib,time\npathlib.Path("
            + repr(str(marker))
            + ").write_text('ready')\ntime.sleep(30)\n",
        )
        env = self.env | {
            "GOMAXPROCS": "4",
            "GOFLAGS": "-p=4",
            "GOMEMLIMIT": "4GiB",
            "GOGC": "20",
        }
        process = self.stack.enter_context(
            subprocess.Popen(  # noqa: S603 - fixed collector and isolated test child
                [
                    self.executable("python3"),
                    str(ROOT / "scripts/profile_provider_runtime.py"),
                    str(evidence),
                    self.source,
                    "github-hosted",
                    "cold",
                    "4",
                    "--",
                    "python3",
                    str(child_script),
                ],
                cwd=self.repo,
                env=env,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                text=True,
            )
        )
        try:
            deadline = time.monotonic() + 5
            while not marker.exists() and time.monotonic() < deadline:
                time.sleep(0.02)
            assert marker.exists(), "profiled child did not start"
            process.send_signal(signal.SIGTERM)
            stdout, stderr = process.communicate(timeout=5)
            assert process.returncode == 143, stdout + stderr
        finally:
            if process.poll() is None:
                process.kill()
                process.communicate()
        assert (
            json.loads((evidence / "runtime-environment.json").read_text())["exit_code"]
            == 143
        )

    def test_generation_failure_cannot_be_hidden_by_later_commands(self):
        for phase in ("provider-generation", "release-preflight"):
            with self.subTest(phase=phase):
                evidence = self.work / f"failure-{phase}"
                later = self.work / f"follow-on-{phase}"
                env = self.env | {
                    "FAIL_FIRST_GENERATOR": "true",
                    "FOLLOW_ON_LOG": str(later),
                }
                result = subprocess.run(  # noqa: S603 - fixed local phase fixture
                    [
                        self.executable("bash"),
                        str(self.repo / "scripts/run-provider-benchmark-phase.sh"),
                        phase,
                        "4",
                        str(evidence),
                    ],
                    cwd=self.repo,
                    env=env,
                    capture_output=True,
                    text=True,
                    check=False,
                )
                assert result.returncode != 0, result.stdout + result.stderr
                assert not later.exists(), "phase continued after failed generation"


if __name__ == "__main__":
    unittest.main()
