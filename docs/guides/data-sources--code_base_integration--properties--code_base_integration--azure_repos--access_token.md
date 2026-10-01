---
page_title: "code_base_integration.azure_repos.access_token"
subcategory: ""
description: "code_base_integration.azure_repos.access_token for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 1968, "body_sha256": "sha256:bd17abdd4eb8eac9b4f5c8189de14377bae58af75d351d29cb402beee20fe9a1", "canonical_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos:access_token", "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos:access_token:blindfold_secret_info", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos:access_token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos:access_token", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:azure_repos", "path": "docs/guides/data-sources--code_base_integration--properties--code_base_integration--azure_repos--access_token.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration", "azure_repos", "access_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/azure_repos/access_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration.azure_repos.access_token for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.azure_repos.access_token

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md)
- [Property reference](data-sources--code_base_integration--reference.md)
- [code_base_integration](data-sources--code_base_integration--properties--code_base_integration.md)
- [code_base_integration.azure_repos](data-sources--code_base_integration--properties--code_base_integration--azure_repos.md)
- code_base_integration.azure_repos.access_token

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

- [blindfold_secret_info](data-sources--code_base_integration--properties--code_base_integration--azure_repos--access_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--properties--code_base_integration--azure_repos--access_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [code_base_integration.azure_repos.access_token.blindfold_secret_info](data-sources--code_base_integration--properties--code_base_integration--azure_repos--access_token--blindfold_secret_info.md)
- [code_base_integration.azure_repos.access_token.clear_secret_info](data-sources--code_base_integration--properties--code_base_integration--azure_repos--access_token--clear_secret_info.md)
- [code_base_integration.azure_repos](data-sources--code_base_integration--properties--code_base_integration--azure_repos.md)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md)
