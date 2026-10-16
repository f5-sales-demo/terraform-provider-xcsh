"""Tests for exhaustive inventory and evidence-preserving refreshes."""

import importlib.util
import unittest
from pathlib import Path

import pytest

SPEC = importlib.util.spec_from_file_location(
    "smsv2_capabilities", Path(__file__).parents[1] / "scripts/smsv2_capabilities.py"
)
MODULE = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(MODULE)


class CapabilityTests(unittest.TestCase):
    def test_nested_attributes_blocks_and_selections(self):
        block = {
            "attributes": {
                "mode": {
                    "optional": True,
                    "description": "[OneOf: mode, other; Default: other] Choice",
                    "nested_type": {
                        "nesting_mode": "list",
                        "attributes": {"name": {"type": "string", "required": True}},
                    },
                },
                "other": {"type": "bool", "optional": True},
            },
            "block_types": {
                "nodes": {
                    "nesting_mode": "list",
                    "block": {"attributes": {"id": {"type": "string"}}},
                }
            },
        }
        rows, selections = MODULE.walk(block)
        assert [row["path"] for row in rows] == [
            "mode",
            "mode[].name",
            "other",
            "nodes",
            "nodes[].id",
        ]
        assert selections == [
            {"parent": "", "arms": ["mode", "other"], "default": "other"}
        ]

    def test_schema_presence_does_not_qualify_runtime(self):
        row = MODULE.capability(
            "resources", "xcsh_securemesh_site_v2", {"path": "aws", "kind": "attribute"}
        )
        assert row["classification"] == "unqualified"
        assert row["runtime"]["aws"] == {"status": "unqualified", "receipts": []}
        assert row["runtime"]["gcp"]["status"] == "deferred"
        assert row["lifecycle"] == {}

    def test_refresh_preserves_evidence_and_removals(self):
        row = MODULE.capability(
            "resources", "xcsh_securemesh_site_v2", {"path": "aws", "kind": "attribute"}
        )
        prior = dict(row, lifecycle={"create": ["receipt-1"]})
        assert MODULE.reconcile([row], [prior])[0]["lifecycle"] == {
            "create": ["receipt-1"]
        }
        removed = MODULE.reconcile([], [prior])[0]
        assert removed["schema"]["presence"] == "removed"
        assert removed["classification"] == "contract-blocked"

    def test_promotion_requires_runtime_receipts(self):
        row = MODULE.capability(
            "resources", "xcsh_securemesh_site_v2", {"path": "aws", "kind": "attribute"}
        )
        row["runtime"]["aws"]["status"] = "qualified"
        with pytest.raises(ValueError, match="requires receipts"):
            MODULE.reconcile([row], [row])

    def test_reconcile_rejects_duplicate_ids(self):
        row = MODULE.capability(
            "resources", "xcsh_securemesh_site_v2", {"path": "aws", "kind": "attribute"}
        )
        with pytest.raises(ValueError, match="duplicate"):
            MODULE.reconcile([row], [row, row])


if __name__ == "__main__":
    unittest.main()
