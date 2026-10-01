---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 25143, "body_sha256": "sha256:9bc999769464e3204ba72c5fa6f167992d6f5a8795d5819c735790db20a1e919", "canonical_id": "xcsh-docs:resources:k8s_cluster:reference", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:cluster_scoped_access_deny", "xcsh-docs:resources:k8s_cluster:properties:cluster_scoped_access_permit", "xcsh-docs:resources:k8s_cluster:properties:cluster_wide_app_list", "xcsh-docs:resources:k8s_cluster:properties:global_access_enable", "xcsh-docs:resources:k8s_cluster:properties:insecure_registry_list", "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "xcsh-docs:resources:k8s_cluster:properties:no_cluster_wide_apps", "xcsh-docs:resources:k8s_cluster:properties:no_global_access", "xcsh-docs:resources:k8s_cluster:properties:no_insecure_registries", "xcsh-docs:resources:k8s_cluster:properties:no_local_access", "xcsh-docs:resources:k8s_cluster:properties:timeouts", "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_bindings", "xcsh-docs:resources:k8s_cluster:properties:use_custom_cluster_role_list", "xcsh-docs:resources:k8s_cluster:properties:use_custom_pod_security_admission", "xcsh-docs:resources:k8s_cluster:properties:use_custom_psp_list", "xcsh-docs:resources:k8s_cluster:properties:use_default_cluster_role_bindings", "xcsh-docs:resources:k8s_cluster:properties:use_default_cluster_roles", "xcsh-docs:resources:k8s_cluster:properties:use_default_pod_security_admission", "xcsh-docs:resources:k8s_cluster:properties:use_default_psp", "xcsh-docs:resources:k8s_cluster:properties:vk8s_namespace_access_deny", "xcsh-docs:resources:k8s_cluster:properties:vk8s_namespace_access_permit"], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:reference", "parent_id": "xcsh-docs:resources:k8s_cluster:fundamentals", "path": "docs/guides/resources--k8s_cluster--reference.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- Property reference

## Direct properties

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

- [cluster_scoped_access_deny](resources--k8s_cluster--properties--cluster_scoped_access_deny.md): complete subsection reference.

- [cluster_scoped_access_permit](resources--k8s_cluster--properties--cluster_scoped_access_permit.md): complete subsection reference.

- [cluster_wide_app_list](resources--k8s_cluster--properties--cluster_wide_app_list.md): complete subsection reference.

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

