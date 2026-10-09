---
page_title: "cloudflare.trusted_clients.http_header"
subcategory: ""
description: "Request header name and value pairs."
xcsh_docs: {"aliases": ["cloudflare trusted clients http header"], "body_bytes": 1293, "body_sha256": "sha256:f0af2a0102bff817bdd57a9a7e3e9b10c7fbb5b0a10bf55544aebb52fcc806d6", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:http_header:headers"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:http_header", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients", "path": "documentation/data-sources/protected_application/properties/cloudflare/trusted_clients/http_header/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3201201012102121-1012202110211001-2233001231022023-3231112123303133-0101231102002120-2112323313002213-3222000033033211-2113110121101121", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "trusted_clients", "http_header"], "schema_version": 1, "sections": [{"aliases": ["cloudflare trusted clients http header headers"], "anchor": "section", "description": "List of HTTP header name and value pairs.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudflare", "trusted_clients", "http_header", "headers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Request header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.trusted_clients.http_header

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- [cloudflare.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/trusted_clients/)
- cloudflare.trusted_clients.http_header

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Additional upstream details:

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/trusted_clients/http_header/headers/): complete subsection reference.
