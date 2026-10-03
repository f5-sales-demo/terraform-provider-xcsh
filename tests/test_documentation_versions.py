# ruff: noqa: INP001
"""Immutable publication selection and version route contracts."""

import importlib.util
import json
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location(
    "versions", ROOT / "scripts/stage_documentation_versions.py"
)
assert SPEC
assert SPEC.loader
VERSIONS = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(VERSIONS)


class VersionTests(unittest.TestCase):
    def test_publication_config_is_loaded_only_from_canonical_root(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            canonical = root / "documentation"
            canonical.mkdir()
            (canonical / "llms-config.json").write_text(
                '{"canonicalCorpus":{"taxonomy":{}}}'
            )
            registry = root / "docs"
            registry.mkdir()
            (registry / "llms-config.json").write_text(
                "conflicting Registry configuration"
            )
            destination = root / "staged"
            destination.mkdir()
            VERSIONS.stage_publication_config(canonical, destination)
            assert json.loads((destination / "llms-config.json").read_text()) == {
                "canonicalCorpus": {"taxonomy": {}}
            }

    def test_stable_selection_requires_published_snapshot(self):
        releases = [
            {
                "tag_name": "v12.0.8",
                "draft": False,
                "prerelease": False,
                "immutable": True,
            },
            {
                "tag_name": "v12.0.7",
                "draft": False,
                "prerelease": False,
                "immutable": True,
            },
            {
                "tag_name": "documentation-v12.0.7",
                "draft": False,
                "prerelease": False,
                "immutable": True,
            },
            {
                "tag_name": "v12.0.6",
                "draft": False,
                "prerelease": False,
                "immutable": True,
            },
            {
                "tag_name": "documentation-v12.0.6",
                "draft": False,
                "prerelease": False,
                "immutable": True,
            },
        ]
        assert VERSIONS.select(releases) == ["v12.0.6", "v12.0.7"]

    def test_version_links_and_source_receipts_are_separate(self):
        text = "https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/"
        assert "/versions/v12.0.6/resources/" in VERSIONS.transform(
            text, "versions/v12.0.6"
        )
        assert "/preview/main/resources/" in VERSIONS.transform(text, "preview/main")
        assert VERSIONS.transform(text, "") == text


if __name__ == "__main__":
    unittest.main()
