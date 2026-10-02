---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address"
subcategory: "Infrastructure"
description: "IP Address used to specify an IPv4 or IPv6 address."
xcsh_docs: {"aliases": ["ingress egress gw inside static routes static route list custom static route nexthop nexthop address"], "body_bytes": 4034, "body_sha256": "sha256:79b8b52bac66d5689e74e2f7fe22ddf25a893c16b552904ee4ee8a6d26873963", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop", "path": "documentation/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "schema_version": 1, "sections": [{"aliases": ["dual stack"], "anchor": "section", "description": "DualStackAddressType represents both IPv4 and IPv6 together.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IP Address used to specify an IPv4 or IPv6 address.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/)
- [ingress_egress_gw.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

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

- [dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/): complete subsection reference.

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
