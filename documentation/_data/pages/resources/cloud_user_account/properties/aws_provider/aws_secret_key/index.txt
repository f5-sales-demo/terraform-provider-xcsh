---
page_title: "aws_provider.aws_secret_key"
subcategory: ""
description: "AWS Programmatic Access Credentials type."
xcsh_docs: {"aliases": ["aws provider aws secret key"], "body_bytes": 2379, "body_sha256": "sha256:c6f01e246b8312f3e5dfa6482f016779baa0c30e488a057ba2645280288609d4", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "path": "documentation/resources/cloud_user_account/properties/aws_provider/aws_secret_key/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130", "registry_path": "docs/guides/resources--cloud_user_account--reference--group-001.md", "relationships": [{"anchor": "schema-aws_provider--aws_secret_key--access_key", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key:RequiredObjectAttributes:access_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_secret_key"], "schema_version": 1, "sections": [{"aliases": ["aws provider aws secret key access key"], "anchor": "schema-aws_provider--aws_secret_key--access_key", "description": "Access key ID for your AWS account.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "access_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws provider aws secret key secret key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key.secret_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key.secret_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["aws_provider", "aws_secret_key", "secret_key"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_secret_key/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "AWS Programmatic Access Credentials type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/)
- aws_provider.aws_secret_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Programmatic Access Credentials type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="schema-aws_provider--aws_secret_key--access_key"></a>

### access_key property

Type: `"string"`. Optional.

Access Key ID. Access key ID for your AWS account.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/): complete subsection reference.
