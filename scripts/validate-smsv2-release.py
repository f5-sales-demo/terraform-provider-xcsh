#!/usr/bin/env python3
# pylint: disable=invalid-name,too-many-branches
# ruff: noqa: TRY003,EM102,PLR2004,FURB162
"""Reject unproven or tampered AWS SMSv2 release capability assets."""

import datetime as dt
import hashlib
import json
import pathlib
import sys
from typing import NoReturn

CONTRACT_ID = "f5xc-ce-automation/v3"
TELEMETRY_SCHEMA_ID = "f5xc-smsv2-aws-tgw-telemetry/v2"
REQUIRED_FACTS = {
    "runtime",
    "gre",
    "bgp",
    "mtu",
    "route",
    "bgp_inside_cidr_block",
}
BASE_CAPABILITIES = {
    "aws_ce_create": "available",
    "runtime_status": "available",
    "tgw_connect": "available",
}
UPGRADE_CAPABILITIES = {**BASE_CAPABILITIES, "site_upgrade": "available"}
BASE_UNAVAILABLE_CAPABILITIES = {
    "aws_ce_create": "unavailable",
    "runtime_status": "unavailable",
    "tgw_connect": "unavailable",
}
UPGRADE_UNAVAILABLE_CAPABILITIES = {
    **BASE_UNAVAILABLE_CAPABILITIES,
    "site_upgrade": "unavailable",
}
BLOCKING_CONDITIONS = [
    "mac_only_interface_rejected_by_live_api",
    "public_ip_empty_string_null_round_trip",
]
BASE_AUTHORITIES = {
    "f5xc": [
        "smsv2_configuration",
        "runtime_health",
        "bgp_peers",
        "bgp_routes",
        "simplified_routes",
    ],
    "aws": [
        "eni",
        "transit_gateway",
        "transit_gateway_connect",
        "gre_endpoints",
        "bgp_inside_cidrs",
        "autonomous_system_numbers",
    ],
}
UPGRADE_AUTHORITIES = {
    **BASE_AUTHORITIES,
    "f5xc": [*BASE_AUTHORITIES["f5xc"], "site_upgrade_observation"],
}
RUNTIME_ENDPOINTS = {
    "configuration": {
        "method": "GET",
        "path": "/api/config/namespaces/{namespace}/securemesh_site_v2s/{site}",
        "operation_id": "ves.io.schema.views.securemesh_site_v2.API.Get",
        "response_schema": "securemesh_site_v2GetResponse",
        "authority": "f5xc",
        "semantics": "configuration",
    },
    "health": {
        "method": "GET",
        "path": "/api/operate/namespaces/system/sites/{site}/vpm/debug/global/health",
        "operation_id": "ves.io.schema.operate.debug.CustomPublicAPI.HealthPublic",
        "response_schema": "debugHealthResponse",
        "authority": "f5xc",
        "semantics": "observational_read_only",
    },
    "bgp_peers": {
        "method": "GET",
        "path": "/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_peers",
        "operation_id": "ves.io.schema.operate.bgp.CustomPublicAPI.ShowBGPPeers",
        "response_schema": "bgpBGPPeersResponse",
        "authority": "f5xc",
        "semantics": "observational_read_only",
    },
    "bgp_routes": {
        "method": "GET",
        "path": "/api/operate/namespaces/{namespace}/sites/{site}/ver/bgp_routes",
        "operation_id": "ves.io.schema.operate.bgp.CustomPublicAPI.ShowBGPRoutes",
        "response_schema": "bgpBGPRoutesResponse",
        "authority": "f5xc",
        "semantics": "observational_read_only",
    },
    "simplified_routes": {
        "method": "POST",
        "path": "/api/operate/namespaces/{namespace}/sites/{site}/ver/simplified_routes",
        "operation_id": "ves.io.schema.operate.route.CustomPublicAPI.ShowSimplifiedRoutes",
        "request_schema": "routeSimplifiedRouteRequest",
        "response_schema": "routeSimplifiedRouteResponse",
        "authority": "f5xc",
        "semantics": "observational_read_only",
    },
}
SITE_UPGRADE_OPERATIONS = {
    "site_status": {
        "method": "GET",
        "path": "/api/config/namespaces/{namespace}/sites/{site}",
        "operation_id": "ves.io.schema.site.API.Get",
        "response_schema": "siteGetResponse",
        "response_mappings": {
            "software_installed_version": "status[].volterra_software_status.last_installed_version",
            "software_available_version": "status[].volterra_software_status.available_version",
            "software_deployment_phase": "status[].volterra_software_status.deployment_state.phase",
            "software_deployment_result": "status[].volterra_software_status.deployment_state.result",
            "os_installed_version": "status[].operating_system_status.deployment_state.version",
            "os_available_version": "status[].operating_system_status.available_version",
            "os_deployment_phase": "status[].operating_system_status.deployment_state.phase",
            "os_deployment_result": "status[].operating_system_status.deployment_state.result",
            "site_state": "spec.site_state",
        },
    },
    "target_discovery": {
        "method": "GET",
        "path": "/api/maurice/upgradable_sw_versions",
        "operation_id": "ves.io.schema.upgrade_status.UpgradeStatusCustomApi.GetUpgradableSWVersions",
        "response_schema": "upgrade_statusGetUpgradableSWVersionsResponse",
        "query_mappings": {
            "installed_os_version": "current_os_version",
            "installed_software_version": "current_sw_version",
        },
        "response_mappings": {"upgradable_software_versions": "sw_versions[]"},
    },
    "precheck": {
        "method": "GET",
        "path": "/api/maurice/namespaces/{namespace}/sites/{site}/pre_upgrade_check",
        "operation_id": "ves.io.schema.upgrade_status.UpgradeStatusCustomApi.PreUpgradeCheck",
        "response_schema": "upgrade_statusPreUpgradeCheckResponse",
        "query_mappings": {"software_version": "sw_version"},
        "response_mappings": {
            "checks": "checklist[]",
            "name": "checklist[].item",
            "status": "checklist[].status",
        },
        "passing_statuses": ["CHECKLIST_PASSED", "CHECKLIST_WARNING"],
        "failure_statuses": ["CHECKLIST_FAILED", "CHECKLIST_UNKNOWN"],
    },
    "upgrade_status": {
        "method": "GET",
        "path": "/api/maurice/namespaces/{namespace}/sites/{site}/upgrade_status",
        "operation_id": "ves.io.schema.upgrade_status.UpgradeStatusCustomApi.GetUpgradeStatus",
        "response_schema": "upgrade_statusGetUpgradeStatusResponse",
        "response_mappings": {
            "version": "upgrade_status.sw_upgrade_progress.version",
            "status": "upgrade_status.sw_upgrade_progress.status",
            "site_level_status": "upgrade_status.sw_upgrade_progress.site_level_upgrade.status",
            "node_level_status": "upgrade_status.sw_upgrade_progress.node_level_upgrade.status",
            "validation_status": "upgrade_status.sw_upgrade_progress.validation.status",
            "os_setup_status": "upgrade_status.sw_upgrade_progress.os_setup.status",
        },
    },
    "software_upgrade": {
        "method": "POST",
        "path": "/api/config/namespaces/{namespace}/sites/{site}/upgrade_sw",
        "operation_id": "ves.io.schema.site.UpgradeAPI.UpgradeSW",
        "request_schema": "siteUpgradeSWRequest",
        "request_mappings": {
            "site": "name",
            "software_version": "version",
            "force": "force",
        },
        "force": False,
        "semantics": "asynchronous",
    },
    "os_upgrade": {
        "method": "POST",
        "path": "/api/config/namespaces/{namespace}/sites/{site}/upgrade_os",
        "operation_id": "ves.io.schema.site.UpgradeAPI.UpgradeOS",
        "request_schema": "siteUpgradeOSRequest",
        "request_mappings": {
            "site": "name",
            "os_version": "version",
            "force": "force",
        },
        "force": False,
        "semantics": "asynchronous",
    },
}
SITE_UPGRADE_SEMANTICS = {
    "eligibility": {
        "site_state": "ONLINE",
        "os_target": "equals_advertised_os_available_version",
        "software_target": "listed_in_upgradable_software_versions",
        "software_prechecks": "all_pass_or_warning",
    },
    "polling": {
        "transient_failure_values": ["UPGRADE_FAILED", "FAILED"],
        "failure_semantics": "transient_until_bounded_timeout",
        "completion": "supplied_targets_installed_and_site_online",
        "timeout_authority": "caller",
    },
    "redaction": {
        "exported": "sanitized_status_fields_only",
        "prohibited": ["raw_api_messages", "node_identifiers", "urls"],
    },
}
SITE_UPGRADE_EVIDENCE_FACTS = {
    "installed_and_available_versions",
    "deployment_phase_and_result",
    "site_state",
    "software_target_advertised",
    "software_prechecks_passed",
    "transient_failure_observed",
}


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


