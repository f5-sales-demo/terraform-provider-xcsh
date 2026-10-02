---
page_title: "code_base_integration.github_enterprise.access_token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["code base integration github enterprise access token"], "body_bytes": 2538, "body_sha256": "sha256:c5c9183f1f11858c49336cc3acad145f4b2743518ce638684979aa9bb5033ed4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise:access_token:blindfold_secret_info", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise:access_token:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise:access_token", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1303332230321330-1023001301112323-2320022023200131-3231100132330131-3210130333030330-3303321030033103-3301331233301210-1330133302131301", "registry_path": "docs/guides/data-sources--code_base_integration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "github_enterprise", "access_token"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise:access_token:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github_enterprise", "access_token", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:github_enterprise:access_token:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "github_enterprise", "access_token", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.github_enterprise.access_token

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- [code_base_integration.github_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/)
- code_base_integration.github_enterprise.access_token

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/clear_secret_info/): complete subsection reference.

## Next pages

- [code_base_integration.github_enterprise.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/blindfold_secret_info/)
- [code_base_integration.github_enterprise.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/clear_secret_info/)
- [code_base_integration.github_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/github_enterprise/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
