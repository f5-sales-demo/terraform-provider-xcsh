---
page_title: "performance_enhancement_mode"
subcategory: ""
description: "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default."
xcsh_docs: {"aliases": ["performance enhancement mode"], "body_bytes": 1201, "body_sha256": "sha256:0166a3635c58bc27a913610ffb61253367287fc222feb2d4508da1fbd9d8e65b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/performance_enhancement_mode/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3133211021322020-2332213213103301-1020331030301223-1210200321112330-0230022112123011-2002100233113211-0213202312111002-0031301033032211", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode"], "schema_version": 1, "sections": [{"aliases": ["performance enhancement mode perf mode l3 enhanced"], "anchor": "section", "description": "L3 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l3_enhanced"], "syntax": "attribute", "type": "object"}, {"aliases": ["performance enhancement mode perf mode l7 enhanced"], "anchor": "section", "description": "L7 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["fleetCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- performance_enhancement_mode

<a id="section"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

## Direct properties

- [perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l3_enhanced/): complete subsection reference.

- [perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/): complete subsection reference.
