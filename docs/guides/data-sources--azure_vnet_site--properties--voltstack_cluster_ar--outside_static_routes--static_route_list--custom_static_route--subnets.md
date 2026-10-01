---
page_title: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3090, "body_sha256": "sha256:de417f5c147a1f6226682f0945e5b111e76d3249b05a5fd2d4828470af6a5df4", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route.md)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="section"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

## Direct properties

- [ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--properties--voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
