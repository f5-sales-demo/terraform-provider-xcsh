---
page_title: "vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address"
subcategory: ""
description: "IP Address used to specify an IPv4 or IPv6 address."
xcsh_docs: {"aliases": ["vn config inside static routes static route list custom static route nexthop nexthop address"], "body_bytes": 3850, "body_sha256": "sha256:4d8a574bea400918676c40d0a2c72c4f648cee9b9a25d5eac5aa59719853130c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "schema_version": 1, "sections": [{"aliases": ["dual stack"], "anchor": "section", "description": "DualStackAddressType represents both IPv4 and IPv6 together.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IP Address used to specify an IPv4 or IPv6 address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- [vn_config.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/)
- [vn_config.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="section"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

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

## Direct properties

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/): complete subsection reference.

## Next pages

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
