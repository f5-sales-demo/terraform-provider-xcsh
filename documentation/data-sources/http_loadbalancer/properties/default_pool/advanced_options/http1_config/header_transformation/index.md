---
page_title: "default_pool.advanced_options.http1_config.header_transformation"
subcategory: "Load Balancing"
description: "Header Transformation OPTIONS for HTTP/1.1 request/response headers."
xcsh_docs: {"aliases": ["default pool advanced options http1 config header transformation"], "body_bytes": 2264, "body_sha256": "sha256:744465dffcdd9aa3cafe2ac6920e54a8b1235bcf178cf6fdb23baff944307e8d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:default_header_transformation", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config", "path": "documentation/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2210232213221001-1003012011230031-0223332332021110-1200121313023211-3331133113002110-0210120312210132-1233232332323013-2312321030211220", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation"], "schema_version": 1, "sections": [{"aliases": ["default pool advanced options http1 config header transformation default header transformation"], "anchor": "section", "description": "Use the platform's current default HTTP header transformation behavior.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:default_header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "default_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool advanced options http1 config header transformation preserve case header transformation"], "anchor": "section", "description": "Preserve HTTP header-name case when upstream case must remain unchanged.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:preserve_case_header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "preserve_case_header_transformation"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool advanced options http1 config header transformation proper case header transformation"], "anchor": "section", "description": "Transform HTTP header names to proper case when explicit transformation is required.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:http1_config:header_transformation:proper_case_header_transformation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "http1_config", "header_transformation", "proper_case_header_transformation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Header Transformation OPTIONS for HTTP/1.1 request/response headers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.http1_config.header_transformation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/)
- [default_pool.advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/)
- default_pool.advanced_options.http1_config.header_transformation

<a id="section"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

## Direct properties

- [default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/default_header_transformation/): complete subsection reference.

- [preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/): complete subsection reference.

- [proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/default_pool/advanced_options/http1_config/header_transformation/proper_case_header_transformation/): complete subsection reference.
