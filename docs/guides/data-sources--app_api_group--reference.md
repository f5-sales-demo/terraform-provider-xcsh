---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 10593, "body_sha256": "sha256:a22f025ee917633d26161c82065e8af45d79d706403f84cad25882389ee2153f", "canonical_id": "xcsh-docs:data-sources:app_api_group:reference", "child_ids": ["xcsh-docs:data-sources:app_api_group:properties:bigip_virtual_server", "xcsh-docs:data-sources:app_api_group:properties:cdn_loadbalancer", "xcsh-docs:data-sources:app_api_group:properties:elements", "xcsh-docs:data-sources:app_api_group:properties:http_loadbalancer"], "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:reference", "parent_id": "xcsh-docs:data-sources:app_api_group:fundamentals", "path": "docs/guides/data-sources--app_api_group--reference.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md)
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

- [bigip_virtual_server](data-sources--app_api_group--properties--bigip_virtual_server.md): complete subsection reference.

- [cdn_loadbalancer](data-sources--app_api_group--properties--cdn_loadbalancer.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AppAPIGroup.

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

- [elements](data-sources--app_api_group--properties--elements.md): complete subsection reference.

- [http_loadbalancer](data-sources--app_api_group--properties--http_loadbalancer.md): complete subsection reference.

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

Name of the AppAPIGroup.

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

Namespace where the AppAPIGroup exists.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--app_api_group--reference.md#schema-annotations) |
| `bigip_virtual_server` | [bigip_virtual_server](data-sources--app_api_group--properties--bigip_virtual_server.md#section) |
| `bigip_virtual_server.bigip_virtual_server` | [bigip_virtual_server.bigip_virtual_server](data-sources--app_api_group--properties--bigip_virtual_server--bigip_virtual_server.md#section) |
| `bigip_virtual_server.bigip_virtual_server.name` | [bigip_virtual_server.bigip_virtual_server.name](data-sources--app_api_group--properties--bigip_virtual_server--bigip_virtual_server.md#schema-bigip_virtual_server--bigip_virtual_server--name) |
| `bigip_virtual_server.bigip_virtual_server.namespace` | [bigip_virtual_server.bigip_virtual_server.namespace](data-sources--app_api_group--properties--bigip_virtual_server--bigip_virtual_server.md#schema-bigip_virtual_server--bigip_virtual_server--namespace) |
| `bigip_virtual_server.bigip_virtual_server.tenant` | [bigip_virtual_server.bigip_virtual_server.tenant](data-sources--app_api_group--properties--bigip_virtual_server--bigip_virtual_server.md#schema-bigip_virtual_server--bigip_virtual_server--tenant) |
| `cdn_loadbalancer` | [cdn_loadbalancer](data-sources--app_api_group--properties--cdn_loadbalancer.md#section) |
| `cdn_loadbalancer.cdn_loadbalancer` | [cdn_loadbalancer.cdn_loadbalancer](data-sources--app_api_group--properties--cdn_loadbalancer--cdn_loadbalancer.md#section) |
| `cdn_loadbalancer.cdn_loadbalancer.name` | [cdn_loadbalancer.cdn_loadbalancer.name](data-sources--app_api_group--properties--cdn_loadbalancer--cdn_loadbalancer.md#schema-cdn_loadbalancer--cdn_loadbalancer--name) |
| `cdn_loadbalancer.cdn_loadbalancer.namespace` | [cdn_loadbalancer.cdn_loadbalancer.namespace](data-sources--app_api_group--properties--cdn_loadbalancer--cdn_loadbalancer.md#schema-cdn_loadbalancer--cdn_loadbalancer--namespace) |
| `cdn_loadbalancer.cdn_loadbalancer.tenant` | [cdn_loadbalancer.cdn_loadbalancer.tenant](data-sources--app_api_group--properties--cdn_loadbalancer--cdn_loadbalancer.md#schema-cdn_loadbalancer--cdn_loadbalancer--tenant) |
| `description` | [description](data-sources--app_api_group--reference.md#schema-description) |
| `elements` | [elements](data-sources--app_api_group--properties--elements.md#section) |
| `elements.methods` | [elements.methods](data-sources--app_api_group--properties--elements.md#schema-elements--methods) |
| `elements.path_regex` | [elements.path_regex](data-sources--app_api_group--properties--elements.md#schema-elements--path_regex) |
| `http_loadbalancer` | [http_loadbalancer](data-sources--app_api_group--properties--http_loadbalancer.md#section) |
| `http_loadbalancer.http_loadbalancer` | [http_loadbalancer.http_loadbalancer](data-sources--app_api_group--properties--http_loadbalancer--http_loadbalancer.md#section) |
| `http_loadbalancer.http_loadbalancer.name` | [http_loadbalancer.http_loadbalancer.name](data-sources--app_api_group--properties--http_loadbalancer--http_loadbalancer.md#schema-http_loadbalancer--http_loadbalancer--name) |
| `http_loadbalancer.http_loadbalancer.namespace` | [http_loadbalancer.http_loadbalancer.namespace](data-sources--app_api_group--properties--http_loadbalancer--http_loadbalancer.md#schema-http_loadbalancer--http_loadbalancer--namespace) |
| `http_loadbalancer.http_loadbalancer.tenant` | [http_loadbalancer.http_loadbalancer.tenant](data-sources--app_api_group--properties--http_loadbalancer--http_loadbalancer.md#schema-http_loadbalancer--http_loadbalancer--tenant) |
| `id` | [id](data-sources--app_api_group--reference.md#schema-id) |
| `labels` | [labels](data-sources--app_api_group--reference.md#schema-labels) |
| `name` | [name](data-sources--app_api_group--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--app_api_group--reference.md#schema-namespace) |

## Next pages

- [bigip_virtual_server](data-sources--app_api_group--properties--bigip_virtual_server.md)
- [cdn_loadbalancer](data-sources--app_api_group--properties--cdn_loadbalancer.md)
- [elements](data-sources--app_api_group--properties--elements.md)
- [http_loadbalancer](data-sources--app_api_group--properties--http_loadbalancer.md)
- [xcsh_app_api_group](../data-sources/app_api_group.md)
