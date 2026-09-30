---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: "Infrastructure"
description: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2895, "body_sha256": "sha256:4279bb6f5115c139c1a79850028a7370fc53d24c126b736a53a1277e51981dfe", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [voltstack_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster.md)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes.md)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

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

- [ipv4](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md): complete subsection reference.

- [ipv6](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
