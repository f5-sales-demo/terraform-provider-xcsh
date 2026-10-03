---
page_title: "ingress_egress_gw_ar.hub.spoke_vnets.labels"
subcategory: "Infrastructure"
description: "Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall policy These labels used must be from known key and label defined in shared namespace."
xcsh_docs: {"aliases": ["ingress egress gw ar hub spoke vnets labels"], "body_bytes": 1926, "body_sha256": "sha256:7681e63ed7c0bed1c44ff5e5cfff5cec9e8746f631048f5653afe0b60e6cd0e5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets:labels", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/labels/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2011112002232222-1032012200102303-0331023233032320-3332223130333003-0003200131221000-3030211313322022-0300001231230211-3221201102223021", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "spoke_vnets", "labels"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/labels/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall policy These labels used must be from known key and label defined in shared namespace.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.spoke_vnets.labels

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/)
- [ingress_egress_gw_ar.hub.spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/)
- ingress_egress_gw_ar.hub.spoke_vnets.labels

<a id="section"></a>

Type: `"single"`. Computed.

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Upstream description:

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw_ar.hub.spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
