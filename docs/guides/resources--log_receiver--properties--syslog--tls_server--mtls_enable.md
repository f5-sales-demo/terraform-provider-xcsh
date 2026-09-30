---
page_title: "syslog.tls_server.mtls_enable"
subcategory: "Monitoring"
description: "syslog.tls_server.mtls_enable for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2560, "body_sha256": "sha256:fdde836deb10d8561c3364bd7b489ab912c5c96d9248fcbc358d05bbeaf3e728", "canonical_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "child_ids": ["xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url"], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "path": "docs/guides/resources--log_receiver--properties--syslog--tls_server--mtls_enable.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_enable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/mtls_enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog.tls_server.mtls_enable for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# syslog.tls_server.mtls_enable

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md)
- [Property reference](resources--log_receiver--reference.md)
- [syslog](resources--log_receiver--properties--syslog.md)
- [syslog.tls_server](resources--log_receiver--properties--syslog--tls_server.md)
- syslog.tls_server.mtls_enable

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for mtls enable.

Upstream description:

TLS config for client.

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
mtls_enable {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-syslog--tls_server--mtls_enable--certificate"></a>

### certificate property

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url.md): complete subsection reference.

## Next pages

- [syslog.tls_server.mtls_enable.key_url](resources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url.md)
- [syslog.tls_server](resources--log_receiver--properties--syslog--tls_server.md)
- [xcsh_log_receiver](../resources/log_receiver.md)
