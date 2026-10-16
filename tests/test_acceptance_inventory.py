# ruff: noqa: INP001 -- Match the existing standalone repository test modules.
"""Verify native policy CRUD is generated with owned routes rather than action exclusions."""

import importlib.util
import json
import tempfile
import unittest
from pathlib import Path

source = Path(__file__).parents[1] / "tools/generate_acceptance_inventory.py"
spec = importlib.util.spec_from_file_location("acceptance_inventory", source)
assert spec is not None
assert spec.loader is not None
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)


class NativePolicyInventoryTests(unittest.TestCase):
    def test_native_policy_keeps_owned_read_verify_delete_handlers(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            (root / "internal/client").mkdir(parents=True)
            (root / "internal/client/secret_policy.go").write_text(
                (source.parents[1] / "internal/client/secret_policy.go").read_text()
            )
            (root / "internal/acctest").mkdir(parents=True)
            (root / "provider-release-surface.json").write_text(
                json.dumps({"resources": ["secret_policy", "public_ip_binding"]})
            )
            module.generate(root)
            handlers = (
                root / "internal/acctest/ownership_handlers_generated.go"
            ).read_text()
            for evidence in (
                "/api/secret_management/namespaces/%s/secret_policys",
                "c.GetSecretPolicy",
                "c.DeleteSecretPolicy",
                "readOwnershipProjection",
            ):
                assert evidence in handlers
            exclusions = json.loads(
                (
                    root / "internal/acctest/ownership_handler_exclusions.json"
                ).read_text()
            )["resources"]
            assert [entry["type"] for entry in exclusions] == ["public_ip_binding"]


if __name__ == "__main__":
    unittest.main()
