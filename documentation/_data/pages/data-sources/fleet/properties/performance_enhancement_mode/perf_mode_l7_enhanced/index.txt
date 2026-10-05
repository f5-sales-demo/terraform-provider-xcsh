---
page_title: "performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: ""
description: "L7 enhanced performance mode OPTIONS."
xcsh_docs: {"aliases": ["performance enhancement mode perf mode l7 enhanced"], "body_bytes": 2131, "body_sha256": "sha256:3ba28aec8c701c67b305c6dcc3239a3e4fcc08c8205a81027622dc8cc8a224cf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode", "path": "documentation/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2100321330111130-0122101110332312-1203131320322323-1321032201332323-0312310300023333-3200302022110021-0130113231022310-0103223230031112", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "sections": [{"aliases": ["performance enhancement mode perf mode l7 enhanced jumbo disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["performance enhancement mode perf mode l7 enhanced jumbo enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_enabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "L7 enhanced performance mode OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

## Direct properties

- [jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/): complete subsection reference.

- [jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/performance_enhancement_mode/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
