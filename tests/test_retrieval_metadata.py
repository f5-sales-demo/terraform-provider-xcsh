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

    def test_aliases_do_not_describe_incidental_certificate_mentions(self):
        rules = MODULE.RetrievalRules.default()
        self.assertNotIn(
            "existing certificates",
            rules.aliases(
                "custom_hash_algorithms", "Algorithms used for TLS certificates."
            ),
        )
        self.assertNotIn(
            "existing certificates",
            rules.aliases(
                "https_auto_cert",
                "Choice for selecting HTTP proxy with bring your own certificates.",
            ),
        )
        self.assertIn("automatic certificates", rules.aliases("https_auto_cert"))
        self.assertIn("existing certificates", rules.aliases("tls_certificates"))

    def test_aliases_keep_credential_and_backend_context_local(self):
        rules = MODULE.RetrievalRules.default()
        self.assertNotIn(
            "credential setup",
            rules.aliases("status", "Status of authentication credentials."),
        )
        self.assertNotIn(
            "backend servers",
            rules.aliases(
                "connection_timeout", "Timeout for connections to origin servers."
            ),
        )
        self.assertIn("backend servers", rules.aliases("origin_servers"))
        self.assertIn("credential setup", rules.aliases("credentials"))
        self.assertNotIn(
            "existing certificates",
            rules.aliases("tls_parameters.tls_certificates.custom_hash_algorithms"),
        )
        self.assertNotIn("backend servers", rules.aliases("origin_servers.public_ip"))
        self.assertNotIn("credential setup", rules.aliases("authentication.status"))
        self.assertIn(
            "automatic certificates", rules.aliases("listener.https_auto_cert")
        )

    def test_reviewed_summary_is_scoped_to_exact_provider_path(self):
        rules = MODULE.RetrievalRules.default()
        text = rules.reviewed_summary(
            "resources", "http_loadbalancer", ["https_auto_cert"]
        )
        self.assertIn("automatic", text.lower())
        self.assertIsNone(
            rules.reviewed_summary(
                "resources", "http_loadbalancer", ["https_auto_cert", "port"]
            )
        )
        self.assertIsNone(
            rules.reviewed_summary("resources", "origin_pool", ["https_auto_cert"])
        )

    def test_invalid_reviewed_summary_rules_fail_before_generation(self):
        rule = {
            "provider_type": "resources",
            "collection": "fixture",
            "schema_path": ["tls"],
            "summary": "Reviewed TLS settings.",
            "review_basis": "Verified schema.",
        }
        with self.assertRaisesRegex(ValueError, "conflicting"):
            MODULE.RetrievalRules(
                {
                    "version": 1,
                    "summaries": [rule, {**rule, "summary": "Different wording."}],
                }
            )
        with self.assertRaisesRegex(ValueError, "invalid"):
            MODULE.RetrievalRules(
                {"version": 1, "summaries": [{**rule, "review_basis": ""}]}
            )

    def test_login_success_alias_requires_authentication_outcome_context(self):
        rules = MODULE.RetrievalRules.default()
        self.assertNotIn(
            "login success",
            rules.aliases(
                "notification_parameters.repeat_interval",
                "Send notification again after it was sent successfully.",
            ),
        )
        self.assertNotIn(
            "login success",
            rules.aliases(
                "tcp_health_check", "Connection succeeds when the backend is healthy."
            ),
        )
        self.assertIn(
            "login success",
            rules.aliases(
                "authentication.login.transaction_result.success_conditions",
                "Success Conditions.",
            ),
        )
        self.assertIn(
            "login success",
            rules.aliases(
                "transaction_result_criteria.transaction_result_success",
                "Success result.",
            ),
        )
        self.assertIn(
            "success",
            rules.aliases(
                "notification_parameters.repeat_interval",
                "Notification sent successfully.",
            ),
        )

    def test_success_alias_ignores_incidental_success_in_login_failure_descriptions(
        self,
    ):
        rules = MODULE.RetrievalRules.default()
        self.assertNotIn(
            "login success",
            rules.aliases(
                "login.transaction_result.failure_conditions",
                "Failure after successful connection.",
            ),
        )
        self.assertIn(
            "login success",
            rules.aliases(
                "login.transaction_result.success_conditions.status",
                "HTTP response code.",
            ),
        )
        with self.assertRaisesRegex(ValueError, "invalid retrieval term match source"):
            MODULE.RetrievalRules(
                {
                    "version": 1,
                    "terms": [
                        {"match_source": "unknown", "pattern": "success", "aliases": []}
                    ],
                }
            )

    def test_summary_preserves_complete_words(self):
        text = "complete " * 40
        result = MODULE.summary(text, "fallback")
        self.assertLessEqual(len(result), 320)
        self.assertEqual(result.split()[-1], "complete")

    def test_list_object_relationships(self):
        coverage = {
            "servers.address": {"document_id": "address", "anchor": "section"},
            "servers.name": {"document_id": "name", "anchor": "section"},
        }
        result = MODULE.constraint_relationships(
            ["servers"],
            {
                "Validators": 'validators.RequiredListObjectAttributes("address"), validators.ConflictingListObjectAttributes("address", "name"), validators.RequiredOneOfListObjectAttributes("address", "name")'
            },
            coverage,
        )
        self.assertEqual(
            {r["type"] for r in result}, {"requires", "conflicts", "choice"}
        )

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
