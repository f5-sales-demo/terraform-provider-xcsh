# ruff: noqa: INP001, PT009
"""Check that current documentation is derived from the active provider surface."""

import json
import sys
import unittest
from pathlib import Path
from typing import ClassVar

ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / "tools"))
from retrieval_metadata import RetrievalRules  # noqa: E402

RETIRED = {
    "aws_vpc_site",
    "azure_vnet_site",
    "cloud_connect",
    "gcp_vpc_site",
    "securemesh_site",
    "voltstack_site",
}


class PublicationTests(unittest.TestCase):
    taxonomy: ClassVar[dict]
    index: ClassVar[dict]
    surface: ClassVar[dict]

    @classmethod
    def setUpClass(cls):
        cls.taxonomy = json.loads(
            (ROOT / "documentation/llms-config.json").read_text()
        )["canonicalCorpus"]["taxonomy"]
        cls.index = json.loads(
            (ROOT / "documentation/terraform-llms-index.json").read_text()
        )
        cls.surface = json.loads((ROOT / "provider-release-surface.json").read_text())

    def test_taxonomy_is_bound_to_final_rules_and_current_surface(self):
        digest = RetrievalRules.default().digest
        self.assertEqual(self.taxonomy["rulesDigest"], digest)
        self.assertTrue(
            all(
                page["classification"]["rules_sha256"] == digest
                for page in self.index["pages"]
            )
        )
        active = {
            name
            for key in ("resources", "data_sources", "actions", "ephemeral_resources")
            for name in self.surface[key]
        }
        active.update(
            {guide.stem for guide in (ROOT / "templates/guides").glob("*.md")}
        )
        active.add("setup")
        present = {page["provider_name"] for page in self.index["pages"]}
        mapped: set[str] = set()
        for group in self.taxonomy["subcategories"]:
            self.assertIn(group["category"], self.taxonomy["topics"])
            self.assertTrue(group["collections"], group["title"])
            for name in group["collections"]:
                self.assertIn(name, active)
                self.assertIn(name, present)
                self.assertNotIn(name, RETIRED)
                self.assertNotIn(name, mapped)
                mapped.add(name)
        self.assertNotIn("release-history", mapped)
        self.assertTrue(RETIRED.isdisjoint(mapped))
        self.assertEqual(len(self.taxonomy["topics"]), 11)

    def test_current_publication_has_no_version_migration_guide(self):
        retired_paths = (
            "docs/guides/release-history.md",
            "documentation/guides/release-history/index.md",
            "documentation/_data/pages/guides/release-history/index.txt",
        )
        for path in retired_paths:
            self.assertFalse((ROOT / path).exists(), path)
        pages = self.index["pages"]
        self.assertFalse(
            any(page["provider_name"] == "release-history" for page in pages)
        )
        for path in (
            "docs/index.md",
            "documentation/index.md",
            "documentation/llms.txt",
        ):
            body = (ROOT / path).read_text()
            self.assertNotIn("release-history", body, path)
            self.assertNotIn("Upgrading from v9 through v11", body, path)
        projection = json.loads(
            (ROOT / "documentation/registry-projection-manifest.json").read_text()
        )
        self.assertNotIn("docs/guides/release-history.md", projection["files"])
        manifest = json.loads(
            (ROOT / "documentation/generated-manifest.json").read_text()
        )
        self.assertFalse(
            any("release-history" in path for path in manifest["files"])
        )


if __name__ == "__main__":
    unittest.main()
