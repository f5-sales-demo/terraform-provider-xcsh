---
page_title: "aws_provider.aws_secret_key"
subcategory: ""
description: "AWS Programmatic Access Credentials type."
xcsh_docs: {"aliases": ["aws provider aws secret key"], "body_bytes": 1922, "body_sha256": "sha256:b19d0dae6c55467db05f5b88264b251672860f101b1ef15ac05e8c1e2333c3d6", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key", "parent_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider", "path": "documentation/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3032231322023332-0223202202222220-3003320111200233-2000232003120003-3021113313232312-1212031313231100-0211021021232023-2213300022133122", "registry_path": "docs/guides/data-sources--cloud_user_account--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_secret_key"], "schema_version": 1, "sections": [{"aliases": ["aws provider aws secret key access key"], "anchor": "schema-aws_provider--aws_secret_key--access_key", "description": "Access key ID for your AWS account.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "access_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws provider aws secret key secret key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "secret_key"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "AWS Programmatic Access Credentials type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/)
- aws_provider.aws_secret_key

<a id="section"></a>

Type: `"single"`. Computed.

AWS Programmatic Access Credentials type.

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

## Direct properties

<a id="schema-aws_provider--aws_secret_key--access_key"></a>

### access_key property

Type: `"string"`. Computed.

Access Key ID. Access key ID for your AWS account.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/): complete subsection reference.
