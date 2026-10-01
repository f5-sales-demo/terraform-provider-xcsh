---
page_title: "code_base_integration.gitlab.access_token"
subcategory: ""
description: "code_base_integration.gitlab.access_token for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 1908, "body_sha256": "sha256:36f8a6d2dececd190b37a70d525a8c393dadbe99ac5241a7d83796e98d14a7b7", "canonical_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab:access_token", "child_ids": ["xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab:access_token:blindfold_secret_info", "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab:access_token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab:access_token", "parent_id": "xcsh-docs:data-sources:code_base_integration:properties:code_base_integration:gitlab", "path": "docs/guides/data-sources--code_base_integration--properties--code_base_integration--gitlab--access_token.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration", "gitlab", "access_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/properties/code_base_integration/gitlab/access_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration.gitlab.access_token for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.gitlab.access_token

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md)
- [Property reference](data-sources--code_base_integration--reference.md)
- [code_base_integration](data-sources--code_base_integration--properties--code_base_integration.md)
- [code_base_integration.gitlab](data-sources--code_base_integration--properties--code_base_integration--gitlab.md)
- code_base_integration.gitlab.access_token

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

- [blindfold_secret_info](data-sources--code_base_integration--properties--code_base_integration--gitlab--access_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--code_base_integration--properties--code_base_integration--gitlab--access_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [code_base_integration.gitlab.access_token.blindfold_secret_info](data-sources--code_base_integration--properties--code_base_integration--gitlab--access_token--blindfold_secret_info.md)
- [code_base_integration.gitlab.access_token.clear_secret_info](data-sources--code_base_integration--properties--code_base_integration--gitlab--access_token--clear_secret_info.md)
- [code_base_integration.gitlab](data-sources--code_base_integration--properties--code_base_integration--gitlab.md)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md)
