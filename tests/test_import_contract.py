# ruff: noqa: INP001, PT009, PT027
"""Import documentation contract regressions."""

import importlib.util
import sys
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / "tools"))
from import_contract import resolve_import_contract  # noqa: E402

SPEC = importlib.util.spec_from_file_location(
    "import_examples", ROOT / "tools/generate-import-examples.py"
)
assert SPEC is not None
assert SPEC.loader is not None
GENERATOR = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(GENERATOR)


class ImportContractTests(unittest.TestCase):
    def test_current_formats_preserve_all_segments(self):
        for syntax in (
            "namespace/name",
            "namespace/name/allowed_domain",
            "namespace/name/mitigated_domain",
            "namespace/name/protected_domain",
        ):
            with self.subTest(syntax=syntax):
                contract = resolve_import_contract(
                    "// Import ID format: "
                    + syntax
                    + "\nfunc (r *Fixture) ImportState() {}",
                    "fixture",
                )
                assert contract is not None
                expected = "system/" + "/".join(
                    ["example"] * (len(syntax.split("/")) - 1)
                )
                self.assertEqual(
                    contract["command"],
                    "terraform import xcsh_fixture.example " + expected + "\n",
                )

    def test_missing_unknown_and_duplicate_metadata_fail(self):
        for metadata in (
            "",
            "// Import ID format: namespace/name/unknown",
            "// Import ID format: name\n// Import ID format: name",
        ):
            with self.subTest(metadata=metadata), self.assertRaises(ValueError):
                resolve_import_contract(
                    metadata + "\nfunc (r *Fixture) ImportState() {}", "fixture"
                )
        self.assertIsNone(
            resolve_import_contract("func (r *Fixture) Read() {}", "fixture")
        )

    def test_generated_script_and_template(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            (root / "provider-release-surface.json").write_text(
                '{"resources":["namespace"]}'
            )
            source = root / "internal/provider/namespace_resource.go"
            source.parent.mkdir(parents=True)
            source.write_text(
                "// Import ID format: name (no namespace for this resource type)\nfunc (r *Namespace) ImportState() {}"
            )
            GENERATOR.generate(root)
            script = (root / "examples/resources/xcsh_namespace/import.sh").read_text()
            self.assertIn(
                "terraform import xcsh_namespace.this example-namespace\n", script
            )
            self.assertNotIn("system/", script)
            self.assertIn("tenant-level", script)
        template = (ROOT / "templates/resources.md.tmpl").read_text()
        self.assertIn("{{ if .HasImport }}", template)
        self.assertIn('{{ codefile "shell" .ImportFile }}', template)
        self.assertNotIn("system/example", template)

    def test_generation_fails_before_writing_unknown_format(self):
        with tempfile.TemporaryDirectory() as temporary:
            root = Path(temporary)
            (root / "provider-release-surface.json").write_text(
                '{"resources":["namespace","fixture"]}'
            )
            source_dir = root / "internal/provider"
            source_dir.mkdir(parents=True)
            (source_dir / "namespace_resource.go").write_text(
                "// Import ID format: name\nfunc (r *Namespace) ImportState() {}"
            )
            (source_dir / "fixture_resource.go").write_text(
                "// Import ID format: namespace/name/unknown\nfunc (r *Fixture) ImportState() {}"
            )
            with self.assertRaisesRegex(ValueError, "unrecognized import"):
                GENERATOR.generate(root)
            self.assertFalse((root / "examples").exists())
