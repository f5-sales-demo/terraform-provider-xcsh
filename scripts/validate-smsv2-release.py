#!/usr/bin/env python3
# pylint: disable=invalid-name
# ruff: noqa: TRY003,EM102,PLR2004,FURB162
"""Reject unproven or tampered AWS SMSv2 release capability assets."""

import datetime as dt
import hashlib
import json
import pathlib
import sys
from typing import NoReturn


def fail(message: str) -> NoReturn:
    """Exit with a release-validation failure."""
    raise SystemExit(f"SMSv2 release validation failed: {message}")


def load_json(path: pathlib.Path, description: str) -> dict:
    """Load a required JSON object or fail with its asset description."""
    try:
        value = json.loads(path.read_text())
    except (OSError, json.JSONDecodeError) as error:
        fail(f"malformed {description}: {error}")
    if not isinstance(value, dict):
        fail(f"malformed {description}: expected an object")
    return value


def validate_manifest(directory: pathlib.Path, tag: str, commit: str) -> dict:
    """Validate the manifest identity, asset set, and checksums."""
    manifest = load_json(directory / "smsv2-contract-manifest.json", "manifest")
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
    required = {"smsv2-contract.json", "smsv2-evidence-receipt.json"}
    if set(manifest["assets"]) != required:
        fail("manifest asset set is incomplete")
    for name in required:
        expected = manifest["assets"][name]
        actual = "sha256:" + hashlib.sha256((directory / name).read_bytes()).hexdigest()
        if expected != actual:
            fail(f"manifest checksum mismatch for {name}")
    return manifest


def validate_contract(contract: dict, contract_id: str) -> None:
    """Validate the evidence-backed SMSv2 provider and CRUD contract."""
    aws = contract.get("providers", {}).get("aws", {})
    if (
        contract.get("contract_id") != contract_id
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


def validate_evidence(evidence: dict, contract_id: str) -> None:
    """Validate the age and sanitization of the behavioral evidence."""
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
        evidence.get("contract_id") != contract_id
        or not isinstance(receipts, list)
        or not receipts
    ):
        fail("evidence receipt is malformed")
    if not all(
        item.get("sanitized") is True and item.get("redaction") for item in receipts
    ):
        fail("evidence receipt is not sanitized")


def validate_concurrency(concurrency: dict, version: str) -> None:
    """Validate complete provider-wide concurrency coverage or exclusion."""
    resources = concurrency.get("resources")
    exclusions = concurrency.get("exclusions")
    if not isinstance(resources, list) or not isinstance(exclusions, list):
        fail("provider-wide concurrency inventory is incomplete")
    counts_match = all(
        concurrency.get(key) == len(resources)
        for key in ("eligible_count", "covered_count")
    )
    exclusions_match = concurrency.get("excluded_count") == len(exclusions)
    if (
        concurrency.get("version") != version
        or not counts_match
        or not exclusions_match
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


def valid_parity_paths(paths: object) -> tuple[list[str], bool]:
    """Return named parity paths and whether the path collection is valid."""
    if not isinstance(paths, list):
        return [], False
    path_names = []
    for item in paths:
        if not isinstance(item, dict):
            return [], False
        path = item.get("path")
        if not isinstance(path, str) or not path:
            return [], False
        path_names.append(path)
    return path_names, True


def validate_parity(parity: dict, version: str) -> None:
    """Validate the exhaustive SMSv2 path classifications."""
    paths = parity.get("paths")
    path_names, paths_are_valid = valid_parity_paths(paths)
    choice_groups = parity.get("choice_groups")
    removals = {
        "spec.segment_vrf[].segment_config.nameserver_v6",
        "spec.segment_vrf[].segment_config.secondary_nameserver_v6",
    }
    identity_is_valid = (
        parity.get("version") == version
        and parity.get("resource") == "securemesh_site_v2"
    )
    path_inventory_is_valid = (
        paths_are_valid
        and len(set(path_names)) == len(path_names)
        and parity.get("path_count") == len(path_names)
        and isinstance(choice_groups, dict)
    )
    classifications_are_valid = (
        set(parity.get("deprecated_exclusions", []))
        == {"spec.log_receiver", "spec.private_adn", "spec.rseries"}
        and set(parity.get("current_platform_removals", [])) == removals
    )
    segment_contract_is_valid = (
        "spec.segment_vrf[].segment_network" in path_names
        and not any(path in path_names for path in removals)
    )
    if not all(
        (
            identity_is_valid,
            path_inventory_is_valid,
            classifications_are_valid,
            segment_contract_is_valid,
        )
    ):
        fail("SMSv2 nested parity manifest is incomplete")


def main() -> None:
    """Validate all release assets supplied on the command line."""
    if len(sys.argv) != 4:
        fail("usage: validate-smsv2-release.py ASSET_DIRECTORY RELEASE_TAG COMMIT")
    directory = pathlib.Path(sys.argv[1])
    tag, commit = sys.argv[2:]
    manifest = validate_manifest(directory, tag, commit)
    contract = load_json(directory / "smsv2-contract.json", "SMSv2 contract")
    evidence = load_json(directory / "smsv2-evidence-receipt.json", "SMSv2 evidence")
    concurrency = load_json(
        directory / "concurrency_contracts.json", "concurrency inventory"
    )
    parity = load_json(
        directory / "smsv2_parity_manifest.json", "SMSv2 parity manifest"
    )
    contract_id = manifest["contract_id"]
    version = tag.removeprefix("v")
    validate_contract(contract, contract_id)
    validate_evidence(evidence, contract_id)
    validate_concurrency(concurrency, version)
    validate_parity(parity, version)


if __name__ == "__main__":
    main()
