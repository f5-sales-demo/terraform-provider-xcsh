---
page_title: "vn_config.inside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: ""
description: "List of route prefixes."
xcsh_docs: {"aliases": ["vn config inside static routes static route list custom static route subnets"], "body_bytes": 2486, "body_sha256": "sha256:757217ede027039a1eed1779f116b378950208e91e2d473c0964a817d054fdda", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3213201111321223-2310320130001123-3031233120100133-3302010200132020-0303203000003031-0202033321130321-1100223323112001-3013313000022331", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "sections": [{"aliases": ["vn config inside static routes static route list custom static route subnets ipv4"], "anchor": "section", "description": "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv4"], "syntax": "attribute", "type": "object"}, {"aliases": ["vn config inside static routes static route list custom static route subnets ipv6"], "anchor": "section", "description": "IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be <= 128.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list:custom_static_route:subnets:ipv6", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vn_config", "inside_static_routes", "static_route_list", "custom_static_route", "subnets", "ipv6"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of route prefixes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- [vn_config.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/)
- [vn_config.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="section"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/): complete subsection reference.
