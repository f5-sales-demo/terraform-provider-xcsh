---
page_title: "aws_provider.aws_secret_key"
subcategory: ""
description: "AWS Programmatic Access Credentials type."
xcsh_docs: {"aliases": ["aws provider aws secret key"], "body_bytes": 2819, "body_sha256": "sha256:e30aedcb188d698580a8dc635abdbaf700878f2502892db9f11bd4a7a704cf61", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider", "path": "documentation/resources/cloud_user_account/properties/aws_provider/aws_secret_key/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130", "registry_path": "docs/guides/resources--cloud_user_account--reference--group-001.md", "relationships": [{"anchor": "schema-aws_provider--aws_secret_key--access_key", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key:RequiredObjectAttributes:access_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_secret_key"], "schema_version": 1, "sections": [{"aliases": ["access key"], "anchor": "schema-aws_provider--aws_secret_key--access_key", "description": "Access key ID for your AWS account.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "access_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["secret key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key.secret_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key.secret_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["aws_provider", "aws_secret_key", "secret_key"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_secret_key/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AWS Programmatic Access Credentials type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/): complete subsection reference.

## Next pages

- [aws_provider.aws_secret_key.secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
