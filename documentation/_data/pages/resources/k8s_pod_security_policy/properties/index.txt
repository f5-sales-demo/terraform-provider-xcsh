---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": ["k8s pod security policy"], "body_bytes": 22953, "body_sha256": "sha256:aad0ab31b3cac9937b96866708ccb94143c101e502bc476bc0105bafe43d7143", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "xcsh-docs:resources:k8s_pod_security_policy:properties:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "parent_id": "xcsh-docs:resources:k8s_pod_security_policy:fundamentals", "path": "documentation/resources/k8s_pod_security_policy/properties/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333", "registry_path": "docs/guides/resources--k8s_pod_security_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["psp spec"], "anchor": "section", "description": "Form based pod security specification.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:allowed_capabilities,no_allowed_capabilities", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:allowed_capabilities", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:default_capabilities,no_default_capabilities", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:default_capabilities", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:drop_capabilities,no_drop_capabilities", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:drop_capabilities", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:fs_group_strategy_options,no_fs_groups", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:fs_group_strategy_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:allowed_capabilities,no_allowed_capabilities", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_allowed_capabilities", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:default_capabilities,no_default_capabilities", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_default_capabilities", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:drop_capabilities,no_drop_capabilities", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_drop_capabilities", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:fs_group_strategy_options,no_fs_groups", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_fs_groups", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:no_run_as_group,run_as_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_run_as_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:no_run_as_user,run_as_user", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_run_as_user", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:no_supplemental_groups,supplemental_groups", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:no_supplemental_groups", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:no_run_as_group,run_as_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:run_as_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:no_run_as_user,run_as_user", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:run_as_user", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "psp_spec:ConflictingObjectAttributes:no_supplemental_groups,supplemental_groups", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:psp_spec:supplemental_groups", "type": "conflicts"}], "schema_path": ["psp_spec"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["yaml"], "anchor": "schema-yaml", "description": "Exclusive with K8s YAML for Pod Security Policy.", "document_id": "xcsh-docs:resources:k8s_pod_security_policy:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["yaml"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_k8s_pod_security_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
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

Name of the K8S Pod Security Policy. Must be unique within the namespace.

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

Namespace where the K8S Pod Security Policy is created.

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

- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/timeouts/): complete subsection reference.

<a id="schema-yaml"></a>

### yaml property

Type: `"string"`. Optional, Computed.

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Upstream description:

