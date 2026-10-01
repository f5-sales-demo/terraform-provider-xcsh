#!/usr/bin/env python3
"""Run a benchmark phase with bounded writes to its Go cache filesystems."""

import json
import os
import shutil
import signal
import subprocess
import sys
import threading
import time
from pathlib import Path


def cache_filesystems(paths: list[str]) -> list[str]:
    """Resolve cache destinations and deduplicate the filesystems to flush."""
    selected: dict[int, str] = {}
    for name in paths:
        path = Path(name).resolve()
        while not path.exists():
            path = path.parent
        selected.setdefault(path.stat().st_dev, str(path))
    return sorted(selected.values())


class CacheWriteback:
    """Flush cache filesystems while leaving disk sampling independent."""

    def __init__(self, paths: list[str], interval: float = 0.2) -> None:
        """Retain exact filesystem targets and explicit writeback outcomes."""
        self.paths = cache_filesystems(paths)
        sync_executable = shutil.which("sync")
        if not sync_executable:
            message = "sync executable is unavailable"
            raise ValueError(message)
        self.sync_executable = sync_executable
        self.interval = interval
        self.calls = 0
        self.errors: list[str] = []
        self.stopped = threading.Event()

    def flush(self) -> None:
        """Flush completed writes without reading or deleting cache contents."""
        for path in self.paths:
            try:
                result = subprocess.run(  # noqa: S603 - resolved sync and exact filesystem path
                    [self.sync_executable, "-f", path],
                    capture_output=True,
                    text=True,
                    timeout=30,
                    check=False,
                )
                self.calls += 1
                if result.returncode:
                    self.errors.append(
                        f"sync -f failed for {path}: {result.stderr.strip()}"
                    )
            except (OSError, subprocess.TimeoutExpired) as error:
                self.errors.append(f"sync -f failed for {path}: {error}")

    def monitor(self) -> None:
        """Bound speculative reservations during the measured workload."""
        while not self.stopped.wait(self.interval):
            self.flush()
            if self.errors:
                return


def run(command: list[str], writeback: CacheWriteback) -> int:
    """Include initial, periodic and final cache writeback in the measured phase."""
    writeback.flush()
    if writeback.errors:
        return 1
    thread = threading.Thread(target=writeback.monitor, daemon=True)
    thread.start()
    try:
        # The benchmark supplies fixed argv; no shell interprets the command.
        with subprocess.Popen(command) as child:  # noqa: S603
            handlers = {
                signum: signal.signal(
                    signum, lambda received, _: child.send_signal(received)
                )
                for signum in (signal.SIGINT, signal.SIGTERM)
            }
            try:
                child.wait()
            finally:
                for signum, handler in handlers.items():
                    signal.signal(signum, handler)
            status = (
                child.returncode if child.returncode >= 0 else 128 - child.returncode
            )
    finally:
        writeback.stopped.set()
        thread.join()
        writeback.flush()
    return status or (1 if writeback.errors else 0)


def main() -> int:
    """Execute the measured phase and retain cache-writeback provenance."""
    evidence, separator, *command = sys.argv[1:]
    if separator != "--" or not command:
        message = "expected benchmark command after --"
        raise ValueError(message)
    writeback = CacheWriteback([os.environ["GOCACHE"], os.environ["GOMODCACHE"]])
    started = time.monotonic()
    status = run(command, writeback)
    report = {
        "schema_version": 1,
        "filesystems": writeback.paths,
        "interval_seconds": writeback.interval,
        "flush_calls": writeback.calls,
        "errors": writeback.errors,
        "duration_seconds": time.monotonic() - started,
        "exit_code": status,
    }
    output = Path(evidence) / "cache-writeback.json"
    temporary = output.with_suffix(".json.tmp")
    temporary.write_text(json.dumps(report, sort_keys=True, indent=2) + chr(10))
    temporary.replace(output)
    for error in writeback.errors:
        print(error, file=sys.stderr)
    return status


if __name__ == "__main__":
    raise SystemExit(main())
