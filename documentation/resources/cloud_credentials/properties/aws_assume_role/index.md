---
page_title: "aws_assume_role"
subcategory: "Infrastructure"
description: "AWS Assume Role to Handle Delegated Access."
xcsh_docs: {"aliases": ["aws assume role"], "body_bytes": 9608, "body_sha256": "sha256:ff57b59240b3285679a74cfacd6c9615bd7d2404a45538592603fa0d60ced606", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "parent_id": "xcsh-docs:resources:cloud_credentials:reference", "path": "documentation/resources/cloud_credentials/properties/aws_assume_role/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3201131310302023-0122011012123001-0130131130232113-3232133100213330-3012223312001231-3021312122320320-3313012123220011-3033311203311103", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [{"anchor": "schema-aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "conflicts"}, {"anchor": "schema-aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "schema-aws_assume_role--duration_seconds", "enforcement": "provider-schema", "group": "aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_assume_role--role_arn", "enforcement": "provider-schema", "group": "aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_assume_role--session_name", "enforcement": "provider-schema", "group": "aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_assume_role"], "schema_version": 1, "sections": [{"aliases": ["aws assume role custom external id"], "anchor": "schema-aws_assume_role--custom_external_id", "description": "Exclusive with External ID is Custom ID.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "custom_external_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws assume role duration seconds"], "anchor": "schema-aws_assume_role--duration_seconds", "description": "The duration, in seconds of the role session.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "duration_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws assume role external id is optional"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "external_id_is_optional"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws assume role external id is tenant id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "external_id_is_tenant_id"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws assume role role arn"], "anchor": "schema-aws_assume_role--role_arn", "description": "IAM Role ARN to assume the role.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "role_arn"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws assume role session name"], "anchor": "schema-aws_assume_role--session_name", "description": "Use the role session name to uniquely identify a session, which will be used for deploy, monitor from F5XC console.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "session_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws assume role session tags"], "anchor": "schema-aws_assume_role--session_tags", "description": "Session tags are key-value pair attributes that you pass when you assume an IAM role.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "session_tags"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/aws_assume_role/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "AWS Assume Role to Handle Delegated Access.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_assume_role

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- aws_assume_role

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_assume\_role, aws\_secret\_key, azure\_client\_secret, azure\_pfx\_certificate,
gcp\_cred\_file\] AWS Assume Role to Handle Delegated Access.

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

OneOf alternatives in this subsection:

- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/#section)
- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/#section)
- [azure_client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/#section)
- [azure_pfx_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/#section)
- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_assume_role {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_assume_role--custom_external_id"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-aws_assume_role--duration_seconds"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [external_id_is_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/): complete subsection reference.

- [external_id_is_tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/external_id_is_tenant_id/): complete subsection reference.

<a id="schema-aws_assume_role--role_arn"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-aws_assume_role--session_name"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-aws_assume_role--session_tags"></a>

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
