---
page_title: "psp_spec"
subcategory: ""
description: "Form based pod security specification."
xcsh_docs: {"aliases": ["psp spec"], "body_bytes": 14823, "body_sha256": "sha256:0b04b5df51e9ab6caf5dd36c464b3679d507a2de172221e18bd83e144689058d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:allowed_capabilities", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:allowed_host_paths", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:default_capabilities", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:drop_capabilities", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_allowed_capabilities", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_default_capabilities", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_drop_capabilities", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_fs_groups", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_run_as_group", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_run_as_user", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_runtime_class", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_se_linux_options", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_supplemental_groups", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_group", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:reference", "path": "documentation/data-sources/k8s_pod_security_policy/properties/psp_spec/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2332030332013022-0112203031313231-0233321210232300-0213323020102212-2031122211212030-1231323030201220-2101210231031020-2020223333122033", "registry_path": "docs/guides/data-sources--k8s_pod_security_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["psp_spec"], "schema_version": 1, "sections": [{"aliases": ["psp spec allow privilege escalation"], "anchor": "schema-psp_spec--allow_privilege_escalation", "description": "Pod can request to privilege escalation.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "allow_privilege_escalation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["psp spec allowed capabilities"], "anchor": "section", "description": "List of capabilities that docker container has.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:allowed_capabilities", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "allowed_capabilities"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec allowed csi drivers"], "anchor": "schema-psp_spec--allowed_csi_drivers", "description": "Restrict the available CSI drivers for POD, default all drivers are available.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "allowed_csi_drivers"], "syntax": "attribute", "type": "list"}, {"aliases": ["psp spec allowed flex volumes"], "anchor": "schema-psp_spec--allowed_flex_volumes", "description": "Restrict list of Flex volumes, default all volumes are allowed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "allowed_flex_volumes"], "syntax": "attribute", "type": "list"}, {"aliases": ["psp spec allowed host paths"], "anchor": "section", "description": "Restrict list of host paths, default all host paths are allowed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:allowed_host_paths", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["psp_spec", "allowed_host_paths"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec allowed proc mounts"], "anchor": "schema-psp_spec--allowed_proc_mounts", "description": "Allowed list of proc mounts, empty list allows default proc mounts.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "allowed_proc_mounts"], "syntax": "attribute", "type": "list"}, {"aliases": ["psp spec allowed unsafe sysctls"], "anchor": "schema-psp_spec--allowed_unsafe_sysctls", "description": "Allowed list of unsafe sysctls, empty list allows none. Supports prefix reg-ex.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "allowed_unsafe_sysctls"], "syntax": "attribute", "type": "list"}, {"aliases": ["psp spec default allow privilege escalation"], "anchor": "schema-psp_spec--default_allow_privilege_escalation", "description": "Pod has permission for privilege escalation by default.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "default_allow_privilege_escalation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["psp spec default capabilities"], "anchor": "section", "description": "List of capabilities that docker container has.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:default_capabilities", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "default_capabilities"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec drop capabilities"], "anchor": "section", "description": "List of capabilities that docker container has.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:drop_capabilities", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "drop_capabilities"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec forbidden sysctls"], "anchor": "schema-psp_spec--forbidden_sysctls", "description": "Forbidden list of sysctls, empty list forbids none. Supports prefix reg-ex.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "forbidden_sysctls"], "syntax": "attribute", "type": "list"}, {"aliases": ["psp spec fs group strategy options"], "anchor": "section", "description": "ID ranges and rules.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "fs_group_strategy_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec host ipc"], "anchor": "schema-psp_spec--host_ipc", "description": "Host IPC determines if the policy allows the use of host IPC in the pod spec.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "host_ipc"], "syntax": "attribute", "type": "bool"}, {"aliases": ["psp spec host network"], "anchor": "schema-psp_spec--host_network", "description": "Host Network determines if the policy allows the use of host network in the pod spec.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "host_network"], "syntax": "attribute", "type": "bool"}, {"aliases": ["psp spec host pid"], "anchor": "schema-psp_spec--host_pid", "description": "Host PID determines if the policy allows the use of host PID in the pod spec.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "host_pid"], "syntax": "attribute", "type": "bool"}, {"aliases": ["psp spec host port ranges"], "anchor": "schema-psp_spec--host_port_ranges", "description": "Host port ranges determines which ports ranges are allowed to be exposed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "host_port_ranges"], "syntax": "attribute", "type": "string"}, {"aliases": ["psp spec no allowed capabilities"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_allowed_capabilities", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "no_allowed_capabilities"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no default capabilities"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_default_capabilities", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "no_default_capabilities"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no drop capabilities"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_drop_capabilities", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "no_drop_capabilities"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no fs groups"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_fs_groups", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "no_fs_groups"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no run as group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_run_as_group", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "no_run_as_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no run as user"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_run_as_user", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "no_run_as_user"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no runtime class"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_runtime_class", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "no_runtime_class"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no se linux options"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_se_linux_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "no_se_linux_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec no supplemental groups"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:no_supplemental_groups", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "no_supplemental_groups"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec privileged"], "anchor": "schema-psp_spec--privileged", "description": "Privileged determines if a pod can request to be run as privileged.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "privileged"], "syntax": "attribute", "type": "bool"}, {"aliases": ["psp spec read only root filesystem"], "anchor": "schema-psp_spec--read_only_root_filesystem", "description": "Containers can only run with read only root filesystem.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "read_only_root_filesystem"], "syntax": "attribute", "type": "bool"}, {"aliases": ["psp spec run as group"], "anchor": "section", "description": "ID ranges and rules.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_group", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "run_as_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec run as user"], "anchor": "section", "description": "ID ranges and rules.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "run_as_user"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec supplemental groups"], "anchor": "section", "description": "ID ranges and rules.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["psp_spec", "supplemental_groups"], "syntax": "attribute", "type": "object"}, {"aliases": ["psp spec volumes"], "anchor": "schema-psp_spec--volumes", "description": "Allow List of volume plugins. Empty no volumes are allowed.", "document_id": "xcsh-docs:data-sources:k8s_pod_security_policy:properties:psp_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["psp_spec", "volumes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/properties/psp_spec/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Form based pod security specification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# psp_spec

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/)
- psp_spec

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: psp\_spec, yaml\] Pod Security Policy Specification. Form based pod security specification.

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

- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/#section)
- [yaml](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/#schema-yaml)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-psp_spec--allow_privilege_escalation"></a>

### allow_privilege_escalation property

Type: `"bool"`. Computed.

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

- [allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/allowed_capabilities/): complete subsection reference.

<a id="schema-psp_spec--allowed_csi_drivers"></a>

### allowed_csi_drivers property

Type: `["list", "string"]`. Computed.

Restrict the available CSI drivers for POD, default all drivers are available.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `["list", "string"]`. Computed.

Restrict list of Flex volumes, default all volumes are allowed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [allowed_host_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/allowed_host_paths/): complete subsection reference.

<a id="schema-psp_spec--allowed_proc_mounts"></a>

### allowed_proc_mounts property

Type: `["list", "string"]`. Computed.

Allowed list of proc mounts, empty list allows default proc mounts.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `["list", "string"]`. Computed.

Allowed list of unsafe sysctls, empty list allows none. Supports prefix reg-ex.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Type: `"bool"`. Computed.

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

- [default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/default_capabilities/): complete subsection reference.

- [drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/drop_capabilities/): complete subsection reference.

<a id="schema-psp_spec--forbidden_sysctls"></a>

### forbidden_sysctls property

Type: `["list", "string"]`. Computed.

Forbidden list of sysctls, empty list forbids none. Supports prefix reg-ex.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [fs_group_strategy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/): complete subsection reference.

<a id="schema-psp_spec--host_ipc"></a>

### host_ipc property

Type: `"bool"`. Computed.

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

Type: `"bool"`. Computed.

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

Type: `"bool"`. Computed.

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

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [no_allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_allowed_capabilities/): complete subsection reference.

- [no_default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_default_capabilities/): complete subsection reference.

- [no_drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_drop_capabilities/): complete subsection reference.

- [no_fs_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_fs_groups/): complete subsection reference.

- [no_run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_run_as_group/): complete subsection reference.

- [no_run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_run_as_user/): complete subsection reference.

- [no_runtime_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_runtime_class/): complete subsection reference.

- [no_se_linux_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_se_linux_options/): complete subsection reference.

- [no_supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/no_supplemental_groups/): complete subsection reference.

<a id="schema-psp_spec--privileged"></a>

### privileged property

Type: `"bool"`. Computed.

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

Type: `"bool"`. Computed.

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

- [run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_group/): complete subsection reference.

- [run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/run_as_user/): complete subsection reference.

- [supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/): complete subsection reference.

<a id="schema-psp_spec--volumes"></a>

### volumes property

Type: `["list", "string"]`. Computed.

Allow List of volume plugins. Empty no volumes are allowed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