Exclusive with \[psp\_spec\] K8s YAML for Pod Security Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

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
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-namespace) |
| `psp_spec` | [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#section) |
| `psp_spec.allow_privilege_escalation` | [psp_spec.allow_privilege_escalation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--allow_privilege_escalation) |
| `psp_spec.allowed_capabilities` | [psp_spec.allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_capabilities/#section) |
| `psp_spec.allowed_capabilities.capabilities` | [psp_spec.allowed_capabilities.capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_capabilities/#schema-psp_spec--allowed_capabilities--capabilities) |
| `psp_spec.allowed_csi_drivers` | [psp_spec.allowed_csi_drivers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--allowed_csi_drivers) |
| `psp_spec.allowed_flex_volumes` | [psp_spec.allowed_flex_volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--allowed_flex_volumes) |
| `psp_spec.allowed_host_paths` | [psp_spec.allowed_host_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_host_paths/#section) |
| `psp_spec.allowed_host_paths.path_prefix` | [psp_spec.allowed_host_paths.path_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_host_paths/#schema-psp_spec--allowed_host_paths--path_prefix) |
| `psp_spec.allowed_host_paths.read_only` | [psp_spec.allowed_host_paths.read_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/allowed_host_paths/#schema-psp_spec--allowed_host_paths--read_only) |
| `psp_spec.allowed_proc_mounts` | [psp_spec.allowed_proc_mounts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--allowed_proc_mounts) |
| `psp_spec.allowed_unsafe_sysctls` | [psp_spec.allowed_unsafe_sysctls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--allowed_unsafe_sysctls) |
| `psp_spec.default_allow_privilege_escalation` | [psp_spec.default_allow_privilege_escalation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--default_allow_privilege_escalation) |
| `psp_spec.default_capabilities` | [psp_spec.default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/default_capabilities/#section) |
| `psp_spec.default_capabilities.capabilities` | [psp_spec.default_capabilities.capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/default_capabilities/#schema-psp_spec--default_capabilities--capabilities) |
| `psp_spec.drop_capabilities` | [psp_spec.drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/drop_capabilities/#section) |
| `psp_spec.drop_capabilities.capabilities` | [psp_spec.drop_capabilities.capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/drop_capabilities/#schema-psp_spec--drop_capabilities--capabilities) |
| `psp_spec.forbidden_sysctls` | [psp_spec.forbidden_sysctls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--forbidden_sysctls) |
| `psp_spec.fs_group_strategy_options` | [psp_spec.fs_group_strategy_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/#section) |
| `psp_spec.fs_group_strategy_options.id_ranges` | [psp_spec.fs_group_strategy_options.id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/id_ranges/#section) |
| `psp_spec.fs_group_strategy_options.id_ranges.max_id` | [psp_spec.fs_group_strategy_options.id_ranges.max_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/id_ranges/#schema-psp_spec--fs_group_strategy_options--id_ranges--max_id) |
| `psp_spec.fs_group_strategy_options.id_ranges.min_id` | [psp_spec.fs_group_strategy_options.id_ranges.min_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/id_ranges/#schema-psp_spec--fs_group_strategy_options--id_ranges--min_id) |
| `psp_spec.fs_group_strategy_options.rule` | [psp_spec.fs_group_strategy_options.rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/fs_group_strategy_options/#schema-psp_spec--fs_group_strategy_options--rule) |
| `psp_spec.host_ipc` | [psp_spec.host_ipc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--host_ipc) |
| `psp_spec.host_network` | [psp_spec.host_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--host_network) |
| `psp_spec.host_pid` | [psp_spec.host_pid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--host_pid) |
| `psp_spec.host_port_ranges` | [psp_spec.host_port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--host_port_ranges) |
| `psp_spec.no_allowed_capabilities` | [psp_spec.no_allowed_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_allowed_capabilities/#section) |
| `psp_spec.no_default_capabilities` | [psp_spec.no_default_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_default_capabilities/#section) |
| `psp_spec.no_drop_capabilities` | [psp_spec.no_drop_capabilities](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_drop_capabilities/#section) |
| `psp_spec.no_fs_groups` | [psp_spec.no_fs_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_fs_groups/#section) |
| `psp_spec.no_run_as_group` | [psp_spec.no_run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_run_as_group/#section) |
| `psp_spec.no_run_as_user` | [psp_spec.no_run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_run_as_user/#section) |
| `psp_spec.no_runtime_class` | [psp_spec.no_runtime_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_runtime_class/#section) |
| `psp_spec.no_se_linux_options` | [psp_spec.no_se_linux_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_se_linux_options/#section) |
| `psp_spec.no_supplemental_groups` | [psp_spec.no_supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/no_supplemental_groups/#section) |
| `psp_spec.privileged` | [psp_spec.privileged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--privileged) |
| `psp_spec.read_only_root_filesystem` | [psp_spec.read_only_root_filesystem](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--read_only_root_filesystem) |
| `psp_spec.run_as_group` | [psp_spec.run_as_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_group/#section) |
| `psp_spec.run_as_group.id_ranges` | [psp_spec.run_as_group.id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_group/id_ranges/#section) |
| `psp_spec.run_as_group.id_ranges.max_id` | [psp_spec.run_as_group.id_ranges.max_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_group/id_ranges/#schema-psp_spec--run_as_group--id_ranges--max_id) |
| `psp_spec.run_as_group.id_ranges.min_id` | [psp_spec.run_as_group.id_ranges.min_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_group/id_ranges/#schema-psp_spec--run_as_group--id_ranges--min_id) |
| `psp_spec.run_as_group.rule` | [psp_spec.run_as_group.rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_group/#schema-psp_spec--run_as_group--rule) |
| `psp_spec.run_as_user` | [psp_spec.run_as_user](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_user/#section) |
| `psp_spec.run_as_user.id_ranges` | [psp_spec.run_as_user.id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_user/id_ranges/#section) |
| `psp_spec.run_as_user.id_ranges.max_id` | [psp_spec.run_as_user.id_ranges.max_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_user/id_ranges/#schema-psp_spec--run_as_user--id_ranges--max_id) |
| `psp_spec.run_as_user.id_ranges.min_id` | [psp_spec.run_as_user.id_ranges.min_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_user/id_ranges/#schema-psp_spec--run_as_user--id_ranges--min_id) |
| `psp_spec.run_as_user.rule` | [psp_spec.run_as_user.rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/run_as_user/#schema-psp_spec--run_as_user--rule) |
| `psp_spec.supplemental_groups` | [psp_spec.supplemental_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/#section) |
| `psp_spec.supplemental_groups.id_ranges` | [psp_spec.supplemental_groups.id_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/#section) |
| `psp_spec.supplemental_groups.id_ranges.max_id` | [psp_spec.supplemental_groups.id_ranges.max_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/#schema-psp_spec--supplemental_groups--id_ranges--max_id) |
| `psp_spec.supplemental_groups.id_ranges.min_id` | [psp_spec.supplemental_groups.id_ranges.min_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/id_ranges/#schema-psp_spec--supplemental_groups--id_ranges--min_id) |
| `psp_spec.supplemental_groups.rule` | [psp_spec.supplemental_groups.rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/supplemental_groups/#schema-psp_spec--supplemental_groups--rule) |
| `psp_spec.volumes` | [psp_spec.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/#schema-psp_spec--volumes) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/timeouts/#schema-timeouts--update) |
| `yaml` | [yaml](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/#schema-yaml) |

## Next pages

- [psp_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/psp_spec/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/timeouts/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/)
