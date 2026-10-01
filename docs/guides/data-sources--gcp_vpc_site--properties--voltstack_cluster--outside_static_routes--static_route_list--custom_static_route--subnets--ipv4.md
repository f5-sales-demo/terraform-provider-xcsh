---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4"
subcategory: "Infrastructure"
description: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3550, "body_sha256": "sha256:98b086bd13f250ae0bf64f7ae2b08132c68febd26e77c2de8c1f24a075df3bd1", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets", "path": "docs/guides/data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes.md)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

<a id="schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen"></a>

### plen property

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix"></a>

### prefix property

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