def version_at_least(version: object, minimum: tuple[int, int]) -> bool:
    """Return whether a dotted contract version meets the required major/minor."""
    if not isinstance(version, str):
        return False
    parts = version.split(".")
    if len(parts) != 3 or not all(part.isdigit() for part in parts):
        return False
    return tuple(int(part) for part in parts[:2]) >= minimum


def validate_site_upgrade(site_upgrade: object, contract_version: object) -> None:
    """Validate exact upgrade operations and fail-closed observation semantics."""
    if not isinstance(site_upgrade, dict) or not version_at_least(
        contract_version, (6, 1)
    ):
        fail("site upgrade requires a v6.1-or-newer contract")
    expected_keys = {
        *SITE_UPGRADE_OPERATIONS,
        *SITE_UPGRADE_SEMANTICS,
        "verified_path",
    }
    if set(site_upgrade) != expected_keys:
        fail("site upgrade operation inventory is incomplete")
    for name, expected in SITE_UPGRADE_OPERATIONS.items():
        if site_upgrade.get(name) != expected:
            fail(f"site upgrade operation {name} is incomplete or legacy")
    for name, expected in SITE_UPGRADE_SEMANTICS.items():
        if site_upgrade.get(name) != expected:
            fail(f"site upgrade {name} contract is incomplete")
    verified_path = site_upgrade.get("verified_path")
    if not isinstance(verified_path, dict) or set(verified_path) != {
        "installed_os",
        "installed_software",
        "software_prechecks",
        "target_os",
        "target_software",
    }:
        fail("site upgrade verified path is incomplete")
    if verified_path.get("software_prechecks") != "passed" or not all(
        isinstance(value, str) and value
        for key, value in verified_path.items()
        if key != "software_prechecks"
    ):
        fail("site upgrade verified path is not evidence-backed")


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
    if manifest["schema_version"] != 1 or manifest["contract_id"] != CONTRACT_ID:
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


