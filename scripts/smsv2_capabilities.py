#!/usr/bin/env python3
"""Inventory installed SMSv2 schemas without inferring runtime acceptance."""

import argparse
import copy
import hashlib
import json
import re
from pathlib import Path
from typing import Any

PROVIDER = "registry.terraform.io/f5-sales-demo/xcsh"
SURFACES = {
    "securemesh_site_v2",
    "enhanced_firewall_policy",
    "forward_proxy_policy",
    "bgp",
    "bgp_asn_set",
    "bgp_routing_policy",
    "route",
    "network_interface",
    "segment",
    "segment_connection",
    "dc_cluster_group",
    "site_mesh_group",
    "log_receiver",
    "network_global_log_receiver",
    "global_log_receiver",
    "network_firewall",
    "token",
    "registration",
    "registration_approval",
    "site",
    "site_bgp_status",
    "site_cloud_init",
    "site_image",
    "site_registration",
    "site_registrations",
    "site_registrations_by_site",
    "site_registrations_by_state",
    "site_upgrade_status",
    "site_signatures_update",
    "site_upgrade_sw",
    "site_upgrade_os",
    "smsv2_contract",
    "smsv2_aws_runtime",
    "smsv2_kvm_runtime",
    "smsv2_kvm_runtime_interface",
    "aws_tgw_site",
}
KINDS = {
    "resource_schemas": "resources",
    "data_source_schemas": "data_sources",
    "action_schemas": "actions",
    "ephemeral_resource_schemas": "ephemeral_resources",
}
EVIDENCE_FIELDS = (
    "classification",
    "validation",
    "lifecycle",
    "runtime",
    "outstanding",
)


def typed_children(value: Any, parent: str) -> list:
    """Enumerate structural object types that omit nested schema descriptions."""
    if not isinstance(value, list):
        return []
    kind, children = value
    if kind in ("list", "set", "map"):
        return typed_children(children, parent + ("{}" if kind == "map" else "[]"))
    if kind != "object":
        return []
    rows = []
    for name, child in sorted(children.items()):
        path = f"{parent}.{name}"
        rows.append({"path": path, "kind": "type_member", "type": child})
        rows.extend(typed_children(child, path))
    return rows


def walk(block: dict[str, Any], parent: str = "") -> tuple[list, list]:
    """Enumerate every attribute and block, including nested object attributes."""
    rows, selections = [], []
    for category in ("attributes", "block_types"):
        for name, node in sorted(block.get(category, {}).items()):
            path = f"{parent}.{name}" if parent else name
            rows.append(
                {
                    "path": path,
                    "kind": "attribute" if category == "attributes" else "block",
                    **{
                        key: node[key]
                        for key in (
                            "required",
                            "optional",
                            "computed",
                            "sensitive",
                            "type",
                            "nesting_mode",
                            "min_items",
                            "max_items",
                        )
                        if key in node
                    },
                }
            )
            description = node.get("description", "")
            for match in re.finditer(
                r"\[OneOf: ([^;\]]+)(?:; Default: ([^\]]+))?\]", description
            ):
                selection = {
                    "parent": parent,
                    "arms": sorted(arm.strip() for arm in match[1].split(",")),
                    "default": match[2],
                }
                if selection not in selections:
                    selections.append(selection)
            nested = node.get("nested_type") or node.get("block")
            if nested is not None:
                mode = node.get("nesting_mode", nested.get("nesting_mode", "single"))
                suffix = (
                    "[]" if mode in ("list", "set") else "{}" if mode == "map" else ""
                )
                children, child_selections = walk(nested, path + suffix)
                rows.extend(children)
                selections.extend(child_selections)
            else:
                rows.extend(typed_children(node.get("type"), path))
    return rows, selections


def capability(kind: str, name: str, row: dict[str, Any]) -> dict[str, Any]:
    """Keep schema presence separate from all acceptance dimensions."""
    return {
        "id": f"{kind}/{name}/{row['path']}",
        "schema": {"presence": "present", **row},
        "classification": "unqualified",
        "validation": {"status": "unqualified", "receipts": []},
        "lifecycle": {},
        "runtime": {
            cloud: {"status": status, "receipts": []}
            for cloud, status in (
                ("aws", "unqualified"),
                ("azure", "unqualified"),
                ("gcp", "deferred"),
                ("oracle", "deferred"),
            )
        },
        "outstanding": [
            "Qualify applicable lifecycle and runtime behavior on disposable topology."
        ],
    }


