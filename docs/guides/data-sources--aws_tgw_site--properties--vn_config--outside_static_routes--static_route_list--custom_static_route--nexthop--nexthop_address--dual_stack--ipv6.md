---
page_title: "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6"
subcategory: ""
description: "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 3781, "body_sha256": "sha256:b55838c81968eeeccaf3252fba0a95aeb022b436a607fc159be0496620879135", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "child_ids": [], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "path": "docs/guides/data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv6"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes.md)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="section"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

## Next pages

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