def validate_contract(contract: dict, contract_id: str, contract_version: str) -> None:
    """Validate the exact MAC-bound SMSv2 runtime and authority contract."""
    aws = contract.get("providers", {}).get("aws", {})
    if (
        contract.get("contract_id") != contract_id
        or contract.get("version") != contract_version
        or contract.get("resource") != "securemesh_site_v2"
    ):
        fail("AWS contract identity is invalid")
    api = contract.get("api", {})
    if api.get("namespace") != "system" or set(api.get("operations", [])) != {
        "create",
        "read",
        "replace",
        "delete",
    }:
        fail("system-namespace CRUD declaration is incomplete")
    if (
        aws.get("node_list_path") != "aws.not_managed.node_list[]"
        or aws.get("interface_list_path")
        != "aws.not_managed.node_list[].interface_list[]"
    ):
        fail("AWS SMSv2 configuration paths are unsupported")
    if aws.get("interface_identity") != {
        "fields": ["node", "ethernet_interface.mac"],
        "guest_device": "rejected",
        "known_value_policy": "reject_null_incomplete_malformed_ambiguous_or_inconsistent",
        "mac": {
            "configuration_path": "spec.aws.not_managed.node_list[].interface_list[].ethernet_interface.mac",
            "input_field": "mac",
            "normalization": "ieee802_lowercase_colon",
            "nullable": False,
        },
        "node": {
            "configuration_path": "spec.aws.not_managed.node_list[].hostname",
            "input_field": "node",
            "normalization": "trim",
            "nullable": False,
        },
        "uniqueness_scope": "node",
        "unknown_value_policy": "defer",
    }:
        fail("AWS interface identity must be MAC-bound")
    if aws.get("roles") != [
        {"name": "slo", "network_option": "site_local_network"},
        {"name": "sli", "network_option": "site_local_inside_network"},
    ]:
        fail("AWS interface role declarations are incomplete")
    intake = aws.get("telemetry_intake", {})
    if (
        intake.get("schema_id") != TELEMETRY_SCHEMA_ID
        or intake.get("unavailable_facts") != []
        or set(intake.get("required_facts", [])) != REQUIRED_FACTS
        or set(intake.get("observed_facts", [])) != REQUIRED_FACTS
    ):
        fail("AWS telemetry declaration is incomplete")
    runtime = aws.get("runtime", {})
    if set(runtime) != set(RUNTIME_ENDPOINTS):
        fail("F5 runtime endpoint inventory is incomplete")
    for name, expected in RUNTIME_ENDPOINTS.items():
        actual = runtime.get(name, {})
        if any(actual.get(key) != value for key, value in expected.items()):
            fail(f"F5 runtime endpoint {name} is incomplete or legacy")
        mappings = actual.get("response_mappings")
        if not isinstance(mappings, dict) or not mappings:
            fail(f"F5 runtime endpoint {name} has no response mappings")
    if runtime["configuration"].get("correlation") != ["node", "normalized_mac"]:
        fail("configuration correlation must use node and normalized MAC")
    if (
        runtime["bgp_peers"].get("response_mappings", {}).get("state_changed_at")
        != "ver[].peer[].up_down_timestamp"
    ):
        fail("BGP state timestamp mapping is incomplete")
    if "observed_at" in runtime["bgp_peers"].get("response_mappings", {}):
        fail("BGP observation freshness is unsupported")
    if runtime["simplified_routes"].get("request_mappings") != {
        "node_scope": "all_nodes",
        "roles": ["slo", "sli"],
    }:
        fail("simplified route request mapping is incomplete")
    site_upgrade = aws.get("site_upgrade")
    has_site_upgrade = site_upgrade is not None
    if has_site_upgrade:
        validate_site_upgrade(site_upgrade, contract_version)
    expected_authorities = UPGRADE_AUTHORITIES if has_site_upgrade else BASE_AUTHORITIES
    if aws.get("authorities") != expected_authorities:
        fail("F5 and AWS authority declarations do not match the v3 contract")
    if aws.get("prohibited_legacy_apis") != ["aws_vpc_site", "aws_tgw_site"]:
        fail("legacy AWS site APIs must remain prohibited")
    availability = aws.get("availability")
    available_capabilities = (
        UPGRADE_CAPABILITIES if has_site_upgrade else BASE_CAPABILITIES
    )
    unavailable_capabilities = (
        UPGRADE_UNAVAILABLE_CAPABILITIES
        if has_site_upgrade
        else BASE_UNAVAILABLE_CAPABILITIES
    )
    if availability == "evidence_backed":
        if (
            aws.get("capabilities") != available_capabilities
            or aws.get("unavailable_capabilities") != []
            or intake.get("availability") != "available"
            or intake.get("complete") is not True
        ):
            fail("evidence-backed AWS capabilities and telemetry are incoherent")
    elif availability == "schema_only":
        if (
            aws.get("capabilities") != unavailable_capabilities
            or set(aws.get("unavailable_capabilities", []))
            != set(unavailable_capabilities)
            or intake.get("availability") != "unavailable"
            or intake.get("complete") is not False
        ):
            fail("schema-only AWS capabilities and telemetry must fail closed")
    else:
        fail("AWS contract availability is unsupported")


