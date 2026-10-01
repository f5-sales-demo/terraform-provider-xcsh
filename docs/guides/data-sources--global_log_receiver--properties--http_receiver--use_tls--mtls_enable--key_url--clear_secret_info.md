---
page_title: "http_receiver.use_tls.mtls_enable.key_url.clear_secret_info"
subcategory: ""
description: "http_receiver.use_tls.mtls_enable.key_url.clear_secret_info for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3459, "body_sha256": "sha256:a2c7991cec8ef4b4f9a7b59134d0a7a0c20e7eebf3f4ea4148ed26fe91a538e5", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable:key_url", "path": "docs/guides/data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url--clear_secret_info.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/http_receiver/use_tls/mtls_enable/key_url/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.use_tls.mtls_enable.key_url.clear_secret_info for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [http_receiver](data-sources--global_log_receiver--properties--http_receiver.md)
- [http_receiver.use_tls](data-sources--global_log_receiver--properties--http_receiver--use_tls.md)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable.md)
- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url.md)
- http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

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

<a id="schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-http_receiver--use_tls--mtls_enable--key_url--clear_secret_info--url"></a>

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

- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable--key_url.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
