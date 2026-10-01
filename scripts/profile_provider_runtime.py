#!/usr/bin/env python3
"""Retain effective benchmark environment and filesystem usage around profiling."""

import json
import os
import signal
import subprocess
import sys
import threading
import time
from pathlib import Path


def existing_parent(path: str | Path) -> Path:
    """Resolve a cache path to the existing filesystem that will contain it."""
    path = Path(path).resolve()
    while not path.exists():
        path = path.parent
    return path


def disk_sample(path: str | Path) -> dict[str, int]:
    """Measure capacity, used space and user-available space in bytes."""
    stat = os.statvfs(existing_parent(path))
    return {
        "total_bytes": stat.f_blocks * stat.f_frsize,
        "used_bytes": (stat.f_blocks - stat.f_bfree) * stat.f_frsize,
        "available_bytes": stat.f_bavail * stat.f_frsize,
    }


class DiskMonitor:
    """Sample the filesystems used by a workload without changing its execution."""

    def __init__(self, paths: list[str]) -> None:
        """Capture initial disk state and initialize independent extrema."""
        self.first = {path: disk_sample(path) for path in paths}
        self.last = dict(self.first)
        self.peak = {path: disk["used_bytes"] for path, disk in self.first.items()}
        self.available = {
            path: disk["available_bytes"] for path, disk in self.first.items()
        }
        self.samples = 1
        self.errors: list[str] = []
        self.stopped = threading.Event()

    def sample(self) -> None:
        """Retain maxima and minima for every filesystem observation."""
        for path in self.first:
            disk = disk_sample(path)
            self.last[path] = disk
            self.peak[path] = max(self.peak[path], disk["used_bytes"])
            self.available[path] = min(self.available[path], disk["available_bytes"])
        self.samples += 1

    def monitor(self) -> None:
        """Record sampling errors so a successful command cannot qualify silently."""
        while not self.stopped.wait(0.2):
            try:
                self.sample()
            except OSError as error:
                self.errors.append(str(error))
                return

    def filesystems(self) -> list[dict[str, str | int]]:
        """Return independent start, peak, end and available-space evidence."""
        return [
            {
                "path": path,
                "total_bytes": first["total_bytes"],
                "start_used_bytes": first["used_bytes"],
                "peak_used_bytes": self.peak[path],
                "end_used_bytes": self.last[path]["used_bytes"],
                "minimum_available_bytes": self.available[path],
            }
            for path, first in self.first.items()
        ]


def run_profile(command: list[str], monitor: DiskMonitor) -> int:
    """Forward signals to the shared profiler while retaining disk observations."""
    thread = threading.Thread(target=monitor.monitor, daemon=True)
    thread.start()
    # The caller supplies its fixed shared-profiler argv; no shell interprets it.
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
            monitor.stopped.set()
            thread.join()
            for signum, handler in handlers.items():
                signal.signal(signum, handler)
        monitor.sample()
        return child.returncode if child.returncode >= 0 else 128 - child.returncode


def main() -> int:
    """Profile the authorized argv command and retain runtime evidence."""
    evidence, source, image, cache, concurrency, separator, *command = sys.argv[1:]
    if separator != "--" or not command:
        message = "expected profiling command after --"
        raise ValueError(message)
    environment = {
        key: os.environ[key]
        for key in (
            "GOMAXPROCS",
            "GOFLAGS",
            "GOMEMLIMIT",
            "GOGC",
            "GOCACHE",
            "GOMODCACHE",
        )
    }
    paths = sorted(
        {
            str(Path(path).resolve())
            for path in (
                Path.cwd(),
                os.environ["RUNNER_TEMP"],
                environment["GOCACHE"],
                environment["GOMODCACHE"],
            )
        }
    )
    os.sync()
    monitor = DiskMonitor(paths)
    started = time.monotonic()
    exit_code = run_profile(command, monitor)
    runtime = {
        "schema_version": 1,
        "source_sha": source,
        "observed_image_digest": image,
        "cache_state": cache,
        "go_concurrency": int(concurrency),
        "go_environment": environment,
        "sample_count": monitor.samples,
        "sample_interval_seconds": 0.2,
        "duration_seconds": time.monotonic() - started,
        "filesystems": monitor.filesystems(),
        "exit_code": exit_code,
        "sampling_errors": monitor.errors,
    }
    output = Path(evidence) / "runtime-environment.json"
    temporary = output.with_suffix(".json.tmp")
    temporary.write_text(json.dumps(runtime, sort_keys=True, indent=2) + "\n")
    temporary.replace(output)
    if monitor.errors:
        print(
            "benchmark filesystem sampling failed: " + "; ".join(monitor.errors),
            file=sys.stderr,
        )
    return exit_code or (1 if monitor.errors else 0)


if __name__ == "__main__":
    raise SystemExit(main())
