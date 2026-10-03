---
page_title: "domains.credentials.basic_auth.password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["domains credentials basic auth password"], "body_bytes": 2388, "body_sha256": "sha256:dbcfde64f061ac5c959bf1ad7b96c0cc4d6a135bab2a0755cc014d5e95d036ac", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password:blindfold_secret_info", "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password", "parent_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth", "path": "documentation/data-sources/api_testing/properties/domains/credentials/basic_auth/password/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1110003032023122-2000201110200230-0320101020230033-2311010103202303-0021230333313321-3113200303121203-0213030333023202-1132021011011031", "registry_path": "docs/guides/data-sources--api_testing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["domains", "credentials", "basic_auth", "password"], "schema_version": 1, "sections": [{"aliases": ["domains credentials basic auth password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "basic_auth", "password", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains credentials basic auth password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:api_testing:properties:domains:credentials:basic_auth:password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["domains", "credentials", "basic_auth", "password", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_testing/properties/domains/credentials/basic_auth/password/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["api_testingCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# domains.credentials.basic_auth.password

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/)
- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/)
- [domains.credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/)
- [domains.credentials.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/)
- domains.credentials.basic_auth.password

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/password/clear_secret_info/): complete subsection reference.

## Next pages

- [domains.credentials.basic_auth.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/password/blindfold_secret_info/)
- [domains.credentials.basic_auth.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/password/clear_secret_info/)
- [domains.credentials.basic_auth](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/properties/domains/credentials/basic_auth/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_testing/)
