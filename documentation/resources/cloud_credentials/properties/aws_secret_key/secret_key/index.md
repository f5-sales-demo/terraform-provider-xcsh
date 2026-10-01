---
page_title: "aws_secret_key.secret_key"
subcategory: "Infrastructure"
description: "aws_secret_key.secret_key for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 2306, "body_sha256": "sha256:b9b1f4d11217ec06e19971294a28eeda31f20da2d0267134116cc1077a3403e6", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:aws_secret_key:secret_key:blindfold_secret_info", "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key:secret_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key:secret_key", "parent_id": "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key", "path": "documentation/resources/cloud_credentials/properties/aws_secret_key/secret_key/index.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["aws_secret_key", "secret_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/aws_secret_key/secret_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_secret_key.secret_key for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_secret_key.secret_key

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/)
- aws_secret_key.secret_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/clear_secret_info/): complete subsection reference.

## Next pages

- [aws_secret_key.secret_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/)
- [aws_secret_key.secret_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/clear_secret_info/)
- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
