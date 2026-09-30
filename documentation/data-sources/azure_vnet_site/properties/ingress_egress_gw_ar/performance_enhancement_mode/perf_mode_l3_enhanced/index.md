---
page_title: "ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2468, "body_sha256": "sha256:0276858cfb334add80507c7adad4bf0aa1cbdcf8baf70e629e74c9230865cc58", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
