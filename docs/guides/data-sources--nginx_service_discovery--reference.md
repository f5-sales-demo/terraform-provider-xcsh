---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 11434, "body_sha256": "sha256:ac76096366df7b5a0eda31bd3639999c568fe65d19eae920294920b8adaf267f", "canonical_id": "xcsh-docs:data-sources:nginx_service_discovery:reference", "child_ids": ["xcsh-docs:data-sources:nginx_service_discovery:properties:discovery_target", "xcsh-docs:data-sources:nginx_service_discovery:properties:server_block_filters"], "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_service_discovery:reference", "parent_id": "xcsh-docs:data-sources:nginx_service_discovery:fundamentals", "path": "docs/guides/data-sources--nginx_service_discovery--reference.md", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_nginx_service_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md)
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

Description of the NginxServiceDiscovery.

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

- [discovery_target](data-sources--nginx_service_discovery--properties--discovery_target.md): complete subsection reference.

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

Name of the NginxServiceDiscovery.

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

Namespace where the NginxServiceDiscovery exists.

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

- [server_block_filters](data-sources--nginx_service_discovery--properties--server_block_filters.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nginx_service_discovery--reference.md#schema-annotations) |
| `description` | [description](data-sources--nginx_service_discovery--reference.md#schema-description) |
| `discovery_target` | [discovery_target](data-sources--nginx_service_discovery--properties--discovery_target.md#section) |
| `discovery_target.config_sync_group` | [discovery_target.config_sync_group](data-sources--nginx_service_discovery--properties--discovery_target--config_sync_group.md#section) |
| `discovery_target.config_sync_group.config_sync_group` | [discovery_target.config_sync_group.config_sync_group](data-sources--nginx_service_discovery--properties--discovery_target--config_sync_group--config_sync_group.md#section) |
| `discovery_target.config_sync_group.config_sync_group.kind` | [discovery_target.config_sync_group.config_sync_group.kind](data-sources--nginx_service_discovery--properties--discovery_target--config_sync_group--config_sync_group.md#schema-discovery_target--config_sync_group--config_sync_group--kind) |
| `discovery_target.config_sync_group.config_sync_group.name` | [discovery_target.config_sync_group.config_sync_group.name](data-sources--nginx_service_discovery--properties--discovery_target--config_sync_group--config_sync_group.md#schema-discovery_target--config_sync_group--config_sync_group--name) |
| `discovery_target.config_sync_group.config_sync_group.namespace` | [discovery_target.config_sync_group.config_sync_group.namespace](data-sources--nginx_service_discovery--properties--discovery_target--config_sync_group--config_sync_group.md#schema-discovery_target--config_sync_group--config_sync_group--namespace) |
| `discovery_target.config_sync_group.config_sync_group.tenant` | [discovery_target.config_sync_group.config_sync_group.tenant](data-sources--nginx_service_discovery--properties--discovery_target--config_sync_group--config_sync_group.md#schema-discovery_target--config_sync_group--config_sync_group--tenant) |
| `discovery_target.config_sync_group.config_sync_group.uid` | [discovery_target.config_sync_group.config_sync_group.uid](data-sources--nginx_service_discovery--properties--discovery_target--config_sync_group--config_sync_group.md#schema-discovery_target--config_sync_group--config_sync_group--uid) |
| `discovery_target.nginx_instance` | [discovery_target.nginx_instance](data-sources--nginx_service_discovery--properties--discovery_target--nginx_instance.md#section) |
| `discovery_target.nginx_instance.nginx_instance` | [discovery_target.nginx_instance.nginx_instance](data-sources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md#section) |
| `discovery_target.nginx_instance.nginx_instance.kind` | [discovery_target.nginx_instance.nginx_instance.kind](data-sources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md#schema-discovery_target--nginx_instance--nginx_instance--kind) |
| `discovery_target.nginx_instance.nginx_instance.name` | [discovery_target.nginx_instance.nginx_instance.name](data-sources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md#schema-discovery_target--nginx_instance--nginx_instance--name) |
| `discovery_target.nginx_instance.nginx_instance.namespace` | [discovery_target.nginx_instance.nginx_instance.namespace](data-sources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md#schema-discovery_target--nginx_instance--nginx_instance--namespace) |
| `discovery_target.nginx_instance.nginx_instance.tenant` | [discovery_target.nginx_instance.nginx_instance.tenant](data-sources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md#schema-discovery_target--nginx_instance--nginx_instance--tenant) |
| `discovery_target.nginx_instance.nginx_instance.uid` | [discovery_target.nginx_instance.nginx_instance.uid](data-sources--nginx_service_discovery--properties--discovery_target--nginx_instance--nginx_instance.md#schema-discovery_target--nginx_instance--nginx_instance--uid) |
| `id` | [id](data-sources--nginx_service_discovery--reference.md#schema-id) |
| `labels` | [labels](data-sources--nginx_service_discovery--reference.md#schema-labels) |
| `name` | [name](data-sources--nginx_service_discovery--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--nginx_service_discovery--reference.md#schema-namespace) |
| `server_block_filters` | [server_block_filters](data-sources--nginx_service_discovery--properties--server_block_filters.md#section) |
| `server_block_filters.name_regex` | [server_block_filters.name_regex](data-sources--nginx_service_discovery--properties--server_block_filters.md#schema-server_block_filters--name_regex) |
| `server_block_filters.port_ranges` | [server_block_filters.port_ranges](data-sources--nginx_service_discovery--properties--server_block_filters.md#schema-server_block_filters--port_ranges) |

## Next pages

- [discovery_target](data-sources--nginx_service_discovery--properties--discovery_target.md)
- [server_block_filters](data-sources--nginx_service_discovery--properties--server_block_filters.md)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md)
