---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_k8s_cluster_role_binding."
xcsh_docs: {"aliases": [], "body_bytes": 8520, "body_sha256": "sha256:9133710394e3b18767520860674b8d8abd03c82008f0f319179d4dc34bc3fb71", "canonical_id": "xcsh-docs:data-sources:k8s_cluster_role_binding:reference", "child_ids": ["xcsh-docs:data-sources:k8s_cluster_role_binding:properties:k8s_cluster_role", "xcsh-docs:data-sources:k8s_cluster_role_binding:properties:subjects"], "collection_id": "xcsh-docs:data-sources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster_role_binding:reference", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role_binding:fundamentals", "path": "docs/guides/data-sources--k8s_cluster_role_binding--reference.md", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role_binding/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_k8s_cluster_role_binding.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md)
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

Description of the K8SClusterRoleBinding.

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

- [k8s_cluster_role](data-sources--k8s_cluster_role_binding--properties--k8s_cluster_role.md): complete subsection reference.

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

Name of the K8SClusterRoleBinding.

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

Namespace where the K8SClusterRoleBinding exists.

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

- [subjects](data-sources--k8s_cluster_role_binding--properties--subjects.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_cluster_role_binding--reference.md#schema-annotations) |
| `description` | [description](data-sources--k8s_cluster_role_binding--reference.md#schema-description) |
| `id` | [id](data-sources--k8s_cluster_role_binding--reference.md#schema-id) |
| `k8s_cluster_role` | [k8s_cluster_role](data-sources--k8s_cluster_role_binding--properties--k8s_cluster_role.md#section) |
| `k8s_cluster_role.name` | [k8s_cluster_role.name](data-sources--k8s_cluster_role_binding--properties--k8s_cluster_role.md#schema-k8s_cluster_role--name) |
| `k8s_cluster_role.namespace` | [k8s_cluster_role.namespace](data-sources--k8s_cluster_role_binding--properties--k8s_cluster_role.md#schema-k8s_cluster_role--namespace) |
| `k8s_cluster_role.tenant` | [k8s_cluster_role.tenant](data-sources--k8s_cluster_role_binding--properties--k8s_cluster_role.md#schema-k8s_cluster_role--tenant) |
| `labels` | [labels](data-sources--k8s_cluster_role_binding--reference.md#schema-labels) |
| `name` | [name](data-sources--k8s_cluster_role_binding--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--k8s_cluster_role_binding--reference.md#schema-namespace) |
| `subjects` | [subjects](data-sources--k8s_cluster_role_binding--properties--subjects.md#section) |
| `subjects.group` | [subjects.group](data-sources--k8s_cluster_role_binding--properties--subjects.md#schema-subjects--group) |
| `subjects.service_account` | [subjects.service_account](data-sources--k8s_cluster_role_binding--properties--subjects--service_account.md#section) |
| `subjects.service_account.name` | [subjects.service_account.name](data-sources--k8s_cluster_role_binding--properties--subjects--service_account.md#schema-subjects--service_account--name) |
| `subjects.service_account.namespace` | [subjects.service_account.namespace](data-sources--k8s_cluster_role_binding--properties--subjects--service_account.md#schema-subjects--service_account--namespace) |
| `subjects.user` | [subjects.user](data-sources--k8s_cluster_role_binding--properties--subjects.md#schema-subjects--user) |

## Next pages

- [k8s_cluster_role](data-sources--k8s_cluster_role_binding--properties--k8s_cluster_role.md)
- [subjects](data-sources--k8s_cluster_role_binding--properties--subjects.md)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md)
