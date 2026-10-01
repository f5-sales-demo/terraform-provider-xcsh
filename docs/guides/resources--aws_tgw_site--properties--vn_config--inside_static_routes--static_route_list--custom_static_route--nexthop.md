---
page_title: "vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop"
subcategory: ""
description: "vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 4059, "body_sha256": "sha256:8250ebdfdbddcce0f93688c8b71ec28c8b0b7c8d07f9fc25982384827ac40472", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:interface", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vn_config.inside_static_routes](resources--aws_tgw_site--properties--vn_config--inside_static_routes.md)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list.md)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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

Terraform syntax:

```terraform
nexthop {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interface](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md): complete subsection reference.

- [nexthop_address](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md): complete subsection reference.

<a id="schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--type"></a>

### type property

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface.md)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--inside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
