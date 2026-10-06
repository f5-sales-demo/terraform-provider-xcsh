# ruff: noqa: INP001, PT009
"""Objective readability rules distinguish prose from IDs and code."""

import importlib.util
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location(
    "doc_readability", ROOT / "tools/audit-doc-readability.py"
)
assert SPEC is not None
assert SPEC.loader is not None
AUDIT = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(AUDIT)


class ReadabilityTests(unittest.TestCase):
    def scan(self, body, relative="docs/resources/fixture.md"):
        with tempfile.TemporaryDirectory() as temporary:
            path = Path(temporary) / relative
            path.parent.mkdir(parents=True)
            path.write_text(body)
            return AUDIT.scan(path, relative)

    def test_visible_defects_are_errors(self):
        findings = self.scan(
            "# Fixture\n\n## Empty\n\n## Repeated\n\n"
            "Manages a HTTP Load Balancer resource.\n\n"
            "## Repeated\n\n"
            "## Detail / 012301230123 / 2\n\nUseful detail.\n"
        )
        categories = {item["category"] for item in findings}
        self.assertGreaterEqual(
            categories,
            {
                "empty_section",
                "duplicate_visible_heading",
                "numeric_heading_artifact",
                "known_grammar_defect",
            },
        )
        self.assertTrue(
            all(item["source"] and item["published_location"] for item in findings)
        )

    def test_invisible_ids_and_code_are_not_visible_defects(self):
        findings = self.scan(
            '# Fixture\n\n<a id="canonical-0123012301230123"></a>\n\n'
            "## Properties\n\n```hcl\n## Repeated / 012301230123 / 2\n```\n"
        )
        self.assertFalse([item for item in findings if item["severity"] == "error"])

    def test_redacted_ledger_keeps_location_and_severity(self):
        finding = {
            "category": "editorial_candidate",
            "detail": "private example prose",
            "published_location": "https://example.test/page",
            "severity": "review",
            "source": "documentation/page.md",
            "line": 4,
        }
        record = AUDIT.redact_detail(finding)
        assert "detail" not in record
        assert record["detail_sha256"].startswith("sha256:")
        assert record["source"] == finding["source"]
        assert record["published_location"] == finding["published_location"]
        assert record["severity"] == finding["severity"]

    def test_entry_boilerplate_is_review_only(self):
        findings = self.scan(
            "# Fixture\n\nThe source describes metadata.namespace configuration.\n\n"
            "## Prerequisites\n\nInstall Terraform.\n",
            "documentation/resources/fixture/index.md",
        )
        assert [item["category"] for item in findings] == ["editorial_candidate"]
        assert findings[0]["severity"] == "review"

    def test_repeated_substantive_paragraph_within_section(self):
        paragraph = "This description explains the exact behavior of the field and its required server-side constraints."
        findings = self.scan("# Fixture\n\n" + paragraph + "\n\n" + paragraph + "\n")
        self.assertIn("repeated_description", {item["category"] for item in findings})


if __name__ == "__main__":
    unittest.main()
