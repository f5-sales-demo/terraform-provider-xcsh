"""Exercise generation recovery through the actual workflow shell steps."""

import os
import re
import subprocess
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
WORKFLOW = ROOT / ".github/workflows/on-merge.yml"


def run_step(name):
    text = WORKFLOW.read_text()
    start = text.index("      - name: " + name + "\n")
    step = text[start:]
    match = re.search(r"        run: \|\n((?:          .*\n|\n)+)", step)
    if match is None:
        raise AssertionError("Missing workflow shell step: " + name)
    return "\n".join(line[10:] for line in match[1].splitlines())


class GenerationRecoveryTests(unittest.TestCase):
    def execute(self, script, **values):
        with tempfile.TemporaryDirectory() as temporary:
            output = Path(temporary) / "outputs"
            output.touch()
            git = Path(temporary) / "git"
            git.write_text(
                '#!/bin/sh\ncase "$1" in\n'
                'rev-parse) printf "%s\\n" "$SOURCE_COMMIT" ;;\n'
                'ls-remote) printf "%s\\t%s\\n" "$MAIN_SHA" refs/heads/main ;;\n'
                '*) printf "%s\\n" "${CHANGED_FILE:-tools/spec-deliveries.json}" ;;\nesac\n'
            )
            git.chmod(0o700)
            env = dict(os.environ, GITHUB_OUTPUT=str(output))
            env["PATH"] = temporary + os.pathsep + env["PATH"]
            env.update(values)
            result = subprocess.run(
                ["/bin/bash", "-euo", "pipefail"],
                input=script,
                cwd=ROOT,
                env=env,
                text=True,
                capture_output=True,
                check=False,
            )
            return result, dict(
                line.split("=", 1) for line in output.read_text().splitlines()
            )

    def decision(self, recovery="false", pending="false", **values):
        env = {
            "IS_REGENERATION": "false",
            "SPECS": "false",
            "CODE": "false",
            "TOOLS": "false",
            "DOCS_SOURCE": "false",
            "HAS_ANY_CHANGES": "false",
            "DELIVERY_METADATA_ONLY": "true",
            "PENDING_DELIVERY": pending,
            "RECONCILE_GENERATED": recovery,
        }
        env.update(values)
        result, outputs = self.execute(
            run_step("Determine if processing should proceed"), **env
        )
        assert result.returncode == 0, result.stderr
        return outputs

    def test_receipt_only_recovery_generates_both(self):
        outputs = self.decision(recovery="true")
        assert outputs["result"] == "true"
        assert outputs["provider_required"] == "true"
        assert outputs["docs_required"] == "true"

    def test_generation_workflow_repair_requires_regeneration(self):
        result, outputs = self.execute(
            run_step("Detect changed paths"),
            PENDING_RESUME="false",
            CHANGED_FILE=".github/workflows/on-merge.yml",
        )
        assert result.returncode == 0, result.stderr
        assert outputs["tools"] == "true"

    def test_normal_receipt_only_commit_skips(self):
        assert self.decision()["result"] == "false"

    def test_pending_recovery_still_generates_both(self):
        outputs = self.decision(pending="true")
        assert outputs["provider_required"] == "true"
        assert outputs["docs_required"] == "true"

    def test_regeneration_merge_still_releases(self):
        outputs = self.decision(IS_REGENERATION="true")
        assert outputs["result"] == "true"
        assert outputs["needs_regeneration"] == "false"

    def test_recovery_requires_main_dispatch(self):
        script = run_step("Identify an exact regeneration completion commit")
        script = script.split("MESSAGE=", 1)[0]
        for event, ref in (
            ("push", "refs/heads/main"),
            ("workflow_dispatch", "refs/heads/feature"),
        ):
            with self.subTest(event=event, ref=ref):
                result, _ = self.execute(
                    script,
                    RECONCILE_GENERATED="true",
                    PENDING_RESUME="false",
                    EVENT_NAME=event,
                    DISPATCH_REF=ref,
                )
                assert result.returncode != 0
                assert "current main" in result.stdout + result.stderr

    def test_main_dispatch_enters_fresh_generation(self):
        script = run_step("Identify an exact regeneration completion commit")
        script = script.split("MESSAGE=", 1)[0]
        result, outputs = self.execute(
            script,
            RECONCILE_GENERATED="true",
            PENDING_RESUME="false",
            EVENT_NAME="workflow_dispatch",
            DISPATCH_REF="refs/heads/main",
        )
        assert result.returncode == 0, result.stderr
        assert outputs.get("is_regeneration") == "false"

    def test_publication_rejects_changed_main(self):
        script = run_step("Verify and apply combined generation artifact")
        script = script.split("digest=", 1)[0]
        result, _ = self.execute(
            script,
            SOURCE_COMMIT="a" * 40,
            EXPECTED_SOURCE_SHA="a" * 40,
            MAIN_SHA="b" * 40,
            RUNNER_TEMP="unused-private-artifact",
        )
        assert result.returncode != 0
        assert "Protected main changed" in result.stdout

    def test_publication_requires_exact_source(self):
        script = run_step("Verify and apply combined generation artifact")
        script = script.split("digest=", 1)[0]
        result, _ = self.execute(
            script,
            SOURCE_COMMIT="a" * 40,
            EXPECTED_SOURCE_SHA="b" * 40,
            MAIN_SHA="a" * 40,
            RUNNER_TEMP="unused-private-artifact",
        )
        assert result.returncode != 0


if __name__ == "__main__":
    unittest.main()
