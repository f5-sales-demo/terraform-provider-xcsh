---
page_title: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3458, "body_sha256": "sha256:f2d3412cf89845a4ef29d3e9cf00aed008a87f176ec58fe08d85257e608ec2d2", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop", "path": "docs/guides/data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="section"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

## Direct properties

- [dual_stack](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md): complete subsection reference.

- [ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
