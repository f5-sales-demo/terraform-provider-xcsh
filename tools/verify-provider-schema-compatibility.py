# ruff: noqa: INP001
"""Verify that a candidate provider accepts the v9.5.2 configuration contract."""

from __future__ import annotations

import argparse
import gzip
import json
import sys
from pathlib import Path
from typing import Any

PROVIDER = "registry.terraform.io/f5-sales-demo/xcsh"
SCHEMA_GROUPS = ("resource_schemas", "data_source_schemas", "action_schemas")


def load_provider_schema(path: Path) -> dict[str, Any]:
    """Load the xcsh portion of a Terraform provider-schema document."""
    if path.suffix == ".gz":
        with gzip.open(path, mode="rt", encoding="utf-8") as schema_file:
            document = json.load(schema_file)
    else:
        document = json.loads(path.read_text(encoding="utf-8"))
    try:
        return document["provider_schemas"][PROVIDER]
    except (KeyError, TypeError) as error:
        message = f"{path}: missing provider schema for {PROVIDER}"
        raise ValueError(message) from error


def is_configurable(attribute: dict[str, Any]) -> bool:
    """Return whether Terraform configuration may set an attribute."""
    return bool(attribute.get("required") or attribute.get("optional"))


def compare_attributes(
    baseline: dict[str, Any], candidate: dict[str, Any], path: str, errors: list[str]
) -> None:
    """Compare configurable attributes, including nested attribute objects."""
    for name, baseline_attribute in sorted(baseline.items()):
        if not is_configurable(baseline_attribute):
            continue
        attribute_path = f"{path}.{name}"
        candidate_attribute = candidate.get(name)
        if candidate_attribute is None:
            errors.append(f"configured attribute removed: {attribute_path}")
            continue
        if not is_configurable(candidate_attribute):
            errors.append(f"configured attribute is now computed-only: {attribute_path}")
            continue
        if not baseline_attribute.get("required") and candidate_attribute.get("required"):
            errors.append(f"optional attribute is now required: {attribute_path}")
        if baseline_attribute.get("type") != candidate_attribute.get("type"):
            errors.append(f"configured attribute type changed: {attribute_path}")
        baseline_nested = baseline_attribute.get("nested_type")
        candidate_nested = candidate_attribute.get("nested_type")
        if baseline_nested is not None:
            if candidate_nested is None:
                errors.append(f"nested configured attribute changed shape: {attribute_path}")
                continue
            if baseline_nested.get("nesting_mode") != candidate_nested.get("nesting_mode"):
                errors.append(f"nested configured attribute mode changed: {attribute_path}")
            compare_attributes(
                baseline_nested.get("attributes", {}),
                candidate_nested.get("attributes", {}),
                attribute_path,
                errors,
            )


def compare_block(
    baseline: dict[str, Any], candidate: dict[str, Any], path: str, errors: list[str]
) -> None:
    """Compare the configurable portion of two Terraform schema blocks."""
    compare_attributes(
        baseline.get("attributes", {}), candidate.get("attributes", {}), path, errors
    )
    for name, baseline_nested in sorted(baseline.get("block_types", {}).items()):
        block_path = f"{path}.{name}"
        candidate_nested = candidate.get("block_types", {}).get(name)
        if candidate_nested is None:
            errors.append(f"configured nested block removed: {block_path}")
            continue
        if baseline_nested.get("nesting_mode") != candidate_nested.get("nesting_mode"):
            errors.append(f"configured nested block mode changed: {block_path}")
        baseline_min = baseline_nested.get("min_items", 0)
        candidate_min = candidate_nested.get("min_items", 0)
        if candidate_min > baseline_min:
            errors.append(f"configured nested block minimum increased: {block_path}")
        baseline_max = baseline_nested.get("max_items", 0)
        candidate_max = candidate_nested.get("max_items", 0)
        if candidate_max and (not baseline_max or candidate_max < baseline_max):
            errors.append(f"configured nested block maximum decreased: {block_path}")
        compare_block(
            baseline_nested.get("block", {}),
            candidate_nested.get("block", {}),
            block_path,
            errors,
        )


def main() -> int:
    """Run the compatibility comparison."""
    parser = argparse.ArgumentParser()
    parser.add_argument("--baseline", type=Path, required=True)
    parser.add_argument("--candidate", type=Path, required=True)
    parser.add_argument("--reviewed-replacements", type=Path)
    args = parser.parse_args()

    baseline = load_provider_schema(args.baseline)
    candidate = load_provider_schema(args.candidate)
    errors: list[str] = []

    compare_block(
        baseline.get("provider", {}).get("block", {}),
        candidate.get("provider", {}).get("block", {}),
        "provider.xcsh",
        errors,
    )
    for group in SCHEMA_GROUPS:
        baseline_schemas = baseline.get(group, {})
        candidate_schemas = candidate.get(group, {})
        for type_name, baseline_schema in sorted(baseline_schemas.items()):
            candidate_schema = candidate_schemas.get(type_name)
            if candidate_schema is None:
                errors.append(f"registered type removed: {type_name}")
                continue
            compare_block(
                baseline_schema.get("block", {}),
                candidate_schema.get("block", {}),
                type_name,
                errors,
            )

    reviewed: dict[str, Any] = {}
    if args.reviewed_replacements:
        reviewed_document = json.loads(
            args.reviewed_replacements.read_text(encoding="utf-8")
        )
        reviewed = {
            item["diagnostic"]: item for item in reviewed_document["replacements"]
        }
        stale = sorted(set(reviewed) - set(errors))
        if stale:
            errors.extend(f"reviewed replacement no longer matches: {item}" for item in stale)
        errors = [error for error in errors if error not in reviewed]

    counts = ", ".join(
        f"{group.removesuffix('_schemas')}={len(baseline.get(group, {}))}"
        for group in SCHEMA_GROUPS
    )
    if errors:
        lines = [f"v9.5.2 provider compatibility failed ({counts}):"]
        lines.extend(f"- {error}" for error in errors)
        sys.stdout.write("\n".join(lines) + "\n")
        return 1
    sys.stdout.write(
        f"v9.5.2 provider compatibility passed ({counts}; "
        f"reviewed replacements={len(reviewed)})\n"
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
