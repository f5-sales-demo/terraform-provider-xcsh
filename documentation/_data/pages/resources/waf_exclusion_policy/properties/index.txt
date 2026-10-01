---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_waf_exclusion_policy."
xcsh_docs: {"aliases": [], "body_bytes": 18717, "body_sha256": "sha256:48ef631a651191c6ba07b1acefeb47bf24716337e1c014af115abbfb62813a99", "child_ids": ["xcsh-docs:resources:waf_exclusion_policy:properties:timeouts", "xcsh-docs:resources:waf_exclusion_policy:properties:waf_exclusion_rules"], "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:waf_exclusion_policy:reference", "parent_id": "xcsh-docs:resources:waf_exclusion_policy:fundamentals", "path": "documentation/resources/waf_exclusion_policy/properties/index.md", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_waf_exclusion_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
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

Name of the WAF Exclusion Policy. Must be unique within the namespace.

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

Namespace where the WAF Exclusion Policy is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/timeouts/): complete subsection reference.

- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/timeouts/#schema-timeouts--update) |
| `waf_exclusion_rules` | [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/#section) |
| `waf_exclusion_rules.any_domain` | [waf_exclusion_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/any_domain/#section) |
| `waf_exclusion_rules.any_path` | [waf_exclusion_rules.any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/any_path/#section) |
| `waf_exclusion_rules.app_firewall_detection_control` | [waf_exclusion_rules.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/#section) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/#section) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_attack_type_contexts--context) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_attack_type_contexts--context_name) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_exclusion_rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_attack_type_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_attack_type_contexts--exclude_attack_type) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/#section) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_bot_name_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_bot_name_contexts--bot_name) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/#section) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--context) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--context_name) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts.signature_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--signature_id) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/#section) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_violation_contexts--context) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.context_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_violation_contexts--context_name) |
| `waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_exclusion_rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_violation_contexts/#schema-waf_exclusion_rules--app_firewall_detection_control--exclude_violation_contexts--exclude_violation) |
| `waf_exclusion_rules.exact_value` | [waf_exclusion_rules.exact_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/#schema-waf_exclusion_rules--exact_value) |
| `waf_exclusion_rules.expiration_timestamp` | [waf_exclusion_rules.expiration_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/#schema-waf_exclusion_rules--expiration_timestamp) |
| `waf_exclusion_rules.metadata` | [waf_exclusion_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/metadata/#section) |
| `waf_exclusion_rules.metadata.description_spec` | [waf_exclusion_rules.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/metadata/#schema-waf_exclusion_rules--metadata--description_spec) |
| `waf_exclusion_rules.metadata.name` | [waf_exclusion_rules.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/metadata/#schema-waf_exclusion_rules--metadata--name) |
| `waf_exclusion_rules.methods` | [waf_exclusion_rules.methods](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/#schema-waf_exclusion_rules--methods) |
| `waf_exclusion_rules.path_prefix` | [waf_exclusion_rules.path_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/#schema-waf_exclusion_rules--path_prefix) |
| `waf_exclusion_rules.path_regex` | [waf_exclusion_rules.path_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/#schema-waf_exclusion_rules--path_regex) |
| `waf_exclusion_rules.suffix_value` | [waf_exclusion_rules.suffix_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/#schema-waf_exclusion_rules--suffix_value) |
| `waf_exclusion_rules.waf_skip_processing` | [waf_exclusion_rules.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/waf_skip_processing/#section) |

## Next pages

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/timeouts/)
- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/waf_exclusion_rules/)
- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/)
