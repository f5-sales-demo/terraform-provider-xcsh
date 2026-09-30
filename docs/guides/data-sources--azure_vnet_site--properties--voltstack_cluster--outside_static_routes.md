---
page_title: "voltstack_cluster.outside_static_routes"
subcategory: "Infrastructure"
description: "voltstack_cluster.outside_static_routes for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1157, "body_sha256": "sha256:ac5cfb37ddc6bb4915ca5aaa8a3be5427ee8e8c69f43d03512c62ecc8fdf7a50", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.outside_static_routes for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.outside_static_routes

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [voltstack_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster.md)
- voltstack_cluster.outside_static_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

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

- [static_route_list](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md)
- [voltstack_cluster](data-sources--azure_vnet_site--properties--voltstack_cluster.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
