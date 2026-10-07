---
page_title: "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop"
subcategory: ""
description: "Identifies the next-hop for a route."
xcsh_docs: {"aliases": ["vn config outside static routes static route list custom static route nexthop"], "body_bytes": 3948, "body_sha256": "sha256:1e8333afedc0f22e34b4f70f24d25ed866e6a67ad84714620b2458a1b2bf8401", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:interface", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route", "path": "documentation/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop"], "schema_version": 1, "sections": [{"aliases": ["vn config outside static routes static route list custom static route nexthop interface"], "anchor": "section", "description": "Nexthop is network interface when type is \"Network-Interface\"", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:interface", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["vn config outside static routes static route list custom static route nexthop nexthop address"], "anchor": "section", "description": "IP Address used to specify an IPv4 or IPv6 address.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "type": "conflicts"}], "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "syntax": "block", "type": "object"}, {"aliases": ["vn config outside static routes static route list custom static route nexthop type"], "anchor": "schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--type", "description": "Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes there is only one local interface on the virtual network. Use the specified address as nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:nexthop", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["NEXT_HOP_DEFAULT_GATEWAY", "NEXT_HOP_NETWORK_INTERFACE", "NEXT_HOP_USE_CONFIGURED"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Identifies the next-hop for a route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [vn_config.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/)
- [vn_config.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

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

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/): complete subsection reference.

- [nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/): complete subsection reference.

<a id="schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--type"></a>

### type property

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Additional upstream details:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Use the specified address as
nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used
in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NEXT_HOP_DEFAULT_GATEWAY","NEXT_HOP_NETWORK_INTERFACE","NEXT_HOP_USE_CONFIGURED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
