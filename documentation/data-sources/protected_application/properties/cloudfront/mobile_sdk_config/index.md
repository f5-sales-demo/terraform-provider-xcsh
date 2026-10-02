---
page_title: "cloudfront.mobile_sdk_config"
subcategory: ""
description: "Mobile SDK configuration."
xcsh_docs: {"aliases": ["cloudfront mobile sdk config"], "body_bytes": 1586, "body_sha256": "sha256:b5eab34af92438dc2b61d29078f797a2fd80ad8c707d23bd7868aa816285b016", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "path": "documentation/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2323133023333211-1301331031223130-0323301221013333-3021333133112332-2002133121021322-1321033012103203-2231303102303230-3321111230000113", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "mobile_sdk_config"], "schema_version": 1, "sections": [{"aliases": ["mobile identifier"], "anchor": "section", "description": "Mobile traffic identifier type.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:mobile_sdk_config:mobile_identifier", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudfront", "mobile_sdk_config", "mobile_identifier"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Mobile SDK configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.mobile_sdk_config

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- cloudfront.mobile_sdk_config

<a id="section"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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

- [mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/): complete subsection reference.

## Next pages

- [cloudfront.mobile_sdk_config.mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/mobile_sdk_config/mobile_identifier/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
