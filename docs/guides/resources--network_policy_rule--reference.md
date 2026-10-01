---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 14587, "body_sha256": "sha256:445d999b2d2afed53b7fc0863bb76074366a92c76bb862ec62c23cb03d6dab3b", "canonical_id": "xcsh-docs:resources:network_policy_rule:reference", "child_ids": ["xcsh-docs:resources:network_policy_rule:properties:advanced_action", "xcsh-docs:resources:network_policy_rule:properties:ip_prefix_set", "xcsh-docs:resources:network_policy_rule:properties:label_matcher", "xcsh-docs:resources:network_policy_rule:properties:prefix", "xcsh-docs:resources:network_policy_rule:properties:prefix_selector", "xcsh-docs:resources:network_policy_rule:properties:timeouts"], "collection_id": "xcsh-docs:resources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_rule:reference", "parent_id": "xcsh-docs:resources:network_policy_rule:fundamentals", "path": "docs/guides/resources--network_policy_rule--reference.md", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_rule/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md)
- Property reference

## Direct properties

<a id="schema-action"></a>

### action property

Type: `"string"`. Optional, Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [advanced_action](resources--network_policy_rule--properties--advanced_action.md): complete subsection reference.

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

- [ip_prefix_set](resources--network_policy_rule--properties--ip_prefix_set.md): complete subsection reference.

- [label_matcher](resources--network_policy_rule--properties--label_matcher.md): complete subsection reference.

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

Name of the Network Policy Rule. Must be unique within the namespace.

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

Namespace where the Network Policy Rule is created.

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

<a id="schema-ports"></a>

### ports property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [prefix](resources--network_policy_rule--properties--prefix.md): complete subsection reference.

- [prefix_selector](resources--network_policy_rule--properties--prefix_selector.md): complete subsection reference.

<a id="schema-protocol"></a>

### protocol property

Type: `"string"`. Optional, Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

- [timeouts](resources--network_policy_rule--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--network_policy_rule--reference.md#schema-action) |
| `advanced_action` | [advanced_action](resources--network_policy_rule--properties--advanced_action.md#section) |
| `advanced_action.action` | [advanced_action.action](resources--network_policy_rule--properties--advanced_action.md#schema-advanced_action--action) |
| `annotations` | [annotations](resources--network_policy_rule--reference.md#schema-annotations) |
| `description` | [description](resources--network_policy_rule--reference.md#schema-description) |
| `disable` | [disable](resources--network_policy_rule--reference.md#schema-disable) |
| `id` | [id](resources--network_policy_rule--reference.md#schema-id) |
| `ip_prefix_set` | [ip_prefix_set](resources--network_policy_rule--properties--ip_prefix_set.md#section) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](resources--network_policy_rule--properties--ip_prefix_set--ref.md#section) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](resources--network_policy_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--kind) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](resources--network_policy_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--name) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](resources--network_policy_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--namespace) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](resources--network_policy_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--tenant) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](resources--network_policy_rule--properties--ip_prefix_set--ref.md#schema-ip_prefix_set--ref--uid) |
| `label_matcher` | [label_matcher](resources--network_policy_rule--properties--label_matcher.md#section) |
| `label_matcher.keys` | [label_matcher.keys](resources--network_policy_rule--properties--label_matcher.md#schema-label_matcher--keys) |
| `labels` | [labels](resources--network_policy_rule--reference.md#schema-labels) |
| `name` | [name](resources--network_policy_rule--reference.md#schema-name) |
| `namespace` | [namespace](resources--network_policy_rule--reference.md#schema-namespace) |
| `ports` | [ports](resources--network_policy_rule--reference.md#schema-ports) |
| `prefix` | [prefix](resources--network_policy_rule--properties--prefix.md#section) |
| `prefix.prefix` | [prefix.prefix](resources--network_policy_rule--properties--prefix.md#schema-prefix--prefix) |
| `prefix_selector` | [prefix_selector](resources--network_policy_rule--properties--prefix_selector.md#section) |
| `prefix_selector.expressions` | [prefix_selector.expressions](resources--network_policy_rule--properties--prefix_selector.md#schema-prefix_selector--expressions) |
| `protocol` | [protocol](resources--network_policy_rule--reference.md#schema-protocol) |
| `timeouts` | [timeouts](resources--network_policy_rule--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--network_policy_rule--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy_rule--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--network_policy_rule--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--network_policy_rule--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [advanced_action](resources--network_policy_rule--properties--advanced_action.md)
- [ip_prefix_set](resources--network_policy_rule--properties--ip_prefix_set.md)
- [label_matcher](resources--network_policy_rule--properties--label_matcher.md)
- [prefix](resources--network_policy_rule--properties--prefix.md)
- [prefix_selector](resources--network_policy_rule--properties--prefix_selector.md)
- [timeouts](resources--network_policy_rule--properties--timeouts.md)
- [xcsh_network_policy_rule](../resources/network_policy_rule.md)
