#!/usr/bin/env python3
"""Verify transport selection and existing retry behavior through executable fixtures."""

import json
import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path

SCRIPT = Path(__file__).resolve().parents[1] / "scripts/go-retry.sh"


class GoRetryTest(unittest.TestCase):
    """Exercise the real wrapper with isolated commands and no network requests."""

    def setUp(self):
        """Create an isolated PATH and a command that records every attempt."""
        self.root = Path(tempfile.mkdtemp())
        self.addCleanup(shutil.rmtree, self.root)
        self.record = self.root / "attempts.jsonl"
        self.python = shutil.which("python3")
        self.bash = shutil.which("bash")
        assert self.python
        assert self.bash
        command = r"""import json, os, pathlib, sys
record = pathlib.Path(os.environ["RECORD"])
rows = record.read_text().splitlines() if record.exists() else []
with record.open("a") as out:
    out.write(json.dumps({"args": sys.argv[1:], "godebug": os.environ.get("GODEBUG")}) + "\n")
if os.environ.get("REQUIRE_HTTP1") == "1":
    settings = dict(item.split("=", 1) for item in os.environ.get("GODEBUG", "").split(",") if "=" in item)
    if settings.get("http2client") != "0":
        print("stream error: stream ID 307; INTERNAL_ERROR; received from peer", file=sys.stderr)
        raise SystemExit(17)
if len(rows) < int(os.environ.get("FAILURES", "0")):
    print("fixture failure", file=sys.stderr)
    raise SystemExit(9)
print("fixture success")
"""
        for name in ("go", "other"):
            path = self.root / name
            path.write_text("#!" + self.python + "\n" + command)
            path.chmod(0o755)
        sleep = self.root / "sleep"
        sleep.write_text("#!/bin/sh\nexit 0\n")
        sleep.chmod(0o755)
        self.environment = os.environ | {
            "PATH": str(self.root) + os.pathsep + os.environ["PATH"],
            "RECORD": str(self.record),
        }
        self.environment.pop("GODEBUG", None)

    def run_wrapper(self, attempts, command="go", subcommand="mod", **environment):
        """Execute the real script and return exact stdout, stderr and records."""
        assert self.bash
        result = subprocess.run(  # noqa: S603 - fixed isolated executable fixture
            [
                self.bash,
                str(SCRIPT),
                str(attempts),
                command,
                subcommand,
                "tidy",
                "argument with spaces",
            ],
            env=self.environment | environment,
            capture_output=True,
            text=True,
            check=False,
        )
        rows = [json.loads(line) for line in self.record.read_text().splitlines()]
        return result, rows

    def test_http2_stream_failure_is_prevented(self):
        """The specific transport failure must disappear before retry is needed."""
        result, rows = self.run_wrapper(1, REQUIRE_HTTP1="1")
        assert result.returncode == 0, result.stderr
        assert rows[0]["godebug"] == "http2client=0"
        assert len(rows) == 1
        assert "INTERNAL_ERROR" not in result.stderr

    def test_existing_debug_settings_and_exact_arguments_survive(self):
        """Append only the transport override and preserve argument boundaries."""
        result, rows = self.run_wrapper(1, GODEBUG="gctrace=1,http2client=1")
        assert result.returncode == 0
        assert rows[0]["godebug"] == "gctrace=1,http2client=1,http2client=0"
        assert rows[0]["args"] == ["mod", "tidy", "argument with spaces"]

    def test_absolute_go_path_gets_transport_override(self):
        """Absolute Go binaries follow the same network transport contract."""
        result, rows = self.run_wrapper(1, command=str(self.root / "go"))
        assert result.returncode == 0
        assert rows[0]["godebug"] == "http2client=0"

    def test_non_go_environment_remains_unchanged(self):
        """Other commands retain their inherited debug settings."""
        result, rows = self.run_wrapper(1, command="other", GODEBUG="http2client=1")
        assert result.returncode == 0
        assert rows[0]["godebug"] == "http2client=1"

    def test_go_test_runtime_settings_remain_unchanged(self):
        """Go test binaries retain their existing HTTP/2 behavior."""
        result, rows = self.run_wrapper(1, subcommand="test", GODEBUG="http2client=1")
        assert result.returncode == 0
        assert rows[0]["godebug"] == "http2client=1"

    def test_recovered_failure_remains_visible(self):
        """Keep raw diagnostics, backoff and attempt count after recovery."""
        result, rows = self.run_wrapper(3, FAILURES="1")
        assert result.returncode == 0
        assert len(rows) == 2
        assert "fixture failure" in result.stderr
        assert "retrying in 20s" in result.stdout
        assert "Attempt 2 of 3" in result.stdout

    def test_terminal_failure_remains_failure(self):
        """Exhausted retries still fail and retain every command diagnostic."""
        result, rows = self.run_wrapper(2, FAILURES="3")
        assert result.returncode == 1
        assert len(rows) == 2
        assert result.stderr.count("fixture failure") == 2
        assert "failed after 2 attempts" in result.stdout


if __name__ == "__main__":
    unittest.main()
