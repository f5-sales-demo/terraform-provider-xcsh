---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 17567, "body_sha256": "sha256:2f217b4317d9e87d52c33f02810dfd481225a3f2dfba0ce8c05f5ccf466a3467", "canonical_id": "xcsh-docs:data-sources:k8s_pod_security_policy:reference", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec"], "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:reference", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:fundamentals", "path": "docs/guides/data-sources--k8s_pod_security_policy--reference.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
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

Description of the K8SPodSecurityPolicy.

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

Name of the K8SPodSecurityPolicy.

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

Namespace where the K8SPodSecurityPolicy exists.

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

- [psp_spec](data-sources--k8s_pod_security_policy--properties--psp_spec.md): complete subsection reference.

<a id="schema-yaml"></a>

### yaml property

Type: `"string"`. Computed.

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Upstream description:

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_pod_security_policy--reference.md#schema-annotations) |
| `description` | [description](data-sources--k8s_pod_security_policy--reference.md#schema-description) |
| `id` | [id](data-sources--k8s_pod_security_policy--reference.md#schema-id) |
| `labels` | [labels](data-sources--k8s_pod_security_policy--reference.md#schema-labels) |
| `name` | [name](data-sources--k8s_pod_security_policy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--k8s_pod_security_policy--reference.md#schema-namespace) |
| `psp_spec` | [psp_spec](data-sources--k8s_pod_security_policy--properties--psp_spec.md#section) |
| `psp_spec.allow_privilege_escalation` | [psp_spec.allow_privilege_escalation](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--allow_privilege_escalation) |
| `psp_spec.allowed_capabilities` | [psp_spec.allowed_capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--allowed_capabilities.md#section) |
| `psp_spec.allowed_capabilities.capabilities` | [psp_spec.allowed_capabilities.capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--allowed_capabilities.md#schema-psp_spec--allowed_capabilities--capabilities) |
| `psp_spec.allowed_csi_drivers` | [psp_spec.allowed_csi_drivers](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--allowed_csi_drivers) |
| `psp_spec.allowed_flex_volumes` | [psp_spec.allowed_flex_volumes](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--allowed_flex_volumes) |
| `psp_spec.allowed_host_paths` | [psp_spec.allowed_host_paths](data-sources--k8s_pod_security_policy--properties--psp_spec--allowed_host_paths.md#section) |
| `psp_spec.allowed_host_paths.path_prefix` | [psp_spec.allowed_host_paths.path_prefix](data-sources--k8s_pod_security_policy--properties--psp_spec--allowed_host_paths.md#schema-psp_spec--allowed_host_paths--path_prefix) |
| `psp_spec.allowed_host_paths.read_only` | [psp_spec.allowed_host_paths.read_only](data-sources--k8s_pod_security_policy--properties--psp_spec--allowed_host_paths.md#schema-psp_spec--allowed_host_paths--read_only) |
| `psp_spec.allowed_proc_mounts` | [psp_spec.allowed_proc_mounts](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--allowed_proc_mounts) |
| `psp_spec.allowed_unsafe_sysctls` | [psp_spec.allowed_unsafe_sysctls](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--allowed_unsafe_sysctls) |
| `psp_spec.default_allow_privilege_escalation` | [psp_spec.default_allow_privilege_escalation](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--default_allow_privilege_escalation) |
| `psp_spec.default_capabilities` | [psp_spec.default_capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--default_capabilities.md#section) |
| `psp_spec.default_capabilities.capabilities` | [psp_spec.default_capabilities.capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--default_capabilities.md#schema-psp_spec--default_capabilities--capabilities) |
| `psp_spec.drop_capabilities` | [psp_spec.drop_capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--drop_capabilities.md#section) |
| `psp_spec.drop_capabilities.capabilities` | [psp_spec.drop_capabilities.capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--drop_capabilities.md#schema-psp_spec--drop_capabilities--capabilities) |
| `psp_spec.forbidden_sysctls` | [psp_spec.forbidden_sysctls](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--forbidden_sysctls) |
| `psp_spec.fs_group_strategy_options` | [psp_spec.fs_group_strategy_options](data-sources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options.md#section) |
| `psp_spec.fs_group_strategy_options.id_ranges` | [psp_spec.fs_group_strategy_options.id_ranges](data-sources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options--id_ranges.md#section) |
| `psp_spec.fs_group_strategy_options.id_ranges.max_id` | [psp_spec.fs_group_strategy_options.id_ranges.max_id](data-sources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options--id_ranges.md#schema-psp_spec--fs_group_strategy_options--id_ranges--max_id) |
| `psp_spec.fs_group_strategy_options.id_ranges.min_id` | [psp_spec.fs_group_strategy_options.id_ranges.min_id](data-sources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options--id_ranges.md#schema-psp_spec--fs_group_strategy_options--id_ranges--min_id) |
| `psp_spec.fs_group_strategy_options.rule` | [psp_spec.fs_group_strategy_options.rule](data-sources--k8s_pod_security_policy--properties--psp_spec--fs_group_strategy_options.md#schema-psp_spec--fs_group_strategy_options--rule) |
| `psp_spec.host_ipc` | [psp_spec.host_ipc](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--host_ipc) |
| `psp_spec.host_network` | [psp_spec.host_network](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--host_network) |
| `psp_spec.host_pid` | [psp_spec.host_pid](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--host_pid) |
| `psp_spec.host_port_ranges` | [psp_spec.host_port_ranges](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--host_port_ranges) |
| `psp_spec.no_allowed_capabilities` | [psp_spec.no_allowed_capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--no_allowed_capabilities.md#section) |
| `psp_spec.no_default_capabilities` | [psp_spec.no_default_capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--no_default_capabilities.md#section) |
| `psp_spec.no_drop_capabilities` | [psp_spec.no_drop_capabilities](data-sources--k8s_pod_security_policy--properties--psp_spec--no_drop_capabilities.md#section) |
| `psp_spec.no_fs_groups` | [psp_spec.no_fs_groups](data-sources--k8s_pod_security_policy--properties--psp_spec--no_fs_groups.md#section) |
| `psp_spec.no_run_as_group` | [psp_spec.no_run_as_group](data-sources--k8s_pod_security_policy--properties--psp_spec--no_run_as_group.md#section) |
| `psp_spec.no_run_as_user` | [psp_spec.no_run_as_user](data-sources--k8s_pod_security_policy--properties--psp_spec--no_run_as_user.md#section) |
| `psp_spec.no_runtime_class` | [psp_spec.no_runtime_class](data-sources--k8s_pod_security_policy--properties--psp_spec--no_runtime_class.md#section) |
| `psp_spec.no_se_linux_options` | [psp_spec.no_se_linux_options](data-sources--k8s_pod_security_policy--properties--psp_spec--no_se_linux_options.md#section) |
| `psp_spec.no_supplemental_groups` | [psp_spec.no_supplemental_groups](data-sources--k8s_pod_security_policy--properties--psp_spec--no_supplemental_groups.md#section) |
| `psp_spec.privileged` | [psp_spec.privileged](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--privileged) |
| `psp_spec.read_only_root_filesystem` | [psp_spec.read_only_root_filesystem](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--read_only_root_filesystem) |
| `psp_spec.run_as_group` | [psp_spec.run_as_group](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_group.md#section) |
| `psp_spec.run_as_group.id_ranges` | [psp_spec.run_as_group.id_ranges](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_group--id_ranges.md#section) |
| `psp_spec.run_as_group.id_ranges.max_id` | [psp_spec.run_as_group.id_ranges.max_id](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_group--id_ranges.md#schema-psp_spec--run_as_group--id_ranges--max_id) |
| `psp_spec.run_as_group.id_ranges.min_id` | [psp_spec.run_as_group.id_ranges.min_id](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_group--id_ranges.md#schema-psp_spec--run_as_group--id_ranges--min_id) |
| `psp_spec.run_as_group.rule` | [psp_spec.run_as_group.rule](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_group.md#schema-psp_spec--run_as_group--rule) |
| `psp_spec.run_as_user` | [psp_spec.run_as_user](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user.md#section) |
| `psp_spec.run_as_user.id_ranges` | [psp_spec.run_as_user.id_ranges](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user--id_ranges.md#section) |
| `psp_spec.run_as_user.id_ranges.max_id` | [psp_spec.run_as_user.id_ranges.max_id](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user--id_ranges.md#schema-psp_spec--run_as_user--id_ranges--max_id) |
| `psp_spec.run_as_user.id_ranges.min_id` | [psp_spec.run_as_user.id_ranges.min_id](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user--id_ranges.md#schema-psp_spec--run_as_user--id_ranges--min_id) |
| `psp_spec.run_as_user.rule` | [psp_spec.run_as_user.rule](data-sources--k8s_pod_security_policy--properties--psp_spec--run_as_user.md#schema-psp_spec--run_as_user--rule) |
| `psp_spec.supplemental_groups` | [psp_spec.supplemental_groups](data-sources--k8s_pod_security_policy--properties--psp_spec--supplemental_groups.md#section) |
| `psp_spec.supplemental_groups.id_ranges` | [psp_spec.supplemental_groups.id_ranges](data-sources--k8s_pod_security_policy--properties--psp_spec--supplemental_groups--id_ranges.md#section) |
| `psp_spec.supplemental_groups.id_ranges.max_id` | [psp_spec.supplemental_groups.id_ranges.max_id](data-sources--k8s_pod_security_policy--properties--psp_spec--supplemental_groups--id_ranges.md#schema-psp_spec--supplemental_groups--id_ranges--max_id) |
| `psp_spec.supplemental_groups.id_ranges.min_id` | [psp_spec.supplemental_groups.id_ranges.min_id](data-sources--k8s_pod_security_policy--properties--psp_spec--supplemental_groups--id_ranges.md#schema-psp_spec--supplemental_groups--id_ranges--min_id) |
| `psp_spec.supplemental_groups.rule` | [psp_spec.supplemental_groups.rule](data-sources--k8s_pod_security_policy--properties--psp_spec--supplemental_groups.md#schema-psp_spec--supplemental_groups--rule) |
| `psp_spec.volumes` | [psp_spec.volumes](data-sources--k8s_pod_security_policy--properties--psp_spec.md#schema-psp_spec--volumes) |
| `yaml` | [yaml](data-sources--k8s_pod_security_policy--reference.md#schema-yaml) |

## Next pages

- [psp_spec](data-sources--k8s_pod_security_policy--properties--psp_spec.md)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
