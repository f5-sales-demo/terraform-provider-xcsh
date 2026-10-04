---
page_title: "voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4"
subcategory: "Infrastructure"
description: "IPv4 Address in dot-decimal notation."
xcsh_docs: {"aliases": ["voltstack cluster ar outside static routes static route list custom static route nexthop nexthop address dual stack ipv4"], "body_bytes": 4304, "body_sha256": "sha256:b8265ff24c0f2e0ee37d0217abe5adba2b65bc89e6e27a29061f0954462e2d4d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1202130211111322-1213033010322200-3200321000003121-3200113022010230-3122232123200100-1213113302122233-3022233011323012-2313322001311130", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["voltstack cluster ar outside static routes static route list custom static route nexthop nexthop address dual stack ipv4 addr"], "anchor": "schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr", "description": "IPv4 Address in string form with dot-decimal notation.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack:ipv4", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "outside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address", "dual_stack", "ipv4", "addr"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "IPv4 Address in dot-decimal notation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [voltstack_cluster_ar.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
