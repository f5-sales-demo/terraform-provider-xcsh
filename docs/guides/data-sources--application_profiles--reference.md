---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 94765, "body_sha256": "sha256:971193ff8dc5738166294c8886e7193b456dde1d4d0d10372de6d243b1d29b9f", "canonical_id": "xcsh-docs:data-sources:application_profiles:reference", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile", "xcsh-docs:data-sources:application_profiles:properties:ddos_profile", "xcsh-docs:data-sources:application_profiles:properties:irules", "xcsh-docs:data-sources:application_profiles:properties:virtual_server"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:reference", "parent_id": "xcsh-docs:data-sources:application_profiles:fundamentals", "path": "docs/guides/data-sources--application_profiles--reference.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- Property reference

## Direct properties

- [advanced_tcp_profile](data-sources--application_profiles--properties--advanced_tcp_profile.md): complete subsection reference.

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

- [ddos_profile](data-sources--application_profiles--properties--ddos_profile.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the ApplicationProfiles.

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

- [irules](data-sources--application_profiles--properties--irules.md): complete subsection reference.

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

Name of the ApplicationProfiles.

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

Namespace where the ApplicationProfiles exists.

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

- [virtual_server](data-sources--application_profiles--properties--virtual_server.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_tcp_profile` | [advanced_tcp_profile](data-sources--application_profiles--properties--advanced_tcp_profile.md#section) |
| `advanced_tcp_profile.disable_tcp_advanced_profile` | [advanced_tcp_profile.disable_tcp_advanced_profile](data-sources--application_profiles--properties--advanced_tcp_profile--disable_tcp_advanced_profile.md#section) |
| `advanced_tcp_profile.enable_tcp_advanced_profile` | [advanced_tcp_profile.enable_tcp_advanced_profile](data-sources--application_profiles--properties--advanced_tcp_profile--enable_tcp_advanced_profile.md#section) |
| `annotations` | [annotations](data-sources--application_profiles--reference.md#schema-annotations) |
| `ddos_profile` | [ddos_profile](data-sources--application_profiles--properties--ddos_profile.md#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--application_profiles--properties--ddos_profile--disable_ddos_mitigation.md#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--application_profiles--properties--ddos_profile--enable_ddos_mitigation.md#section) |
| `description` | [description](data-sources--application_profiles--reference.md#schema-description) |
| `id` | [id](data-sources--application_profiles--reference.md#schema-id) |
| `irules` | [irules](data-sources--application_profiles--properties--irules.md#section) |
| `irules.kind` | [irules.kind](data-sources--application_profiles--properties--irules.md#schema-irules--kind) |
| `irules.name` | [irules.name](data-sources--application_profiles--properties--irules.md#schema-irules--name) |
| `irules.namespace` | [irules.namespace](data-sources--application_profiles--properties--irules.md#schema-irules--namespace) |
| `irules.tenant` | [irules.tenant](data-sources--application_profiles--properties--irules.md#schema-irules--tenant) |
| `irules.uid` | [irules.uid](data-sources--application_profiles--properties--irules.md#schema-irules--uid) |
| `labels` | [labels](data-sources--application_profiles--reference.md#schema-labels) |
| `name` | [name](data-sources--application_profiles--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--application_profiles--reference.md#schema-namespace) |
| `virtual_server` | [virtual_server](data-sources--application_profiles--properties--virtual_server.md#section) |
| `virtual_server.access_profile` | [virtual_server.access_profile](data-sources--application_profiles--properties--virtual_server--access_profile.md#section) |
| `virtual_server.access_profile.kind` | [virtual_server.access_profile.kind](data-sources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--kind) |
| `virtual_server.access_profile.name` | [virtual_server.access_profile.name](data-sources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--name) |
| `virtual_server.access_profile.namespace` | [virtual_server.access_profile.namespace](data-sources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--namespace) |
| `virtual_server.access_profile.tenant` | [virtual_server.access_profile.tenant](data-sources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--tenant) |
| `virtual_server.access_profile.uid` | [virtual_server.access_profile.uid](data-sources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--uid) |
| `virtual_server.address_translation` | [virtual_server.address_translation](data-sources--application_profiles--properties--virtual_server--address_translation.md#section) |
| `virtual_server.address_translation.address_translation_disable` | [virtual_server.address_translation.address_translation_disable](data-sources--application_profiles--properties--virtual_server--address_translation--address_translation_disable.md#section) |
| `virtual_server.address_translation.address_translation_enable` | [virtual_server.address_translation.address_translation_enable](data-sources--application_profiles--properties--virtual_server--address_translation--address_translation_enable.md#section) |
| `virtual_server.auto_last_hop` | [virtual_server.auto_last_hop](data-sources--application_profiles--properties--virtual_server--auto_last_hop.md#section) |
| `virtual_server.auto_last_hop.auto_last_hop_default` | [virtual_server.auto_last_hop.auto_last_hop_default](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_default.md#section) |
| `virtual_server.auto_last_hop.auto_last_hop_disable` | [virtual_server.auto_last_hop.auto_last_hop_disable](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_disable.md#section) |
| `virtual_server.auto_last_hop.auto_last_hop_enable` | [virtual_server.auto_last_hop.auto_last_hop_enable](data-sources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_enable.md#section) |
| `virtual_server.clone_pool_client` | [virtual_server.clone_pool_client](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md#section) |
| `virtual_server.clone_pool_client.kind` | [virtual_server.clone_pool_client.kind](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--kind) |
| `virtual_server.clone_pool_client.name` | [virtual_server.clone_pool_client.name](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--name) |
| `virtual_server.clone_pool_client.namespace` | [virtual_server.clone_pool_client.namespace](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--namespace) |
| `virtual_server.clone_pool_client.tenant` | [virtual_server.clone_pool_client.tenant](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--tenant) |
| `virtual_server.clone_pool_client.uid` | [virtual_server.clone_pool_client.uid](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--uid) |
| `virtual_server.clone_pool_server` | [virtual_server.clone_pool_server](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md#section) |
| `virtual_server.clone_pool_server.kind` | [virtual_server.clone_pool_server.kind](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--kind) |
| `virtual_server.clone_pool_server.name` | [virtual_server.clone_pool_server.name](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--name) |
| `virtual_server.clone_pool_server.namespace` | [virtual_server.clone_pool_server.namespace](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--namespace) |
| `virtual_server.clone_pool_server.tenant` | [virtual_server.clone_pool_server.tenant](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--tenant) |
| `virtual_server.clone_pool_server.uid` | [virtual_server.clone_pool_server.uid](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--uid) |
| `virtual_server.connection_limit` | [virtual_server.connection_limit](data-sources--application_profiles--properties--virtual_server.md#schema-virtual_server--connection_limit) |
| `virtual_server.connection_rate_limit` | [virtual_server.connection_rate_limit](data-sources--application_profiles--properties--virtual_server.md#schema-virtual_server--connection_rate_limit) |
| `virtual_server.connection_rate_limit_mode` | [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md#section) |
| `virtual_server.connection_rate_limit_mode.per_destination_address` | [virtual_server.connection_rate_limit_mode.per_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_address` | [virtual_server.connection_rate_limit_mode.per_source_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_address.source_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md#schema-virtual_server--connection_rate_limit_mode--per_source_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_source_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server` | [virtual_server.connection_rate_limit_mode.per_virtual_server](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address--source_mask) |
| `virtual_server.default_persistence_profile` | [virtual_server.default_persistence_profile](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md#section) |
| `virtual_server.default_persistence_profile.kind` | [virtual_server.default_persistence_profile.kind](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--kind) |
| `virtual_server.default_persistence_profile.name` | [virtual_server.default_persistence_profile.name](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--name) |
| `virtual_server.default_persistence_profile.namespace` | [virtual_server.default_persistence_profile.namespace](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--namespace) |
| `virtual_server.default_persistence_profile.tenant` | [virtual_server.default_persistence_profile.tenant](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--tenant) |
| `virtual_server.default_persistence_profile.uid` | [virtual_server.default_persistence_profile.uid](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--uid) |
| `virtual_server.default_pool` | [virtual_server.default_pool](data-sources--application_profiles--properties--virtual_server--default_pool.md#section) |
| `virtual_server.default_pool.kind` | [virtual_server.default_pool.kind](data-sources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--kind) |
| `virtual_server.default_pool.name` | [virtual_server.default_pool.name](data-sources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--name) |
| `virtual_server.default_pool.namespace` | [virtual_server.default_pool.namespace](data-sources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--namespace) |
| `virtual_server.default_pool.tenant` | [virtual_server.default_pool.tenant](data-sources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--tenant) |
| `virtual_server.default_pool.uid` | [virtual_server.default_pool.uid](data-sources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--uid) |
| `virtual_server.fallback_persistence_profile` | [virtual_server.fallback_persistence_profile](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#section) |
| `virtual_server.fallback_persistence_profile.kind` | [virtual_server.fallback_persistence_profile.kind](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--kind) |
| `virtual_server.fallback_persistence_profile.name` | [virtual_server.fallback_persistence_profile.name](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--name) |
| `virtual_server.fallback_persistence_profile.namespace` | [virtual_server.fallback_persistence_profile.namespace](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--namespace) |
| `virtual_server.fallback_persistence_profile.tenant` | [virtual_server.fallback_persistence_profile.tenant](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--tenant) |
| `virtual_server.fallback_persistence_profile.uid` | [virtual_server.fallback_persistence_profile.uid](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--uid) |
| `virtual_server.fix_profile` | [virtual_server.fix_profile](data-sources--application_profiles--properties--virtual_server--fix_profile.md#section) |
| `virtual_server.fix_profile.kind` | [virtual_server.fix_profile.kind](data-sources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--kind) |
| `virtual_server.fix_profile.name` | [virtual_server.fix_profile.name](data-sources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--name) |
| `virtual_server.fix_profile.namespace` | [virtual_server.fix_profile.namespace](data-sources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--namespace) |
| `virtual_server.fix_profile.tenant` | [virtual_server.fix_profile.tenant](data-sources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--tenant) |
| `virtual_server.fix_profile.uid` | [virtual_server.fix_profile.uid](data-sources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--uid) |
| `virtual_server.http` | [virtual_server.http](data-sources--application_profiles--properties--virtual_server--http.md#section) |
| `virtual_server.http.client_ssl_profile` | [virtual_server.http.client_ssl_profile](data-sources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#section) |
| `virtual_server.http.client_ssl_profile.kind` | [virtual_server.http.client_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--kind) |
| `virtual_server.http.client_ssl_profile.name` | [virtual_server.http.client_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--name) |
| `virtual_server.http.client_ssl_profile.namespace` | [virtual_server.http.client_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--namespace) |
| `virtual_server.http.client_ssl_profile.tenant` | [virtual_server.http.client_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--tenant) |
| `virtual_server.http.client_ssl_profile.uid` | [virtual_server.http.client_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--uid) |
| `virtual_server.http.http2_client_profile` | [virtual_server.http.http2_client_profile](data-sources--application_profiles--properties--virtual_server--http--http2_client_profile.md#section) |
| `virtual_server.http.http2_client_profile.kind` | [virtual_server.http.http2_client_profile.kind](data-sources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--kind) |
| `virtual_server.http.http2_client_profile.name` | [virtual_server.http.http2_client_profile.name](data-sources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--name) |
| `virtual_server.http.http2_client_profile.namespace` | [virtual_server.http.http2_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--namespace) |
| `virtual_server.http.http2_client_profile.tenant` | [virtual_server.http.http2_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--tenant) |
| `virtual_server.http.http2_client_profile.uid` | [virtual_server.http.http2_client_profile.uid](data-sources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--uid) |
| `virtual_server.http.http2_server_profile` | [virtual_server.http.http2_server_profile](data-sources--application_profiles--properties--virtual_server--http--http2_server_profile.md#section) |
| `virtual_server.http.http2_server_profile.kind` | [virtual_server.http.http2_server_profile.kind](data-sources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--kind) |
| `virtual_server.http.http2_server_profile.name` | [virtual_server.http.http2_server_profile.name](data-sources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--name) |
| `virtual_server.http.http2_server_profile.namespace` | [virtual_server.http.http2_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--namespace) |
| `virtual_server.http.http2_server_profile.tenant` | [virtual_server.http.http2_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--tenant) |
| `virtual_server.http.http2_server_profile.uid` | [virtual_server.http.http2_server_profile.uid](data-sources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--uid) |
| `virtual_server.http.http_client_profile` | [virtual_server.http.http_client_profile](data-sources--application_profiles--properties--virtual_server--http--http_client_profile.md#section) |
| `virtual_server.http.http_client_profile.kind` | [virtual_server.http.http_client_profile.kind](data-sources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--kind) |
| `virtual_server.http.http_client_profile.name` | [virtual_server.http.http_client_profile.name](data-sources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--name) |
| `virtual_server.http.http_client_profile.namespace` | [virtual_server.http.http_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--namespace) |
| `virtual_server.http.http_client_profile.tenant` | [virtual_server.http.http_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--tenant) |
| `virtual_server.http.http_client_profile.uid` | [virtual_server.http.http_client_profile.uid](data-sources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--uid) |
| `virtual_server.http.http_server_profile` | [virtual_server.http.http_server_profile](data-sources--application_profiles--properties--virtual_server--http--http_server_profile.md#section) |
| `virtual_server.http.http_server_profile.kind` | [virtual_server.http.http_server_profile.kind](data-sources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--kind) |
| `virtual_server.http.http_server_profile.name` | [virtual_server.http.http_server_profile.name](data-sources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--name) |
| `virtual_server.http.http_server_profile.namespace` | [virtual_server.http.http_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--namespace) |
| `virtual_server.http.http_server_profile.tenant` | [virtual_server.http.http_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--tenant) |
| `virtual_server.http.http_server_profile.uid` | [virtual_server.http.http_server_profile.uid](data-sources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--uid) |
| `virtual_server.http.ocsp_profile` | [virtual_server.http.ocsp_profile](data-sources--application_profiles--properties--virtual_server--http--ocsp_profile.md#section) |
| `virtual_server.http.ocsp_profile.kind` | [virtual_server.http.ocsp_profile.kind](data-sources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--kind) |
| `virtual_server.http.ocsp_profile.name` | [virtual_server.http.ocsp_profile.name](data-sources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--name) |
| `virtual_server.http.ocsp_profile.namespace` | [virtual_server.http.ocsp_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--namespace) |
| `virtual_server.http.ocsp_profile.tenant` | [virtual_server.http.ocsp_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--tenant) |
| `virtual_server.http.ocsp_profile.uid` | [virtual_server.http.ocsp_profile.uid](data-sources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--uid) |
| `virtual_server.http.server_ssl_profile` | [virtual_server.http.server_ssl_profile](data-sources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#section) |
| `virtual_server.http.server_ssl_profile.kind` | [virtual_server.http.server_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--kind) |
| `virtual_server.http.server_ssl_profile.name` | [virtual_server.http.server_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--name) |
| `virtual_server.http.server_ssl_profile.namespace` | [virtual_server.http.server_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--namespace) |
| `virtual_server.http.server_ssl_profile.tenant` | [virtual_server.http.server_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--tenant) |
| `virtual_server.http.server_ssl_profile.uid` | [virtual_server.http.server_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--uid) |
| `virtual_server.http.stream_profile` | [virtual_server.http.stream_profile](data-sources--application_profiles--properties--virtual_server--http--stream_profile.md#section) |
| `virtual_server.http.stream_profile.kind` | [virtual_server.http.stream_profile.kind](data-sources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--kind) |
| `virtual_server.http.stream_profile.name` | [virtual_server.http.stream_profile.name](data-sources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--name) |
| `virtual_server.http.stream_profile.namespace` | [virtual_server.http.stream_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--namespace) |
| `virtual_server.http.stream_profile.tenant` | [virtual_server.http.stream_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--tenant) |
| `virtual_server.http.stream_profile.uid` | [virtual_server.http.stream_profile.uid](data-sources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--uid) |
| `virtual_server.http.tcp_client_profile` | [virtual_server.http.tcp_client_profile](data-sources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#section) |
| `virtual_server.http.tcp_client_profile.kind` | [virtual_server.http.tcp_client_profile.kind](data-sources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--kind) |
| `virtual_server.http.tcp_client_profile.name` | [virtual_server.http.tcp_client_profile.name](data-sources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--name) |
| `virtual_server.http.tcp_client_profile.namespace` | [virtual_server.http.tcp_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--namespace) |
| `virtual_server.http.tcp_client_profile.tenant` | [virtual_server.http.tcp_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--tenant) |
| `virtual_server.http.tcp_client_profile.uid` | [virtual_server.http.tcp_client_profile.uid](data-sources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--uid) |
| `virtual_server.http.tcp_server_profile` | [virtual_server.http.tcp_server_profile](data-sources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#section) |
| `virtual_server.http.tcp_server_profile.kind` | [virtual_server.http.tcp_server_profile.kind](data-sources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--kind) |
| `virtual_server.http.tcp_server_profile.name` | [virtual_server.http.tcp_server_profile.name](data-sources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--name) |
| `virtual_server.http.tcp_server_profile.namespace` | [virtual_server.http.tcp_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--namespace) |
| `virtual_server.http.tcp_server_profile.tenant` | [virtual_server.http.tcp_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--tenant) |
| `virtual_server.http.tcp_server_profile.uid` | [virtual_server.http.tcp_server_profile.uid](data-sources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--uid) |
| `virtual_server.http.websocket_client_profile` | [virtual_server.http.websocket_client_profile](data-sources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#section) |
| `virtual_server.http.websocket_client_profile.kind` | [virtual_server.http.websocket_client_profile.kind](data-sources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--kind) |
| `virtual_server.http.websocket_client_profile.name` | [virtual_server.http.websocket_client_profile.name](data-sources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--name) |
| `virtual_server.http.websocket_client_profile.namespace` | [virtual_server.http.websocket_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--namespace) |
| `virtual_server.http.websocket_client_profile.tenant` | [virtual_server.http.websocket_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--tenant) |
| `virtual_server.http.websocket_client_profile.uid` | [virtual_server.http.websocket_client_profile.uid](data-sources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--uid) |
| `virtual_server.http.websocket_server_profile` | [virtual_server.http.websocket_server_profile](data-sources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#section) |
| `virtual_server.http.websocket_server_profile.kind` | [virtual_server.http.websocket_server_profile.kind](data-sources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--kind) |
| `virtual_server.http.websocket_server_profile.name` | [virtual_server.http.websocket_server_profile.name](data-sources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--name) |
| `virtual_server.http.websocket_server_profile.namespace` | [virtual_server.http.websocket_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--namespace) |
| `virtual_server.http.websocket_server_profile.tenant` | [virtual_server.http.websocket_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--tenant) |
| `virtual_server.http.websocket_server_profile.uid` | [virtual_server.http.websocket_server_profile.uid](data-sources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--uid) |
| `virtual_server.http3` | [virtual_server.http3](data-sources--application_profiles--properties--virtual_server--http3.md#section) |
| `virtual_server.http3.client_ssl_profile` | [virtual_server.http3.client_ssl_profile](data-sources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#section) |
| `virtual_server.http3.client_ssl_profile.kind` | [virtual_server.http3.client_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--kind) |
| `virtual_server.http3.client_ssl_profile.name` | [virtual_server.http3.client_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--name) |
| `virtual_server.http3.client_ssl_profile.namespace` | [virtual_server.http3.client_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--namespace) |
| `virtual_server.http3.client_ssl_profile.tenant` | [virtual_server.http3.client_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--tenant) |
| `virtual_server.http3.client_ssl_profile.uid` | [virtual_server.http3.client_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--uid) |
| `virtual_server.http3.http3_profile` | [virtual_server.http3.http3_profile](data-sources--application_profiles--properties--virtual_server--http3--http3_profile.md#section) |
| `virtual_server.http3.http3_profile.kind` | [virtual_server.http3.http3_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--kind) |
| `virtual_server.http3.http3_profile.name` | [virtual_server.http3.http3_profile.name](data-sources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--name) |
| `virtual_server.http3.http3_profile.namespace` | [virtual_server.http3.http3_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--namespace) |
| `virtual_server.http3.http3_profile.tenant` | [virtual_server.http3.http3_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--tenant) |
| `virtual_server.http3.http3_profile.uid` | [virtual_server.http3.http3_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--uid) |
| `virtual_server.http3.http_client_profile` | [virtual_server.http3.http_client_profile](data-sources--application_profiles--properties--virtual_server--http3--http_client_profile.md#section) |
| `virtual_server.http3.http_client_profile.kind` | [virtual_server.http3.http_client_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--kind) |
| `virtual_server.http3.http_client_profile.name` | [virtual_server.http3.http_client_profile.name](data-sources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--name) |
| `virtual_server.http3.http_client_profile.namespace` | [virtual_server.http3.http_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--namespace) |
| `virtual_server.http3.http_client_profile.tenant` | [virtual_server.http3.http_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--tenant) |
| `virtual_server.http3.http_client_profile.uid` | [virtual_server.http3.http_client_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--uid) |
| `virtual_server.http3.http_server_profile` | [virtual_server.http3.http_server_profile](data-sources--application_profiles--properties--virtual_server--http3--http_server_profile.md#section) |
| `virtual_server.http3.http_server_profile.kind` | [virtual_server.http3.http_server_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--kind) |
| `virtual_server.http3.http_server_profile.name` | [virtual_server.http3.http_server_profile.name](data-sources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--name) |
| `virtual_server.http3.http_server_profile.namespace` | [virtual_server.http3.http_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--namespace) |
| `virtual_server.http3.http_server_profile.tenant` | [virtual_server.http3.http_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--tenant) |
| `virtual_server.http3.http_server_profile.uid` | [virtual_server.http3.http_server_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--uid) |
| `virtual_server.http3.quic_profile` | [virtual_server.http3.quic_profile](data-sources--application_profiles--properties--virtual_server--http3--quic_profile.md#section) |
| `virtual_server.http3.quic_profile.kind` | [virtual_server.http3.quic_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--kind) |
| `virtual_server.http3.quic_profile.name` | [virtual_server.http3.quic_profile.name](data-sources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--name) |
| `virtual_server.http3.quic_profile.namespace` | [virtual_server.http3.quic_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--namespace) |
| `virtual_server.http3.quic_profile.tenant` | [virtual_server.http3.quic_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--tenant) |
| `virtual_server.http3.quic_profile.uid` | [virtual_server.http3.quic_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--uid) |
| `virtual_server.http3.server_ssl_profile` | [virtual_server.http3.server_ssl_profile](data-sources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#section) |
| `virtual_server.http3.server_ssl_profile.kind` | [virtual_server.http3.server_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--kind) |
| `virtual_server.http3.server_ssl_profile.name` | [virtual_server.http3.server_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--name) |
| `virtual_server.http3.server_ssl_profile.namespace` | [virtual_server.http3.server_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--namespace) |
| `virtual_server.http3.server_ssl_profile.tenant` | [virtual_server.http3.server_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--tenant) |
| `virtual_server.http3.server_ssl_profile.uid` | [virtual_server.http3.server_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--uid) |
| `virtual_server.http3.tcp_server_profile` | [virtual_server.http3.tcp_server_profile](data-sources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#section) |
| `virtual_server.http3.tcp_server_profile.kind` | [virtual_server.http3.tcp_server_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--kind) |
| `virtual_server.http3.tcp_server_profile.name` | [virtual_server.http3.tcp_server_profile.name](data-sources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--name) |
| `virtual_server.http3.tcp_server_profile.namespace` | [virtual_server.http3.tcp_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--namespace) |
| `virtual_server.http3.tcp_server_profile.tenant` | [virtual_server.http3.tcp_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--tenant) |
| `virtual_server.http3.tcp_server_profile.uid` | [virtual_server.http3.tcp_server_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--uid) |
| `virtual_server.http3.udp_client_profile` | [virtual_server.http3.udp_client_profile](data-sources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#section) |
| `virtual_server.http3.udp_client_profile.kind` | [virtual_server.http3.udp_client_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--kind) |
| `virtual_server.http3.udp_client_profile.name` | [virtual_server.http3.udp_client_profile.name](data-sources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--name) |
| `virtual_server.http3.udp_client_profile.namespace` | [virtual_server.http3.udp_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--namespace) |
| `virtual_server.http3.udp_client_profile.tenant` | [virtual_server.http3.udp_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--tenant) |
| `virtual_server.http3.udp_client_profile.uid` | [virtual_server.http3.udp_client_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--uid) |
| `virtual_server.http3.udp_server_profile` | [virtual_server.http3.udp_server_profile](data-sources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#section) |
| `virtual_server.http3.udp_server_profile.kind` | [virtual_server.http3.udp_server_profile.kind](data-sources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--kind) |
| `virtual_server.http3.udp_server_profile.name` | [virtual_server.http3.udp_server_profile.name](data-sources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--name) |
| `virtual_server.http3.udp_server_profile.namespace` | [virtual_server.http3.udp_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--namespace) |
| `virtual_server.http3.udp_server_profile.tenant` | [virtual_server.http3.udp_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--tenant) |
| `virtual_server.http3.udp_server_profile.uid` | [virtual_server.http3.udp_server_profile.uid](data-sources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--uid) |
| `virtual_server.https` | [virtual_server.https](data-sources--application_profiles--properties--virtual_server--https.md#section) |
| `virtual_server.https.client_ssl_profile` | [virtual_server.https.client_ssl_profile](data-sources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#section) |
| `virtual_server.https.client_ssl_profile.kind` | [virtual_server.https.client_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--kind) |
| `virtual_server.https.client_ssl_profile.name` | [virtual_server.https.client_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--name) |
| `virtual_server.https.client_ssl_profile.namespace` | [virtual_server.https.client_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--namespace) |
| `virtual_server.https.client_ssl_profile.tenant` | [virtual_server.https.client_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--tenant) |
| `virtual_server.https.client_ssl_profile.uid` | [virtual_server.https.client_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--uid) |
| `virtual_server.https.http2_client_profile` | [virtual_server.https.http2_client_profile](data-sources--application_profiles--properties--virtual_server--https--http2_client_profile.md#section) |
| `virtual_server.https.http2_client_profile.kind` | [virtual_server.https.http2_client_profile.kind](data-sources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--kind) |
| `virtual_server.https.http2_client_profile.name` | [virtual_server.https.http2_client_profile.name](data-sources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--name) |
| `virtual_server.https.http2_client_profile.namespace` | [virtual_server.https.http2_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--namespace) |
| `virtual_server.https.http2_client_profile.tenant` | [virtual_server.https.http2_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--tenant) |
| `virtual_server.https.http2_client_profile.uid` | [virtual_server.https.http2_client_profile.uid](data-sources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--uid) |
| `virtual_server.https.http2_server_profile` | [virtual_server.https.http2_server_profile](data-sources--application_profiles--properties--virtual_server--https--http2_server_profile.md#section) |
| `virtual_server.https.http2_server_profile.kind` | [virtual_server.https.http2_server_profile.kind](data-sources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--kind) |
| `virtual_server.https.http2_server_profile.name` | [virtual_server.https.http2_server_profile.name](data-sources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--name) |
| `virtual_server.https.http2_server_profile.namespace` | [virtual_server.https.http2_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--namespace) |
| `virtual_server.https.http2_server_profile.tenant` | [virtual_server.https.http2_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--tenant) |
| `virtual_server.https.http2_server_profile.uid` | [virtual_server.https.http2_server_profile.uid](data-sources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--uid) |
| `virtual_server.https.http_client_profile` | [virtual_server.https.http_client_profile](data-sources--application_profiles--properties--virtual_server--https--http_client_profile.md#section) |
| `virtual_server.https.http_client_profile.kind` | [virtual_server.https.http_client_profile.kind](data-sources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--kind) |
| `virtual_server.https.http_client_profile.name` | [virtual_server.https.http_client_profile.name](data-sources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--name) |
| `virtual_server.https.http_client_profile.namespace` | [virtual_server.https.http_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--namespace) |
| `virtual_server.https.http_client_profile.tenant` | [virtual_server.https.http_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--tenant) |
| `virtual_server.https.http_client_profile.uid` | [virtual_server.https.http_client_profile.uid](data-sources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--uid) |
| `virtual_server.https.http_server_profile` | [virtual_server.https.http_server_profile](data-sources--application_profiles--properties--virtual_server--https--http_server_profile.md#section) |
| `virtual_server.https.http_server_profile.kind` | [virtual_server.https.http_server_profile.kind](data-sources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--kind) |
| `virtual_server.https.http_server_profile.name` | [virtual_server.https.http_server_profile.name](data-sources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--name) |
| `virtual_server.https.http_server_profile.namespace` | [virtual_server.https.http_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--namespace) |
| `virtual_server.https.http_server_profile.tenant` | [virtual_server.https.http_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--tenant) |
| `virtual_server.https.http_server_profile.uid` | [virtual_server.https.http_server_profile.uid](data-sources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--uid) |
| `virtual_server.https.ocsp_profile` | [virtual_server.https.ocsp_profile](data-sources--application_profiles--properties--virtual_server--https--ocsp_profile.md#section) |
| `virtual_server.https.ocsp_profile.kind` | [virtual_server.https.ocsp_profile.kind](data-sources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--kind) |
| `virtual_server.https.ocsp_profile.name` | [virtual_server.https.ocsp_profile.name](data-sources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--name) |
| `virtual_server.https.ocsp_profile.namespace` | [virtual_server.https.ocsp_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--namespace) |
| `virtual_server.https.ocsp_profile.tenant` | [virtual_server.https.ocsp_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--tenant) |
| `virtual_server.https.ocsp_profile.uid` | [virtual_server.https.ocsp_profile.uid](data-sources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--uid) |
| `virtual_server.https.server_ssl_profile` | [virtual_server.https.server_ssl_profile](data-sources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#section) |
| `virtual_server.https.server_ssl_profile.kind` | [virtual_server.https.server_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--kind) |
| `virtual_server.https.server_ssl_profile.name` | [virtual_server.https.server_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--name) |
| `virtual_server.https.server_ssl_profile.namespace` | [virtual_server.https.server_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--namespace) |
| `virtual_server.https.server_ssl_profile.tenant` | [virtual_server.https.server_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--tenant) |
| `virtual_server.https.server_ssl_profile.uid` | [virtual_server.https.server_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--uid) |
| `virtual_server.https.stream_profile` | [virtual_server.https.stream_profile](data-sources--application_profiles--properties--virtual_server--https--stream_profile.md#section) |
| `virtual_server.https.stream_profile.kind` | [virtual_server.https.stream_profile.kind](data-sources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--kind) |
| `virtual_server.https.stream_profile.name` | [virtual_server.https.stream_profile.name](data-sources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--name) |
| `virtual_server.https.stream_profile.namespace` | [virtual_server.https.stream_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--namespace) |
| `virtual_server.https.stream_profile.tenant` | [virtual_server.https.stream_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--tenant) |
| `virtual_server.https.stream_profile.uid` | [virtual_server.https.stream_profile.uid](data-sources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--uid) |
| `virtual_server.https.tcp_client_profile` | [virtual_server.https.tcp_client_profile](data-sources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#section) |
| `virtual_server.https.tcp_client_profile.kind` | [virtual_server.https.tcp_client_profile.kind](data-sources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--kind) |
| `virtual_server.https.tcp_client_profile.name` | [virtual_server.https.tcp_client_profile.name](data-sources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--name) |
| `virtual_server.https.tcp_client_profile.namespace` | [virtual_server.https.tcp_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--namespace) |
| `virtual_server.https.tcp_client_profile.tenant` | [virtual_server.https.tcp_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--tenant) |
| `virtual_server.https.tcp_client_profile.uid` | [virtual_server.https.tcp_client_profile.uid](data-sources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--uid) |
| `virtual_server.https.tcp_server_profile` | [virtual_server.https.tcp_server_profile](data-sources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#section) |
| `virtual_server.https.tcp_server_profile.kind` | [virtual_server.https.tcp_server_profile.kind](data-sources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--kind) |
| `virtual_server.https.tcp_server_profile.name` | [virtual_server.https.tcp_server_profile.name](data-sources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--name) |
| `virtual_server.https.tcp_server_profile.namespace` | [virtual_server.https.tcp_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--namespace) |
| `virtual_server.https.tcp_server_profile.tenant` | [virtual_server.https.tcp_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--tenant) |
| `virtual_server.https.tcp_server_profile.uid` | [virtual_server.https.tcp_server_profile.uid](data-sources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--uid) |
| `virtual_server.https.websocket_client_profile` | [virtual_server.https.websocket_client_profile](data-sources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#section) |
| `virtual_server.https.websocket_client_profile.kind` | [virtual_server.https.websocket_client_profile.kind](data-sources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--kind) |
| `virtual_server.https.websocket_client_profile.name` | [virtual_server.https.websocket_client_profile.name](data-sources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--name) |
| `virtual_server.https.websocket_client_profile.namespace` | [virtual_server.https.websocket_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--namespace) |
| `virtual_server.https.websocket_client_profile.tenant` | [virtual_server.https.websocket_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--tenant) |
| `virtual_server.https.websocket_client_profile.uid` | [virtual_server.https.websocket_client_profile.uid](data-sources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--uid) |
| `virtual_server.https.websocket_server_profile` | [virtual_server.https.websocket_server_profile](data-sources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#section) |
| `virtual_server.https.websocket_server_profile.kind` | [virtual_server.https.websocket_server_profile.kind](data-sources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--kind) |
| `virtual_server.https.websocket_server_profile.name` | [virtual_server.https.websocket_server_profile.name](data-sources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--name) |
| `virtual_server.https.websocket_server_profile.namespace` | [virtual_server.https.websocket_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--namespace) |
| `virtual_server.https.websocket_server_profile.tenant` | [virtual_server.https.websocket_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--tenant) |
| `virtual_server.https.websocket_server_profile.uid` | [virtual_server.https.websocket_server_profile.uid](data-sources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--uid) |
| `virtual_server.immediate_action_on_service_down` | [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--properties--virtual_server--immediate_action_on_service_down.md#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](data-sources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_drop.md#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](data-sources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_none.md#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](data-sources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_reset.md#section) |
| `virtual_server.last_hop_pool` | [virtual_server.last_hop_pool](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md#section) |
| `virtual_server.last_hop_pool.kind` | [virtual_server.last_hop_pool.kind](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--kind) |
| `virtual_server.last_hop_pool.name` | [virtual_server.last_hop_pool.name](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--name) |
| `virtual_server.last_hop_pool.namespace` | [virtual_server.last_hop_pool.namespace](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--namespace) |
| `virtual_server.last_hop_pool.tenant` | [virtual_server.last_hop_pool.tenant](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--tenant) |
| `virtual_server.last_hop_pool.uid` | [virtual_server.last_hop_pool.uid](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--uid) |
| `virtual_server.nat64` | [virtual_server.nat64](data-sources--application_profiles--properties--virtual_server--nat64.md#section) |
| `virtual_server.nat64.nat64_disable` | [virtual_server.nat64.nat64_disable](data-sources--application_profiles--properties--virtual_server--nat64--nat64_disable.md#section) |
| `virtual_server.nat64.nat64_enable` | [virtual_server.nat64.nat64_enable](data-sources--application_profiles--properties--virtual_server--nat64--nat64_enable.md#section) |
| `virtual_server.port_translation` | [virtual_server.port_translation](data-sources--application_profiles--properties--virtual_server--port_translation.md#section) |
| `virtual_server.port_translation.port_translation_disable` | [virtual_server.port_translation.port_translation_disable](data-sources--application_profiles--properties--virtual_server--port_translation--port_translation_disable.md#section) |
| `virtual_server.port_translation.port_translation_enable` | [virtual_server.port_translation.port_translation_enable](data-sources--application_profiles--properties--virtual_server--port_translation--port_translation_enable.md#section) |
| `virtual_server.request_logging_profile` | [virtual_server.request_logging_profile](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md#section) |
| `virtual_server.request_logging_profile.kind` | [virtual_server.request_logging_profile.kind](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--kind) |
| `virtual_server.request_logging_profile.name` | [virtual_server.request_logging_profile.name](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--name) |
| `virtual_server.request_logging_profile.namespace` | [virtual_server.request_logging_profile.namespace](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--namespace) |
| `virtual_server.request_logging_profile.tenant` | [virtual_server.request_logging_profile.tenant](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--tenant) |
| `virtual_server.request_logging_profile.uid` | [virtual_server.request_logging_profile.uid](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--uid) |
| `virtual_server.source_port` | [virtual_server.source_port](data-sources--application_profiles--properties--virtual_server--source_port.md#section) |
| `virtual_server.source_port.source_port_change` | [virtual_server.source_port.source_port_change](data-sources--application_profiles--properties--virtual_server--source_port--source_port_change.md#section) |
| `virtual_server.source_port.source_port_preserve` | [virtual_server.source_port.source_port_preserve](data-sources--application_profiles--properties--virtual_server--source_port--source_port_preserve.md#section) |
| `virtual_server.source_port.source_port_preserve_strict` | [virtual_server.source_port.source_port_preserve_strict](data-sources--application_profiles--properties--virtual_server--source_port--source_port_preserve_strict.md#section) |
| `virtual_server.statistics_profile` | [virtual_server.statistics_profile](data-sources--application_profiles--properties--virtual_server--statistics_profile.md#section) |
| `virtual_server.statistics_profile.kind` | [virtual_server.statistics_profile.kind](data-sources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--kind) |
| `virtual_server.statistics_profile.name` | [virtual_server.statistics_profile.name](data-sources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--name) |
| `virtual_server.statistics_profile.namespace` | [virtual_server.statistics_profile.namespace](data-sources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--namespace) |
| `virtual_server.statistics_profile.tenant` | [virtual_server.statistics_profile.tenant](data-sources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--tenant) |
| `virtual_server.statistics_profile.uid` | [virtual_server.statistics_profile.uid](data-sources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--uid) |
| `virtual_server.tcp` | [virtual_server.tcp](data-sources--application_profiles--properties--virtual_server--tcp.md#section) |
| `virtual_server.tcp.client_ssl_profile` | [virtual_server.tcp.client_ssl_profile](data-sources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#section) |
| `virtual_server.tcp.client_ssl_profile.kind` | [virtual_server.tcp.client_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--kind) |
| `virtual_server.tcp.client_ssl_profile.name` | [virtual_server.tcp.client_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--name) |
| `virtual_server.tcp.client_ssl_profile.namespace` | [virtual_server.tcp.client_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--namespace) |
| `virtual_server.tcp.client_ssl_profile.tenant` | [virtual_server.tcp.client_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--tenant) |
| `virtual_server.tcp.client_ssl_profile.uid` | [virtual_server.tcp.client_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--uid) |
| `virtual_server.tcp.ocsp_profile` | [virtual_server.tcp.ocsp_profile](data-sources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#section) |
| `virtual_server.tcp.ocsp_profile.kind` | [virtual_server.tcp.ocsp_profile.kind](data-sources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--kind) |
| `virtual_server.tcp.ocsp_profile.name` | [virtual_server.tcp.ocsp_profile.name](data-sources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--name) |
| `virtual_server.tcp.ocsp_profile.namespace` | [virtual_server.tcp.ocsp_profile.namespace](data-sources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--namespace) |
| `virtual_server.tcp.ocsp_profile.tenant` | [virtual_server.tcp.ocsp_profile.tenant](data-sources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--tenant) |
| `virtual_server.tcp.ocsp_profile.uid` | [virtual_server.tcp.ocsp_profile.uid](data-sources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--uid) |
| `virtual_server.tcp.server_ssl_profile` | [virtual_server.tcp.server_ssl_profile](data-sources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#section) |
| `virtual_server.tcp.server_ssl_profile.kind` | [virtual_server.tcp.server_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--kind) |
| `virtual_server.tcp.server_ssl_profile.name` | [virtual_server.tcp.server_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--name) |
| `virtual_server.tcp.server_ssl_profile.namespace` | [virtual_server.tcp.server_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--namespace) |
| `virtual_server.tcp.server_ssl_profile.tenant` | [virtual_server.tcp.server_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--tenant) |
| `virtual_server.tcp.server_ssl_profile.uid` | [virtual_server.tcp.server_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--uid) |
| `virtual_server.tcp.tcp_client_profile` | [virtual_server.tcp.tcp_client_profile](data-sources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#section) |
| `virtual_server.tcp.tcp_client_profile.kind` | [virtual_server.tcp.tcp_client_profile.kind](data-sources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--kind) |
| `virtual_server.tcp.tcp_client_profile.name` | [virtual_server.tcp.tcp_client_profile.name](data-sources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--name) |
| `virtual_server.tcp.tcp_client_profile.namespace` | [virtual_server.tcp.tcp_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--namespace) |
| `virtual_server.tcp.tcp_client_profile.tenant` | [virtual_server.tcp.tcp_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--tenant) |
| `virtual_server.tcp.tcp_client_profile.uid` | [virtual_server.tcp.tcp_client_profile.uid](data-sources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--uid) |
| `virtual_server.tcp.tcp_server_profile` | [virtual_server.tcp.tcp_server_profile](data-sources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#section) |
| `virtual_server.tcp.tcp_server_profile.kind` | [virtual_server.tcp.tcp_server_profile.kind](data-sources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--kind) |
| `virtual_server.tcp.tcp_server_profile.name` | [virtual_server.tcp.tcp_server_profile.name](data-sources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--name) |
| `virtual_server.tcp.tcp_server_profile.namespace` | [virtual_server.tcp.tcp_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--namespace) |
| `virtual_server.tcp.tcp_server_profile.tenant` | [virtual_server.tcp.tcp_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--tenant) |
| `virtual_server.tcp.tcp_server_profile.uid` | [virtual_server.tcp.tcp_server_profile.uid](data-sources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--uid) |
| `virtual_server.udp` | [virtual_server.udp](data-sources--application_profiles--properties--virtual_server--udp.md#section) |
| `virtual_server.udp.client_ssl_profile` | [virtual_server.udp.client_ssl_profile](data-sources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#section) |
| `virtual_server.udp.client_ssl_profile.kind` | [virtual_server.udp.client_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--kind) |
| `virtual_server.udp.client_ssl_profile.name` | [virtual_server.udp.client_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--name) |
| `virtual_server.udp.client_ssl_profile.namespace` | [virtual_server.udp.client_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--namespace) |
| `virtual_server.udp.client_ssl_profile.tenant` | [virtual_server.udp.client_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--tenant) |
| `virtual_server.udp.client_ssl_profile.uid` | [virtual_server.udp.client_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--uid) |
| `virtual_server.udp.server_ssl_profile` | [virtual_server.udp.server_ssl_profile](data-sources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#section) |
| `virtual_server.udp.server_ssl_profile.kind` | [virtual_server.udp.server_ssl_profile.kind](data-sources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--kind) |
| `virtual_server.udp.server_ssl_profile.name` | [virtual_server.udp.server_ssl_profile.name](data-sources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--name) |
| `virtual_server.udp.server_ssl_profile.namespace` | [virtual_server.udp.server_ssl_profile.namespace](data-sources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--namespace) |
| `virtual_server.udp.server_ssl_profile.tenant` | [virtual_server.udp.server_ssl_profile.tenant](data-sources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--tenant) |
| `virtual_server.udp.server_ssl_profile.uid` | [virtual_server.udp.server_ssl_profile.uid](data-sources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--uid) |
| `virtual_server.udp.udp_client_profile` | [virtual_server.udp.udp_client_profile](data-sources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#section) |
| `virtual_server.udp.udp_client_profile.kind` | [virtual_server.udp.udp_client_profile.kind](data-sources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--kind) |
| `virtual_server.udp.udp_client_profile.name` | [virtual_server.udp.udp_client_profile.name](data-sources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--name) |
| `virtual_server.udp.udp_client_profile.namespace` | [virtual_server.udp.udp_client_profile.namespace](data-sources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--namespace) |
| `virtual_server.udp.udp_client_profile.tenant` | [virtual_server.udp.udp_client_profile.tenant](data-sources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--tenant) |
| `virtual_server.udp.udp_client_profile.uid` | [virtual_server.udp.udp_client_profile.uid](data-sources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--uid) |
| `virtual_server.udp.udp_server_profile` | [virtual_server.udp.udp_server_profile](data-sources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#section) |
| `virtual_server.udp.udp_server_profile.kind` | [virtual_server.udp.udp_server_profile.kind](data-sources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--kind) |
| `virtual_server.udp.udp_server_profile.name` | [virtual_server.udp.udp_server_profile.name](data-sources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--name) |
| `virtual_server.udp.udp_server_profile.namespace` | [virtual_server.udp.udp_server_profile.namespace](data-sources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--namespace) |
| `virtual_server.udp.udp_server_profile.tenant` | [virtual_server.udp.udp_server_profile.tenant](data-sources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--tenant) |
| `virtual_server.udp.udp_server_profile.uid` | [virtual_server.udp.udp_server_profile.uid](data-sources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--uid) |
| `virtual_server.virtual_server_state` | [virtual_server.virtual_server_state](data-sources--application_profiles--properties--virtual_server--virtual_server_state.md#section) |
| `virtual_server.virtual_server_state.state_disabled` | [virtual_server.virtual_server_state.state_disabled](data-sources--application_profiles--properties--virtual_server--virtual_server_state--state_disabled.md#section) |
| `virtual_server.virtual_server_state.state_enabled` | [virtual_server.virtual_server_state.state_enabled](data-sources--application_profiles--properties--virtual_server--virtual_server_state--state_enabled.md#section) |
| `virtual_server.vs_score` | [virtual_server.vs_score](data-sources--application_profiles--properties--virtual_server.md#schema-virtual_server--vs_score) |

## Next pages

- [advanced_tcp_profile](data-sources--application_profiles--properties--advanced_tcp_profile.md)
- [ddos_profile](data-sources--application_profiles--properties--ddos_profile.md)
- [irules](data-sources--application_profiles--properties--irules.md)
- [virtual_server](data-sources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
