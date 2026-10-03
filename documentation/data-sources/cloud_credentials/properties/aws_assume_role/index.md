---
page_title: "aws_assume_role"
subcategory: "Infrastructure"
description: "AWS Assume Role to Handle Delegated Access."
xcsh_docs: {"aliases": ["aws assume role"], "body_bytes": 8263, "body_sha256": "sha256:76ac3399e6f3545193fa430884c2f740785425ea992b99bfbbc63ad7a6321ca0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "parent_id": "xcsh-docs:data-sources:cloud_credentials:reference", "path": "documentation/data-sources/cloud_credentials/properties/aws_assume_role/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3213022131213113-2301210111023231-2312122222210231-1011301133122301-3100211101032221-3200113202110032-3133322013030213-2101311121002303", "registry_path": "docs/guides/data-sources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_assume_role"], "schema_version": 1, "sections": [{"aliases": ["aws assume role custom external id"], "anchor": "schema-aws_assume_role--custom_external_id", "description": "Exclusive with External ID is Custom ID.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "custom_external_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws assume role duration seconds"], "anchor": "schema-aws_assume_role--duration_seconds", "description": "The duration, in seconds of the role session.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "duration_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws assume role external id is optional"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "external_id_is_optional"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws assume role external id is tenant id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "external_id_is_tenant_id"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws assume role role arn"], "anchor": "schema-aws_assume_role--role_arn", "description": "IAM Role ARN to assume the role.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "role_arn"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws assume role session name"], "anchor": "schema-aws_assume_role--session_name", "description": "Use the role session name to uniquely identify a session, which will be used for deploy, monitor from F5XC console.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "session_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws assume role session tags"], "anchor": "schema-aws_assume_role--session_tags", "description": "Session tags are key-value pair attributes that you pass when you assume an IAM role.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:aws_assume_role", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_assume_role", "session_tags"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/aws_assume_role/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AWS Assume Role to Handle Delegated Access.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_assume_role

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- aws_assume_role

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_assume\_role, aws\_secret\_key, azure\_client\_secret, azure\_pfx\_certificate,
gcp\_cred\_file\] AWS Assume Role to Handle Delegated Access.

Upstream description:

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

OneOf alternatives in this subsection:

- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_assume_role/#section)
- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_secret_key/#section)
- [azure_client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_client_secret/#section)
- [azure_pfx_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/azure_pfx_certificate/#section)
- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/gcp_cred_file/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-aws_assume_role--custom_external_id"></a>

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

<a id="schema-aws_assume_role--duration_seconds"></a>

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

- [external_id_is_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/): complete subsection reference.

- [external_id_is_tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_assume_role/external_id_is_tenant_id/): complete subsection reference.

<a id="schema-aws_assume_role--role_arn"></a>

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

<a id="schema-aws_assume_role--session_name"></a>

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

<a id="schema-aws_assume_role--session_tags"></a>

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

- [aws_assume_role.external_id_is_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/)
- [aws_assume_role.external_id_is_tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/aws_assume_role/external_id_is_tenant_id/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
