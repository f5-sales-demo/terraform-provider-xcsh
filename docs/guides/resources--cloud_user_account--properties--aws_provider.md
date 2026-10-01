---
page_title: "aws_provider"
subcategory: ""
description: "aws_provider for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 2617, "body_sha256": "sha256:7c2aa0efbc8e37287cfa585e65956f11293f1ae54c0ba4c24d431eca8bd57342", "canonical_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key"], "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "parent_id": "xcsh-docs:resources:cloud_user_account:reference", "path": "docs/guides/resources--cloud_user_account--properties--aws_provider.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
- [Property reference](resources--cloud_user_account--reference.md)
- aws_provider

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws provider.

Upstream description:

Create AWS Provider Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_account_number"),
  validators.ConflictingObjectAttributes("aws_assume_role",
    "aws_secret_key")}
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
  "x-ves-oneof-field-aws_authentication_type": "[\"aws_assume_role\",\"aws_secret_key\"]"
}
```

Terraform syntax:

```terraform
aws_provider {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_provider--aws_account_number"></a>

### aws_account_number property

Type: `"string"`. Optional.

Account Number. 12 Digit Account Number.

Upstream description:

12 Digit Account Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  }
}
```

- [aws_assume_role](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md): complete subsection reference.

- [aws_secret_key](resources--cloud_user_account--properties--aws_provider--aws_secret_key.md): complete subsection reference.

## Next pages

- [aws_provider.aws_assume_role](resources--cloud_user_account--properties--aws_provider--aws_assume_role.md)
- [aws_provider.aws_secret_key](resources--cloud_user_account--properties--aws_provider--aws_secret_key.md)
- [Property reference](resources--cloud_user_account--reference.md)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
