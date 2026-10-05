---
page_title: "code_base_integration.github.access_token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["code base integration github access token"], "body_bytes": 2406, "body_sha256": "sha256:8feefc8c9113eb9c5597ec0e19509362c7fd86043733cdc2ee47b4c426417ac7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github:access_token:blindfold_secret_info", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github:access_token:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github:access_token", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/github/access_token/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0031203003130223-2331110102032102-0230201121010323-2313233122220111-0213333311303010-1120232310331033-3023103133102301-3113310111220102", "registry_path": "docs/guides/data-sources--code_base_integration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "github", "access_token"], "schema_version": 1, "sections": [{"aliases": ["code base integration github access token blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github:access_token:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github", "access_token", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration github access token clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github:access_token:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github", "access_token", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/github/access_token/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.github.access_token

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- [code_base_integration.github](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/)
- code_base_integration.github.access_token

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/clear_secret_info/): complete subsection reference.

## Next pages

- [code_base_integration.github.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/blindfold_secret_info/)
- [code_base_integration.github.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/access_token/clear_secret_info/)
- [code_base_integration.github](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