- [global_access_enable](resources--k8s_cluster--properties--global_access_enable.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [insecure_registry_list](resources--k8s_cluster--properties--insecure_registry_list.md): complete subsection reference.

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

- [local_access_config](resources--k8s_cluster--properties--local_access_config.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the K8S Cluster. Must be unique within the namespace.

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

Type: `"string"`. Optional, Computed.

Namespace for the K8S Cluster. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [no_cluster_wide_apps](resources--k8s_cluster--properties--no_cluster_wide_apps.md): complete subsection reference.

- [no_global_access](resources--k8s_cluster--properties--no_global_access.md): complete subsection reference.

- [no_insecure_registries](resources--k8s_cluster--properties--no_insecure_registries.md): complete subsection reference.

- [no_local_access](resources--k8s_cluster--properties--no_local_access.md): complete subsection reference.

- [timeouts](resources--k8s_cluster--properties--timeouts.md): complete subsection reference.

- [use_custom_cluster_role_bindings](resources--k8s_cluster--properties--use_custom_cluster_role_bindings.md): complete subsection reference.

- [use_custom_cluster_role_list](resources--k8s_cluster--properties--use_custom_cluster_role_list.md): complete subsection reference.

- [use_custom_pod_security_admission](resources--k8s_cluster--properties--use_custom_pod_security_admission.md): complete subsection reference.

- [use_custom_psp_list](resources--k8s_cluster--properties--use_custom_psp_list.md): complete subsection reference.

- [use_default_cluster_role_bindings](resources--k8s_cluster--properties--use_default_cluster_role_bindings.md): complete subsection reference.

- [use_default_cluster_roles](resources--k8s_cluster--properties--use_default_cluster_roles.md): complete subsection reference.

- [use_default_pod_security_admission](resources--k8s_cluster--properties--use_default_pod_security_admission.md): complete subsection reference.

- [use_default_psp](resources--k8s_cluster--properties--use_default_psp.md): complete subsection reference.

- [vk8s_namespace_access_deny](resources--k8s_cluster--properties--vk8s_namespace_access_deny.md): complete subsection reference.

- [vk8s_namespace_access_permit](resources--k8s_cluster--properties--vk8s_namespace_access_permit.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_cluster--reference.md#schema-annotations) |
| `cluster_scoped_access_deny` | [cluster_scoped_access_deny](resources--k8s_cluster--properties--cluster_scoped_access_deny.md#section) |
| `cluster_scoped_access_permit` | [cluster_scoped_access_permit](resources--k8s_cluster--properties--cluster_scoped_access_permit.md#section) |
| `cluster_wide_app_list` | [cluster_wide_app_list](resources--k8s_cluster--properties--cluster_wide_app_list.md#section) |
| `cluster_wide_app_list.cluster_wide_apps` | [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd` | [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--default_port.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain.md#schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--local_domain) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info.md#schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info--decryption_provider) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info.md#schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info--location) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info.md#schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--blindfold_secret_info--store_provider) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--clear_secret_info.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--clear_secret_info.md#schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--clear_secret_info--provider_ref) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--clear_secret_info.md#schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--password--clear_secret_info--url) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain.md#schema-cluster_wide_app_list--cluster_wide_apps--argo_cd--local_domain--port) |
| `cluster_wide_app_list.cluster_wide_apps.dashboard` | [cluster_wide_app_list.cluster_wide_apps.dashboard](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--dashboard.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.metrics_server` | [cluster_wide_app_list.cluster_wide_apps.metrics_server](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--metrics_server.md#section) |
| `cluster_wide_app_list.cluster_wide_apps.prometheus` | [cluster_wide_app_list.cluster_wide_apps.prometheus](resources--k8s_cluster--properties--cluster_wide_app_list--cluster_wide_apps--prometheus.md#section) |
| `description` | [description](resources--k8s_cluster--reference.md#schema-description) |
| `disable` | [disable](resources--k8s_cluster--reference.md#schema-disable) |
| `global_access_enable` | [global_access_enable](resources--k8s_cluster--properties--global_access_enable.md#section) |
| `id` | [id](resources--k8s_cluster--reference.md#schema-id) |
| `insecure_registry_list` | [insecure_registry_list](resources--k8s_cluster--properties--insecure_registry_list.md#section) |
| `insecure_registry_list.insecure_registries` | [insecure_registry_list.insecure_registries](resources--k8s_cluster--properties--insecure_registry_list.md#schema-insecure_registry_list--insecure_registries) |
| `labels` | [labels](resources--k8s_cluster--reference.md#schema-labels) |
| `local_access_config` | [local_access_config](resources--k8s_cluster--properties--local_access_config.md#section) |
| `local_access_config.default_port` | [local_access_config.default_port](resources--k8s_cluster--properties--local_access_config--default_port.md#section) |
| `local_access_config.local_domain` | [local_access_config.local_domain](resources--k8s_cluster--properties--local_access_config.md#schema-local_access_config--local_domain) |
| `local_access_config.port` | [local_access_config.port](resources--k8s_cluster--properties--local_access_config.md#schema-local_access_config--port) |
| `name` | [name](resources--k8s_cluster--reference.md#schema-name) |
| `namespace` | [namespace](resources--k8s_cluster--reference.md#schema-namespace) |
| `no_cluster_wide_apps` | [no_cluster_wide_apps](resources--k8s_cluster--properties--no_cluster_wide_apps.md#section) |
| `no_global_access` | [no_global_access](resources--k8s_cluster--properties--no_global_access.md#section) |
| `no_insecure_registries` | [no_insecure_registries](resources--k8s_cluster--properties--no_insecure_registries.md#section) |
| `no_local_access` | [no_local_access](resources--k8s_cluster--properties--no_local_access.md#section) |
| `timeouts` | [timeouts](resources--k8s_cluster--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--k8s_cluster--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_cluster--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--k8s_cluster--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--k8s_cluster--properties--timeouts.md#schema-timeouts--update) |
| `use_custom_cluster_role_bindings` | [use_custom_cluster_role_bindings](resources--k8s_cluster--properties--use_custom_cluster_role_bindings.md#section) |
| `use_custom_cluster_role_bindings.cluster_role_bindings` | [use_custom_cluster_role_bindings.cluster_role_bindings](resources--k8s_cluster--properties--use_custom_cluster_role_bindings--cluster_role_bindings.md#section) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.name` | [use_custom_cluster_role_bindings.cluster_role_bindings.name](resources--k8s_cluster--properties--use_custom_cluster_role_bindings--cluster_role_bindings.md#schema-use_custom_cluster_role_bindings--cluster_role_bindings--name) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.namespace` | [use_custom_cluster_role_bindings.cluster_role_bindings.namespace](resources--k8s_cluster--properties--use_custom_cluster_role_bindings--cluster_role_bindings.md#schema-use_custom_cluster_role_bindings--cluster_role_bindings--namespace) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.tenant` | [use_custom_cluster_role_bindings.cluster_role_bindings.tenant](resources--k8s_cluster--properties--use_custom_cluster_role_bindings--cluster_role_bindings.md#schema-use_custom_cluster_role_bindings--cluster_role_bindings--tenant) |
| `use_custom_cluster_role_list` | [use_custom_cluster_role_list](resources--k8s_cluster--properties--use_custom_cluster_role_list.md#section) |
| `use_custom_cluster_role_list.cluster_roles` | [use_custom_cluster_role_list.cluster_roles](resources--k8s_cluster--properties--use_custom_cluster_role_list--cluster_roles.md#section) |
| `use_custom_cluster_role_list.cluster_roles.name` | [use_custom_cluster_role_list.cluster_roles.name](resources--k8s_cluster--properties--use_custom_cluster_role_list--cluster_roles.md#schema-use_custom_cluster_role_list--cluster_roles--name) |
| `use_custom_cluster_role_list.cluster_roles.namespace` | [use_custom_cluster_role_list.cluster_roles.namespace](resources--k8s_cluster--properties--use_custom_cluster_role_list--cluster_roles.md#schema-use_custom_cluster_role_list--cluster_roles--namespace) |
| `use_custom_cluster_role_list.cluster_roles.tenant` | [use_custom_cluster_role_list.cluster_roles.tenant](resources--k8s_cluster--properties--use_custom_cluster_role_list--cluster_roles.md#schema-use_custom_cluster_role_list--cluster_roles--tenant) |
| `use_custom_pod_security_admission` | [use_custom_pod_security_admission](resources--k8s_cluster--properties--use_custom_pod_security_admission.md#section) |
| `use_custom_pod_security_admission.name` | [use_custom_pod_security_admission.name](resources--k8s_cluster--properties--use_custom_pod_security_admission.md#schema-use_custom_pod_security_admission--name) |
| `use_custom_pod_security_admission.namespace` | [use_custom_pod_security_admission.namespace](resources--k8s_cluster--properties--use_custom_pod_security_admission.md#schema-use_custom_pod_security_admission--namespace) |
| `use_custom_pod_security_admission.tenant` | [use_custom_pod_security_admission.tenant](resources--k8s_cluster--properties--use_custom_pod_security_admission.md#schema-use_custom_pod_security_admission--tenant) |
| `use_custom_psp_list` | [use_custom_psp_list](resources--k8s_cluster--properties--use_custom_psp_list.md#section) |
| `use_custom_psp_list.pod_security_policies` | [use_custom_psp_list.pod_security_policies](resources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md#section) |
| `use_custom_psp_list.pod_security_policies.name` | [use_custom_psp_list.pod_security_policies.name](resources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md#schema-use_custom_psp_list--pod_security_policies--name) |
| `use_custom_psp_list.pod_security_policies.namespace` | [use_custom_psp_list.pod_security_policies.namespace](resources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md#schema-use_custom_psp_list--pod_security_policies--namespace) |
| `use_custom_psp_list.pod_security_policies.tenant` | [use_custom_psp_list.pod_security_policies.tenant](resources--k8s_cluster--properties--use_custom_psp_list--pod_security_policies.md#schema-use_custom_psp_list--pod_security_policies--tenant) |
| `use_default_cluster_role_bindings` | [use_default_cluster_role_bindings](resources--k8s_cluster--properties--use_default_cluster_role_bindings.md#section) |
| `use_default_cluster_roles` | [use_default_cluster_roles](resources--k8s_cluster--properties--use_default_cluster_roles.md#section) |
| `use_default_pod_security_admission` | [use_default_pod_security_admission](resources--k8s_cluster--properties--use_default_pod_security_admission.md#section) |
| `use_default_psp` | [use_default_psp](resources--k8s_cluster--properties--use_default_psp.md#section) |
| `vk8s_namespace_access_deny` | [vk8s_namespace_access_deny](resources--k8s_cluster--properties--vk8s_namespace_access_deny.md#section) |
| `vk8s_namespace_access_permit` | [vk8s_namespace_access_permit](resources--k8s_cluster--properties--vk8s_namespace_access_permit.md#section) |

## Next pages

- [cluster_scoped_access_deny](resources--k8s_cluster--properties--cluster_scoped_access_deny.md)
- [cluster_scoped_access_permit](resources--k8s_cluster--properties--cluster_scoped_access_permit.md)
- [cluster_wide_app_list](resources--k8s_cluster--properties--cluster_wide_app_list.md)
- [global_access_enable](resources--k8s_cluster--properties--global_access_enable.md)
- [insecure_registry_list](resources--k8s_cluster--properties--insecure_registry_list.md)
- [local_access_config](resources--k8s_cluster--properties--local_access_config.md)
- [no_cluster_wide_apps](resources--k8s_cluster--properties--no_cluster_wide_apps.md)
- [no_global_access](resources--k8s_cluster--properties--no_global_access.md)
- [no_insecure_registries](resources--k8s_cluster--properties--no_insecure_registries.md)
- [no_local_access](resources--k8s_cluster--properties--no_local_access.md)
- [timeouts](resources--k8s_cluster--properties--timeouts.md)
- [use_custom_cluster_role_bindings](resources--k8s_cluster--properties--use_custom_cluster_role_bindings.md)
- [use_custom_cluster_role_list](resources--k8s_cluster--properties--use_custom_cluster_role_list.md)
- [use_custom_pod_security_admission](resources--k8s_cluster--properties--use_custom_pod_security_admission.md)
- [use_custom_psp_list](resources--k8s_cluster--properties--use_custom_psp_list.md)
- [use_default_cluster_role_bindings](resources--k8s_cluster--properties--use_default_cluster_role_bindings.md)
- [use_default_cluster_roles](resources--k8s_cluster--properties--use_default_cluster_roles.md)
- [use_default_pod_security_admission](resources--k8s_cluster--properties--use_default_pod_security_admission.md)
- [use_default_psp](resources--k8s_cluster--properties--use_default_psp.md)
- [vk8s_namespace_access_deny](resources--k8s_cluster--properties--vk8s_namespace_access_deny.md)
- [vk8s_namespace_access_permit](resources--k8s_cluster--properties--vk8s_namespace_access_permit.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
