---
page_title: "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address"
subcategory: ""
description: "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 3527, "body_sha256": "sha256:52c0d12cb30469ea81dbc1b31da57fd2135caeed76f70ed6d2c82419505ca19e", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vn_config.outside_static_routes](resources--aws_tgw_site--properties--vn_config--outside_static_routes.md)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
```

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

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dual_stack](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md): complete subsection reference.

- [ipv4](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md): complete subsection reference.

- [ipv6](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md): complete subsection reference.

## Next pages

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
