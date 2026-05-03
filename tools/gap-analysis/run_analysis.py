#!/usr/bin/env python3
"""Run the cross-repository gap analysis and generate the report."""

import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))

from extension_map import build_extension_map
from spec_coverage import audit_all_resources, audit_operation_extensions
from generate_report import create_gap_items, generate_markdown_report


def main():
    provider_root = Path(__file__).parent.parent.parent
    specs_root = Path("/workspace/api-specs-enriched")
    specs_dir = specs_root / "docs" / "specifications" / "api"
    output_path = provider_root / "docs" / "gap-analysis" / "cross-repo-gap-report.md"

    print("Phase 1: Building extension consumption map...")
    ext_map = build_extension_map(provider_root, specs_root)
    print(f"  Classified {len(ext_map)} extensions")

    print("Phase 1: Auditing spec coverage per resource...")
    coverage_data = audit_all_resources(specs_dir)
    print(f"  Audited {len(coverage_data)} resources")

    print("Phase 1: Auditing operation-level extensions...")
    op_data_list = []
    op_data = {
        "total_operations": 0,
        "ops_with_operation_metadata": 0,
        "ops_with_danger_level": 0,
        "ops_with_confirmation_required": 0,
        "ops_with_side_effects": 0,
        "ops_with_required_fields": 0,
    }
    domain_count = 0
    for spec_file in sorted(specs_dir.glob("*.json")):
        if spec_file.name == "index.json":
            continue
        ops = audit_operation_extensions(spec_file)
        for key in op_data:
            op_data[key] += ops.get(key, 0)
        ops["domain_file"] = spec_file.name
        op_data_list.append(ops)
        domain_count += 1
    print(f"  Audited {domain_count} domain specs")

    print("Phase 2: Creating gap items and scoring...")
    gap_items = create_gap_items(ext_map, coverage_data)
    print(f"  Created {len(gap_items)} gap items")

    print("Phase 2: Generating report...")
    report = generate_markdown_report(ext_map, coverage_data, gap_items, op_data)

    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text(report)
    print(f"Report written to {output_path}")

    print("\nTop 5 gaps by priority:")
    for i, item in enumerate(gap_items[:5], 1):
        print(f"  {i}. [{item['priority_score']:.1f}] {item['title']} ({item['repo']})")


if __name__ == "__main__":
    main()
