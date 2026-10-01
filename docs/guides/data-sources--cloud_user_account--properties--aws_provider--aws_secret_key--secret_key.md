---
page_title: "aws_provider.aws_secret_key.secret_key"
subcategory: ""
description: "aws_provider.aws_secret_key.secret_key for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 1826, "body_sha256": "sha256:95da03175032a1b4887409798e9b3fdfd900a624ad7b11281ace4ab040653c95", "canonical_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "child_ids": ["xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:blindfold_secret_info", "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "parent_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key", "path": "docs/guides/data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_secret_key", "secret_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_secret_key.secret_key for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key.secret_key

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
- [Property reference](data-sources--cloud_user_account--reference.md)
- [aws_provider](data-sources--cloud_user_account--properties--aws_provider.md)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key.md)
- aws_provider.aws_secret_key.secret_key

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

- [blindfold_secret_info](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--blindfold_secret_info.md)
- [aws_provider.aws_secret_key.secret_key.clear_secret_info](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--clear_secret_info.md)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key.md)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
