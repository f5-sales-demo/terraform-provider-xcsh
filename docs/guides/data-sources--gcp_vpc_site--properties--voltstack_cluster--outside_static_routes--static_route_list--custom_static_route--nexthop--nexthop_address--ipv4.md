---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4"
subcategory: "Infrastructure"
description: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3086, "body_sha256": "sha256:c4d17cbf1970e8050b49ed3658192698573d89707b4da5d8a86de22c50b69c4d", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "path": "docs/guides/data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes.md)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

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

<a id="schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr"></a>

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

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
