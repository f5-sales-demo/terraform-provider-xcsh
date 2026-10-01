---
page_title: "code_base_integration.github_enterprise.access_token"
subcategory: ""
description: "code_base_integration.github_enterprise.access_token for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 2319, "body_sha256": "sha256:d85210c626ad5c15cef609263e1fee2c3fbd1cdffbfee0192e8e353b5dba12ca", "canonical_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github_enterprise:access_token", "child_ids": ["xcsh-docs:resources:code_base_integration:properties:code_base_integration:github_enterprise:access_token:blindfold_secret_info", "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github_enterprise:access_token:clear_secret_info"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github_enterprise:access_token", "parent_id": "xcsh-docs:resources:code_base_integration:properties:code_base_integration:github_enterprise", "path": "docs/guides/resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["code_base_integration", "github_enterprise", "access_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/properties/code_base_integration/github_enterprise/access_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "code_base_integration.github_enterprise.access_token for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# code_base_integration.github_enterprise.access_token

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md)
- [Property reference](resources--code_base_integration--reference.md)
- [code_base_integration](resources--code_base_integration--properties--code_base_integration.md)
- [code_base_integration.github_enterprise](resources--code_base_integration--properties--code_base_integration--github_enterprise.md)
- code_base_integration.github_enterprise.access_token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
access_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [code_base_integration.github_enterprise.access_token.blindfold_secret_info](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--blindfold_secret_info.md)
- [code_base_integration.github_enterprise.access_token.clear_secret_info](resources--code_base_integration--properties--code_base_integration--github_enterprise--access_token--clear_secret_info.md)
- [code_base_integration.github_enterprise](resources--code_base_integration--properties--code_base_integration--github_enterprise.md)
- [xcsh_code_base_integration](../resources/code_base_integration.md)
