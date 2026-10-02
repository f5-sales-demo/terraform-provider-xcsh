---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4"
subcategory: "Infrastructure"
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["voltstack cluster outside static routes static route list custom static route nexthop nexthop address dual stack ipv4"], "body_bytes": 3959, "body_sha256": "sha256:d823579b1e68f7ac06bab5314d31c5497b80ec165723807605bb6a2bb1d0b194", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "path": "documentation/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3101310331031013-3331323310020121-1132200330000333-0101023132001000-3131210333132101-3002000102202303-2311223112100111-3320202320210023", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["addr"], "anchor": "schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv4", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/)
- [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/)
- [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr"></a>

### addr property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
