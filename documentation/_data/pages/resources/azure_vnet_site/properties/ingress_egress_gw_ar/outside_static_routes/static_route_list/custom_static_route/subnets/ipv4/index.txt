---
page_title: "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4"
subcategory: "Infrastructure"
description: "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32."
xcsh_docs: {"aliases": ["ingress egress gw ar outside static routes static route list custom static route subnets ipv4"], "body_bytes": 4462, "body_sha256": "sha256:2fef2ca03bee9343d6a09c98b3a64dad146f63ea42bfd4d3f43b68f5c1b29848", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0220032111330030-3131033312322303-1310022011201331-1313233102320022-0102231210113203-0120311010000303-2230030320113110-0330302123101233", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw ar outside static routes static route list custom static route subnets ipv4 plen"], "anchor": "schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen", "description": "Prefix-length of the IPv4 subnet. Must be <= 32.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw_ar", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4", "plen"], "syntax": "attribute", "type": "number"}, {"aliases": ["ingress egress gw ar outside static routes static route list custom static route subnets ipv4 prefix"], "anchor": "schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix", "description": "Prefix part of the IPv4 subnet in string form with dot-decimal notation.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw_ar", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4", "prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

<a id="schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen"></a>

### plen property

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix"></a>

### prefix property

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
