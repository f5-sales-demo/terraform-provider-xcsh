---
page_title: "aws_provider.aws_assume_role"
subcategory: ""
description: "AWS Assume Role to Handle Delegated Access."
xcsh_docs: {"aliases": ["aws provider aws assume role"], "body_bytes": 7565, "body_sha256": "sha256:c8e7110a7acc0dde5c2110743fe247beff6f8fde50a5385bd303037220f97b69", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "parent_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider", "path": "documentation/data-sources/cloud_user_account/properties/aws_provider/aws_assume_role/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2220311311210332-2112023333100233-2311303110212122-2110301013321303-1131211211030133-0031333223321010-3331013110000201-2202113120120211", "registry_path": "docs/guides/data-sources--cloud_user_account--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_assume_role"], "schema_version": 1, "sections": [{"aliases": ["custom external id"], "anchor": "schema-aws_provider--aws_assume_role--custom_external_id", "description": "Exclusive with External ID is Custom ID.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "custom_external_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration seconds"], "anchor": "schema-aws_provider--aws_assume_role--duration_seconds", "description": "The duration, in seconds of the role session.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "duration_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["external id is optional"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "external_id_is_optional"], "syntax": "attribute", "type": "object"}, {"aliases": ["external id is tenant id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "external_id_is_tenant_id"], "syntax": "attribute", "type": "object"}, {"aliases": ["role arn"], "anchor": "schema-aws_provider--aws_assume_role--role_arn", "description": "IAM Role ARN to assume the role.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "role_arn"], "syntax": "attribute", "type": "string"}, {"aliases": ["session name"], "anchor": "schema-aws_provider--aws_assume_role--session_name", "description": "Use the role session name to uniquely identify a session, which will be used for deploy, monitor from F5XC console.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "session_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["session tags"], "anchor": "schema-aws_provider--aws_assume_role--session_tags", "description": "Session tags are key-value pair attributes that you pass when you assume an IAM role.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_assume_role", "session_tags"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/properties/aws_provider/aws_assume_role/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AWS Assume Role to Handle Delegated Access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_assume_role

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/)
- aws_provider.aws_assume_role

<a id="section"></a>

Type: `"single"`. Computed.

AWS Assume Role to Handle Delegated Access.

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

## Direct properties

<a id="schema-aws_provider--aws_assume_role--custom_external_id"></a>

### custom_external_id property

Type: `"string"`. Computed.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Upstream description:

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

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

Type: `"number"`. Computed.

The duration, in seconds of the role session.

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

- [external_id_is_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_optional/): complete subsection reference.

- [external_id_is_tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_tenant_id/): complete subsection reference.

<a id="schema-aws_provider--aws_assume_role--role_arn"></a>

### role_arn property

Type: `"string"`. Computed.

IAM Role ARN. IAM Role ARN to assume the role.

Upstream description:

IAM Role ARN to assume the role.

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

Type: `"string"`. Computed.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Upstream description:

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

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

Type: `["map", "string"]`. Computed.

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

- [aws_provider.aws_assume_role.external_id_is_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_optional/)
- [aws_provider.aws_assume_role.external_id_is_tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_assume_role/external_id_is_tenant_id/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
