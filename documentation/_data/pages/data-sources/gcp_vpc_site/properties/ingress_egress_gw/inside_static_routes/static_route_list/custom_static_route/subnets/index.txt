---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: "Infrastructure"
description: "List of route prefixes."
xcsh_docs: {"aliases": ["ingress egress gw inside static routes static route list custom static route subnets"], "body_bytes": 3563, "body_sha256": "sha256:a9ab279a0c4ec0c0780bbb39904a8adcc9926ffe4aec96f71e2489d3e1902599", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route", "path": "documentation/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "sections": [{"aliases": ["ipv4"], "anchor": "section", "description": "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6"], "anchor": "section", "description": "IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be <= 128.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:subnets:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of route prefixes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/)
- [ingress_egress_gw.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="section"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

## Direct properties

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
