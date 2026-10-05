---
page_title: "cloudfront.protected_endpoints.web_client"
subcategory: ""
description: "Web client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudfront protected endpoints web client"], "body_bytes": 2712, "body_sha256": "sha256:7ba8b5db81553ca65481993c27f52a0cbb046091e536c5404cf33c9f8ad817db", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:continue", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:redirect"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0021131221101020-3323020220213013-3221300121212033-3022302311332020-0101211313203023-0133110222010302-0221020001111012-0011131003021220", "registry_path": "docs/guides/data-sources--protected_application--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "web_client"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints web client block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_client", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints web client continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:continue", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_client", "continue"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints web client redirect"], "anchor": "section", "description": "Redirect.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:web_client:redirect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "web_client", "redirect"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/web_client/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Web client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
