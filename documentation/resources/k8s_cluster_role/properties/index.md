---
page_title: "Property reference"
subcategory: "Container"
description: "Property reference for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 14512, "body_sha256": "sha256:d78deb5b852e79b241e2b604ea24954c9cccfa64433fea3e059a14ada585c3b8", "child_ids": ["xcsh-docs:resources:k8s_cluster_role:properties:k8s_cluster_role_selector", "xcsh-docs:resources:k8s_cluster_role:properties:policy_rule_list", "xcsh-docs:resources:k8s_cluster_role:properties:timeouts"], "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:reference", "parent_id": "xcsh-docs:resources:k8s_cluster_role:fundamentals", "path": "documentation/resources/k8s_cluster_role/properties/index.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
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

- [k8s_cluster_role_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/k8s_cluster_role_selector/): complete subsection reference.

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

Name of the K8S Cluster Role. Must be unique within the namespace.

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

Namespace for the K8S Cluster Role. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

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

- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/timeouts/): complete subsection reference.

<a id="schema-yaml"></a>

### yaml property

Type: `"string"`. Optional, Computed.

Exclusive with \[k8s\_cluster\_role\_selector policy\_rule\_list\] K8s YAML for ClusterRole.

Upstream description:

Exclusive with \[k8s\_cluster\_role\_selector policy\_rule\_list\] K8s YAML for ClusterRole.

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
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-id) |
| `k8s_cluster_role_selector` | [k8s_cluster_role_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/k8s_cluster_role_selector/#section) |
| `k8s_cluster_role_selector.expressions` | [k8s_cluster_role_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/k8s_cluster_role_selector/#schema-k8s_cluster_role_selector--expressions) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-namespace) |
| `policy_rule_list` | [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/#section) |
| `policy_rule_list.policy_rule` | [policy_rule_list.policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/#section) |
| `policy_rule_list.policy_rule.non_resource_url_list` | [policy_rule_list.policy_rule.non_resource_url_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/#section) |
| `policy_rule_list.policy_rule.non_resource_url_list.urls` | [policy_rule_list.policy_rule.non_resource_url_list.urls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/#schema-policy_rule_list--policy_rule--non_resource_url_list--urls) |
| `policy_rule_list.policy_rule.non_resource_url_list.verbs` | [policy_rule_list.policy_rule.non_resource_url_list.verbs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/non_resource_url_list/#schema-policy_rule_list--policy_rule--non_resource_url_list--verbs) |
| `policy_rule_list.policy_rule.resource_list` | [policy_rule_list.policy_rule.resource_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/#section) |
| `policy_rule_list.policy_rule.resource_list.api_groups` | [policy_rule_list.policy_rule.resource_list.api_groups](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/#schema-policy_rule_list--policy_rule--resource_list--api_groups) |
| `policy_rule_list.policy_rule.resource_list.resource_instances` | [policy_rule_list.policy_rule.resource_list.resource_instances](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/#schema-policy_rule_list--policy_rule--resource_list--resource_instances) |
| `policy_rule_list.policy_rule.resource_list.resource_types` | [policy_rule_list.policy_rule.resource_list.resource_types](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/#schema-policy_rule_list--policy_rule--resource_list--resource_types) |
| `policy_rule_list.policy_rule.resource_list.verbs` | [policy_rule_list.policy_rule.resource_list.verbs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/policy_rule/resource_list/#schema-policy_rule_list--policy_rule--resource_list--verbs) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/timeouts/#schema-timeouts--update) |
| `yaml` | [yaml](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/#schema-yaml) |

## Next pages

- [k8s_cluster_role_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/k8s_cluster_role_selector/)
- [policy_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/policy_rule_list/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/timeouts/)
- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
