---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route"
subcategory: "Infrastructure"
description: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3837, "body_sha256": "sha256:8df868248fe52e61f0dfce15404761feb14b5205ffecf83517b8239ce3490456", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:labels", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list", "path": "docs/guides/data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [voltstack_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster.md)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes.md)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="section"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

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

<a id="schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--attrs"></a>

### attrs property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--labels.md): complete subsection reference.

- [nexthop](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md): complete subsection reference.

- [subnets](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--labels.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
