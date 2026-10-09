---
page_title: "aws_provider.aws_assume_role"
subcategory: ""
description: "AWS Assume Role to Handle Delegated Access."
xcsh_docs: {"aliases": ["aws provider aws assume role"], "body_bytes": 8926, "body_sha256": "sha256:90be574ddb627976694187c2c20672e8a111a567c6296c0dcf66dbd3e058e09a", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "path": "documentation/resources/cloud_user_account/properties/aws_provider/aws_assume_role/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031", "registry_path": "docs/guides/resources--cloud_user_account--reference--group-001.md", "relationships": [{"anchor": "schema-aws_provider--aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "conflicts"}, {"anchor": "schema-aws_provider--aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "schema-aws_provider--aws_assume_role--duration_seconds", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_provider--aws_assume_role--role_arn", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_provider--aws_assume_role--session_name", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_assume_role"], "schema_version": 1, "sections": [{"aliases": ["aws provider aws assume role custom external id"], "anchor": "schema-aws_provider--aws_assume_role--custom_external_id", "description": "Exclusive with External ID is Custom ID.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "custom_external_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws provider aws assume role duration seconds"], "anchor": "schema-aws_provider--aws_assume_role--duration_seconds", "description": "The duration, in seconds of the role session.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "duration_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws provider aws assume role external id is optional"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "external_id_is_optional"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws provider aws assume role external id is tenant id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "external_id_is_tenant_id"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws provider aws assume role role arn"], "anchor": "schema-aws_provider--aws_assume_role--role_arn", "description": "IAM Role ARN to assume the role.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "role_arn"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws provider aws assume role session name"], "anchor": "schema-aws_provider--aws_assume_role--session_name", "description": "Use the role session name to uniquely identify a session, which will be used for deploy, monitor from F5XC console.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "session_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws provider aws assume role session tags"], "anchor": "schema-aws_provider--aws_assume_role--session_tags", "description": "Session tags are key-value pair attributes that you pass when you assume an IAM role.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "session_tags"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_assume_role/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "AWS Assume Role to Handle Delegated Access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_assume_role

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/)
- aws_provider.aws_assume_role

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Assume Role to Handle Delegated Access.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [external_id_is_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_optional/): complete subsection reference.

- [external_id_is_tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_tenant_id/): complete subsection reference.

<a id="schema-aws_provider--aws_assume_role--role_arn"></a>

### role_arn property

Type: `"string"`. Optional.

IAM Role ARN. IAM Role ARN to assume the role.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":40},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":127,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"127\",\"ves.io.schema.rules.map.max_pairs\":\"40\",\"ves.io.schema.rules.map.values.string.max_len\":\"255\"},\"values\":{\"maxLength\":255,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
      "type": "string"
    }
  },
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
