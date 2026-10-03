---
page_title: "code_base_integration.gitlab_enterprise.access_token"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["code base integration gitlab enterprise access token"], "body_bytes": 2538, "body_sha256": "sha256:42656e39ec1c3185aab6327feae76d3827a2998c59e217a858b0a3d444bc2ca9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:blindfold_secret_info", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise", "path": "documentation/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3011122023332211-2213323302222103-2022023313232233-3210101231330221-1301120110112000-2201322000220113-3131312223112103-0003313032000030", "registry_path": "docs/guides/data-sources--code_base_integration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["code_base_integration", "gitlab_enterprise", "access_token"], "schema_version": 1, "sections": [{"aliases": ["code base integration gitlab enterprise access token blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "gitlab_enterprise", "access_token", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["code base integration gitlab enterprise access token clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab_enterprise:access_token:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["code_base_integration", "gitlab_enterprise", "access_token", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.gitlab_enterprise.access_token

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/)
- [code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/)
- [code_base_integration.gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/)
- code_base_integration.gitlab_enterprise.access_token

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/clear_secret_info/): complete subsection reference.

## Next pages

- [code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/blindfold_secret_info/)
- [code_base_integration.gitlab_enterprise.access_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/access_token/clear_secret_info/)
- [code_base_integration.gitlab_enterprise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/properties/code_base_integration/gitlab_enterprise/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
