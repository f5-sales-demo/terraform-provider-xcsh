#!/usr/bin/env python3
"""Exercise measured cache writeback through isolated executable fixtures."""

import json
import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path

SCRIPT = Path(__file__).resolve().parents[1] / "scripts/run_with_cache_writeback.py"


class CacheWritebackTest(unittest.TestCase):
    """Check filesystem scope, command status and independent disk sampling."""

    def setUp(self):
        """Create temporary caches and a sync spy with deterministic failures."""
        self.root = Path(tempfile.mkdtemp())
        self.addCleanup(shutil.rmtree, self.root)
        self.python = shutil.which("python3")
        assert self.python
        self.bin = self.root / "bin"
        self.bin.mkdir()
        self.evidence = self.root / "evidence"
        self.evidence.mkdir()
        self.calls = self.root / "calls.jsonl"
        sync = self.bin / "sync"
        sync.write_text(
            "#!"
            + self.python
            + chr(10)
            + r"""import json, os, pathlib, sys
log = pathlib.Path(os.environ["SYNC_LOG"])
calls = log.read_text().splitlines() if log.exists() else []
with log.open("a") as out:
    out.write(json.dumps(sys.argv[1:]) + chr(10))
if len(calls) >= int(os.environ.get("FAIL_AFTER", "999")):
    print("fixture writeback failure", file=sys.stderr)
    raise SystemExit(7)
"""
        )
        sync.chmod(0o755)
        self.environment = os.environ | {
            "PATH": str(self.bin) + os.pathsep + os.environ["PATH"],
            "GOCACHE": str(self.root / "cache/build"),
            "GOMODCACHE": str(self.root / "cache/mod"),
            "SYNC_LOG": str(self.calls),
        }

    def execute(self, code, **environment):
        """Run the actual wrapper and read its durable measurement record."""
        assert self.python
        result = subprocess.run(  # noqa: S603 - fixed isolated command fixture
            [
                self.python,
                str(SCRIPT),
                str(self.evidence),
                "--",
                self.python,
                "-c",
                code,
            ],
            env=self.environment | environment,
            capture_output=True,
            text=True,
            check=False,
        )
        report = json.loads((self.evidence / "cache-writeback.json").read_text())
        calls = [json.loads(line) for line in self.calls.read_text().splitlines()]
        return result, report, calls

    def test_cache_filesystem_scope_and_periodic_writeback(self):
        """Flush one shared filesystem before, during and after the command."""
        result, report, calls = self.execute("import time; time.sleep(0.65)")
        assert result.returncode == 0
        assert report["errors"] == []
        assert report["exit_code"] == 0
        assert report["duration_seconds"] >= 0.65
        assert len(calls) >= 4
        assert all(call == ["-f", str(self.root)] for call in calls)
        assert report["filesystems"] == [str(self.root)]

    def test_initial_failure_prevents_command(self):
        """An unavailable writeback boundary cannot silently qualify."""
        marker = self.root / "ran"
        result, report, calls = self.execute(
            "import pathlib; pathlib.Path(" + repr(str(marker)) + ").touch()",
            FAIL_AFTER="0",
        )
        assert result.returncode == 1
        assert not marker.exists()
        assert report["errors"]
        assert len(calls) == 1

    def test_periodic_failure_is_reported(self):
        """A successful command cannot hide a failed periodic sync."""
        result, report, calls = self.execute(
            "import time; time.sleep(0.45)", FAIL_AFTER="1"
        )
        assert result.returncode == 1
        assert report["errors"]
        assert len(calls) >= 2

    def test_command_failure_preserves_status(self):
        """Keep the workload's failure code and still flush completed writes."""
        result, report, calls = self.execute("raise SystemExit(19)")
        assert result.returncode == 19
        assert report["exit_code"] == 19
        assert len(calls) == 2

    def test_signal_exit_preserves_status(self):
        """A signalled workload remains failed after final writeback."""
        result, report, calls = self.execute(
            "import os, signal; os.kill(os.getpid(), signal.SIGTERM)"
        )
        assert result.returncode == 143
        assert report["exit_code"] == 143
        assert len(calls) == 2


if __name__ == "__main__":
    unittest.main()
