---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack"
subcategory: "Infrastructure"
description: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3278, "body_sha256": "sha256:1973ad85ce348aabc4bcc66b82f4d806023ff962985a06975226c3ab594fee82", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes.md)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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
dual_stack {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md): complete subsection reference.

- [ipv6](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
