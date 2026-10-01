---
page_title: "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3173, "body_sha256": "sha256:8c4f7f74c333d8cfe77d249cfceaed0f6eaf27470a09e156b31e9edae67763aa", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.outside_static_routes](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