def reconcile(rows: list, previous: list) -> list:
    """Preserve receipts and flag removals rather than silently dropping them."""
    old = {row["id"]: row for row in previous}
    if len(old) != len(previous) or len({row["id"] for row in rows}) != len(rows):
        message = "duplicate capability IDs"
        raise ValueError(message)
    result = []
    for current in rows:
        row = copy.deepcopy(current)
        prior = old.pop(row["id"], {})
        for key in EVIDENCE_FIELDS:
            if key in prior:
                row[key] = copy.deepcopy(prior[key])
        result.append(row)
    for prior_row in old.values():
        row = copy.deepcopy(prior_row)
        row["schema"]["presence"] = "removed"
        row["classification"] = "contract-blocked"
        row["outstanding"] = ["Review removed schema path against upstream contract."]
        result.append(row)
    for row in result:
        for evidence in row["runtime"].values():
            if evidence["status"] == "qualified" and not evidence["receipts"]:
                message = "runtime qualification requires receipts: " + row["id"]
                raise ValueError(message)
    return sorted(result, key=lambda row: row["id"])


def inventory(document: dict, pins: dict, previous: dict | None = None) -> dict:
    """Build a frozen ledger of current operation surfaces and selection groups."""
    provider = document["provider_schemas"][PROVIDER]
    rows: list[dict[str, Any]] = []
    operations: list[dict[str, Any]] = []
    selections: list[dict[str, Any]] = []
    for schema_kind, kind in KINDS.items():
        for name, schema in sorted(provider.get(schema_kind, {}).items()):
            short = name.removeprefix("xcsh_")
            if short not in SURFACES:
                continue
            attributes, groups = walk(schema.get("block", {}))
            if "config_schema" in schema:
                attributes, groups = walk(schema["config_schema"])
            source_kind = {
                "resources": "resource",
                "data_sources": "data_source",
                "actions": "action",
            }.get(kind, "ephemeral_resource")
            operations.append(
                {
                    "id": f"{kind}/{name}",
                    "schema_owner": f"internal/provider/{short}_{source_kind}.go",
                    "wire_owner": f"internal/client/{short}_types.go",
                    "owner_status": "candidate paths; verify source before defect repair",
                    "cloud_applicability": "scope-dependent; aws/azure qualification, gcp/oracle live deferred",
                }
            )
            rows.extend(capability(kind, name, row) for row in attributes)
            selections.extend(
                {"operation": f"{kind}/{name}", **group} for group in groups
            )
    present = {operation["id"] for operation in operations}
    required = {
        "resources/xcsh_securemesh_site_v2",
        "resources/xcsh_enhanced_firewall_policy",
        "resources/xcsh_forward_proxy_policy",
        "data_sources/xcsh_site_registration",
        "actions/xcsh_site_upgrade_sw",
        "actions/xcsh_site_upgrade_os",
    }
    if not required <= present:
        message = "required SMSv2 surfaces absent: " + str(sorted(required - present))
        raise ValueError(message)
    return {
        "ledger_version": 1,
        "umbrella_issue": "https://github.com/f5-sales-demo/terraform-provider-xcsh/issues/2586",
        "pins": pins,
        "operations": operations,
        "selections": selections,
        "capabilities": reconcile(rows, (previous or {}).get("capabilities", [])),
    }


def source_audit(root: Path, operations: list) -> list:
    """Record actual source ownership and constraints omitted by schema JSON."""
    records = []
    for operation in operations:
        name = operation["id"].split("/")[1].removeprefix("xcsh_")
        candidates = [root / operation[key] for key in ("schema_owner", "wire_owner")]
        candidates.extend(sorted((root / "internal/client").glob(name + "*.go")))
        for source in sorted(set(candidates)):
            if not source.is_file() or source.name.endswith("_test.go"):
                continue
            contents = source.read_bytes()
            constraints = [
                {"line": number, "expression": line.strip()}
                for number, line in enumerate(contents.decode().splitlines(), 1)
                if any(
                    term in line
                    for term in (
                        "OneOf:",
                        "OneOf(",
                        "OneOfObject",
                        "OneOfList",
                        "Conflicting",
                        "mutually exclusive",
                        "RequiresReplace",
                    )
                )
            ]
            records.append(
                {
                    "operation": operation["id"],
                    "source": source.relative_to(root).as_posix(),
                    "sha256": hashlib.sha256(contents).hexdigest(),
                    "selection_and_lifecycle_constraints": constraints,
                }
            )
    return records


def main() -> None:
    """Read explicit schema and revision pins, then write a deterministic ledger."""
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--schema", type=Path, required=True)
    parser.add_argument("--pins", type=Path, required=True)
    parser.add_argument("--previous", type=Path)
    parser.add_argument("--source-root", type=Path, required=True)
    parser.add_argument("--output", type=Path, required=True)
    args = parser.parse_args()
    schema_bytes = args.schema.read_bytes()
    pins = json.loads(args.pins.read_text())
    actual = hashlib.sha256(schema_bytes).hexdigest()
    if pins["schema_sha256"] != actual:
        message = "schema digest differs from frozen pins"
        raise ValueError(message)
    previous = json.loads(args.previous.read_text()) if args.previous else None
    ledger = inventory(json.loads(schema_bytes), pins, previous)
    ledger["source_audit"] = source_audit(args.source_root, ledger["operations"])
    args.output.write_text(json.dumps(ledger, indent=2, sort_keys=True) + "\n")


if __name__ == "__main__":
    main()
