---
page_title: "aws_provider.aws_secret_key.secret_key.clear_secret_info"
subcategory: ""
description: "aws_provider.aws_secret_key.secret_key.clear_secret_info for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 3307, "body_sha256": "sha256:a83748de78bc7b4df44c322726c07781ea8b82f6df0a12f9ea78fbd247f31e01", "canonical_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "parent_id": "xcsh-docs:data-sources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "path": "docs/guides/data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key--clear_secret_info.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_secret_key", "secret_key", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_secret_key.secret_key.clear_secret_info for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key.secret_key.clear_secret_info

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
- [Property reference](data-sources--cloud_user_account--reference.md)
- [aws_provider](data-sources--cloud_user_account--properties--aws_provider.md)
- [aws_provider.aws_secret_key](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key.md)
- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key.md)
- aws_provider.aws_secret_key.secret_key.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [aws_provider.aws_secret_key.secret_key](data-sources--cloud_user_account--properties--aws_provider--aws_secret_key--secret_key.md)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
