---
page_title: "vn_config.inside_static_routes"
subcategory: ""
description: "List of static routes."
xcsh_docs: {"aliases": ["vn config inside static routes"], "body_bytes": 1506, "body_sha256": "sha256:a56bc5651751d436995f005b5ee5ca329e65cb33df21642f1f0e740cb8f2c6a2", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "path": "documentation/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "inside_static_routes"], "schema_version": 1, "sections": [{"aliases": ["vn config inside static routes static route list"], "anchor": "section", "description": "List of Static routes.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:inside_static_routes:static_route_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["vn_config", "inside_static_routes", "static_route_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.inside_static_routes

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- vn_config.inside_static_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Upstream description:

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

- [static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/): complete subsection reference.

## Next pages

- [vn_config.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/vn_config/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
