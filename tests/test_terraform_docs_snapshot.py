# ruff: noqa: INP001
"""Exact tagged snapshot contract tests."""

import importlib.util
import json
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location(
    "snapshot", ROOT / "scripts/terraform_docs_snapshot.py"
)
assert SPEC is not None
assert SPEC.loader is not None
SNAPSHOT = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(SNAPSHOT)


class SnapshotTests(unittest.TestCase):
    """Verify complete deterministic coverage and frontmatter validation."""

    def test_exact_tag_full_coverage_and_determinism(self) -> None:
        """The published baseline reproduces all Markdown bytes."""
        with tempfile.TemporaryDirectory() as folder:
            root = Path(folder)
            manifest = SNAPSHOT.snapshot("v12.0.4", root / "first")
            SNAPSHOT.snapshot("v12.0.4", root / "second")
            expected = {
                p
                for p in SNAPSHOT.run(
                    "git", "ls-tree", "-r", "--name-only", "v12.0.4", "docs"
                )
                .decode()
                .splitlines()
                if p.endswith(".md")
            }
            assert {d["path"] for d in manifest["documents"]} == expected
            assert manifest["document_count"] == 18965
            for file in (root / "first").iterdir():
                assert file.read_bytes() == (root / "second" / file.name).read_bytes()

    def test_plain_markdown_and_invalid_metadata(self) -> None:
        """Plain Markdown gets fallback metadata; enriched hashes fail closed."""
        body, meta = SNAPSHOT.metadata("docs/guides/plain.md", b"# Plain\n")
        assert body == "# Plain\n"
        assert meta["role"] == "overview"
        bad = {
            "path": "docs/bad.md",
            "body_sha256": "sha256:" + "0" * 64,
            "body_bytes": 3,
        }
        data = ("---\nxcsh_docs: " + json.dumps(bad) + "\n---\n\nBad").encode()
        with self.assertRaisesRegex(ValueError, "metadata body mismatch"):  # noqa: PT027 - standard-library CI runner
            SNAPSHOT.metadata("docs/bad.md", data)
        with self.assertRaisesRegex(ValueError, "only stable"):  # noqa: PT027 - standard-library CI runner
            SNAPSHOT.snapshot("docs-v12.0.4", Path("unused"))


if __name__ == "__main__":
    unittest.main()
