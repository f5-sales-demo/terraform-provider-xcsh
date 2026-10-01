---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 95259, "body_sha256": "sha256:b1b1d339614df4a7275c53f6746484450fbe71495ea032f35a5402ca248122b4", "canonical_id": "xcsh-docs:resources:application_profiles:reference", "child_ids": ["xcsh-docs:resources:application_profiles:properties:advanced_tcp_profile", "xcsh-docs:resources:application_profiles:properties:ddos_profile", "xcsh-docs:resources:application_profiles:properties:irules", "xcsh-docs:resources:application_profiles:properties:timeouts", "xcsh-docs:resources:application_profiles:properties:virtual_server"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:reference", "parent_id": "xcsh-docs:resources:application_profiles:fundamentals", "path": "docs/guides/resources--application_profiles--reference.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md)
- Property reference

## Direct properties

- [advanced_tcp_profile](resources--application_profiles--properties--advanced_tcp_profile.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [ddos_profile](resources--application_profiles--properties--ddos_profile.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--application_profiles--properties--irules.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the Application Profiles. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

Namespace where the Application Profiles is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--application_profiles--properties--timeouts.md): complete subsection reference.

- [virtual_server](resources--application_profiles--properties--virtual_server.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_tcp_profile` | [advanced_tcp_profile](resources--application_profiles--properties--advanced_tcp_profile.md#section) |
| `advanced_tcp_profile.disable_tcp_advanced_profile` | [advanced_tcp_profile.disable_tcp_advanced_profile](resources--application_profiles--properties--advanced_tcp_profile--disable_tcp_advanced_profile.md#section) |
| `advanced_tcp_profile.enable_tcp_advanced_profile` | [advanced_tcp_profile.enable_tcp_advanced_profile](resources--application_profiles--properties--advanced_tcp_profile--enable_tcp_advanced_profile.md#section) |
| `annotations` | [annotations](resources--application_profiles--reference.md#schema-annotations) |
| `ddos_profile` | [ddos_profile](resources--application_profiles--properties--ddos_profile.md#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--application_profiles--properties--ddos_profile--disable_ddos_mitigation.md#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--application_profiles--properties--ddos_profile--enable_ddos_mitigation.md#section) |
| `description` | [description](resources--application_profiles--reference.md#schema-description) |
| `disable` | [disable](resources--application_profiles--reference.md#schema-disable) |
| `id` | [id](resources--application_profiles--reference.md#schema-id) |
| `irules` | [irules](resources--application_profiles--properties--irules.md#section) |
| `irules.kind` | [irules.kind](resources--application_profiles--properties--irules.md#schema-irules--kind) |
| `irules.name` | [irules.name](resources--application_profiles--properties--irules.md#schema-irules--name) |
| `irules.namespace` | [irules.namespace](resources--application_profiles--properties--irules.md#schema-irules--namespace) |
| `irules.tenant` | [irules.tenant](resources--application_profiles--properties--irules.md#schema-irules--tenant) |
| `irules.uid` | [irules.uid](resources--application_profiles--properties--irules.md#schema-irules--uid) |
| `labels` | [labels](resources--application_profiles--reference.md#schema-labels) |
| `name` | [name](resources--application_profiles--reference.md#schema-name) |
| `namespace` | [namespace](resources--application_profiles--reference.md#schema-namespace) |
| `timeouts` | [timeouts](resources--application_profiles--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--application_profiles--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--application_profiles--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--application_profiles--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--application_profiles--properties--timeouts.md#schema-timeouts--update) |
| `virtual_server` | [virtual_server](resources--application_profiles--properties--virtual_server.md#section) |
| `virtual_server.access_profile` | [virtual_server.access_profile](resources--application_profiles--properties--virtual_server--access_profile.md#section) |
| `virtual_server.access_profile.kind` | [virtual_server.access_profile.kind](resources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--kind) |
| `virtual_server.access_profile.name` | [virtual_server.access_profile.name](resources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--name) |
| `virtual_server.access_profile.namespace` | [virtual_server.access_profile.namespace](resources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--namespace) |
| `virtual_server.access_profile.tenant` | [virtual_server.access_profile.tenant](resources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--tenant) |
| `virtual_server.access_profile.uid` | [virtual_server.access_profile.uid](resources--application_profiles--properties--virtual_server--access_profile.md#schema-virtual_server--access_profile--uid) |
| `virtual_server.address_translation` | [virtual_server.address_translation](resources--application_profiles--properties--virtual_server--address_translation.md#section) |
| `virtual_server.address_translation.address_translation_disable` | [virtual_server.address_translation.address_translation_disable](resources--application_profiles--properties--virtual_server--address_translation--address_translation_disable.md#section) |
| `virtual_server.address_translation.address_translation_enable` | [virtual_server.address_translation.address_translation_enable](resources--application_profiles--properties--virtual_server--address_translation--address_translation_enable.md#section) |
| `virtual_server.auto_last_hop` | [virtual_server.auto_last_hop](resources--application_profiles--properties--virtual_server--auto_last_hop.md#section) |
| `virtual_server.auto_last_hop.auto_last_hop_default` | [virtual_server.auto_last_hop.auto_last_hop_default](resources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_default.md#section) |
| `virtual_server.auto_last_hop.auto_last_hop_disable` | [virtual_server.auto_last_hop.auto_last_hop_disable](resources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_disable.md#section) |
| `virtual_server.auto_last_hop.auto_last_hop_enable` | [virtual_server.auto_last_hop.auto_last_hop_enable](resources--application_profiles--properties--virtual_server--auto_last_hop--auto_last_hop_enable.md#section) |
| `virtual_server.clone_pool_client` | [virtual_server.clone_pool_client](resources--application_profiles--properties--virtual_server--clone_pool_client.md#section) |
| `virtual_server.clone_pool_client.kind` | [virtual_server.clone_pool_client.kind](resources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--kind) |
| `virtual_server.clone_pool_client.name` | [virtual_server.clone_pool_client.name](resources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--name) |
| `virtual_server.clone_pool_client.namespace` | [virtual_server.clone_pool_client.namespace](resources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--namespace) |
| `virtual_server.clone_pool_client.tenant` | [virtual_server.clone_pool_client.tenant](resources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--tenant) |
| `virtual_server.clone_pool_client.uid` | [virtual_server.clone_pool_client.uid](resources--application_profiles--properties--virtual_server--clone_pool_client.md#schema-virtual_server--clone_pool_client--uid) |
| `virtual_server.clone_pool_server` | [virtual_server.clone_pool_server](resources--application_profiles--properties--virtual_server--clone_pool_server.md#section) |
| `virtual_server.clone_pool_server.kind` | [virtual_server.clone_pool_server.kind](resources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--kind) |
| `virtual_server.clone_pool_server.name` | [virtual_server.clone_pool_server.name](resources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--name) |
| `virtual_server.clone_pool_server.namespace` | [virtual_server.clone_pool_server.namespace](resources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--namespace) |
| `virtual_server.clone_pool_server.tenant` | [virtual_server.clone_pool_server.tenant](resources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--tenant) |
| `virtual_server.clone_pool_server.uid` | [virtual_server.clone_pool_server.uid](resources--application_profiles--properties--virtual_server--clone_pool_server.md#schema-virtual_server--clone_pool_server--uid) |
| `virtual_server.connection_limit` | [virtual_server.connection_limit](resources--application_profiles--properties--virtual_server.md#schema-virtual_server--connection_limit) |
| `virtual_server.connection_rate_limit` | [virtual_server.connection_rate_limit](resources--application_profiles--properties--virtual_server.md#schema-virtual_server--connection_rate_limit) |
| `virtual_server.connection_rate_limit_mode` | [virtual_server.connection_rate_limit_mode](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md#section) |
| `virtual_server.connection_rate_limit_mode.per_destination_address` | [virtual_server.connection_rate_limit_mode.per_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_address` | [virtual_server.connection_rate_limit_mode.per_source_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_address.source_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_address.md#schema-virtual_server--connection_rate_limit_mode--per_source_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_source_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server` | [virtual_server.connection_rate_limit_mode.per_virtual_server](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask](resources--application_profiles--properties--virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address.md#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address--source_mask) |
| `virtual_server.default_persistence_profile` | [virtual_server.default_persistence_profile](resources--application_profiles--properties--virtual_server--default_persistence_profile.md#section) |
| `virtual_server.default_persistence_profile.kind` | [virtual_server.default_persistence_profile.kind](resources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--kind) |
| `virtual_server.default_persistence_profile.name` | [virtual_server.default_persistence_profile.name](resources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--name) |
| `virtual_server.default_persistence_profile.namespace` | [virtual_server.default_persistence_profile.namespace](resources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--namespace) |
| `virtual_server.default_persistence_profile.tenant` | [virtual_server.default_persistence_profile.tenant](resources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--tenant) |
| `virtual_server.default_persistence_profile.uid` | [virtual_server.default_persistence_profile.uid](resources--application_profiles--properties--virtual_server--default_persistence_profile.md#schema-virtual_server--default_persistence_profile--uid) |
| `virtual_server.default_pool` | [virtual_server.default_pool](resources--application_profiles--properties--virtual_server--default_pool.md#section) |
| `virtual_server.default_pool.kind` | [virtual_server.default_pool.kind](resources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--kind) |
| `virtual_server.default_pool.name` | [virtual_server.default_pool.name](resources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--name) |
| `virtual_server.default_pool.namespace` | [virtual_server.default_pool.namespace](resources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--namespace) |
| `virtual_server.default_pool.tenant` | [virtual_server.default_pool.tenant](resources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--tenant) |
| `virtual_server.default_pool.uid` | [virtual_server.default_pool.uid](resources--application_profiles--properties--virtual_server--default_pool.md#schema-virtual_server--default_pool--uid) |
| `virtual_server.fallback_persistence_profile` | [virtual_server.fallback_persistence_profile](resources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#section) |
| `virtual_server.fallback_persistence_profile.kind` | [virtual_server.fallback_persistence_profile.kind](resources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--kind) |
| `virtual_server.fallback_persistence_profile.name` | [virtual_server.fallback_persistence_profile.name](resources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--name) |
| `virtual_server.fallback_persistence_profile.namespace` | [virtual_server.fallback_persistence_profile.namespace](resources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--namespace) |
| `virtual_server.fallback_persistence_profile.tenant` | [virtual_server.fallback_persistence_profile.tenant](resources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--tenant) |
| `virtual_server.fallback_persistence_profile.uid` | [virtual_server.fallback_persistence_profile.uid](resources--application_profiles--properties--virtual_server--fallback_persistence_profile.md#schema-virtual_server--fallback_persistence_profile--uid) |
| `virtual_server.fix_profile` | [virtual_server.fix_profile](resources--application_profiles--properties--virtual_server--fix_profile.md#section) |
| `virtual_server.fix_profile.kind` | [virtual_server.fix_profile.kind](resources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--kind) |
| `virtual_server.fix_profile.name` | [virtual_server.fix_profile.name](resources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--name) |
| `virtual_server.fix_profile.namespace` | [virtual_server.fix_profile.namespace](resources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--namespace) |
| `virtual_server.fix_profile.tenant` | [virtual_server.fix_profile.tenant](resources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--tenant) |
| `virtual_server.fix_profile.uid` | [virtual_server.fix_profile.uid](resources--application_profiles--properties--virtual_server--fix_profile.md#schema-virtual_server--fix_profile--uid) |
| `virtual_server.http` | [virtual_server.http](resources--application_profiles--properties--virtual_server--http.md#section) |
| `virtual_server.http.client_ssl_profile` | [virtual_server.http.client_ssl_profile](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#section) |
| `virtual_server.http.client_ssl_profile.kind` | [virtual_server.http.client_ssl_profile.kind](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--kind) |
| `virtual_server.http.client_ssl_profile.name` | [virtual_server.http.client_ssl_profile.name](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--name) |
| `virtual_server.http.client_ssl_profile.namespace` | [virtual_server.http.client_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--namespace) |
| `virtual_server.http.client_ssl_profile.tenant` | [virtual_server.http.client_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--tenant) |
| `virtual_server.http.client_ssl_profile.uid` | [virtual_server.http.client_ssl_profile.uid](resources--application_profiles--properties--virtual_server--http--client_ssl_profile.md#schema-virtual_server--http--client_ssl_profile--uid) |
| `virtual_server.http.http2_client_profile` | [virtual_server.http.http2_client_profile](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md#section) |
| `virtual_server.http.http2_client_profile.kind` | [virtual_server.http.http2_client_profile.kind](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--kind) |
| `virtual_server.http.http2_client_profile.name` | [virtual_server.http.http2_client_profile.name](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--name) |
| `virtual_server.http.http2_client_profile.namespace` | [virtual_server.http.http2_client_profile.namespace](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--namespace) |
| `virtual_server.http.http2_client_profile.tenant` | [virtual_server.http.http2_client_profile.tenant](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--tenant) |
| `virtual_server.http.http2_client_profile.uid` | [virtual_server.http.http2_client_profile.uid](resources--application_profiles--properties--virtual_server--http--http2_client_profile.md#schema-virtual_server--http--http2_client_profile--uid) |
| `virtual_server.http.http2_server_profile` | [virtual_server.http.http2_server_profile](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md#section) |
| `virtual_server.http.http2_server_profile.kind` | [virtual_server.http.http2_server_profile.kind](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--kind) |
| `virtual_server.http.http2_server_profile.name` | [virtual_server.http.http2_server_profile.name](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--name) |
| `virtual_server.http.http2_server_profile.namespace` | [virtual_server.http.http2_server_profile.namespace](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--namespace) |
| `virtual_server.http.http2_server_profile.tenant` | [virtual_server.http.http2_server_profile.tenant](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--tenant) |
| `virtual_server.http.http2_server_profile.uid` | [virtual_server.http.http2_server_profile.uid](resources--application_profiles--properties--virtual_server--http--http2_server_profile.md#schema-virtual_server--http--http2_server_profile--uid) |
| `virtual_server.http.http_client_profile` | [virtual_server.http.http_client_profile](resources--application_profiles--properties--virtual_server--http--http_client_profile.md#section) |
| `virtual_server.http.http_client_profile.kind` | [virtual_server.http.http_client_profile.kind](resources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--kind) |
| `virtual_server.http.http_client_profile.name` | [virtual_server.http.http_client_profile.name](resources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--name) |
| `virtual_server.http.http_client_profile.namespace` | [virtual_server.http.http_client_profile.namespace](resources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--namespace) |
| `virtual_server.http.http_client_profile.tenant` | [virtual_server.http.http_client_profile.tenant](resources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--tenant) |
| `virtual_server.http.http_client_profile.uid` | [virtual_server.http.http_client_profile.uid](resources--application_profiles--properties--virtual_server--http--http_client_profile.md#schema-virtual_server--http--http_client_profile--uid) |
| `virtual_server.http.http_server_profile` | [virtual_server.http.http_server_profile](resources--application_profiles--properties--virtual_server--http--http_server_profile.md#section) |
| `virtual_server.http.http_server_profile.kind` | [virtual_server.http.http_server_profile.kind](resources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--kind) |
| `virtual_server.http.http_server_profile.name` | [virtual_server.http.http_server_profile.name](resources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--name) |
| `virtual_server.http.http_server_profile.namespace` | [virtual_server.http.http_server_profile.namespace](resources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--namespace) |
| `virtual_server.http.http_server_profile.tenant` | [virtual_server.http.http_server_profile.tenant](resources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--tenant) |
| `virtual_server.http.http_server_profile.uid` | [virtual_server.http.http_server_profile.uid](resources--application_profiles--properties--virtual_server--http--http_server_profile.md#schema-virtual_server--http--http_server_profile--uid) |
| `virtual_server.http.ocsp_profile` | [virtual_server.http.ocsp_profile](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md#section) |
| `virtual_server.http.ocsp_profile.kind` | [virtual_server.http.ocsp_profile.kind](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--kind) |
| `virtual_server.http.ocsp_profile.name` | [virtual_server.http.ocsp_profile.name](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--name) |
| `virtual_server.http.ocsp_profile.namespace` | [virtual_server.http.ocsp_profile.namespace](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--namespace) |
| `virtual_server.http.ocsp_profile.tenant` | [virtual_server.http.ocsp_profile.tenant](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--tenant) |
| `virtual_server.http.ocsp_profile.uid` | [virtual_server.http.ocsp_profile.uid](resources--application_profiles--properties--virtual_server--http--ocsp_profile.md#schema-virtual_server--http--ocsp_profile--uid) |
| `virtual_server.http.server_ssl_profile` | [virtual_server.http.server_ssl_profile](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#section) |
| `virtual_server.http.server_ssl_profile.kind` | [virtual_server.http.server_ssl_profile.kind](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--kind) |
| `virtual_server.http.server_ssl_profile.name` | [virtual_server.http.server_ssl_profile.name](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--name) |
| `virtual_server.http.server_ssl_profile.namespace` | [virtual_server.http.server_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--namespace) |
| `virtual_server.http.server_ssl_profile.tenant` | [virtual_server.http.server_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--tenant) |
| `virtual_server.http.server_ssl_profile.uid` | [virtual_server.http.server_ssl_profile.uid](resources--application_profiles--properties--virtual_server--http--server_ssl_profile.md#schema-virtual_server--http--server_ssl_profile--uid) |
| `virtual_server.http.stream_profile` | [virtual_server.http.stream_profile](resources--application_profiles--properties--virtual_server--http--stream_profile.md#section) |
| `virtual_server.http.stream_profile.kind` | [virtual_server.http.stream_profile.kind](resources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--kind) |
| `virtual_server.http.stream_profile.name` | [virtual_server.http.stream_profile.name](resources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--name) |
| `virtual_server.http.stream_profile.namespace` | [virtual_server.http.stream_profile.namespace](resources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--namespace) |
| `virtual_server.http.stream_profile.tenant` | [virtual_server.http.stream_profile.tenant](resources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--tenant) |
| `virtual_server.http.stream_profile.uid` | [virtual_server.http.stream_profile.uid](resources--application_profiles--properties--virtual_server--http--stream_profile.md#schema-virtual_server--http--stream_profile--uid) |
| `virtual_server.http.tcp_client_profile` | [virtual_server.http.tcp_client_profile](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#section) |
| `virtual_server.http.tcp_client_profile.kind` | [virtual_server.http.tcp_client_profile.kind](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--kind) |
| `virtual_server.http.tcp_client_profile.name` | [virtual_server.http.tcp_client_profile.name](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--name) |
| `virtual_server.http.tcp_client_profile.namespace` | [virtual_server.http.tcp_client_profile.namespace](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--namespace) |
| `virtual_server.http.tcp_client_profile.tenant` | [virtual_server.http.tcp_client_profile.tenant](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--tenant) |
| `virtual_server.http.tcp_client_profile.uid` | [virtual_server.http.tcp_client_profile.uid](resources--application_profiles--properties--virtual_server--http--tcp_client_profile.md#schema-virtual_server--http--tcp_client_profile--uid) |
| `virtual_server.http.tcp_server_profile` | [virtual_server.http.tcp_server_profile](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#section) |
| `virtual_server.http.tcp_server_profile.kind` | [virtual_server.http.tcp_server_profile.kind](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--kind) |
| `virtual_server.http.tcp_server_profile.name` | [virtual_server.http.tcp_server_profile.name](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--name) |
| `virtual_server.http.tcp_server_profile.namespace` | [virtual_server.http.tcp_server_profile.namespace](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--namespace) |
| `virtual_server.http.tcp_server_profile.tenant` | [virtual_server.http.tcp_server_profile.tenant](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--tenant) |
| `virtual_server.http.tcp_server_profile.uid` | [virtual_server.http.tcp_server_profile.uid](resources--application_profiles--properties--virtual_server--http--tcp_server_profile.md#schema-virtual_server--http--tcp_server_profile--uid) |
| `virtual_server.http.websocket_client_profile` | [virtual_server.http.websocket_client_profile](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#section) |
| `virtual_server.http.websocket_client_profile.kind` | [virtual_server.http.websocket_client_profile.kind](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--kind) |
| `virtual_server.http.websocket_client_profile.name` | [virtual_server.http.websocket_client_profile.name](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--name) |
| `virtual_server.http.websocket_client_profile.namespace` | [virtual_server.http.websocket_client_profile.namespace](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--namespace) |
| `virtual_server.http.websocket_client_profile.tenant` | [virtual_server.http.websocket_client_profile.tenant](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--tenant) |
| `virtual_server.http.websocket_client_profile.uid` | [virtual_server.http.websocket_client_profile.uid](resources--application_profiles--properties--virtual_server--http--websocket_client_profile.md#schema-virtual_server--http--websocket_client_profile--uid) |
| `virtual_server.http.websocket_server_profile` | [virtual_server.http.websocket_server_profile](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#section) |
| `virtual_server.http.websocket_server_profile.kind` | [virtual_server.http.websocket_server_profile.kind](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--kind) |
| `virtual_server.http.websocket_server_profile.name` | [virtual_server.http.websocket_server_profile.name](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--name) |
| `virtual_server.http.websocket_server_profile.namespace` | [virtual_server.http.websocket_server_profile.namespace](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--namespace) |
| `virtual_server.http.websocket_server_profile.tenant` | [virtual_server.http.websocket_server_profile.tenant](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--tenant) |
| `virtual_server.http.websocket_server_profile.uid` | [virtual_server.http.websocket_server_profile.uid](resources--application_profiles--properties--virtual_server--http--websocket_server_profile.md#schema-virtual_server--http--websocket_server_profile--uid) |
| `virtual_server.http3` | [virtual_server.http3](resources--application_profiles--properties--virtual_server--http3.md#section) |
| `virtual_server.http3.client_ssl_profile` | [virtual_server.http3.client_ssl_profile](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#section) |
| `virtual_server.http3.client_ssl_profile.kind` | [virtual_server.http3.client_ssl_profile.kind](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--kind) |
| `virtual_server.http3.client_ssl_profile.name` | [virtual_server.http3.client_ssl_profile.name](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--name) |
| `virtual_server.http3.client_ssl_profile.namespace` | [virtual_server.http3.client_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--namespace) |
| `virtual_server.http3.client_ssl_profile.tenant` | [virtual_server.http3.client_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--tenant) |
| `virtual_server.http3.client_ssl_profile.uid` | [virtual_server.http3.client_ssl_profile.uid](resources--application_profiles--properties--virtual_server--http3--client_ssl_profile.md#schema-virtual_server--http3--client_ssl_profile--uid) |
| `virtual_server.http3.http3_profile` | [virtual_server.http3.http3_profile](resources--application_profiles--properties--virtual_server--http3--http3_profile.md#section) |
| `virtual_server.http3.http3_profile.kind` | [virtual_server.http3.http3_profile.kind](resources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--kind) |
| `virtual_server.http3.http3_profile.name` | [virtual_server.http3.http3_profile.name](resources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--name) |
| `virtual_server.http3.http3_profile.namespace` | [virtual_server.http3.http3_profile.namespace](resources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--namespace) |
| `virtual_server.http3.http3_profile.tenant` | [virtual_server.http3.http3_profile.tenant](resources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--tenant) |
| `virtual_server.http3.http3_profile.uid` | [virtual_server.http3.http3_profile.uid](resources--application_profiles--properties--virtual_server--http3--http3_profile.md#schema-virtual_server--http3--http3_profile--uid) |
| `virtual_server.http3.http_client_profile` | [virtual_server.http3.http_client_profile](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md#section) |
| `virtual_server.http3.http_client_profile.kind` | [virtual_server.http3.http_client_profile.kind](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--kind) |
| `virtual_server.http3.http_client_profile.name` | [virtual_server.http3.http_client_profile.name](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--name) |
| `virtual_server.http3.http_client_profile.namespace` | [virtual_server.http3.http_client_profile.namespace](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--namespace) |
| `virtual_server.http3.http_client_profile.tenant` | [virtual_server.http3.http_client_profile.tenant](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--tenant) |
| `virtual_server.http3.http_client_profile.uid` | [virtual_server.http3.http_client_profile.uid](resources--application_profiles--properties--virtual_server--http3--http_client_profile.md#schema-virtual_server--http3--http_client_profile--uid) |
| `virtual_server.http3.http_server_profile` | [virtual_server.http3.http_server_profile](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md#section) |
| `virtual_server.http3.http_server_profile.kind` | [virtual_server.http3.http_server_profile.kind](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--kind) |
| `virtual_server.http3.http_server_profile.name` | [virtual_server.http3.http_server_profile.name](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--name) |
| `virtual_server.http3.http_server_profile.namespace` | [virtual_server.http3.http_server_profile.namespace](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--namespace) |
| `virtual_server.http3.http_server_profile.tenant` | [virtual_server.http3.http_server_profile.tenant](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--tenant) |
| `virtual_server.http3.http_server_profile.uid` | [virtual_server.http3.http_server_profile.uid](resources--application_profiles--properties--virtual_server--http3--http_server_profile.md#schema-virtual_server--http3--http_server_profile--uid) |
| `virtual_server.http3.quic_profile` | [virtual_server.http3.quic_profile](resources--application_profiles--properties--virtual_server--http3--quic_profile.md#section) |
| `virtual_server.http3.quic_profile.kind` | [virtual_server.http3.quic_profile.kind](resources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--kind) |
| `virtual_server.http3.quic_profile.name` | [virtual_server.http3.quic_profile.name](resources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--name) |
| `virtual_server.http3.quic_profile.namespace` | [virtual_server.http3.quic_profile.namespace](resources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--namespace) |
| `virtual_server.http3.quic_profile.tenant` | [virtual_server.http3.quic_profile.tenant](resources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--tenant) |
| `virtual_server.http3.quic_profile.uid` | [virtual_server.http3.quic_profile.uid](resources--application_profiles--properties--virtual_server--http3--quic_profile.md#schema-virtual_server--http3--quic_profile--uid) |
| `virtual_server.http3.server_ssl_profile` | [virtual_server.http3.server_ssl_profile](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#section) |
| `virtual_server.http3.server_ssl_profile.kind` | [virtual_server.http3.server_ssl_profile.kind](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--kind) |
| `virtual_server.http3.server_ssl_profile.name` | [virtual_server.http3.server_ssl_profile.name](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--name) |
| `virtual_server.http3.server_ssl_profile.namespace` | [virtual_server.http3.server_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--namespace) |
| `virtual_server.http3.server_ssl_profile.tenant` | [virtual_server.http3.server_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--tenant) |
| `virtual_server.http3.server_ssl_profile.uid` | [virtual_server.http3.server_ssl_profile.uid](resources--application_profiles--properties--virtual_server--http3--server_ssl_profile.md#schema-virtual_server--http3--server_ssl_profile--uid) |
| `virtual_server.http3.tcp_server_profile` | [virtual_server.http3.tcp_server_profile](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#section) |
| `virtual_server.http3.tcp_server_profile.kind` | [virtual_server.http3.tcp_server_profile.kind](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--kind) |
| `virtual_server.http3.tcp_server_profile.name` | [virtual_server.http3.tcp_server_profile.name](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--name) |
| `virtual_server.http3.tcp_server_profile.namespace` | [virtual_server.http3.tcp_server_profile.namespace](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--namespace) |
| `virtual_server.http3.tcp_server_profile.tenant` | [virtual_server.http3.tcp_server_profile.tenant](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--tenant) |
| `virtual_server.http3.tcp_server_profile.uid` | [virtual_server.http3.tcp_server_profile.uid](resources--application_profiles--properties--virtual_server--http3--tcp_server_profile.md#schema-virtual_server--http3--tcp_server_profile--uid) |
| `virtual_server.http3.udp_client_profile` | [virtual_server.http3.udp_client_profile](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#section) |
| `virtual_server.http3.udp_client_profile.kind` | [virtual_server.http3.udp_client_profile.kind](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--kind) |
| `virtual_server.http3.udp_client_profile.name` | [virtual_server.http3.udp_client_profile.name](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--name) |
| `virtual_server.http3.udp_client_profile.namespace` | [virtual_server.http3.udp_client_profile.namespace](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--namespace) |
| `virtual_server.http3.udp_client_profile.tenant` | [virtual_server.http3.udp_client_profile.tenant](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--tenant) |
| `virtual_server.http3.udp_client_profile.uid` | [virtual_server.http3.udp_client_profile.uid](resources--application_profiles--properties--virtual_server--http3--udp_client_profile.md#schema-virtual_server--http3--udp_client_profile--uid) |
| `virtual_server.http3.udp_server_profile` | [virtual_server.http3.udp_server_profile](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#section) |
| `virtual_server.http3.udp_server_profile.kind` | [virtual_server.http3.udp_server_profile.kind](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--kind) |
| `virtual_server.http3.udp_server_profile.name` | [virtual_server.http3.udp_server_profile.name](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--name) |
| `virtual_server.http3.udp_server_profile.namespace` | [virtual_server.http3.udp_server_profile.namespace](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--namespace) |
| `virtual_server.http3.udp_server_profile.tenant` | [virtual_server.http3.udp_server_profile.tenant](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--tenant) |
| `virtual_server.http3.udp_server_profile.uid` | [virtual_server.http3.udp_server_profile.uid](resources--application_profiles--properties--virtual_server--http3--udp_server_profile.md#schema-virtual_server--http3--udp_server_profile--uid) |
| `virtual_server.https` | [virtual_server.https](resources--application_profiles--properties--virtual_server--https.md#section) |
| `virtual_server.https.client_ssl_profile` | [virtual_server.https.client_ssl_profile](resources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#section) |
| `virtual_server.https.client_ssl_profile.kind` | [virtual_server.https.client_ssl_profile.kind](resources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--kind) |
| `virtual_server.https.client_ssl_profile.name` | [virtual_server.https.client_ssl_profile.name](resources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--name) |
| `virtual_server.https.client_ssl_profile.namespace` | [virtual_server.https.client_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--namespace) |
| `virtual_server.https.client_ssl_profile.tenant` | [virtual_server.https.client_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--tenant) |
| `virtual_server.https.client_ssl_profile.uid` | [virtual_server.https.client_ssl_profile.uid](resources--application_profiles--properties--virtual_server--https--client_ssl_profile.md#schema-virtual_server--https--client_ssl_profile--uid) |
| `virtual_server.https.http2_client_profile` | [virtual_server.https.http2_client_profile](resources--application_profiles--properties--virtual_server--https--http2_client_profile.md#section) |
| `virtual_server.https.http2_client_profile.kind` | [virtual_server.https.http2_client_profile.kind](resources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--kind) |
| `virtual_server.https.http2_client_profile.name` | [virtual_server.https.http2_client_profile.name](resources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--name) |
| `virtual_server.https.http2_client_profile.namespace` | [virtual_server.https.http2_client_profile.namespace](resources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--namespace) |
| `virtual_server.https.http2_client_profile.tenant` | [virtual_server.https.http2_client_profile.tenant](resources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--tenant) |
| `virtual_server.https.http2_client_profile.uid` | [virtual_server.https.http2_client_profile.uid](resources--application_profiles--properties--virtual_server--https--http2_client_profile.md#schema-virtual_server--https--http2_client_profile--uid) |
| `virtual_server.https.http2_server_profile` | [virtual_server.https.http2_server_profile](resources--application_profiles--properties--virtual_server--https--http2_server_profile.md#section) |
| `virtual_server.https.http2_server_profile.kind` | [virtual_server.https.http2_server_profile.kind](resources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--kind) |
| `virtual_server.https.http2_server_profile.name` | [virtual_server.https.http2_server_profile.name](resources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--name) |
| `virtual_server.https.http2_server_profile.namespace` | [virtual_server.https.http2_server_profile.namespace](resources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--namespace) |
| `virtual_server.https.http2_server_profile.tenant` | [virtual_server.https.http2_server_profile.tenant](resources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--tenant) |
| `virtual_server.https.http2_server_profile.uid` | [virtual_server.https.http2_server_profile.uid](resources--application_profiles--properties--virtual_server--https--http2_server_profile.md#schema-virtual_server--https--http2_server_profile--uid) |
| `virtual_server.https.http_client_profile` | [virtual_server.https.http_client_profile](resources--application_profiles--properties--virtual_server--https--http_client_profile.md#section) |
| `virtual_server.https.http_client_profile.kind` | [virtual_server.https.http_client_profile.kind](resources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--kind) |
| `virtual_server.https.http_client_profile.name` | [virtual_server.https.http_client_profile.name](resources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--name) |
| `virtual_server.https.http_client_profile.namespace` | [virtual_server.https.http_client_profile.namespace](resources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--namespace) |
| `virtual_server.https.http_client_profile.tenant` | [virtual_server.https.http_client_profile.tenant](resources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--tenant) |
| `virtual_server.https.http_client_profile.uid` | [virtual_server.https.http_client_profile.uid](resources--application_profiles--properties--virtual_server--https--http_client_profile.md#schema-virtual_server--https--http_client_profile--uid) |
| `virtual_server.https.http_server_profile` | [virtual_server.https.http_server_profile](resources--application_profiles--properties--virtual_server--https--http_server_profile.md#section) |
| `virtual_server.https.http_server_profile.kind` | [virtual_server.https.http_server_profile.kind](resources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--kind) |
| `virtual_server.https.http_server_profile.name` | [virtual_server.https.http_server_profile.name](resources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--name) |
| `virtual_server.https.http_server_profile.namespace` | [virtual_server.https.http_server_profile.namespace](resources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--namespace) |
| `virtual_server.https.http_server_profile.tenant` | [virtual_server.https.http_server_profile.tenant](resources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--tenant) |
| `virtual_server.https.http_server_profile.uid` | [virtual_server.https.http_server_profile.uid](resources--application_profiles--properties--virtual_server--https--http_server_profile.md#schema-virtual_server--https--http_server_profile--uid) |
| `virtual_server.https.ocsp_profile` | [virtual_server.https.ocsp_profile](resources--application_profiles--properties--virtual_server--https--ocsp_profile.md#section) |
| `virtual_server.https.ocsp_profile.kind` | [virtual_server.https.ocsp_profile.kind](resources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--kind) |
| `virtual_server.https.ocsp_profile.name` | [virtual_server.https.ocsp_profile.name](resources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--name) |
| `virtual_server.https.ocsp_profile.namespace` | [virtual_server.https.ocsp_profile.namespace](resources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--namespace) |
| `virtual_server.https.ocsp_profile.tenant` | [virtual_server.https.ocsp_profile.tenant](resources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--tenant) |
| `virtual_server.https.ocsp_profile.uid` | [virtual_server.https.ocsp_profile.uid](resources--application_profiles--properties--virtual_server--https--ocsp_profile.md#schema-virtual_server--https--ocsp_profile--uid) |
| `virtual_server.https.server_ssl_profile` | [virtual_server.https.server_ssl_profile](resources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#section) |
| `virtual_server.https.server_ssl_profile.kind` | [virtual_server.https.server_ssl_profile.kind](resources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--kind) |
| `virtual_server.https.server_ssl_profile.name` | [virtual_server.https.server_ssl_profile.name](resources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--name) |
| `virtual_server.https.server_ssl_profile.namespace` | [virtual_server.https.server_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--namespace) |
| `virtual_server.https.server_ssl_profile.tenant` | [virtual_server.https.server_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--tenant) |
| `virtual_server.https.server_ssl_profile.uid` | [virtual_server.https.server_ssl_profile.uid](resources--application_profiles--properties--virtual_server--https--server_ssl_profile.md#schema-virtual_server--https--server_ssl_profile--uid) |
| `virtual_server.https.stream_profile` | [virtual_server.https.stream_profile](resources--application_profiles--properties--virtual_server--https--stream_profile.md#section) |
| `virtual_server.https.stream_profile.kind` | [virtual_server.https.stream_profile.kind](resources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--kind) |
| `virtual_server.https.stream_profile.name` | [virtual_server.https.stream_profile.name](resources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--name) |
| `virtual_server.https.stream_profile.namespace` | [virtual_server.https.stream_profile.namespace](resources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--namespace) |
| `virtual_server.https.stream_profile.tenant` | [virtual_server.https.stream_profile.tenant](resources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--tenant) |
| `virtual_server.https.stream_profile.uid` | [virtual_server.https.stream_profile.uid](resources--application_profiles--properties--virtual_server--https--stream_profile.md#schema-virtual_server--https--stream_profile--uid) |
| `virtual_server.https.tcp_client_profile` | [virtual_server.https.tcp_client_profile](resources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#section) |
| `virtual_server.https.tcp_client_profile.kind` | [virtual_server.https.tcp_client_profile.kind](resources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--kind) |
| `virtual_server.https.tcp_client_profile.name` | [virtual_server.https.tcp_client_profile.name](resources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--name) |
| `virtual_server.https.tcp_client_profile.namespace` | [virtual_server.https.tcp_client_profile.namespace](resources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--namespace) |
| `virtual_server.https.tcp_client_profile.tenant` | [virtual_server.https.tcp_client_profile.tenant](resources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--tenant) |
| `virtual_server.https.tcp_client_profile.uid` | [virtual_server.https.tcp_client_profile.uid](resources--application_profiles--properties--virtual_server--https--tcp_client_profile.md#schema-virtual_server--https--tcp_client_profile--uid) |
| `virtual_server.https.tcp_server_profile` | [virtual_server.https.tcp_server_profile](resources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#section) |
| `virtual_server.https.tcp_server_profile.kind` | [virtual_server.https.tcp_server_profile.kind](resources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--kind) |
| `virtual_server.https.tcp_server_profile.name` | [virtual_server.https.tcp_server_profile.name](resources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--name) |
| `virtual_server.https.tcp_server_profile.namespace` | [virtual_server.https.tcp_server_profile.namespace](resources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--namespace) |
| `virtual_server.https.tcp_server_profile.tenant` | [virtual_server.https.tcp_server_profile.tenant](resources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--tenant) |
| `virtual_server.https.tcp_server_profile.uid` | [virtual_server.https.tcp_server_profile.uid](resources--application_profiles--properties--virtual_server--https--tcp_server_profile.md#schema-virtual_server--https--tcp_server_profile--uid) |
| `virtual_server.https.websocket_client_profile` | [virtual_server.https.websocket_client_profile](resources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#section) |
| `virtual_server.https.websocket_client_profile.kind` | [virtual_server.https.websocket_client_profile.kind](resources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--kind) |
| `virtual_server.https.websocket_client_profile.name` | [virtual_server.https.websocket_client_profile.name](resources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--name) |
| `virtual_server.https.websocket_client_profile.namespace` | [virtual_server.https.websocket_client_profile.namespace](resources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--namespace) |
| `virtual_server.https.websocket_client_profile.tenant` | [virtual_server.https.websocket_client_profile.tenant](resources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--tenant) |
| `virtual_server.https.websocket_client_profile.uid` | [virtual_server.https.websocket_client_profile.uid](resources--application_profiles--properties--virtual_server--https--websocket_client_profile.md#schema-virtual_server--https--websocket_client_profile--uid) |
| `virtual_server.https.websocket_server_profile` | [virtual_server.https.websocket_server_profile](resources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#section) |
| `virtual_server.https.websocket_server_profile.kind` | [virtual_server.https.websocket_server_profile.kind](resources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--kind) |
| `virtual_server.https.websocket_server_profile.name` | [virtual_server.https.websocket_server_profile.name](resources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--name) |
| `virtual_server.https.websocket_server_profile.namespace` | [virtual_server.https.websocket_server_profile.namespace](resources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--namespace) |
| `virtual_server.https.websocket_server_profile.tenant` | [virtual_server.https.websocket_server_profile.tenant](resources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--tenant) |
| `virtual_server.https.websocket_server_profile.uid` | [virtual_server.https.websocket_server_profile.uid](resources--application_profiles--properties--virtual_server--https--websocket_server_profile.md#schema-virtual_server--https--websocket_server_profile--uid) |
| `virtual_server.immediate_action_on_service_down` | [virtual_server.immediate_action_on_service_down](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down.md#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_drop.md#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_none.md#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](resources--application_profiles--properties--virtual_server--immediate_action_on_service_down--immediate_action_on_service_down_reset.md#section) |
| `virtual_server.last_hop_pool` | [virtual_server.last_hop_pool](resources--application_profiles--properties--virtual_server--last_hop_pool.md#section) |
| `virtual_server.last_hop_pool.kind` | [virtual_server.last_hop_pool.kind](resources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--kind) |
| `virtual_server.last_hop_pool.name` | [virtual_server.last_hop_pool.name](resources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--name) |
| `virtual_server.last_hop_pool.namespace` | [virtual_server.last_hop_pool.namespace](resources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--namespace) |
| `virtual_server.last_hop_pool.tenant` | [virtual_server.last_hop_pool.tenant](resources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--tenant) |
| `virtual_server.last_hop_pool.uid` | [virtual_server.last_hop_pool.uid](resources--application_profiles--properties--virtual_server--last_hop_pool.md#schema-virtual_server--last_hop_pool--uid) |
| `virtual_server.nat64` | [virtual_server.nat64](resources--application_profiles--properties--virtual_server--nat64.md#section) |
| `virtual_server.nat64.nat64_disable` | [virtual_server.nat64.nat64_disable](resources--application_profiles--properties--virtual_server--nat64--nat64_disable.md#section) |
| `virtual_server.nat64.nat64_enable` | [virtual_server.nat64.nat64_enable](resources--application_profiles--properties--virtual_server--nat64--nat64_enable.md#section) |
| `virtual_server.port_translation` | [virtual_server.port_translation](resources--application_profiles--properties--virtual_server--port_translation.md#section) |
| `virtual_server.port_translation.port_translation_disable` | [virtual_server.port_translation.port_translation_disable](resources--application_profiles--properties--virtual_server--port_translation--port_translation_disable.md#section) |
| `virtual_server.port_translation.port_translation_enable` | [virtual_server.port_translation.port_translation_enable](resources--application_profiles--properties--virtual_server--port_translation--port_translation_enable.md#section) |
| `virtual_server.request_logging_profile` | [virtual_server.request_logging_profile](resources--application_profiles--properties--virtual_server--request_logging_profile.md#section) |
| `virtual_server.request_logging_profile.kind` | [virtual_server.request_logging_profile.kind](resources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--kind) |
| `virtual_server.request_logging_profile.name` | [virtual_server.request_logging_profile.name](resources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--name) |
| `virtual_server.request_logging_profile.namespace` | [virtual_server.request_logging_profile.namespace](resources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--namespace) |
| `virtual_server.request_logging_profile.tenant` | [virtual_server.request_logging_profile.tenant](resources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--tenant) |
| `virtual_server.request_logging_profile.uid` | [virtual_server.request_logging_profile.uid](resources--application_profiles--properties--virtual_server--request_logging_profile.md#schema-virtual_server--request_logging_profile--uid) |
| `virtual_server.source_port` | [virtual_server.source_port](resources--application_profiles--properties--virtual_server--source_port.md#section) |
| `virtual_server.source_port.source_port_change` | [virtual_server.source_port.source_port_change](resources--application_profiles--properties--virtual_server--source_port--source_port_change.md#section) |
| `virtual_server.source_port.source_port_preserve` | [virtual_server.source_port.source_port_preserve](resources--application_profiles--properties--virtual_server--source_port--source_port_preserve.md#section) |
| `virtual_server.source_port.source_port_preserve_strict` | [virtual_server.source_port.source_port_preserve_strict](resources--application_profiles--properties--virtual_server--source_port--source_port_preserve_strict.md#section) |
| `virtual_server.statistics_profile` | [virtual_server.statistics_profile](resources--application_profiles--properties--virtual_server--statistics_profile.md#section) |
| `virtual_server.statistics_profile.kind` | [virtual_server.statistics_profile.kind](resources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--kind) |
| `virtual_server.statistics_profile.name` | [virtual_server.statistics_profile.name](resources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--name) |
| `virtual_server.statistics_profile.namespace` | [virtual_server.statistics_profile.namespace](resources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--namespace) |
| `virtual_server.statistics_profile.tenant` | [virtual_server.statistics_profile.tenant](resources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--tenant) |
| `virtual_server.statistics_profile.uid` | [virtual_server.statistics_profile.uid](resources--application_profiles--properties--virtual_server--statistics_profile.md#schema-virtual_server--statistics_profile--uid) |
| `virtual_server.tcp` | [virtual_server.tcp](resources--application_profiles--properties--virtual_server--tcp.md#section) |
| `virtual_server.tcp.client_ssl_profile` | [virtual_server.tcp.client_ssl_profile](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#section) |
| `virtual_server.tcp.client_ssl_profile.kind` | [virtual_server.tcp.client_ssl_profile.kind](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--kind) |
| `virtual_server.tcp.client_ssl_profile.name` | [virtual_server.tcp.client_ssl_profile.name](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--name) |
| `virtual_server.tcp.client_ssl_profile.namespace` | [virtual_server.tcp.client_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--namespace) |
| `virtual_server.tcp.client_ssl_profile.tenant` | [virtual_server.tcp.client_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--tenant) |
| `virtual_server.tcp.client_ssl_profile.uid` | [virtual_server.tcp.client_ssl_profile.uid](resources--application_profiles--properties--virtual_server--tcp--client_ssl_profile.md#schema-virtual_server--tcp--client_ssl_profile--uid) |
| `virtual_server.tcp.ocsp_profile` | [virtual_server.tcp.ocsp_profile](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#section) |
| `virtual_server.tcp.ocsp_profile.kind` | [virtual_server.tcp.ocsp_profile.kind](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--kind) |
| `virtual_server.tcp.ocsp_profile.name` | [virtual_server.tcp.ocsp_profile.name](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--name) |
| `virtual_server.tcp.ocsp_profile.namespace` | [virtual_server.tcp.ocsp_profile.namespace](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--namespace) |
| `virtual_server.tcp.ocsp_profile.tenant` | [virtual_server.tcp.ocsp_profile.tenant](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--tenant) |
| `virtual_server.tcp.ocsp_profile.uid` | [virtual_server.tcp.ocsp_profile.uid](resources--application_profiles--properties--virtual_server--tcp--ocsp_profile.md#schema-virtual_server--tcp--ocsp_profile--uid) |
| `virtual_server.tcp.server_ssl_profile` | [virtual_server.tcp.server_ssl_profile](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#section) |
| `virtual_server.tcp.server_ssl_profile.kind` | [virtual_server.tcp.server_ssl_profile.kind](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--kind) |
| `virtual_server.tcp.server_ssl_profile.name` | [virtual_server.tcp.server_ssl_profile.name](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--name) |
| `virtual_server.tcp.server_ssl_profile.namespace` | [virtual_server.tcp.server_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--namespace) |
| `virtual_server.tcp.server_ssl_profile.tenant` | [virtual_server.tcp.server_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--tenant) |
| `virtual_server.tcp.server_ssl_profile.uid` | [virtual_server.tcp.server_ssl_profile.uid](resources--application_profiles--properties--virtual_server--tcp--server_ssl_profile.md#schema-virtual_server--tcp--server_ssl_profile--uid) |
| `virtual_server.tcp.tcp_client_profile` | [virtual_server.tcp.tcp_client_profile](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#section) |
| `virtual_server.tcp.tcp_client_profile.kind` | [virtual_server.tcp.tcp_client_profile.kind](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--kind) |
| `virtual_server.tcp.tcp_client_profile.name` | [virtual_server.tcp.tcp_client_profile.name](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--name) |
| `virtual_server.tcp.tcp_client_profile.namespace` | [virtual_server.tcp.tcp_client_profile.namespace](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--namespace) |
| `virtual_server.tcp.tcp_client_profile.tenant` | [virtual_server.tcp.tcp_client_profile.tenant](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--tenant) |
| `virtual_server.tcp.tcp_client_profile.uid` | [virtual_server.tcp.tcp_client_profile.uid](resources--application_profiles--properties--virtual_server--tcp--tcp_client_profile.md#schema-virtual_server--tcp--tcp_client_profile--uid) |
| `virtual_server.tcp.tcp_server_profile` | [virtual_server.tcp.tcp_server_profile](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#section) |
| `virtual_server.tcp.tcp_server_profile.kind` | [virtual_server.tcp.tcp_server_profile.kind](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--kind) |
| `virtual_server.tcp.tcp_server_profile.name` | [virtual_server.tcp.tcp_server_profile.name](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--name) |
| `virtual_server.tcp.tcp_server_profile.namespace` | [virtual_server.tcp.tcp_server_profile.namespace](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--namespace) |
| `virtual_server.tcp.tcp_server_profile.tenant` | [virtual_server.tcp.tcp_server_profile.tenant](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--tenant) |
| `virtual_server.tcp.tcp_server_profile.uid` | [virtual_server.tcp.tcp_server_profile.uid](resources--application_profiles--properties--virtual_server--tcp--tcp_server_profile.md#schema-virtual_server--tcp--tcp_server_profile--uid) |
| `virtual_server.udp` | [virtual_server.udp](resources--application_profiles--properties--virtual_server--udp.md#section) |
| `virtual_server.udp.client_ssl_profile` | [virtual_server.udp.client_ssl_profile](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#section) |
| `virtual_server.udp.client_ssl_profile.kind` | [virtual_server.udp.client_ssl_profile.kind](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--kind) |
| `virtual_server.udp.client_ssl_profile.name` | [virtual_server.udp.client_ssl_profile.name](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--name) |
| `virtual_server.udp.client_ssl_profile.namespace` | [virtual_server.udp.client_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--namespace) |
| `virtual_server.udp.client_ssl_profile.tenant` | [virtual_server.udp.client_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--tenant) |
| `virtual_server.udp.client_ssl_profile.uid` | [virtual_server.udp.client_ssl_profile.uid](resources--application_profiles--properties--virtual_server--udp--client_ssl_profile.md#schema-virtual_server--udp--client_ssl_profile--uid) |
| `virtual_server.udp.server_ssl_profile` | [virtual_server.udp.server_ssl_profile](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#section) |
| `virtual_server.udp.server_ssl_profile.kind` | [virtual_server.udp.server_ssl_profile.kind](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--kind) |
| `virtual_server.udp.server_ssl_profile.name` | [virtual_server.udp.server_ssl_profile.name](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--name) |
| `virtual_server.udp.server_ssl_profile.namespace` | [virtual_server.udp.server_ssl_profile.namespace](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--namespace) |
| `virtual_server.udp.server_ssl_profile.tenant` | [virtual_server.udp.server_ssl_profile.tenant](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--tenant) |
| `virtual_server.udp.server_ssl_profile.uid` | [virtual_server.udp.server_ssl_profile.uid](resources--application_profiles--properties--virtual_server--udp--server_ssl_profile.md#schema-virtual_server--udp--server_ssl_profile--uid) |
| `virtual_server.udp.udp_client_profile` | [virtual_server.udp.udp_client_profile](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#section) |
| `virtual_server.udp.udp_client_profile.kind` | [virtual_server.udp.udp_client_profile.kind](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--kind) |
| `virtual_server.udp.udp_client_profile.name` | [virtual_server.udp.udp_client_profile.name](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--name) |
| `virtual_server.udp.udp_client_profile.namespace` | [virtual_server.udp.udp_client_profile.namespace](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--namespace) |
| `virtual_server.udp.udp_client_profile.tenant` | [virtual_server.udp.udp_client_profile.tenant](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--tenant) |
| `virtual_server.udp.udp_client_profile.uid` | [virtual_server.udp.udp_client_profile.uid](resources--application_profiles--properties--virtual_server--udp--udp_client_profile.md#schema-virtual_server--udp--udp_client_profile--uid) |
| `virtual_server.udp.udp_server_profile` | [virtual_server.udp.udp_server_profile](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#section) |
| `virtual_server.udp.udp_server_profile.kind` | [virtual_server.udp.udp_server_profile.kind](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--kind) |
| `virtual_server.udp.udp_server_profile.name` | [virtual_server.udp.udp_server_profile.name](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--name) |
| `virtual_server.udp.udp_server_profile.namespace` | [virtual_server.udp.udp_server_profile.namespace](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--namespace) |
| `virtual_server.udp.udp_server_profile.tenant` | [virtual_server.udp.udp_server_profile.tenant](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--tenant) |
| `virtual_server.udp.udp_server_profile.uid` | [virtual_server.udp.udp_server_profile.uid](resources--application_profiles--properties--virtual_server--udp--udp_server_profile.md#schema-virtual_server--udp--udp_server_profile--uid) |
| `virtual_server.virtual_server_state` | [virtual_server.virtual_server_state](resources--application_profiles--properties--virtual_server--virtual_server_state.md#section) |
| `virtual_server.virtual_server_state.state_disabled` | [virtual_server.virtual_server_state.state_disabled](resources--application_profiles--properties--virtual_server--virtual_server_state--state_disabled.md#section) |
| `virtual_server.virtual_server_state.state_enabled` | [virtual_server.virtual_server_state.state_enabled](resources--application_profiles--properties--virtual_server--virtual_server_state--state_enabled.md#section) |
| `virtual_server.vs_score` | [virtual_server.vs_score](resources--application_profiles--properties--virtual_server.md#schema-virtual_server--vs_score) |

## Next pages

- [advanced_tcp_profile](resources--application_profiles--properties--advanced_tcp_profile.md)
- [ddos_profile](resources--application_profiles--properties--ddos_profile.md)
- [irules](resources--application_profiles--properties--irules.md)
- [timeouts](resources--application_profiles--properties--timeouts.md)
- [virtual_server](resources--application_profiles--properties--virtual_server.md)
- [xcsh_application_profiles](../resources/application_profiles.md)
