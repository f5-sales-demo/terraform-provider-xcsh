# ruff: noqa: INP001, PT009, PT027
"""Contract tests for complete schema traversal and progressive retrieval."""

import importlib.util
import json
import tempfile
import unittest
from itertools import pairwise
from pathlib import Path
from typing import ClassVar

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

    def test_namespace_import_contract(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            source = root / "internal/provider/namespace_resource.go"
            source.parent.mkdir(parents=True)
            source.write_text(
                "// Import ID format: name (no namespace for this resource type)\nfunc (r *NamespaceResource) ImportState() {}"
            )
            example = root / "examples/resources/xcsh_namespace/resource.tf"
            example.parent.mkdir(parents=True)
            example.write_text('resource "xcsh_namespace" "this" {}\n')
            collection = DOCS.Collection(
                "resources",
                "namespace",
                {"block": {}},
                FixtureSpecs(),
                {},
                root,
                "sha256:" + "c" * 64,
            )
            page = next(p for p in collection.pages.values() if p["role"] == "import")
            self.assertEqual(
                page["import"],
                "terraform import xcsh_namespace.this example-namespace\n",
            )
            self.assertIn("tenant-level", collection.body(page))
            self.assertIn("omitted", collection.body(page))

    def test_namespace_registry_projection_uses_bare_name(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            source = root / "internal/provider/namespace_resource.go"
            source.parent.mkdir(parents=True)
            source.write_text(
                "// Import ID format: name\nfunc (r *NamespaceResource) ImportState() {}"
            )
            example = root / "examples/resources/xcsh_namespace/resource.tf"
            example.parent.mkdir(parents=True)
            example.write_text(
                'provider "xcsh" {}\nresource "xcsh_namespace" "this" { name = "example-namespace" }\n'
            )
            collection = DOCS.Collection(
                "resources",
                "namespace",
                {"block": {}},
                FixtureSpecs(),
                {},
                root,
                "sha256:" + "c" * 64,
            )
            pages = list(collection.pages.values())
            for page in pages:
                page["body"] = collection.body(page)
            outputs = DOCS.registry_project(
                pages, {"xcsh-docs:resources:namespace:collection": ""}
            )
            imports = [
                content
                for content in outputs.values()
                if "terraform import " in content
            ]
            self.assertTrue(imports)
            self.assertTrue(
                all(
                    "terraform import xcsh_namespace.this example-namespace" in content
                    for content in imports
                )
            )
            self.assertTrue(all("system/example" not in content for content in imports))

    def test_unknown_import_contract_fails(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            source = root / "internal/provider/fixture_resource.go"
            source.parent.mkdir(parents=True)
            source.write_text(
                "// Import ID format: made-up\nfunc (r *FixtureResource) ImportState() {}"
            )
            with self.assertRaisesRegex(ValueError, "unrecognized import"):
                self.collection(root, {})

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
                text for path, text in outputs.items() if "reference--group-" in path
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
        with self.assertRaisesRegex(ValueError, "missing navigation target"):
            DOCS.registry_navigation(surface, outputs)

    def test_complete_registry_projection_coverage(self):
        manifest_path = ROOT / "documentation/registry-projection-manifest.json"
        if not manifest_path.exists():
            self.skipTest("grouped generation precedes corpus acceptance")
        projection = json.loads(manifest_path.read_text())
        index = json.loads(
            (ROOT / "documentation/terraform-llms-index.json").read_text()
        )
        pages = {page["id"]: page for page in index["pages"]}
        represented = {
            section["canonical_id"]
            for shard in projection["section_manifests"]
            for section in json.loads((ROOT / shard["path"]).read_text())["sections"]
        }
        self.assertEqual(
            represented,
            {
                identifier
                for identifier, page in pages.items()
                if page["role"] != "navigation"
            },
        )
        self.assertEqual(
            sum(page["role"] == "navigation" for page in pages.values()), 5
        )
        self.assertTrue(
            any(page["provider_type"] == "provider" for page in pages.values())
        )
        self.assertTrue(
            any(page["provider_type"] == "guides" for page in pages.values())
        )
        self.assertEqual(
            sum(len(collection["properties"]) for collection in index["collections"]),
            39557,
        )
        sections = [
            section
            for shard in projection["section_manifests"]
            for section in json.loads((ROOT / shard["path"]).read_text())["sections"]
        ]
        section_by_page: dict[str, list[dict]] = {}
        for section in sections:
            section_by_page.setdefault(section["canonical_id"], []).append(section)
            page = pages[section["canonical_id"]]
            self.assertEqual(section["canonical_source_sha256"], page["body_sha256"])
            self.assertIn(section["mode"], ("embedded", "canonical-link"))
            if section["mode"] == "canonical-link":
                self.assertIn("/versions/", section["canonical_url"])
                self.assertTrue(section["reason"])
        texts = {}
        for path, evidence in projection["files"].items():
            data = (ROOT / path).read_bytes()
            self.assertEqual(len(data), evidence["bytes"])
            self.assertEqual(DOCS.digest(data), evidence["sha256"])
            self.assertLessEqual(len(data), 500000)
            texts[path] = data.decode()
        for collection in index["collections"]:
            for target in collection["properties"].values():
                anchors = [
                    section
                    for section in section_by_page[target["document_id"]]
                    if target["anchor"] in dict(section["anchor_map"])
                ]
                self.assertTrue(anchors, target)
                for section in anchors:
                    self.assertIn(
                        'id="' + dict(section["anchor_map"])[target["anchor"]] + '"',
                        texts[section["registry_path"]],
                    )

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
        self.assertEqual(
            route[0]["id"], DOCS.stable_id("provider", "xcsh", "navigation")
        )
        fundamentals = next(
            page for page in route if page["id"] == collection["fundamentals_id"]
        )
        self.assertLess(fundamentals["body_bytes"], 8192)
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


class ImmutableSelectionLifecycleTests(unittest.TestCase):
    def test_lifecycle_guidance_links_to_each_selection(self):
        class SelectionSpecs(FixtureSpecs):
            def immutable_oneof_groups(self, name):
                return {"loadbalancer_type": ["http", "https", "https_auto_cert"]}

        with tempfile.TemporaryDirectory() as temporary:
            example = Path(temporary) / "examples/resources/xcsh_fixture/resource.tf"
            example.parent.mkdir(parents=True)
            example.write_text('resource "xcsh_fixture" "example" {}\n')
            collection = DOCS.Collection(
                "resources",
                "fixture",
                {
                    "block": {
                        "block_types": {
                            name: {
                                "nesting_mode": "single",
                                "block": {"attributes": {}},
                            }
                            for name in ("http", "https", "https_auto_cert")
                        }
                    }
                },
                SelectionSpecs(),
                {},
                Path(temporary),
                "sha256:" + "c" * 64,
            )
            page = collection.pages[DOCS.stable_id("resources", "fixture", "lifecycle")]
            body = collection.body(page)
            self.assertTrue(all(len(line) <= 400 for line in body.splitlines()))
            for phrase in (
                "requires recreation",
                "may interrupt service",
                "prevent_destroy",
                "create_before_destroy",
                "same selected type",
                "omitted",
            ):
                self.assertIn(phrase, body)
            for member in ("http", "https", "https_auto_cert"):
                self.assertIn(
                    collection.link(
                        DOCS.stable_id("resources", "fixture", "properties", (member,))
                    ),
                    body,
                )
