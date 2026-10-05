# ruff: noqa: INP001
"""Exact tagged snapshot contract tests."""

import importlib.util
import json
import tempfile
import unittest
from pathlib import Path
from unittest.mock import patch

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
                    "git", "ls-tree", "-r", "--name-only", "v12.0.4", "documentation"
                )
                .decode()
                .splitlines()
                if p.endswith(".md")
            }
            assert {d["path"] for d in manifest["documents"]} == expected
            assert manifest["document_count"] == len(expected)
            assert manifest["source_root"] == "documentation"
            assert all(
                d["path"].startswith("documentation/") for d in manifest["documents"]
            )
            assert all(
                d["metadata"].get("canonical_id") == d["metadata"]["id"]
                for d in manifest["documents"]
            )
            for file in (root / "first").iterdir():
                assert file.read_bytes() == (root / "second" / file.name).read_bytes()

    def test_complete_canonical_archive_includes_discovery(self) -> None:
        """Canonical receipts include source Markdown and progressive indexes."""
        with tempfile.TemporaryDirectory():
            assets, manifest = SNAPSHOT.canonical_assets("v12.0.6")
            expected = set(
                SNAPSHOT.run(
                    "git", "ls-tree", "-r", "--name-only", "v12.0.6", "documentation"
                )
                .decode()
                .splitlines()
            )
            assert {item["path"] for item in manifest["files"]} == expected
            assert "canonical-documentation.tar.gz" in assets
            assert (
                manifest["source_commit"]
                == SNAPSHOT.run("git", "rev-parse", "v12.0.6^{commit}").decode().strip()
            )

    def test_plain_markdown_and_invalid_metadata(self) -> None:
        """Plain Markdown gets fallback metadata; enriched hashes fail closed."""
        body, meta = SNAPSHOT.metadata(
            "documentation/guides/plain/index.md", b"# Plain\n"
        )
        assert body == "# Plain\n"
        assert meta["role"] == "overview"
        bad = {
            "path": "documentation/bad/index.md",
            "body_sha256": "sha256:" + "0" * 64,
            "body_bytes": 3,
        }
        data = ("---\nxcsh_docs: " + json.dumps(bad) + "\n---\n\nBad").encode()
        with self.assertRaisesRegex(  # noqa: PT027 - standard-library CI runner
            ValueError, "metadata body mismatch"
        ):
            SNAPSHOT.metadata("documentation/bad/index.md", data)
        with self.assertRaisesRegex(  # noqa: PT027 - standard-library CI runner
            ValueError, "only stable"
        ):
            SNAPSHOT.snapshot("docs-v12.0.4", Path("unused"))

    def test_snapshot_create_retries_only_without_remote_state(self) -> None:
        """A failed upload retries only after the release and tag are absent."""
        with (
            patch.object(
                SNAPSHOT,
                "run",
                side_effect=[RuntimeError("gh exited 1: HTTP 502"), b""],
            ) as command,
            patch.object(SNAPSHOT, "assert_no_partial_snapshot") as absent,
            patch.object(SNAPSHOT.time, "sleep") as delay,
        ):
            SNAPSHOT.create_snapshot_draft("documentation-v15.0.0", ["--draft"])
        assert command.call_count == 2
        absent.assert_called_once_with("documentation-v15.0.0")
        delay.assert_called_once_with(1)

        with (
            patch.object(
                SNAPSHOT, "run", side_effect=RuntimeError("gh exited 1")
            ) as command,
            patch.object(
                SNAPSHOT,
                "assert_no_partial_snapshot",
                side_effect=ValueError("snapshot creation left a draft"),
            ),
            self.assertRaisesRegex(ValueError, "left a draft"),  # noqa: PT027
        ):
            SNAPSHOT.create_snapshot_draft("documentation-v15.0.0", ["--draft"])
        assert command.call_count == 1

    def test_snapshot_rejects_unsealed_provider_release(self) -> None:
        """Documentation publication requires a final immutable provider."""
        provider = {
            "tag_name": "v15.0.0",
            "draft": False,
            "prerelease": False,
            "immutable": False,
        }
        with (
            patch.object(SNAPSHOT, "run", return_value=json.dumps(provider).encode()),
            self.assertRaisesRegex(ValueError, "stable, and immutable"),  # noqa: PT027
        ):
            SNAPSHOT.publish("v15.0.0", Path("unused"))


if __name__ == "__main__":
    unittest.main()
