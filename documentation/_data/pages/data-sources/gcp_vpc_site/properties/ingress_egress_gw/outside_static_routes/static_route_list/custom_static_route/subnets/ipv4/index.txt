---
page_title: "ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4"
subcategory: "Infrastructure"
description: "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32."
xcsh_docs: {"aliases": ["ingress egress gw outside static routes static route list custom static route subnets ipv4"], "body_bytes": 3997, "body_sha256": "sha256:ed5e87ce9acaef9e5d7913354de2082a6ba4df5a403214c35080d8d13327f468", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:subnets", "path": "documentation/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1000331232020313-3031213301022210-2122200300022200-3231130030032110-0133012331300013-0123202202322022-2032233223101333-1203223100110012", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4"], "schema_version": 1, "sections": [{"aliases": ["plen"], "anchor": "schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen", "description": "Prefix-length of the IPv4 subnet. Must be <= 32.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4", "plen"], "syntax": "attribute", "type": "number"}, {"aliases": ["prefix"], "anchor": "schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix", "description": "Prefix part of the IPv4 subnet in string form with dot-decimal notation.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "outside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4", "prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/)
- [ingress_egress_gw.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen"></a>

### plen property

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix"></a>

### prefix property

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
