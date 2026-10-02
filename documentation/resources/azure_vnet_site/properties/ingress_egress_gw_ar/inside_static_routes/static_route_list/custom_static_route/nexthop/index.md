---
page_title: "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop"
subcategory: "Infrastructure"
description: "Identifies the next-hop for a route."
xcsh_docs: {"aliases": ["ingress egress gw ar inside static routes static route list custom static route nexthop"], "body_bytes": 4895, "body_sha256": "sha256:26c1013f7d688b6d97e8a294f6b70b27fd9a01dbee35d935f24a7733b6a3654b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:interface", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop"], "schema_version": 1, "sections": [{"aliases": ["interface"], "anchor": "section", "description": "Nexthop is network interface when type is \"Network-Interface\"", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:interface", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ingress_egress_gw_ar", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["nexthop address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "type": "conflicts"}], "schema_path": ["ingress_egress_gw_ar", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "syntax": "block", "type": "object"}, {"aliases": ["type"], "anchor": "schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--type", "description": "Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes there is only one local interface on the virtual network. Use the specified address as nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw_ar", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Identifies the next-hop for a route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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

Terraform syntax:

```terraform
nexthop {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/): complete subsection reference.

- [nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/): complete subsection reference.

<a id="schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--type"></a>

### type property

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
