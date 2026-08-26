#!/usr/bin/env python3
# pylint: disable=invalid-name,too-many-branches
# ruff: noqa: D103,TRY003,EM102,PLR2004,FURB162
"""Reject unproven or tampered AWS SMSv2 release capability assets."""

import datetime as dt
import hashlib
import json
import pathlib
import sys


def fail(message: str) -> None:
    raise SystemExit(f"SMSv2 release validation failed: {message}")


def main() -> None:
    if len(sys.argv) != 4:
        fail("usage: validate-smsv2-release.py ASSET_DIRECTORY RELEASE_TAG COMMIT")
    directory = pathlib.Path(sys.argv[1])
    tag, commit = sys.argv[2:]
    required = {"smsv2-contract.json", "smsv2-evidence-receipt.json"}
    try:
        manifest = json.loads((directory / "smsv2-contract-manifest.json").read_text())
    except (OSError, json.JSONDecodeError) as error:
        fail(f"malformed manifest: {error}")
    if set(manifest) != {
        "assets",
        "contract_id",
        "contract_version",
        "release",
        "schema_version",
    }:
        fail("manifest fields are malformed")
    if (
        manifest["schema_version"] != 1
        or manifest["contract_id"] != "f5xc-ce-automation/v1"
    ):
        fail("manifest contract identity is unsupported")
    if manifest["release"] != {"tag": tag, "commit": commit}:
        fail("manifest release identity does not match the resolved tag")
    if set(manifest["assets"]) != required:
        fail("manifest asset set is incomplete")
    for name in required:
        expected = manifest["assets"][name]
        actual = "sha256:" + hashlib.sha256((directory / name).read_bytes()).hexdigest()
        if expected != actual:
            fail(f"manifest checksum mismatch for {name}")

    try:
        contract = json.loads((directory / "smsv2-contract.json").read_text())
        evidence = json.loads((directory / "smsv2-evidence-receipt.json").read_text())
        concurrency = json.loads((directory / "concurrency_contracts.json").read_text())
        parity = json.loads((directory / "smsv2_parity_manifest.json").read_text())
    except (OSError, json.JSONDecodeError) as error:
        fail(f"malformed SMSv2 asset: {error}")
    aws = contract.get("providers", {}).get("aws", {})
    if (
        contract.get("contract_id") != manifest["contract_id"]
        or aws.get("availability") != "evidence_backed"
    ):
        fail("AWS contract identity or evidence state is invalid")
    if aws.get("capabilities") != {
        "aws_ce_create": "available",
        "runtime_status": "unavailable",
        "tgw_connect": "unavailable",
    }:
        fail("AWS capability declaration is unsupported")
    api = contract.get("api", {})
    if api.get("namespace") != "system" or set(api.get("operations", [])) != {
        "create",
        "read",
        "replace",
        "delete",
    }:
        fail("system-namespace CRUD declaration is incomplete")
    try:
        observed_at = dt.datetime.fromisoformat(
            evidence["observed_at"].replace("Z", "+00:00")
        )
        if observed_at.tzinfo is None or observed_at.utcoffset() is None:
            fail("evidence observation timestamp must include a timezone")
    except (KeyError, AttributeError, ValueError):
        fail("evidence observation timestamp is invalid")
    if dt.datetime.now(dt.UTC) - observed_at > dt.timedelta(days=90):
        fail("evidence is stale")
    receipts = evidence.get("receipts")
    if (
        evidence.get("contract_id") != manifest["contract_id"]
        or not isinstance(receipts, list)
        or not receipts
    ):
        fail("evidence receipt is malformed")
    if not all(
        item.get("sanitized") is True and item.get("redaction") for item in receipts
    ):
        fail("evidence receipt is not sanitized")

    version = tag.removeprefix("v")
    resources = concurrency.get("resources")
    exclusions = concurrency.get("exclusions")
    if (
        concurrency.get("version") != version
        or not isinstance(resources, list)
        or concurrency.get("eligible_count") != len(resources)
        or concurrency.get("covered_count") != len(resources)
        or not isinstance(exclusions, list)
        or concurrency.get("excluded_count") != len(exclusions)
    ):
        fail("provider-wide concurrency inventory is incomplete")
    if not all(
        item.get("token") == "resource_version"
        and item.get("get", {}).get("schema")
        and item.get("replace", {}).get("schema")
        for item in resources
    ):
        fail("provider-wide concurrency inventory contains an invalid resource")
    if not all(item.get("api_identity") and item.get("reason") for item in exclusions):
        fail("concurrency exclusion lacks evidence-backed identity and reason")

    paths = parity.get("paths")
    path_names = (
        [item.get("path") for item in paths if isinstance(item, dict)]
        if isinstance(paths, list)
        else []
    )
    choice_groups = parity.get("choice_groups")
    removals = {
        "spec.segment_vrf[].segment_config.nameserver_v6",
        "spec.segment_vrf[].segment_config.secondary_nameserver_v6",
    }
    if (
        parity.get("version") != version
        or parity.get("resource") != "securemesh_site_v2"
        or not isinstance(paths, list)
        or len(path_names) != len(paths)
        or not all(isinstance(path, str) and path for path in path_names)
        or len(set(path_names)) != len(path_names)
        or parity.get("path_count") != len(path_names)
        or not isinstance(choice_groups, dict)
        or set(parity.get("deprecated_exclusions", []))
        != {"spec.log_receiver", "spec.private_adn", "spec.rseries"}
        or set(parity.get("current_platform_removals", [])) != removals
        or "spec.segment_vrf[].segment_network" not in path_names
        or any(path in path_names for path in removals)
    ):
        fail("SMSv2 nested parity manifest is incomplete")


if __name__ == "__main__":
    main()
