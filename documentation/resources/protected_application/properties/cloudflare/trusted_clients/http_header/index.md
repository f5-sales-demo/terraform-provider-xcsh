---
page_title: "cloudflare.trusted_clients.http_header"
subcategory: ""
description: "Request header name and value pairs."
xcsh_docs: {"aliases": ["cloudflare trusted clients http header"], "body_bytes": 2032, "body_sha256": "sha256:3430fa485bdbf16121da69cde6489d7f0a3cd297f0f96128787f431f81dee389", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients", "path": "documentation/resources/protected_application/properties/cloudflare/trusted_clients/http_header/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header:RequiredObjectAttributes:headers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "trusted_clients", "http_header"], "schema_version": 1, "sections": [{"aliases": ["cloudflare trusted clients http header headers"], "anchor": "section", "description": "List of HTTP header name and value pairs.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cloudflare--trusted_clients--http_header--headers--exact", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-cloudflare--trusted_clients--http_header--headers--regex", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "conflicts"}, {"anchor": "schema-cloudflare--trusted_clients--http_header--headers--name", "enforcement": "provider-schema", "group": "cloudflare.trusted_clients.http_header.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:trusted_clients:http_header:headers", "type": "requires"}], "schema_path": ["cloudflare", "trusted_clients", "http_header", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/trusted_clients/http_header/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Request header name and value pairs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.trusted_clients.http_header

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/)
- cloudflare.trusted_clients.http_header

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
```

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
http_header {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/http_header/headers/): complete subsection reference.

## Next pages

- [cloudflare.trusted_clients.http_header.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/http_header/headers/)
- [cloudflare.trusted_clients](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/trusted_clients/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
