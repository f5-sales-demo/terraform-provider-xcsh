---
page_title: "aws_provider.aws_secret_key.secret_key.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["aws provider aws secret key secret key clear secret info"], "body_bytes": 4053, "body_sha256": "sha256:5fcdf733f3fb8b88a5278efa7a66ce281e16e0c8d4070acda7e3f611ab7b72f0", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "parent_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key", "path": "documentation/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1133313101003120-3013100103310303-2231312031111233-0222222011322110-3233320133321320-2211032113123301-1002331322322331-2320323030303113", "registry_path": "docs/guides/resources--cloud_user_account--reference--group-001.md", "relationships": [{"anchor": "schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "aws_provider.aws_secret_key.secret_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_secret_key", "secret_key", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "secret_key", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:resources:cloud_user_account:properties:aws_provider:aws_secret_key:secret_key:clear_secret_info", "flags": ["optional", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_secret_key", "secret_key", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_secret_key.secret_key.clear_secret_info

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/)
- [aws_provider.aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/)
- [aws_provider.aws_secret_key.secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/)
- aws_provider.aws_secret_key.secret_key.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-aws_provider--aws_secret_key--secret_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [aws_provider.aws_secret_key.secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/aws_provider/aws_secret_key/secret_key/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
