---
page_title: "vn_config.outside_static_routes"
subcategory: ""
description: "List of static routes."
xcsh_docs: {"aliases": ["vn config outside static routes"], "body_bytes": 1079, "body_sha256": "sha256:5585094194da1b2cb62cf9be09103849a573dac668276e989b585ed5706f2648", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "outside_static_routes"], "schema_version": 1, "sections": [{"aliases": ["vn config outside static routes static route list"], "anchor": "section", "description": "List of Static routes.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["vn_config", "outside_static_routes", "static_route_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- vn_config.outside_static_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Additional upstream details:

List of static routes.

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

- [static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/): complete subsection reference.
