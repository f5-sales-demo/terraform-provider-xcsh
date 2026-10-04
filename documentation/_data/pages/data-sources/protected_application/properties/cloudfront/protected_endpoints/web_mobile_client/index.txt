---
page_title: "cloudfront.protected_endpoints.web_mobile_client"
subcategory: ""
description: "Web and Mobile client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudfront protected endpoints web mobile client"], "body_bytes": 3785, "body_sha256": "sha256:43b3a88eb6b4077bf7e78b0a6ace883c274f70f9844f0f8b6fffd90a1c9206d5", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_web", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:redirect_web"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1221222210233203-0121321330221033-3221101000112032-3233233303213100-2313001001111110-3313022302233212-3003311322313220-1200102012332120", "registry_path": "docs/guides/data-sources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints web mobile client block mobile"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_mobile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "block_mobile"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints web mobile client block web"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:block_web", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "block_web"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints web mobile client continue mobile"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_mobile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "continue_mobile"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints web mobile client continue web"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:continue_web", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "continue_web"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints web mobile client redirect web"], "anchor": "section", "description": "Redirect.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_mobile_client:redirect_web", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_mobile_client", "redirect_web"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Web and Mobile client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.web_mobile_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- cloudfront.protected_endpoints.web_mobile_client

<a id="section"></a>

Type: `"single"`. Computed.

Web and Mobile client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mobile_mitigation": "[\"block_mobile\",\"continue_mobile\"]",
  "x-ves-oneof-field-web_mitigation": "[\"block_web\",\"continue_web\",\"redirect_web\"]"
}
```

## Direct properties

- [block_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_mobile/): complete subsection reference.

- [block_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_web/): complete subsection reference.

- [continue_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/): complete subsection reference.

- [continue_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/): complete subsection reference.

- [redirect_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/redirect_web/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.web_mobile_client.block_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_mobile/)
- [cloudfront.protected_endpoints.web_mobile_client.block_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/block_web/)
- [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_mobile/)
- [cloudfront.protected_endpoints.web_mobile_client.continue_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/continue_web/)
- [cloudfront.protected_endpoints.web_mobile_client.redirect_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_mobile_client/redirect_web/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
