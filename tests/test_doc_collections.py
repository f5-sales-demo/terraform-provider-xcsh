# ruff: noqa: INP001, PT009
"""Contract tests for complete schema traversal and progressive retrieval."""

import importlib.util
import json
import tempfile
import unittest
from itertools import pairwise
from pathlib import Path
from typing import ClassVar

import pytest

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location(
    "doc_collections", ROOT / "tools/generate-doc-collections.py"
)
assert SPEC is not None
assert SPEC.loader is not None
DOCS = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(DOCS)


class FixtureSpecs:
    pin: ClassVar[dict] = {"release_tag": "v1.0.0", "target_commit": "a" * 40}
    pin_digest = "sha256:" + "b" * 64

    def resource(self, name):
        return {}, {}, {}, []

    def field(self, roots, schemas, path):
        return {}


class CollectionTests(unittest.TestCase):
    def collection(self, root, block):
        example = root / "examples/resources/xcsh_fixture/resource.tf"
        example.parent.mkdir(parents=True)
        example.write_text('resource "xcsh_fixture" "example" {}\n')
        return DOCS.Collection(
            "resources",
            "fixture",
            {"block": block},
            FixtureSpecs(),
            {},
            root,
            "sha256:" + "c" * 64,
        )

    def test_deep_repeated_empty_choices_and_sensitive_objects(self):
        block = {
            "attributes": {
                "empty": {
                    "type": ["object", {}],
                    "optional": True,
                    "description": "[OneOf: empty, feature] Choose one.",
                }
            },
            "block_types": {
                "feature": {
                    "nesting_mode": "list",
                    "min_items": 1,
                    "max_items": 3,
                    "block": {
                        "attributes": {
                            "secret": {
                                "nested_type": {
                                    "nesting_mode": "single",
                                    "attributes": {
                                        "value": {"type": "string", "optional": True}
                                    },
                                },
                                "sensitive": True,
                                "optional": True,
                            }
                        },
                        "block_types": {
                            "feature": {
                                "nesting_mode": "single",
                                "block": {
                                    "attributes": {
                                        "value": {"type": "string", "required": True}
                                    },
                                    "block_types": {
                                        "empty": {"nesting_mode": "single", "block": {}}
                                    },
                                },
                            }
                        },
                    },
                }
            },
        }
        with tempfile.TemporaryDirectory() as tmp:
            collection = self.collection(Path(tmp), block)
            expected = {
                "empty",
                "feature",
                "feature.secret",
                "feature.secret.value",
                "feature.feature",
                "feature.feature.value",
                "feature.feature.empty",
            }
            self.assertEqual(set(collection.coverage), expected)
            target = collection.coverage["feature.secret.value"]
            page = collection.pages[target["document_id"]]
            body = collection.body(page)
            self.assertIn("Sensitive", body)
            empty = collection.pages[collection.coverage["empty"]["document_id"]]
            self.assertIn("empty = {}", collection.body(empty))
            nested = collection.pages[
                collection.coverage["feature.feature.empty"]["document_id"]
            ]
            self.assertIn("empty {}", collection.body(nested))
            feature = collection.pages[collection.coverage["feature"]["document_id"]]
            self.assertIn("min_items: `1`", collection.body(feature))
            self.assertIn("max_items: `3`", collection.body(feature))
            self.assertIn("OneOf alternatives", collection.body(empty))
            self.assertNotEqual(
                collection.coverage["feature"]["document_id"],
                collection.coverage["feature.feature"]["document_id"],
            )

    def test_large_coherent_leaf_is_complete(self):
        sentinel = "END_OF_COMPLETE_PROPERTY"
        description = "A complete contextual sentence. " * 6000 + sentinel
        with tempfile.TemporaryDirectory() as tmp:
            collection = self.collection(
                Path(tmp),
                {
                    "attributes": {
                        "large": {
                            "type": "string",
                            "description": description,
                            "optional": True,
                        }
                    }
                },
            )
            reference = collection.pages[collection.coverage["large"]["document_id"]]
            body = collection.body(reference)
            self.assertIn(
                " ".join(DOCS.description_markdown(description).split()),
                " ".join(body.split()),
            )
            self.assertGreater(len(body.encode()), 8192)
            self.assertIn(DOCS.description_markdown(sentinel), body)

    def test_ids_are_content_independent_and_preserve_full_path(self):
        path = ("repeated", "repeated", "long_schema_identifier")
        identifier = DOCS.stable_id("resources", "fixture", "properties", path)
        self.assertEqual(
            identifier,
            "xcsh-docs:resources:fixture:properties:repeated:repeated:long_schema_identifier",
        )
        self.assertNotEqual(
            identifier, DOCS.stable_id("data-sources", "fixture", "properties", path)
        )

    def test_long_registry_filename_preserves_identifier(self):
        page = {
            "provider_type": "resources",
            "provider_name": "fixture",
            "role": "properties",
            "schema_path": ["long_contextual_subsection_name"] * 12,
            "id": DOCS.stable_id(
                "resources",
                "fixture",
                "properties",
                ("long_contextual_subsection_name",) * 12,
            ),
        }
        projected = DOCS.projection_name(page)
        self.assertLess(len((projected + ".md").encode()), 255)
        self.assertEqual(projected, DOCS.projection_name(page))
        self.assertIn("long_contextual_subsection_name", page["id"])

    def test_registry_semantic_subdivision_preserves_rows(self):
        with tempfile.TemporaryDirectory() as tmp:
            collection = self.collection(
                Path(tmp),
                {"attributes": {"name": {"type": "string", "optional": True}}},
            )
            pages = list(collection.pages.values())
            for page in pages:
                page["body"] = collection.body(page)
            reference = collection.pages[
                DOCS.stable_id("resources", "fixture", "reference")
            ]
            rows = [
                f"| `field_{i}` | Complete contextual reference {chr(120) * 200} |\n"
                for i in range(3000)
            ]
            reference["body"] += (
                "\n## All schema paths\n\n| Schema path | Complete reference |\n| --- | --- |\n"
                + "".join(rows)
            )
            outputs = DOCS.registry_project(pages, {reference["collection_id"]: ""})
            parts = [
                text for path, text in outputs.items() if "reference--part-" in path
            ]
            self.assertGreater(len(parts), 1)
            for row in rows:
                self.assertEqual(sum(text.count(row) for text in parts), 1)
            for path, text in outputs.items():
                self.assertLessEqual(len(text.encode()), DOCS.REGISTRY_LIMIT, path)

    def test_registry_navigation_uses_exact_surface_and_valid_targets(self):
        surface = {key: ["fixture"] for _, key, _, _ in DOCS.TYPES.values()}
        outputs = {f"docs/{kind}/fixture.md": "# fixture\n" for kind in DOCS.TYPES}
        DOCS.registry_navigation(surface, outputs)
        for kind in DOCS.TYPES:
            page = outputs[f"docs/{kind}/index.md"]
            self.assertIn("includes 1 ", page)
            self.assertIn("[xcsh_fixture](fixture.md)", page)
            self.assertNotIn("functions", page)
        del outputs["docs/resources/fixture.md"]
        with pytest.raises(ValueError, match="missing navigation target"):
            DOCS.registry_navigation(surface, outputs)

    def test_generated_manifest_and_progressive_http_navigation(self):
        manifest_file = ROOT / "documentation/generated-manifest.json"
        if not manifest_file.exists():
            self.skipTest(
                "canonical generation is performed before this acceptance test"
            )
        manifest = json.loads(manifest_file.read_text())
        for relative, evidence in manifest["files"].items():
            data = (ROOT / relative).read_bytes()
            self.assertEqual(len(data), evidence["bytes"], relative)
            self.assertEqual(DOCS.digest(data), evidence["sha256"], relative)
        index = json.loads(
            (ROOT / "documentation/terraform-llms-index.json").read_text()
        )
        collection = next(
            c
            for c in index["collections"]
            if c["provider_type"] == "resources"
            and c["provider_name"] == "http_loadbalancer"
        )
        pages = {p["id"]: p for p in index["pages"]}
        target_path = (
            "advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name"
        )
        self.assertIn(target_path, collection["properties"])
        target = pages[collection["properties"][target_path]["document_id"]]
        route = []
        current = target
        while current:
            route.append(current)
            current = pages.get(current["parent_id"])
        route.reverse()
        self.assertEqual(route[0]["id"], collection["fundamentals_id"])
        self.assertLess(route[0]["body_bytes"], 8192)
        for parent, child in pairwise(route):
            self.assertIn(child["id"], parent["child_ids"])
            self.assertTrue(
                set(child["schema_path"]).issubset(set(target["schema_path"]))
            )
        self.assertFalse(
            any(
                "bot_defense" in p["schema_path"] or "waf" in p["schema_path"]
                for p in route
            )
        )
        # Retrieve only this route, and verify whole leaf bytes against metadata.
        for page in route:
            markdown = (ROOT / page["path"]).read_text()
            body = markdown.split("---\n\n", 1)[1]
            self.assertEqual(DOCS.digest(body.encode()), page["body_sha256"])
        self.assertIn("### name", (ROOT / target["path"]).read_text())


if __name__ == "__main__":
    unittest.main()