def validate_openapi(openapi: dict) -> None:
    """Require the SMSv2 node wire schema used by the fail-closed evidence."""
    public_ip = (
        openapi.get("components", {})
        .get("schemas", {})
        .get("viewssecuremesh_site_v2Node", {})
        .get("properties", {})
        .get("public_ip")
    )
    if not isinstance(public_ip, dict) or public_ip.get("type") != "string":
        fail("SMSv2 node public_ip schema is missing or malformed")
    if public_ip.get("nullable") is not True:
        fail("SMSv2 node public_ip must preserve the null wire value")


def validate_evidence(
    evidence: dict, contract_id: str, availability: str, site_upgrade: object
) -> None:
    """Validate the age and sanitization of the behavioral evidence."""
    try:
        observed_at = dt.datetime.fromisoformat(
            evidence["recorded_at"].replace("Z", "+00:00")
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
    if isinstance(site_upgrade, dict) and availability == "evidence_backed":
        upgrade_receipts = [
            item
            for item in receipts
            if item.get("operations")
            == [
                "site_status",
                "target_discovery",
                "precheck",
                "software_upgrade",
                "os_upgrade",
            ]
        ]
        if len(upgrade_receipts) != 1:
            fail("site upgrade evidence receipt is missing or ambiguous")
        receipt = upgrade_receipts[0]
        if (
            receipt.get("result") != "accepted"
            or receipt.get("upgrade_path") != site_upgrade.get("verified_path")
            or set(receipt.get("validated_facts", [])) != SITE_UPGRADE_EVIDENCE_FACTS
        ):
            fail("site upgrade evidence does not match the contract")
    if availability == "schema_only":
        if len(receipts) != 1:
            fail("schema-only evidence must contain one blocking receipt")
        receipt = receipts[0]
        if (
            receipt.get("operations") != ["replace"]
            or receipt.get("result") != "rejected"
            or receipt.get("blocking_conditions") != BLOCKING_CONDITIONS
        ):
            fail("schema-only evidence does not match the live API blockers")


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
    openapi = load_json(directory / "openapi.json", "OpenAPI specification")
    concurrency = load_json(
        directory / "concurrency_contracts.json", "concurrency inventory"
    )
    parity = load_json(
        directory / "smsv2_parity_manifest.json", "SMSv2 parity manifest"
    )
    contract_id = manifest["contract_id"]
    version = tag.removeprefix("v")
    validate_contract(contract, contract_id, manifest["contract_version"])
    validate_openapi(openapi)
    validate_evidence(
        evidence,
        contract_id,
        contract["providers"]["aws"]["availability"],
        contract["providers"]["aws"].get("site_upgrade"),
    )
    validate_concurrency(concurrency, version)
    validate_parity(parity, version)


if __name__ == "__main__":
    main()
