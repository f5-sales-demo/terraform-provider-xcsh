# ruff: noqa: INP001, PT027
"""Complete, byte-bounded Registry projection contracts."""

import sys
import unittest
from pathlib import Path

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
            records[0]["anchor_map"]["section"] != records[1]["anchor_map"]["section"]
        )
        text = next(iter(outputs.values()))
        assert "#" + records[0]["anchor_map"]["section"] in text
        assert all(record["mode"] == "embedded" for record in records)
        assert (outputs, manifest) == projection.project(
            [first, second], {"fixture": ""}, "https://example.test"
        )

    def test_oversized_exact_canonical_exception(self):
        page = self.page("huge", "```hcl\n" + "x" * 510000 + "\n```\n")
        outputs, manifest = projection.project(
            [page], {"fixture": ""}, "https://example.test", version="v12.0.7"
        )
        record = manifest["sections"][0]
        assert record["mode"] == "canonical-link"
        assert "/versions/v12.0.7/" in record["canonical_url"]
        assert record["canonical_url"] in next(iter(outputs.values()))
        assert record["reason"]
        assert max(len(text.encode()) for text in outputs.values()) <= 500000


if __name__ == "__main__":
    unittest.main()
