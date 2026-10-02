# ruff: noqa: INP001, PT009, PT027
"""Retrieval metadata contracts independent of generated corpora."""

import importlib.util
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SPEC = importlib.util.spec_from_file_location(
    "retrieval_metadata", ROOT / "tools/retrieval_metadata.py"
)
assert SPEC is not None
assert SPEC.loader is not None
MODULE = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(MODULE)


class RetrievalMetadataTests(unittest.TestCase):
    def test_specific_classification_and_task_separation(self):
        rules = MODULE.RetrievalRules.default()
        result = rules.classify(
            "resources",
            "http_loadbalancer",
            ["bot_defense", "login"],
            "properties",
            "Load Balancing",
        )
        self.assertEqual(result["retrieval_version"], 1)
        self.assertEqual(result["category"], "load-balancing")
        self.assertIn("security.bot-defense", result["capabilities"])
        self.assertIn("configuration", result["tasks"])
        self.assertNotIn("configuration", result["capabilities"])
        self.assertIn("reviewed-rule", result["classification"]["sources"])

    def test_unknown_classification_is_explicit(self):
        result = MODULE.RetrievalRules.default().classify(
            "resources", "unknown", [], "fundamentals", ""
        )
        self.assertIsNone(result["category"])
        self.assertEqual(result["classification"]["status"], "unresolved")

    def test_conflicting_rules_fail(self):
        with self.assertRaisesRegex(ValueError, "conflicting"):
            MODULE.RetrievalRules(
                {
                    "version": 1,
                    "rules": [
                        {
                            "provider_type": "resources",
                            "collection": "fixture",
                            "schema_prefix": [],
                            "category": "dns",
                        },
                        {
                            "provider_type": "resources",
                            "collection": "fixture",
                            "schema_prefix": [],
                            "category": "security",
                        },
                    ],
                    "terms": [],
                }
            )

    def test_aliases_are_reviewed_terminology(self):
        aliases = MODULE.RetrievalRules.default().aliases(
            "login_result", "Reports whether login was successful."
        )
        self.assertIn("login success", aliases)
        self.assertNotIn("benchmark", " ".join(aliases))

    def test_relations_distinguish_schema_and_advice(self):
        coverage = {
            "tls.certificates": {"document_id": "cert", "anchor": "section"},
            "tls.inline": {"document_id": "inline", "anchor": "section"},
            "tls.existing": {"document_id": "existing", "anchor": "section"},
        }
        result = MODULE.constraint_relationships(
            ["tls"],
            {
                "Validators": 'validators.RequiredObjectAttributes("certificates"), validators.ConflictingObjectAttributes("inline", "existing")'
            },
            coverage,
        )
        self.assertEqual({r["type"] for r in result}, {"requires", "conflicts"})
        self.assertTrue(all(r["enforcement"] == "provider-schema" for r in result))
        self.assertTrue(all(r["target_id"] and r["anchor"] for r in result))
        self.assertEqual(
            MODULE.constraint_relationships(
                ["tls"], {"description": "requires a namespace"}, coverage
            ),
            [],
        )


if __name__ == "__main__":
    unittest.main()
