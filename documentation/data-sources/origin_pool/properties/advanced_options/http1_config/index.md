---
page_title: "advanced_options.http1_config"
subcategory: "Load Balancing"
description: "HTTP/1.1 Protocol OPTIONS for upstream connections."
xcsh_docs: {"aliases": ["advanced options http1 config"], "body_bytes": 1491, "body_sha256": "sha256:19ea593061503135cc654bbe2e1cd57a43b1c3fbb50be4b85aac7a24bfa3d664", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "documentation/data-sources/origin_pool/properties/advanced_options/http1_config/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "http1_config"], "schema_version": 1, "sections": [{"aliases": ["advanced options http1 config header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "http1_config", "header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/http1_config/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "HTTP/1.1 Protocol OPTIONS for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.http1_config

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/)
- advanced_options.http1_config

<a id="section"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for upstream connections.

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

- [header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/): complete subsection reference.

## Next pages

- [advanced_options.http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/header_transformation/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
