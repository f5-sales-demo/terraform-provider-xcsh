---
page_title: "voltstack_cluster.outside_static_routes"
subcategory: "Infrastructure"
description: "voltstack_cluster.outside_static_routes for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1229, "body_sha256": "sha256:496ea395f0875dd33caab9fbd040928906dbbe56070e2050a04aa12690b5de7c", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.outside_static_routes for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
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

- [static_route_list](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
