---
page_title: "ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: "Infrastructure"
description: "L7 enhanced performance mode OPTIONS."
xcsh_docs: {"aliases": ["ingress gw ar performance enhancement mode perf mode l7 enhanced"], "body_bytes": 2539, "body_sha256": "sha256:ae99d5b3dd5132846ac9d0471b11287c48d5ff44321723f2c9b39430b072bf30", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3333221323302321-0100102122333003-2320122122020332-2300011331011020-1231200003103021-3212102021311311-2030233222233303-2002320123032133", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "sections": [{"aliases": ["jumbo disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["jumbo enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_enabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "L7 enhanced performance mode OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/)
- [ingress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced

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

- [jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/): complete subsection reference.

- [jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/): complete subsection reference.

## Next pages

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/)
- [ingress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
