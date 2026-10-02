---
page_title: "cloudfront.protected_endpoints.web_client"
subcategory: ""
description: "Web client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudfront protected endpoints web client"], "body_bytes": 2712, "body_sha256": "sha256:7ba8b5db81553ca65481993c27f52a0cbb046091e536c5404cf33c9f8ad817db", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:continue", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:redirect"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0021131221101020-3323020220213013-3221300121212033-3022302311332020-0101211313203023-0133110222010302-0221020001111012-0011131003021220", "registry_path": "docs/guides/data-sources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_client"], "schema_version": 1, "sections": [{"aliases": ["block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_client", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:continue", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_client", "continue"], "syntax": "attribute", "type": "object"}, {"aliases": ["redirect"], "anchor": "section", "description": "Redirect.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:redirect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_client", "redirect"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Web client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.web_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- cloudfront.protected_endpoints.web_client

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

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/block/): complete subsection reference.

- [continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/continue/): complete subsection reference.

- [redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/redirect/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.web_client.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/block/)
- [cloudfront.protected_endpoints.web_client.continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/continue/)
- [cloudfront.protected_endpoints.web_client.redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/redirect/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
