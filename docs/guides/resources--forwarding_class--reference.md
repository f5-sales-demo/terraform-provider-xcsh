---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_forwarding_class."
xcsh_docs: {"aliases": [], "body_bytes": 14697, "body_sha256": "sha256:8b674edb7fb599bc6d4b1b0b1acf2e8f34156792c951a050c363f6fe966e0716", "canonical_id": "xcsh-docs:resources:forwarding_class:reference", "child_ids": ["xcsh-docs:resources:forwarding_class:properties:dscp", "xcsh-docs:resources:forwarding_class:properties:dscp_based_queue", "xcsh-docs:resources:forwarding_class:properties:no_marking", "xcsh-docs:resources:forwarding_class:properties:no_policer", "xcsh-docs:resources:forwarding_class:properties:policer", "xcsh-docs:resources:forwarding_class:properties:timeouts"], "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:resources:forwarding_class:reference", "parent_id": "xcsh-docs:resources:forwarding_class:fundamentals", "path": "docs/guides/resources--forwarding_class--reference.md", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_forwarding_class.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md)
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

- [dscp](resources--forwarding_class--properties--dscp.md): complete subsection reference.

- [dscp_based_queue](resources--forwarding_class--properties--dscp_based_queue.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-interface_group"></a>

### interface_group property

Type: `"string"`. Optional, Computed.

\[Enum: ANY\_AVAILABLE\_INTERFACE|INTERFACE\_GROUP1|INTERFACE\_GROUP2|INTERFACE\_GROUP3\] Interface
group, group membership by adding group label to interface Choose any of the available interfaces
Choose all interfaces with label group1 Choose all interfaces with label group2 Choose all
interfaces with label group3. Possible values are \`ANY\_AVAILABLE\_INTERFACE\`,
\`INTERFACE\_GROUP1\`, \`INTERFACE\_GROUP2\`, \`INTERFACE\_GROUP3\`. Defaults to
\`ANY\_AVAILABLE\_INTERFACE\`.

Upstream description:

Interface group, group membership by adding group label to interface

Choose any of the available interfaces Choose all interfaces with label group1 Choose all interfaces
with label group2 Choose all interfaces with label group3.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY_AVAILABLE_INTERFACE",
    "INTERFACE_GROUP1",
    "INTERFACE_GROUP2",
    "INTERFACE_GROUP3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY_AVAILABLE_INTERFACE",
  "enum": [
    "ANY_AVAILABLE_INTERFACE",
    "INTERFACE_GROUP1",
    "INTERFACE_GROUP2",
    "INTERFACE_GROUP3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

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

Name of the Forwarding Class. Must be unique within the namespace.

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

Namespace where the Forwarding Class is created.

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

- [no_marking](resources--forwarding_class--properties--no_marking.md): complete subsection reference.

- [no_policer](resources--forwarding_class--properties--no_policer.md): complete subsection reference.

- [policer](resources--forwarding_class--properties--policer.md): complete subsection reference.

<a id="schema-queue_id_to_use"></a>

### queue_id_to_use property

Type: `"string"`. Optional, Computed.

\[Enum:
DSCP\_BEST\_EFFORT|DSCP\_CLASS1|DSCP\_CLASS2|DSCP\_CLASS3|DSCP\_CLASS4|DSCP\_EXPRESS\_FORWARDING|DSCP\_CONTROL\_L3|DSCP\_CONTROL\_L2\]
DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1
service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used
for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic
like.. Possible values are \`DSCP\_BEST\_EFFORT\`, \`DSCP\_CLASS1\`, \`DSCP\_CLASS2\`,
\`DSCP\_CLASS3\`, \`DSCP\_CLASS4\`, \`DSCP\_EXPRESS\_FORWARDING\`, \`DSCP\_CONTROL\_L3\`,
\`DSCP\_CONTROL\_L2\`. Defaults to \`DSCP\_BEST\_EFFORT\`.

Upstream description:

DSCP Precedence Level Values

Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP
Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is
used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not
recommended.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DSCP_BEST_EFFORT",
  "enum": [
    "DSCP_BEST_EFFORT",
    "DSCP_CLASS1",
    "DSCP_CLASS2",
    "DSCP_CLASS3",
    "DSCP_CLASS4",
    "DSCP_EXPRESS_FORWARDING",
    "DSCP_CONTROL_L3",
    "DSCP_CONTROL_L2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--forwarding_class--properties--timeouts.md): complete subsection reference.

<a id="schema-tos_value"></a>

### tos_value property

Type: `"number"`. Optional, Computed.

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Upstream description:

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--forwarding_class--reference.md#schema-annotations) |
| `description` | [description](resources--forwarding_class--reference.md#schema-description) |
| `disable` | [disable](resources--forwarding_class--reference.md#schema-disable) |
| `dscp` | [dscp](resources--forwarding_class--properties--dscp.md#section) |
| `dscp.drop_precedence` | [dscp.drop_precedence](resources--forwarding_class--properties--dscp.md#schema-dscp--drop_precedence) |
| `dscp.dscp_class` | [dscp.dscp_class](resources--forwarding_class--properties--dscp.md#schema-dscp--dscp_class) |
| `dscp_based_queue` | [dscp_based_queue](resources--forwarding_class--properties--dscp_based_queue.md#section) |
| `id` | [id](resources--forwarding_class--reference.md#schema-id) |
| `interface_group` | [interface_group](resources--forwarding_class--reference.md#schema-interface_group) |
| `labels` | [labels](resources--forwarding_class--reference.md#schema-labels) |
| `name` | [name](resources--forwarding_class--reference.md#schema-name) |
| `namespace` | [namespace](resources--forwarding_class--reference.md#schema-namespace) |
| `no_marking` | [no_marking](resources--forwarding_class--properties--no_marking.md#section) |
| `no_policer` | [no_policer](resources--forwarding_class--properties--no_policer.md#section) |
| `policer` | [policer](resources--forwarding_class--properties--policer.md#section) |
| `policer.name` | [policer.name](resources--forwarding_class--properties--policer.md#schema-policer--name) |
| `policer.namespace` | [policer.namespace](resources--forwarding_class--properties--policer.md#schema-policer--namespace) |
| `policer.tenant` | [policer.tenant](resources--forwarding_class--properties--policer.md#schema-policer--tenant) |
| `queue_id_to_use` | [queue_id_to_use](resources--forwarding_class--reference.md#schema-queue_id_to_use) |
| `timeouts` | [timeouts](resources--forwarding_class--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--forwarding_class--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--forwarding_class--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--forwarding_class--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--forwarding_class--properties--timeouts.md#schema-timeouts--update) |
| `tos_value` | [tos_value](resources--forwarding_class--reference.md#schema-tos_value) |

## Next pages

- [dscp](resources--forwarding_class--properties--dscp.md)
- [dscp_based_queue](resources--forwarding_class--properties--dscp_based_queue.md)
- [no_marking](resources--forwarding_class--properties--no_marking.md)
- [no_policer](resources--forwarding_class--properties--no_policer.md)
- [policer](resources--forwarding_class--properties--policer.md)
- [timeouts](resources--forwarding_class--properties--timeouts.md)
- [xcsh_forwarding_class](../resources/forwarding_class.md)
