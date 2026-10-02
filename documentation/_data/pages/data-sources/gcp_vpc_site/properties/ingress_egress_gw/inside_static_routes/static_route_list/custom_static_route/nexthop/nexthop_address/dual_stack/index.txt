---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack"
subcategory: "Infrastructure"
description: "DualStackAddressType represents both IPv4 and IPv6 together."
xcsh_docs: {"aliases": ["ingress egress gw inside static routes static route list custom static route nexthop nexthop address dual stack"], "body_bytes": 3821, "body_sha256": "sha256:8f85c4789cc0c539ee6c0026a4d8581a6ad259bf5c72b9be6eef7f435227c8a6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "path": "documentation/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "DualStackAddressType represents both IPv4 and IPv6 together.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/)
- [ingress_egress_gw.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="section"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
