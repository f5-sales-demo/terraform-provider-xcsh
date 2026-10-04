---
page_title: "ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced"
subcategory: "Infrastructure"
description: "L3 enhanced performance mode OPTIONS."
xcsh_docs: {"aliases": ["ingress egress gw ar performance enhancement mode perf mode l3 enhanced"], "body_bytes": 2567, "body_sha256": "sha256:aab1f96a58e7a97c08ece73bc1e4af4d2b7d10fa2c3547cf72dfcdce9a5d1ac1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3030003201331101-2310200132330321-1020301102103210-1002332112231230-1013030230321033-3302022310302320-3213323131021030-3120002331113101", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw ar performance enhancement mode perf mode l3 enhanced jumbo"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced", "jumbo"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw ar performance enhancement mode perf mode l3 enhanced no jumbo"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced", "no_jumbo"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "L3 enhanced performance mode OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

## Direct properties

- [jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/): complete subsection reference.

- [no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/)
- [ingress_egress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
