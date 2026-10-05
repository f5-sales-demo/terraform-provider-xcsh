# ruff: noqa: INP001, PT009
"""Check the authored taxonomy and release record against generated docs."""

import json
import re
import sys
import unittest
from pathlib import Path
from typing import ClassVar

ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / "tools"))
from retrieval_metadata import RetrievalRules  # noqa: E402

VERSIONS = (
    "v12.0.0",
    "v12.0.1",
    "v12.0.2",
    "v12.0.3",
    "v12.0.4",
    "v12.0.5",
    "v12.0.6",
    "v12.0.7",
    "v12.1.0",
    "v12.1.1",
    "v12.1.2",
    "v12.2.0",
    "v12.2.1",
    "v12.3.0",
    "v12.3.1",
    "v12.3.2",
    "v12.4.0",
    "v13.0.0",
    "v13.0.1",
    "v13.0.2",
    "v13.0.3",
    "v13.1.0",
    "v13.1.1",
    "v14.0.0",
)
RETIRED = {
    "aws_vpc_site",
    "azure_vnet_site",
    "cloud_connect",
    "gcp_vpc_site",
    "securemesh_site",
}


class PublicationTests(unittest.TestCase):
    taxonomy: ClassVar[dict]
    index: ClassVar[dict]
    surface: ClassVar[dict]
    changelog: ClassVar[str]

    @classmethod
    def setUpClass(cls):
        cls.taxonomy = json.loads(
            (ROOT / "documentation/llms-config.json").read_text()
        )["canonicalCorpus"]["taxonomy"]
        cls.index = json.loads(
            (ROOT / "documentation/terraform-llms-index.json").read_text()
        )
        cls.surface = json.loads((ROOT / "provider-release-surface.json").read_text())
        cls.changelog = (ROOT / "CHANGELOG.md").read_text()

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
        active.update({"setup", "release-history"})
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
        self.assertIn("release-history", mapped)
        self.assertTrue(RETIRED.isdisjoint(mapped))
        self.assertEqual(len(self.taxonomy["topics"]), 11)

    def test_release_record_and_generated_guide_links(self):
        headings = re.findall(
            r"^## \[(v(?:12|13|14)\.\d+\.\d+)\].*$", self.changelog, re.MULTILINE
        )
        self.assertEqual(set(headings), set(VERSIONS))
        self.assertEqual(len(headings), 24)
        for version in VERSIONS:
            self.assertRegex(
                self.changelog,
                rf"(?m)^## \[{re.escape(version)}\]\([^\n]+\) - 20\d\d-\d\d-\d\d$",
            )
        section = self.changelog.split("## [v14.0.0]", 1)[1].split("## [v13.1.1]", 1)[0]
        for kind in ("resource", "data source"):
            for name in RETIRED:
                self.assertIn(f"- {kind} `xcsh_{name}`", section)
        self.assertIn("does not establish a drop-in replacement", section)
        expected = self.changelog.replace("# Changelog\n", "# Release history\n", 1)
        self.assertEqual(
            (ROOT / "docs/guides/release-history.md").read_text(), expected
        )
        canonical = (ROOT / "documentation/guides/release-history/index.md").read_text()
        self.assertIn(expected, canonical)
        self.assertIn(
            "/guides/release-history/", (ROOT / "documentation/llms.txt").read_text()
        )
        self.assertIn(
            "/guides/release-history/", (ROOT / "documentation/index.md").read_text()
        )
        guide = [
            page
            for page in self.index["pages"]
            if page["id"] == "xcsh-docs:guides:release-history:overview"
        ]
        self.assertEqual(len(guide), 1)
        self.assertEqual(
            guide[0]["path"], "documentation/guides/release-history/index.md"
        )
        self.assertEqual(guide[0]["registry_path"], "docs/guides/release-history.md")
        self.assertEqual(
            guide[0]["aliases"],
            ["changelog", "provider releases", "release history"],
        )
        self.assertEqual(guide[0]["tasks"], [])
        self.assertEqual(guide[0]["category"], "administration")
        self.assertIn(
            "reviewed-publication-taxonomy",
            guide[0]["classification"]["sources"],
        )


if __name__ == "__main__":
    unittest.main()
