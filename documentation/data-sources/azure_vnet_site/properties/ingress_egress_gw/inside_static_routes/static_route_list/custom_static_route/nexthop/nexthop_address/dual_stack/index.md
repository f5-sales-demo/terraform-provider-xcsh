---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack"
subcategory: "Infrastructure"
description: "DualStackAddressType represents both IPv4 and IPv6 together."
xcsh_docs: {"aliases": ["ingress egress gw inside static routes static route list custom static route nexthop nexthop address dual stack"], "body_bytes": 3869, "body_sha256": "sha256:3d9daa46a583c960558b832971dfa3873353e34e04c1408cdacab6c8f5de5a75", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3111113212010333-0202101130232302-2323321321131200-1131123121202222-1223323112111122-3020321002011110-2223131013132103-2210132113323323", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw inside static routes static route list custom static route nexthop nexthop address dual stack ipv4"], "anchor": "section", "description": "IPv4 Address in dot-decimal notation.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw inside static routes static route list custom static route nexthop nexthop address dual stack ipv6"], "anchor": "section", "description": "IPv6 Address specified as hexadecimal numbers separated by ':'", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "DualStackAddressType represents both IPv4 and IPv6 together.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/)
- [ingress_egress_gw.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
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

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
