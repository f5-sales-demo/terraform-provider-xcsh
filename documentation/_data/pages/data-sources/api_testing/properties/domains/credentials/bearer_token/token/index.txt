---
page_title: "domains.credentials.bearer_token.token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "domains credentials bearer token token"], "body_bytes": 2388, "body_sha256": "sha256:42ef8890b7c2d06f0f4c36f8e62bfb0f7b5f212e8ca9054f27bc2e4a2716bab0", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token:blindfold_secret_info", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token", "path": "documentation/data-sources/api_testing/properties/domains/credentials/bearer_token/token/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2321300003011131-3322013221100103-0031022101130311-3101120310232002-0011110233101202-1132233333232000-2030130212012000-3022013123230123", "registry_path": "docs/guides/data-sources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials", "bearer_token", "token"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "bearer_token", "token", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:bearer_token:token:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "bearer_token", "token", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/bearer_token/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.bearer_token.token

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/)
- [domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/)
- domains.credentials.bearer_token.token

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/token/clear_secret_info/): complete subsection reference.

## Next pages

- [domains.credentials.bearer_token.token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/token/blindfold_secret_info/)
- [domains.credentials.bearer_token.token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/token/clear_secret_info/)
- [domains.credentials.bearer_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/bearer_token/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
