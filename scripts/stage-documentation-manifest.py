#!/usr/bin/env python3
# pylint: disable=invalid-name
"""Verify and stage exact manifest-owned documentation, including ignored paths."""

from __future__ import annotations

import argparse
import hashlib
import json
import shutil
import subprocess
import sys
from pathlib import Path, PurePosixPath


def verified_paths(root: Path) -> list[str]:
    """Reject unsafe or changed output before touching the Git index."""
    manifest = json.loads((root / "documentation/generated-manifest.json").read_text())
    paths = []
    for relative, evidence in manifest["files"].items():
        candidate = PurePosixPath(relative)
        target = root / relative
        if (
            candidate.is_absolute()
            or ".." in candidate.parts
            or candidate.parts[0] not in ("docs", "documentation")
        ):
            message = f"unsafe documentation manifest path: {relative}"
            raise ValueError(message)
        if target.is_symlink() or root.resolve() not in target.resolve().parents:
            message = f"documentation manifest path escapes repository: {relative}"
            raise ValueError(message)
        data = target.read_bytes()
        if (
            len(data) != evidence["bytes"]
            or "sha256:" + hashlib.sha256(data).hexdigest() != evidence["sha256"]
        ):
            message = f"documentation output differs from manifest: {relative}"
            raise ValueError(message)
        paths.append(relative)
    paths.append("documentation/generated-manifest.json")
    return paths


def main() -> None:
    """Verify output and either force-stage it or require complete tracked coverage."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--check-tracked", action="store_true")
    args = parser.parse_args()
    root = Path.cwd()
    paths = verified_paths(root)
    git = shutil.which("git") or "/usr/bin/git"
    if args.check_tracked:
        tracked = set(
            subprocess.run([git, "ls-files", "-z"], check=True, capture_output=True)  # noqa: S603 - resolved Git, fixed arguments
            .stdout.decode()
            .split("\0")
        )
        missing = sorted(set(paths) - tracked)
        if missing:
            message = f"documentation manifest contains {len(missing)} untracked outputs: {missing[:5]}"
            raise ValueError(message)
    else:
        # NUL-delimited stdin avoids argv limits and shell interpretation. The
        # complete allowlist passed byte/digest and repository-boundary checks.
        payload = "\0".join(paths).encode() + b"\0"
        subprocess.run(  # noqa: S603 - resolved Git and fixed argument array
            [git, "add", "-f", "--pathspec-from-file=-", "--pathspec-file-nul"],
            input=payload,
            check=True,
        )
    print(f"Verified {len(paths)} manifest-owned documentation paths.")


if __name__ == "__main__":
    try:
        main()
    except (OSError, ValueError, subprocess.CalledProcessError) as error:
        print(str(error), file=sys.stderr)
        sys.exit(1)
