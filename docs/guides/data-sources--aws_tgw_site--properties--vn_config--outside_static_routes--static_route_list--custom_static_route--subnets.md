---
page_title: "vn_config.outside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: ""
description: "vn_config.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2850, "body_sha256": "sha256:d1d8ee2f1df866e3413fd8643b5606ce019cd4fb40b2feac42df97e12dc71eb0", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes.md)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

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

- [ipv4](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md): complete subsection reference.

## Next pages

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
