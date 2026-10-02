---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address"
subcategory: "Infrastructure"
description: "IP Address used to specify an IPv4 or IPv6 address."
xcsh_docs: {"aliases": ["voltstack cluster outside static routes static route list custom static route nexthop nexthop address"], "body_bytes": 4428, "body_sha256": "sha256:d03a6c54a1a7d24a998eb31d99869690752991e9942e8960fb152a93e3f1d37a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop", "path": "documentation/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv4", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:dual_stack,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address:ConflictingObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "schema_version": 1, "sections": [{"aliases": ["dual stack"], "anchor": "section", "description": "DualStackAddressType represents both IPv4 and IPv6 together.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack"], "syntax": "block", "type": "object"}, {"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv4"], "syntax": "block", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv6"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IP Address used to specify an IPv4 or IPv6 address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/)
- [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/)
- [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

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

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
