---
page_title: "aws_provider.aws_secret_key"
subcategory: ""
description: "aws_provider.aws_secret_key for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 2075, "body_sha256": "sha256:3f6ff48ce049e68cbd9b4d5ff67a3633d04bc277e5da336cbb782d8f8f7f9447", "canonical_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key", "child_ids": ["xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key"], "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key", "parent_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider", "path": "docs/guides/data-sources--cloud_user_account--properties--aws_provider--aws_secret_key.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_secret_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_secret_key for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
- [Property reference](data-sources--cloud_user_account--reference.md)
- [aws_provider](data-sources--cloud_user_account--properties--aws_provider.md)
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

Upstream description:

Access key ID for your AWS account.

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

- [secret_key](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key.md): complete subsection reference.

## Next pages

- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key.md)
- [aws_provider](data-sources--cloud_user_account--properties--aws_provider.md)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
