# ruff: noqa: INP001, PT027
"""Complete, byte-bounded Registry projection contracts."""

import re
import sys
import unittest
from pathlib import Path
from unittest.mock import patch

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "tools"))
import registry_projection as projection


class ProjectionTests(unittest.TestCase):
    def test_exact_utf8_boundary(self):
        assert projection.pack(["é" * 5], "", 10, 10) == [["é" * 5]]
        assert projection.pack(["12345", "67890"], "", 9, 10) == [["12345"], ["67890"]]
        with self.assertRaisesRegex(ValueError, "indivisible"):
            projection.pack(["é" * 6], "", 10, 10)

    def test_tables_split_between_rows_with_headers(self):
        header = "| Property | Description |\n| --- | --- |\n"
        rows = [f"| p{i} | {'x' * 30} |\n" for i in range(10)]
        pieces = projection.sections("## Properties\n\n" + header + "".join(rows), 180)
        assert len(pieces) > 1
        for piece in pieces:
            assert header in piece
        for row in rows:
            assert sum(piece.count(row) for piece in pieces) == 1

    def test_fences_and_paragraphs_are_indivisible(self):
        body = "# Example\n\n```hcl\n## not a section\n" + "x" * 100 + "\n```\n\n"
        assert projection.sections(body, 50) == [body]

    def page(self, identifier, body, role="properties"):
        return {
            "id": identifier,
            "collection_id": "fixture",
            "provider_type": "resources",
            "provider_name": "fixture",
            "role": role,
            "schema_path": ["same"],
            "path": "documentation/resources/fixture/" + identifier + "/index.md",
            "title": identifier,
            "body": body,
        }

    def test_grouped_collision_safe_fragments_and_coverage(self):
        first = self.page("first", '<a id="section"></a>\nfirst\n')
        second = self.page(
            "second",
            '<a id="section"></a>\n[other](https://example.test/resources/fixture/first/#section)\n',
        )
        outputs, manifest = projection.project(
            [first, second], {"fixture": ""}, "https://example.test"
        )
        assert len(outputs) == 1
        records = manifest["sections"]
        assert len(records) == 2
        assert (
            dict(records[0]["anchor_map"])["section"]
            != dict(records[1]["anchor_map"])["section"]
        )
        text = next(iter(outputs.values()))
        assert "#" + dict(records[0]["anchor_map"])["section"] in text
        assert all(record["mode"] == "embedded" for record in records)
        assert (outputs, manifest) == projection.project(
            [first, second], {"fixture": ""}, "https://example.test"
        )

    def test_landing_heading_and_schema_path_context(self):
        landing = self.page("landing", "# fixture\n\nOverview prose.\n", "fundamentals")
        property_page = self.page(
            "property",
            "# name\n\n## Direct properties\n\n### value property\n\nDescription.\n",
        )
        property_page["schema_path"] = ["parent", "name"]
        outputs, _ = projection.project(
            [landing, property_page], {"fixture": ""}, "https://example.test"
        )
        landing_text = outputs["docs/resources/fixture.md"]
        reference_text = next(
            text for path, text in outputs.items() if path.startswith("docs/guides/")
        )
        assert "# xcsh_fixture\n" in landing_text
        assert "## Overview\n" in landing_text
        assert "## `parent.name` properties" in reference_text
        assert "#### `parent.name.value` property" in reference_text
        assert re.search(r"/[0-3]{12}/[0-9]+", reference_text) is None

    def test_comment_preamble_stays_with_first_heading(self):
        body = "<!-- hidden -->\n\n# Title\n\nProse.\n"
        assert projection.sections(body, 1000) == [body]

    def test_parent_heading_stays_with_first_child_across_groups(self):
        body = (
            '## Direct properties\n\n<a id="property"></a>\n\n'
            "### name property\n\nComplete description.\n"
        )
        assert projection.sections(body, 1000) == [body]

    def test_published_group_routes_are_retained(self):
        pages = [
            self.page(str(index), f"# Section {index}\n\nDetails.\n")
            for index in range(3)
        ]
        with patch.dict(projection.ROUTE_LAYOUT, {"resources/fixture/reference": 3}):
            outputs, _ = projection.project(
                pages, {"fixture": ""}, "https://example.test"
            )
        assert len(outputs) == 3
        assert all(f"group-{index:03}.md" in " ".join(outputs) for index in range(1, 4))

    def test_oversized_exact_canonical_exception(self):
        page = self.page("huge", "```hcl\n" + "x" * 510000 + "\n```\n")
        outputs, manifest = projection.project(
            [page],
            {"fixture": ""},
            "https://example.test",
        )
        record = manifest["sections"][0]
        assert record["mode"] == "canonical-link"
        assert "/versions/" not in record["canonical_url"]
        assert record["canonical_url"] in next(iter(outputs.values()))
        assert record["reason"]
        assert max(len(text.encode()) for text in outputs.values()) <= 500000


if __name__ == "__main__":
    unittest.main()
