---
page_title: "aws_provider.aws_secret_key.secret_key"
subcategory: ""
description: "aws_provider.aws_secret_key.secret_key for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 2601, "body_sha256": "sha256:f783a8c141519993d98de6b161ff71991551ba22c11721b8ad72018da9ee9c44", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:blindfold_secret_info", "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "path": "documentation/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/index.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["aws_provider", "aws_secret_key", "secret_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_secret_key.secret_key for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key.secret_key

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/)
- [aws_provider.aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/)
- aws_provider.aws_secret_key.secret_key

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
secret_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/clear_secret_info/): complete subsection reference.

## Next pages

- [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/blindfold_secret_info/)
- [aws_provider.aws_secret_key.secret_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/clear_secret_info/)
- [aws_provider.aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
