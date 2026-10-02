---
page_title: "cloudfront.mobile_sdk_config.mobile_identifier"
subcategory: ""
description: "Mobile traffic identifier type."
xcsh_docs: {"aliases": ["cloudfront mobile sdk config mobile identifier"], "body_bytes": 1851, "body_sha256": "sha256:6a4e54803d50dcf28437f3182d688154e41d63b9af84da17ec95c23f9405bd01", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config", "path": "documentation/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3111022200003202-3213033223231322-3001030302002231-1021330202220111-0022300303131033-3121101331221213-0031123031121320-3020332023002031", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "section", "description": "A list of headers that can be used to identify mobile traffic.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier", "headers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Mobile traffic identifier type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.mobile_sdk_config.mobile_identifier

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [cloudfront.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/)
- cloudfront.mobile_sdk_config.mobile_identifier

<a id="section"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/headers/): complete subsection reference.

## Next pages

- [cloudfront.mobile_sdk_config.mobile_identifier.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/headers/)
- [cloudfront.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
