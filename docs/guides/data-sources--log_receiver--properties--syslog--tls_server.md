---
page_title: "syslog.tls_server"
subcategory: "Monitoring"
description: "syslog.tls_server for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 5803, "body_sha256": "sha256:bdf1350a8c65a85e844eb6f991c10587dd96df67efea6de8655e896c53347724", "canonical_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "child_ids": ["xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:default_https_port", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_disabled", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_enable", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:volterra_ca"], "collection_id": "xcsh-docs:data-sources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "parent_id": "xcsh-docs:data-sources:log_receiver:properties:syslog", "path": "docs/guides/data-sources--log_receiver--properties--syslog--tls_server.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["syslog", "tls_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/log_receiver/properties/syslog/tls_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog.tls_server for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md)
- [Property reference](data-sources--log_receiver--reference.md)
- [syslog](data-sources--log_receiver--properties--syslog.md)
- syslog.tls_server

<a id="section"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"trusted_ca_url\",\"volterra_ca\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"default_syslog_tls_port\",\"port\"]"
}
```

## Direct properties

- [default_https_port](data-sources--log_receiver--properties--syslog--tls_server--default_https_port.md): complete subsection reference.

- [default_syslog_tls_port](data-sources--log_receiver--properties--syslog--tls_server--default_syslog_tls_port.md): complete subsection reference.

- [mtls_disabled](data-sources--log_receiver--properties--syslog--tls_server--mtls_disabled.md): complete subsection reference.

- [mtls_enable](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable.md): complete subsection reference.

<a id="schema-syslog--tls_server--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Upstream description:

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-syslog--tls_server--server_name"></a>

### server_name property

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-syslog--tls_server--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [volterra_ca](data-sources--log_receiver--properties--syslog--tls_server--volterra_ca.md): complete subsection reference.

## Next pages

- [syslog.tls_server.default_https_port](data-sources--log_receiver--properties--syslog--tls_server--default_https_port.md)
- [syslog.tls_server.default_syslog_tls_port](data-sources--log_receiver--properties--syslog--tls_server--default_syslog_tls_port.md)
- [syslog.tls_server.mtls_disabled](data-sources--log_receiver--properties--syslog--tls_server--mtls_disabled.md)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable.md)
- [syslog.tls_server.volterra_ca](data-sources--log_receiver--properties--syslog--tls_server--volterra_ca.md)
- [syslog](data-sources--log_receiver--properties--syslog.md)
- [xcsh_log_receiver](../data-sources/log_receiver.md)
