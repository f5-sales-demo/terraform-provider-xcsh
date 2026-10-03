---
page_title: "cloudfront.trusted_clients.http_header"
subcategory: ""
description: "Request header name and value pairs."
xcsh_docs: {"aliases": ["cloudfront trusted clients http header"], "body_bytes": 1790, "body_sha256": "sha256:539880a034f58260437bcb0f90e68f7fa088ea8975db59f8e39d741607a002b1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header:headers"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients", "path": "documentation/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1033022210101312-3011123312230132-1310310020002123-2001123113112112-3130232113230121-1300102012321113-3332313312131212-1132123323322302", "registry_path": "docs/guides/data-sources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "trusted_clients", "http_header"], "schema_version": 1, "sections": [{"aliases": ["cloudfront trusted clients http header headers"], "anchor": "section", "description": "List of HTTP header name and value pairs.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:trusted_clients:http_header:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "trusted_clients", "http_header", "headers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Request header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.trusted_clients.http_header

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/)
- cloudfront.trusted_clients.http_header

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/headers/): complete subsection reference.

## Next pages

- [cloudfront.trusted_clients.http_header.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/http_header/headers/)
- [cloudfront.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/trusted_clients/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
