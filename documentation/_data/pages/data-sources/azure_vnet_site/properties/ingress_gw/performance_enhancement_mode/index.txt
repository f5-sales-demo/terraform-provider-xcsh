---
page_title: "ingress_gw.performance_enhancement_mode"
subcategory: "Infrastructure"
description: "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default."
xcsh_docs: {"aliases": ["ingress gw performance enhancement mode"], "body_bytes": 2264, "body_sha256": "sha256:74f9fe4715be88ff7d73900094ecb2af13225edbc5b8e242b73a6ec9b0341059", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1030221111321100-2132331010000112-2111302102120320-2201020232000230-3121323222333020-2311221211003020-3021113113033122-2311102233222210", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "performance_enhancement_mode"], "schema_version": 1, "sections": [{"aliases": ["ingress gw performance enhancement mode perf mode l3 enhanced"], "anchor": "section", "description": "L3 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress gw performance enhancement mode perf mode l7 enhanced"], "anchor": "section", "description": "L7 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw", "performance_enhancement_mode", "perf_mode_l7_enhanced"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.performance_enhancement_mode

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/)
- ingress_gw.performance_enhancement_mode

<a id="section"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. Server applies
default when omitted.

Upstream description:

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

- [perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/): complete subsection reference.

- [perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/): complete subsection reference.

## Next pages

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
