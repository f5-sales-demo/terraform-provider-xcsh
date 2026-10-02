---
page_title: "ingress_egress_gw.hub"
subcategory: "Infrastructure"
description: "Hub VNet type."
xcsh_docs: {"aliases": ["ingress egress gw hub"], "body_bytes": 2356, "body_sha256": "sha256:9604b8267d6a1fb74f620499b4c03bada9eade0defbd7922b3c11bb2701a4698", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_disabled", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0300000102201333-3130020231023031-3010111000030120-3032232232303101-3312101332313322-1121132310023303-1022002212133221-1320232203211211", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "hub"], "schema_version": 1, "sections": [{"aliases": ["express route disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "express_route_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["express route enabled"], "anchor": "section", "description": "Express Route Configuration.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["spoke vnets"], "anchor": "section", "description": "Spoke VNet Peering.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Hub VNet type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- ingress_egress_gw.hub

<a id="section"></a>

Type: `"single"`. Computed.

Hub VNet type. Hub VNet type.

Upstream description:

Hub VNet type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-express_route_choice": "[\"express_route_disabled\",\"express_route_enabled\"]"
}
```

## Direct properties

- [express_route_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_disabled/): complete subsection reference.

- [express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/): complete subsection reference.

- [spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/): complete subsection reference.

## Next pages

- [ingress_egress_gw.hub.express_route_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_disabled/)
- [ingress_egress_gw.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/)
- [ingress_egress_gw.hub.spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
