---
page_title: "aws_provider"
subcategory: ""
description: "Create AWS Provider Type."
xcsh_docs: {"aliases": ["aws provider"], "body_bytes": 2443, "body_sha256": "sha256:32e9cfa159efbef60f0fb343619a4f6b4d52a95ed6762ce3af64b8246c5aebde", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "parent_id": "xcsh-docs:resources:cloud_user_account:reference", "path": "documentation/resources/cloud_user_account/properties/aws_provider/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110", "registry_path": "docs/guides/resources--cloud_user_account--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider:ConflictingObjectAttributes:aws_assume_role,aws_secret_key", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider:ConflictingObjectAttributes:aws_assume_role,aws_secret_key", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "type": "conflicts"}, {"anchor": "schema-aws_provider--aws_account_number", "enforcement": "provider-schema", "group": "aws_provider:RequiredObjectAttributes:aws_account_number", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider"], "schema_version": 1, "sections": [{"aliases": ["aws provider aws account number"], "anchor": "schema-aws_provider--aws_account_number", "description": "12 Digit Account Number.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_account_number"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws provider aws assume role"], "anchor": "section", "description": "AWS Assume Role to Handle Delegated Access.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_provider--aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "conflicts"}, {"anchor": "schema-aws_provider--aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "schema-aws_provider--aws_assume_role--duration_seconds", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_provider--aws_assume_role--role_arn", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_provider--aws_assume_role--session_name", "enforcement": "provider-schema", "group": "aws_provider.aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_assume_role", "type": "requires"}], "schema_path": ["aws_provider", "aws_assume_role"], "syntax": "block", "type": "object"}, {"aliases": ["aws provider aws secret key"], "anchor": "section", "description": "AWS Programmatic Access Credentials type.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_provider--aws_secret_key--access_key", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key:RequiredObjectAttributes:access_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "type": "requires"}], "schema_path": ["aws_provider", "aws_secret_key"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Create AWS Provider Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/)
- aws_provider

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws provider.

Additional upstream details:

Create AWS Provider Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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

- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_assume_role/): complete subsection reference.

- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/): complete subsection reference.
