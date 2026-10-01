---
page_title: "aws_provider.aws_assume_role"
subcategory: ""
description: "aws_provider.aws_assume_role for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 8257, "body_sha256": "sha256:7be9888da51de928f367f53ec8c3e51d4038c394a7c3da81adf836b8c0636917", "canonical_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id"], "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "path": "docs/guides/resources--cloud_user_account--properties--aws_provider--aws_assume_role.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_assume_role"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_assume_role/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_assume_role for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_assume_role

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
- [Property reference](resources--cloud_user_account--reference.md)
- [aws_provider](resources--cloud_user_account--properties--aws_provider.md)
- aws_provider.aws_assume_role

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Assume Role to Handle Delegated Access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration_seconds",
    "role_arn",
    "session_name"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_optional"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_tenant_id"),
  validators.ConflictingObjectAttributes("external_id_is_optional",
    "external_id_is_tenant_id")}
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
  "x-ves-oneof-field-external_id": "[\"custom_external_id\",\"external_id_is_optional\",\"external_id_is_tenant_id\"]"
}
```

Terraform syntax:

```terraform
aws_assume_role {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_provider--aws_assume_role--custom_external_id"></a>

### custom_external_id property

Type: `"string"`. Optional.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Upstream description:

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  }
}
```

<a id="schema-aws_provider--aws_assume_role--duration_seconds"></a>

### duration_seconds property

Type: `"number"`. Optional.

The duration, in seconds of the role session.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(3600, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

- [external_id_is_optional](resources--cloud_user_account--properties--aws_provider--aws_assume_role--external_id_is_optional.md): complete subsection reference.

- [external_id_is_tenant_id](resources--cloud_user_account--properties--aws_provider--aws_assume_role--external_id_is_tenant_id.md): complete subsection reference.

<a id="schema-aws_provider--aws_assume_role--role_arn"></a>

### role_arn property

Type: `"string"`. Optional.

IAM Role ARN. IAM Role ARN to assume the role.

Upstream description:

IAM Role ARN to assume the role.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 20,
    "pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  }
}
```

<a id="schema-aws_provider--aws_assume_role--session_name"></a>

### session_name property

Type: `"string"`. Optional.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Upstream description:

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2,
    "pattern": "[\\\\w+=,.@-]*"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  }
}
```

<a id="schema-aws_provider--aws_assume_role--session_tags"></a>

### session_tags property

Type: `["map", "string"]`. Optional.

Session tags are key-value pair attributes that you pass when you assume an IAM role.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

## Next pages

- [aws_provider.aws_assume_role.external_id_is_optional](resources--cloud_user_account--properties--aws_provider--aws_assume_role--external_id_is_optional.md)
- [aws_provider.aws_assume_role.external_id_is_tenant_id](resources--cloud_user_account--properties--aws_provider--aws_assume_role--external_id_is_tenant_id.md)
- [aws_provider](resources--cloud_user_account--properties--aws_provider.md)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
