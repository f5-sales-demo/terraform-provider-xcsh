---
page_title: "syslog.tls_server.mtls_enable.key_url.clear_secret_info"
subcategory: "Monitoring"
description: "syslog.tls_server.mtls_enable.key_url.clear_secret_info for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3617, "body_sha256": "sha256:fa3b0cb11c2e19de0c44e015d3f341be933959f53f133a2f1809aa73c0461403", "canonical_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url:clear_secret_info", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url", "path": "docs/guides/resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--clear_secret_info.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_enable", "key_url", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog.tls_server.mtls_enable.key_url.clear_secret_info for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# syslog.tls_server.mtls_enable.key_url.clear_secret_info

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md)
- [Property reference](resources--log_receiver--reference.md)
- [syslog](resources--log_receiver--properties--syslog.md)
- [syslog.tls_server](resources--log_receiver--properties--syslog--tls_server.md)
- [syslog.tls_server.mtls_enable](resources--log_receiver--properties--syslog--tls_server--mtls_enable.md)
- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url.md)
- syslog.tls_server.mtls_enable.key_url.clear_secret_info

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

<a id="schema-syslog--tls_server--mtls_enable--key_url--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-syslog--tls_server--mtls_enable--key_url--clear_secret_info--url"></a>

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

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url.md)
- [xcsh_log_receiver](../resources/log_receiver.md)
