---
page_title: "advanced_options.http1_config"
subcategory: "Load Balancing"
description: "HTTP/1.1 Protocol OPTIONS for upstream connections."
xcsh_docs: {"aliases": ["advanced options http1 config"], "body_bytes": 1037, "body_sha256": "sha256:b8d8ba0bef709358f329270668f987587465e9a3c3c015560ab06e5663f6dc0f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "documentation/data-sources/origin_pool/properties/advanced_options/http1_config/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2333202000210100-2223123212311300-0121112332320203-3313101312222030-3120331312233212-3213301021131320-2102101023112012-1201110021332032", "registry_path": "docs/guides/data-sources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "http1_config"], "schema_version": 1, "sections": [{"aliases": ["advanced options http1 config header transformation"], "anchor": "section", "description": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "document_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config:header_transformation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "http1_config", "header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/http1_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "HTTP/1.1 Protocol OPTIONS for upstream connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
