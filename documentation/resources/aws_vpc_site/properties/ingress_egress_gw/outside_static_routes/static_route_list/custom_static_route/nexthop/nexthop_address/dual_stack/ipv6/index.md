---
page_title: "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6"
subcategory: "Infrastructure"
description: "IPv6 Address specified as hexadecimal numbers separated by ':'"
xcsh_docs: {"aliases": ["ingress egress gw outside static routes static route list custom static route nexthop nexthop address dual stack ipv6"], "body_bytes": 4718, "body_sha256": "sha256:cd7d904188f921db1611efebe7977191a2738af65b2fc50dfbd6c840d8f51dc8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "path": "documentation/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0011011030211222-1221031013031123-0210121311013332-3100011023223001-0200223110011000-0103302102332200-0100002212123303-0012210031221130", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv6"], "schema_version": 1, "sections": [{"aliases": ["addr"], "anchor": "schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr", "description": "IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by ':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv6", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IPv6 Address specified as hexadecimal numbers separated by ':'", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/)
- [ingress_egress_gw.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
