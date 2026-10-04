---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["ingress egress gw ar hub express route enabled gateway subnet"], "body_bytes": 3078, "body_sha256": "sha256:bc5db8abcdec6b660ed1026716a0ceb1ec297ce58301a87c019213794262bf34", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:auto", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw ar hub express route enabled gateway subnet auto"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:auto", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet", "auto"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw ar hub express route enabled gateway subnet subnet"], "anchor": "section", "description": "Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet", "subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw ar hub express route enabled gateway subnet subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "gateway_subnet", "subnet_param"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/)
- [ingress_egress_gw_ar.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for gateway subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

## Direct properties

- [auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/auto/): complete subsection reference.

- [subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet/): complete subsection reference.

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet_param/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/auto/)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet/)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet_param/)
- [ingress_egress_gw_ar.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
