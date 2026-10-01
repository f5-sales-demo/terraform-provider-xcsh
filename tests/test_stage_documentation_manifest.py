# ruff: noqa: INP001, PT009, PT027, S108, S603, S607
"""Verify manifest staging retains ignored schema paths and rejects unsafe output."""

import hashlib
import importlib.util
import json
import subprocess
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location(
    "stage_docs", ROOT / "scripts/stage-documentation-manifest.py"
)
assert SPEC is not None
assert SPEC.loader is not None
STAGE = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(STAGE)


class ManifestStagingTests(unittest.TestCase):
    def test_safe_ignored_output_is_verified_before_staging(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            subprocess.run(["git", "init", "-q", str(root)], check=True)
            (root / ".gitignore").write_text("site/\ntarget/\n")
            relative = "documentation/resources/example/properties/site/target/index.md"
            output = root / relative
            output.parent.mkdir(parents=True)
            output.write_bytes(b"complete reference\n")
            evidence = {
                "bytes": output.stat().st_size,
                "sha256": "sha256:" + hashlib.sha256(output.read_bytes()).hexdigest(),
            }
            manifest = root / "documentation/generated-manifest.json"
            manifest.write_text(json.dumps({"files": {relative: evidence}}))
            self.assertEqual(
                set(STAGE.verified_paths(root)),
                {relative, "documentation/generated-manifest.json"},
            )
            subprocess.run(
                ["python3", str(ROOT / "scripts/stage-documentation-manifest.py")],
                cwd=root,
                check=True,
                capture_output=True,
            )
            subprocess.run(
                [
                    "python3",
                    str(ROOT / "scripts/stage-documentation-manifest.py"),
                    "--check-tracked",
                ],
                cwd=root,
                check=True,
                capture_output=True,
            )
            output.write_bytes(b"edited reference\n")
            with self.assertRaises(ValueError):
                STAGE.verified_paths(root)

    def test_manifest_path_cannot_escape_or_follow_symlink(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            (root / "documentation").mkdir()
            manifest = root / "documentation/generated-manifest.json"
            for relative in ("../unrelated", "/tmp/unrelated", "scripts/unrelated"):
                manifest.write_text(
                    json.dumps({"files": {relative: {"bytes": 0, "sha256": ""}}})
                )
                with self.assertRaises(ValueError):
                    STAGE.verified_paths(root)
            target = root / "documentation/page.md"
            target.symlink_to(root / "unrelated")
            manifest.write_text(
                json.dumps(
                    {"files": {"documentation/page.md": {"bytes": 0, "sha256": ""}}}
                )
            )
            with self.assertRaises(ValueError):
                STAGE.verified_paths(root)


if __name__ == "__main__":
    unittest.main()
