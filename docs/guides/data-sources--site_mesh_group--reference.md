---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 11656, "body_sha256": "sha256:ec7e932e3e88fa7fca1522760a9b9dbde87456dbb88546a87aac651487c74aa0", "canonical_id": "xcsh-docs:data-sources:site_mesh_group:reference", "child_ids": ["xcsh-docs:data-sources:site_mesh_group:properties:bfd_disabled", "xcsh-docs:data-sources:site_mesh_group:properties:bfd_enabled", "xcsh-docs:data-sources:site_mesh_group:properties:disable_re_fallback", "xcsh-docs:data-sources:site_mesh_group:properties:enable_re_fallback", "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:spoke_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:virtual_site"], "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:reference", "parent_id": "xcsh-docs:data-sources:site_mesh_group:fundamentals", "path": "docs/guides/data-sources--site_mesh_group--reference.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
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

- [bfd_disabled](data-sources--site_mesh_group--properties--bfd_disabled.md): complete subsection reference.

- [bfd_enabled](data-sources--site_mesh_group--properties--bfd_enabled.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the SiteMeshGroup.

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

- [disable_re_fallback](data-sources--site_mesh_group--properties--disable_re_fallback.md): complete subsection reference.

- [enable_re_fallback](data-sources--site_mesh_group--properties--enable_re_fallback.md): complete subsection reference.

- [full_mesh](data-sources--site_mesh_group--properties--full_mesh.md): complete subsection reference.

- [hub_mesh](data-sources--site_mesh_group--properties--hub_mesh.md): complete subsection reference.

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

Name of the SiteMeshGroup.

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

Namespace where the SiteMeshGroup exists.

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

- [spoke_mesh](data-sources--site_mesh_group--properties--spoke_mesh.md): complete subsection reference.

- [virtual_site](data-sources--site_mesh_group--properties--virtual_site.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--site_mesh_group--reference.md#schema-annotations) |
| `bfd_disabled` | [bfd_disabled](data-sources--site_mesh_group--properties--bfd_disabled.md#section) |
| `bfd_enabled` | [bfd_enabled](data-sources--site_mesh_group--properties--bfd_enabled.md#section) |
| `bfd_enabled.multiplier` | [bfd_enabled.multiplier](data-sources--site_mesh_group--properties--bfd_enabled.md#schema-bfd_enabled--multiplier) |
| `bfd_enabled.receive_interval_milliseconds` | [bfd_enabled.receive_interval_milliseconds](data-sources--site_mesh_group--properties--bfd_enabled.md#schema-bfd_enabled--receive_interval_milliseconds) |
| `bfd_enabled.transmit_interval_milliseconds` | [bfd_enabled.transmit_interval_milliseconds](data-sources--site_mesh_group--properties--bfd_enabled.md#schema-bfd_enabled--transmit_interval_milliseconds) |
| `description` | [description](data-sources--site_mesh_group--reference.md#schema-description) |
| `disable_re_fallback` | [disable_re_fallback](data-sources--site_mesh_group--properties--disable_re_fallback.md#section) |
| `enable_re_fallback` | [enable_re_fallback](data-sources--site_mesh_group--properties--enable_re_fallback.md#section) |
| `full_mesh` | [full_mesh](data-sources--site_mesh_group--properties--full_mesh.md#section) |
| `full_mesh.control_and_data_plane_mesh` | [full_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--properties--full_mesh--control_and_data_plane_mesh.md#section) |
| `full_mesh.data_plane_mesh` | [full_mesh.data_plane_mesh](data-sources--site_mesh_group--properties--full_mesh--data_plane_mesh.md#section) |
| `hub_mesh` | [hub_mesh](data-sources--site_mesh_group--properties--hub_mesh.md#section) |
| `hub_mesh.control_and_data_plane_mesh` | [hub_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--properties--hub_mesh--control_and_data_plane_mesh.md#section) |
| `hub_mesh.data_plane_mesh` | [hub_mesh.data_plane_mesh](data-sources--site_mesh_group--properties--hub_mesh--data_plane_mesh.md#section) |
| `id` | [id](data-sources--site_mesh_group--reference.md#schema-id) |
| `labels` | [labels](data-sources--site_mesh_group--reference.md#schema-labels) |
| `name` | [name](data-sources--site_mesh_group--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--site_mesh_group--reference.md#schema-namespace) |
| `spoke_mesh` | [spoke_mesh](data-sources--site_mesh_group--properties--spoke_mesh.md#section) |
| `spoke_mesh.control_and_data_plane_mesh` | [spoke_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--properties--spoke_mesh--control_and_data_plane_mesh.md#section) |
| `spoke_mesh.data_plane_mesh` | [spoke_mesh.data_plane_mesh](data-sources--site_mesh_group--properties--spoke_mesh--data_plane_mesh.md#section) |
| `spoke_mesh.hub_mesh_group` | [spoke_mesh.hub_mesh_group](data-sources--site_mesh_group--properties--spoke_mesh--hub_mesh_group.md#section) |
| `spoke_mesh.hub_mesh_group.name` | [spoke_mesh.hub_mesh_group.name](data-sources--site_mesh_group--properties--spoke_mesh--hub_mesh_group.md#schema-spoke_mesh--hub_mesh_group--name) |
| `spoke_mesh.hub_mesh_group.namespace` | [spoke_mesh.hub_mesh_group.namespace](data-sources--site_mesh_group--properties--spoke_mesh--hub_mesh_group.md#schema-spoke_mesh--hub_mesh_group--namespace) |
| `spoke_mesh.hub_mesh_group.tenant` | [spoke_mesh.hub_mesh_group.tenant](data-sources--site_mesh_group--properties--spoke_mesh--hub_mesh_group.md#schema-spoke_mesh--hub_mesh_group--tenant) |
| `virtual_site` | [virtual_site](data-sources--site_mesh_group--properties--virtual_site.md#section) |
| `virtual_site.kind` | [virtual_site.kind](data-sources--site_mesh_group--properties--virtual_site.md#schema-virtual_site--kind) |
| `virtual_site.name` | [virtual_site.name](data-sources--site_mesh_group--properties--virtual_site.md#schema-virtual_site--name) |
| `virtual_site.namespace` | [virtual_site.namespace](data-sources--site_mesh_group--properties--virtual_site.md#schema-virtual_site--namespace) |
| `virtual_site.tenant` | [virtual_site.tenant](data-sources--site_mesh_group--properties--virtual_site.md#schema-virtual_site--tenant) |
| `virtual_site.uid` | [virtual_site.uid](data-sources--site_mesh_group--properties--virtual_site.md#schema-virtual_site--uid) |

## Next pages

- [bfd_disabled](data-sources--site_mesh_group--properties--bfd_disabled.md)
- [bfd_enabled](data-sources--site_mesh_group--properties--bfd_enabled.md)
- [disable_re_fallback](data-sources--site_mesh_group--properties--disable_re_fallback.md)
- [enable_re_fallback](data-sources--site_mesh_group--properties--enable_re_fallback.md)
- [full_mesh](data-sources--site_mesh_group--properties--full_mesh.md)
- [hub_mesh](data-sources--site_mesh_group--properties--hub_mesh.md)
- [spoke_mesh](data-sources--site_mesh_group--properties--spoke_mesh.md)
- [virtual_site](data-sources--site_mesh_group--properties--virtual_site.md)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
