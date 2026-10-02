---
page_title: "ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced"
subcategory: "Infrastructure"
description: "L3 enhanced performance mode OPTIONS."
xcsh_docs: {"aliases": ["ingress gw ar performance enhancement mode perf mode l3 enhanced"], "body_bytes": 2469, "body_sha256": "sha256:2302c151b33b19535ddb6372889be1b5fb373eb909edf1cf87550779b9d04993", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1002201022332303-3111331111022301-3010030102013212-0202222130210110-2021103312112100-2101121332331231-3331301220121203-0222032313001203", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "schema_version": 1, "sections": [{"aliases": ["jumbo"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced", "jumbo"], "syntax": "attribute", "type": "object"}, {"aliases": ["no jumbo"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced", "no_jumbo"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "L3 enhanced performance mode OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/)
- [ingress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced

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

- [jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/): complete subsection reference.

- [no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/): complete subsection reference.

## Next pages

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/)
- [ingress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
