---
page_title: "voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: "Infrastructure"
description: "List of route prefixes."
xcsh_docs: {"aliases": ["voltstack cluster outside static routes static route list custom static route subnets"], "body_bytes": 3621, "body_sha256": "sha256:28641ab118dfd5d1578e360c2a0e9c72fdbb318d0da5fdbd255d5184eed20190", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "sections": [{"aliases": ["voltstack cluster outside static routes static route list custom static route subnets ipv4"], "anchor": "section", "description": "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["voltstack cluster outside static routes static route list custom static route subnets ipv6"], "anchor": "section", "description": "IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be <= 128.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of route prefixes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/)
- [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

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

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/): complete subsection reference.

## Next pages

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
