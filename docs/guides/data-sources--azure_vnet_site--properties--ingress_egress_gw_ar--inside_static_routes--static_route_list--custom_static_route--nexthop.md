---
page_title: "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3912, "body_sha256": "sha256:22f061813f3dc02a10eb0591fb8f25619e8c34d0252f0b8ab9880630043e38cf", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:interface", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route.md)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [interface](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md): complete subsection reference.

<a id="schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--type"></a>

### type property

Type: `"string"`. Computed.

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

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
