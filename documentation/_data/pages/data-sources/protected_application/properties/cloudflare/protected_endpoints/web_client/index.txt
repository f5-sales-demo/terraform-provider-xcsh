---
page_title: "cloudflare.protected_endpoints.web_client"
subcategory: ""
description: "Web client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudflare protected endpoints web client"], "body_bytes": 2712, "body_sha256": "sha256:8cb3640c9b097e5b9f2fb91a097a696b5d76a9a38fe5abe5eee4cf0319d710ca", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_client"], "schema_version": 1, "sections": [{"aliases": ["cloudflare protected endpoints web client block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_client", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare protected endpoints web client continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_client", "continue"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare protected endpoints web client redirect"], "anchor": "section", "description": "Redirect.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:protected_endpoints:web_client:redirect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_client", "redirect"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Web client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/)
- cloudflare.protected_endpoints.web_client

<a id="section"></a>

Type: `"single"`. Computed.

Web Client. Web client configuration OPTIONS.

Upstream description:

Web client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\",\"redirect\"]"
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/block/): complete subsection reference.

- [continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/continue/): complete subsection reference.

- [redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/redirect/): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_client.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/block/)
- [cloudflare.protected_endpoints.web_client.continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/continue/)
- [cloudflare.protected_endpoints.web_client.redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/web_client/redirect/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudflare/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
