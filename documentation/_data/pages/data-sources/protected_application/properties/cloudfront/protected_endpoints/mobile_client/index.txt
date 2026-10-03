---
page_title: "cloudfront.protected_endpoints.mobile_client"
subcategory: ""
description: "Mobile client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudfront protected endpoints mobile client"], "body_bytes": 2321, "body_sha256": "sha256:2a14c378100215ce698b3279274ffb640d0fb369b5c05bbbd223968e57a56845", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints", "path": "documentation/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322", "registry_path": "docs/guides/data-sources--protected_application--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "mobile_client"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints mobile client block"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints mobile client continue"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:mobile_client:continue", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "mobile_client", "continue"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Mobile client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.mobile_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- cloudfront.protected_endpoints.mobile_client

<a id="section"></a>

Type: `"single"`. Computed.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/block/): complete subsection reference.

- [continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.mobile_client.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/block/)
- [cloudfront.protected_endpoints.mobile_client.continue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/mobile_client/continue/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
