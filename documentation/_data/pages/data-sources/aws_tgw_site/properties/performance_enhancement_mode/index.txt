---
page_title: "performance_enhancement_mode"
subcategory: ""
description: "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default."
xcsh_docs: {"aliases": ["performance enhancement mode"], "body_bytes": 1236, "body_sha256": "sha256:7101cdeeeb2f11c05c6645455e043b02650d61989bcad8e25f654d002f6bed75", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "documentation/data-sources/aws_tgw_site/properties/performance_enhancement_mode/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode"], "schema_version": 1, "sections": [{"aliases": ["performance enhancement mode perf mode l3 enhanced"], "anchor": "section", "description": "L3 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l3_enhanced"], "syntax": "attribute", "type": "object"}, {"aliases": ["performance enhancement mode perf mode l7 enhanced"], "anchor": "section", "description": "L7 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
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

- [perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/): complete subsection reference.

- [perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/): complete subsection reference.
