---
page_title: "psp_spec"
subcategory: ""
description: "psp_spec for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 19487, "body_sha256": "sha256:03dbaf8b24e4123f4bcc276c7523a409d78c57779cf42d4234dc19010388fc55", "child_ids": ["xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:allowed_capabilities", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:allowed_host_paths", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:default_capabilities", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:drop_capabilities", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_allowed_capabilities", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_default_capabilities", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_drop_capabilities", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_fs_groups", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_run_as_group", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_run_as_user", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_runtime_class", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_se_linux_options", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_supplemental_groups", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:run_as_group", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups"], "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "path": "documentation/resources/k8s_pod_security_policy/properties/psp_spec/index.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["psp_spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/psp_spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "psp_spec for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/)
- psp_spec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: psp\_spec, yaml\] Pod Security Policy Specification. Form based pod security specification.

Upstream description:

Form based pod security specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allowed_capabilities",
    "no_allowed_capabilities"),
  validators.ConflictingObjectAttributes("default_capabilities",
    "no_default_capabilities"),
  validators.ConflictingObjectAttributes("drop_capabilities",
    "no_drop_capabilities"),
  validators.ConflictingObjectAttributes("fs_group_strategy_options",
    "no_fs_groups"),
  validators.ConflictingObjectAttributes("no_run_as_group",
    "run_as_group"),
  validators.ConflictingObjectAttributes("no_run_as_user",
    "run_as_user"),
  validators.ConflictingObjectAttributes("no_supplemental_groups",
    "supplemental_groups")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_capabilities_choice": "[\"allowed_capabilities\",\"no_allowed_capabilities\"]",
  "x-ves-oneof-field-default_capabilities_choice": "[\"default_capabilities\",\"no_default_capabilities\"]",
  "x-ves-oneof-field-drop_capabilities_choice": "[\"drop_capabilities\",\"no_drop_capabilities\"]",
  "x-ves-oneof-field-fs_group_choice": "[\"fs_group_strategy_options\",\"no_fs_groups\"]",
  "x-ves-oneof-field-group_choice": "[\"no_run_as_group\",\"run_as_group\"]",
  "x-ves-oneof-field-runtime_class_choice": "[\"no_runtime_class\"]",
  "x-ves-oneof-field-se_linux_choice": "[\"no_se_linux_options\"]",
  "x-ves-oneof-field-supplemental_group_choice": "[\"no_supplemental_groups\",\"supplemental_groups\"]",
  "x-ves-oneof-field-user_choice": "[\"no_run_as_user\",\"run_as_user\"]"
}
```

OneOf alternatives in this subsection:

- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#section)
- [yaml](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-yaml)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
psp_spec {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-psp_spec--allow_privilege_escalation"></a>

### allow_privilege_escalation property

Type: `"bool"`. Optional.

Pod can request to privilege escalation.

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

- [allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_capabilities/): complete subsection reference.

<a id="schema-psp_spec--allowed_csi_drivers"></a>

### allowed_csi_drivers property

Type: `["list", "string"]`. Optional.

Restrict the available CSI drivers for POD, default all drivers are available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-psp_spec--allowed_flex_volumes"></a>

### allowed_flex_volumes property

Type: `["list", "string"]`. Optional.

Restrict list of Flex volumes, default all volumes are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [allowed_host_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_host_paths/): complete subsection reference.

<a id="schema-psp_spec--allowed_proc_mounts"></a>

### allowed_proc_mounts property

Type: `["list", "string"]`. Optional.

Allowed list of proc mounts, empty list allows default proc mounts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-psp_spec--allowed_unsafe_sysctls"></a>

### allowed_unsafe_sysctls property

Type: `["list", "string"]`. Optional.

Allowed list of unsafe sysctls, empty list allows none. Supports prefix reg-ex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-psp_spec--default_allow_privilege_escalation"></a>

### default_allow_privilege_escalation property

Type: `"bool"`. Optional.

Pod has permission for privilege escalation by default.

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

- [default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/default_capabilities/): complete subsection reference.

- [drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/drop_capabilities/): complete subsection reference.

<a id="schema-psp_spec--forbidden_sysctls"></a>

### forbidden_sysctls property

Type: `["list", "string"]`. Optional.

Forbidden list of sysctls, empty list forbids none. Supports prefix reg-ex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [fs_group_strategy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/): complete subsection reference.

<a id="schema-psp_spec--host_ipc"></a>

### host_ipc property

Type: `"bool"`. Optional.

Host IPC determines if the policy allows the use of host IPC in the pod spec.

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

<a id="schema-psp_spec--host_network"></a>

### host_network property

Type: `"bool"`. Optional.

Host Network determines if the policy allows the use of host network in the pod spec.

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

<a id="schema-psp_spec--host_pid"></a>

### host_pid property

Type: `"bool"`. Optional.

Host PID determines if the policy allows the use of host PID in the pod spec.

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

<a id="schema-psp_spec--host_port_ranges"></a>

### host_port_ranges property

Type: `"string"`. Optional.

Host port ranges determines which ports ranges are allowed to be exposed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

- [no_allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_allowed_capabilities/): complete subsection reference.

- [no_default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_default_capabilities/): complete subsection reference.

- [no_drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_drop_capabilities/): complete subsection reference.

- [no_fs_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_fs_groups/): complete subsection reference.

- [no_run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_run_as_group/): complete subsection reference.

- [no_run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_run_as_user/): complete subsection reference.

- [no_runtime_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_runtime_class/): complete subsection reference.

- [no_se_linux_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_se_linux_options/): complete subsection reference.

- [no_supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_supplemental_groups/): complete subsection reference.

<a id="schema-psp_spec--privileged"></a>

### privileged property

Type: `"bool"`. Optional.

Privileged determines if a pod can request to be run as privileged.

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

<a id="schema-psp_spec--read_only_root_filesystem"></a>

### read_only_root_filesystem property

Type: `"bool"`. Optional.

Containers can only run with read only root filesystem.

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

- [run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_group/): complete subsection reference.

- [run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_user/): complete subsection reference.

- [supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/): complete subsection reference.

<a id="schema-psp_spec--volumes"></a>

### volumes property

Type: `["list", "string"]`. Optional.

Allow List of volume plugins. Empty no volumes are allowed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [psp_spec.allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_capabilities/)
- [psp_spec.allowed_host_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_host_paths/)
- [psp_spec.default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/default_capabilities/)
- [psp_spec.drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/drop_capabilities/)
- [psp_spec.fs_group_strategy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/)
- [psp_spec.no_allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_allowed_capabilities/)
- [psp_spec.no_default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_default_capabilities/)
- [psp_spec.no_drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_drop_capabilities/)
- [psp_spec.no_fs_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_fs_groups/)
- [psp_spec.no_run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_run_as_group/)
- [psp_spec.no_run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_run_as_user/)
- [psp_spec.no_runtime_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_runtime_class/)
- [psp_spec.no_se_linux_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_se_linux_options/)
- [psp_spec.no_supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_supplemental_groups/)
- [psp_spec.run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_group/)
- [psp_spec.run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_user/)
- [psp_spec.supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
