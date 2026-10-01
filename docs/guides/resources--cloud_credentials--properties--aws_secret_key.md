---
page_title: "aws_secret_key"
subcategory: "Infrastructure"
description: "aws_secret_key for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 2289, "body_sha256": "sha256:a7041e48687fcaf2c387ec6133d6c7ae47aa1b722ee19ada7fb1942273d0a543", "canonical_id": "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:aws_secret_key:secret_key"], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key", "parent_id": "xcsh-docs:resources:cloud_credentials:reference", "path": "docs/guides/resources--cloud_credentials--properties--aws_secret_key.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_secret_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/aws_secret_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_secret_key for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_secret_key

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
- [Property reference](resources--cloud_credentials--reference.md)
- aws_secret_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Programmatic Access Credentials type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("access_key")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
aws_secret_key {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_secret_key--access_key"></a>

### access_key property

Type: `"string"`. Optional.

Access Key ID. Access key ID for your AWS account.

Upstream description:

Access key ID for your AWS account.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [secret_key](resources--cloud_credentials--properties--aws_secret_key--secret_key.md): complete subsection reference.

## Next pages

- [aws_secret_key.secret_key](resources--cloud_credentials--properties--aws_secret_key--secret_key.md)
- [Property reference](resources--cloud_credentials--reference.md)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
