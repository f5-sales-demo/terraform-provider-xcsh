---
page_title: "Property reference"
subcategory: "Monitoring"
description: "Property reference for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 11902, "body_sha256": "sha256:cc791313ca613db57b74abd14e232e1cb52ea88aac5d35852aaf7b884e2d4e97", "canonical_id": "xcsh-docs:data-sources:log_receiver:reference", "child_ids": ["xcsh-docs:data-sources:log_receiver:properties:site_local", "xcsh-docs:data-sources:log_receiver:properties:syslog"], "collection_id": "xcsh-docs:data-sources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:log_receiver:reference", "parent_id": "xcsh-docs:data-sources:log_receiver:fundamentals", "path": "docs/guides/data-sources--log_receiver--reference.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/log_receiver/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the LogReceiver.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the LogReceiver.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the LogReceiver exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site_local](data-sources--log_receiver--properties--site_local.md): complete subsection reference.

- [syslog](data-sources--log_receiver--properties--syslog.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--log_receiver--reference.md#schema-annotations) |
| `description` | [description](data-sources--log_receiver--reference.md#schema-description) |
| `id` | [id](data-sources--log_receiver--reference.md#schema-id) |
| `labels` | [labels](data-sources--log_receiver--reference.md#schema-labels) |
| `name` | [name](data-sources--log_receiver--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--log_receiver--reference.md#schema-namespace) |
| `site_local` | [site_local](data-sources--log_receiver--properties--site_local.md#section) |
| `syslog` | [syslog](data-sources--log_receiver--properties--syslog.md#section) |
| `syslog.syslog_rfc5424` | [syslog.syslog_rfc5424](data-sources--log_receiver--properties--syslog.md#schema-syslog--syslog_rfc5424) |
| `syslog.tcp_server` | [syslog.tcp_server](data-sources--log_receiver--properties--syslog--tcp_server.md#section) |
| `syslog.tcp_server.port` | [syslog.tcp_server.port](data-sources--log_receiver--properties--syslog--tcp_server.md#schema-syslog--tcp_server--port) |
| `syslog.tcp_server.server_name` | [syslog.tcp_server.server_name](data-sources--log_receiver--properties--syslog--tcp_server.md#schema-syslog--tcp_server--server_name) |
| `syslog.tls_server` | [syslog.tls_server](data-sources--log_receiver--properties--syslog--tls_server.md#section) |
| `syslog.tls_server.default_https_port` | [syslog.tls_server.default_https_port](data-sources--log_receiver--properties--syslog--tls_server--default_https_port.md#section) |
| `syslog.tls_server.default_syslog_tls_port` | [syslog.tls_server.default_syslog_tls_port](data-sources--log_receiver--properties--syslog--tls_server--default_syslog_tls_port.md#section) |
| `syslog.tls_server.mtls_disabled` | [syslog.tls_server.mtls_disabled](data-sources--log_receiver--properties--syslog--tls_server--mtls_disabled.md#section) |
| `syslog.tls_server.mtls_enable` | [syslog.tls_server.mtls_enable](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable.md#section) |
| `syslog.tls_server.mtls_enable.certificate` | [syslog.tls_server.mtls_enable.certificate](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable.md#schema-syslog--tls_server--mtls_enable--certificate) |
| `syslog.tls_server.mtls_enable.key_url` | [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url.md#section) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--blindfold_secret_info.md#section) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--blindfold_secret_info.md#schema-syslog--tls_server--mtls_enable--key_url--blindfold_secret_info--decryption_provider) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--blindfold_secret_info.md#schema-syslog--tls_server--mtls_enable--key_url--blindfold_secret_info--location) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--blindfold_secret_info.md#schema-syslog--tls_server--mtls_enable--key_url--blindfold_secret_info--store_provider) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--clear_secret_info.md#section) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--clear_secret_info.md#schema-syslog--tls_server--mtls_enable--key_url--clear_secret_info--provider_ref) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.url` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.url](data-sources--log_receiver--properties--syslog--tls_server--mtls_enable--key_url--clear_secret_info.md#schema-syslog--tls_server--mtls_enable--key_url--clear_secret_info--url) |
| `syslog.tls_server.port` | [syslog.tls_server.port](data-sources--log_receiver--properties--syslog--tls_server.md#schema-syslog--tls_server--port) |
| `syslog.tls_server.server_name` | [syslog.tls_server.server_name](data-sources--log_receiver--properties--syslog--tls_server.md#schema-syslog--tls_server--server_name) |
| `syslog.tls_server.trusted_ca_url` | [syslog.tls_server.trusted_ca_url](data-sources--log_receiver--properties--syslog--tls_server.md#schema-syslog--tls_server--trusted_ca_url) |
| `syslog.tls_server.volterra_ca` | [syslog.tls_server.volterra_ca](data-sources--log_receiver--properties--syslog--tls_server--volterra_ca.md#section) |
| `syslog.udp_server` | [syslog.udp_server](data-sources--log_receiver--properties--syslog--udp_server.md#section) |
| `syslog.udp_server.port` | [syslog.udp_server.port](data-sources--log_receiver--properties--syslog--udp_server.md#schema-syslog--udp_server--port) |
| `syslog.udp_server.server_name` | [syslog.udp_server.server_name](data-sources--log_receiver--properties--syslog--udp_server.md#schema-syslog--udp_server--server_name) |

## Next pages

- [site_local](data-sources--log_receiver--properties--site_local.md)
- [syslog](data-sources--log_receiver--properties--syslog.md)
- [xcsh_log_receiver](../data-sources/log_receiver.md)
