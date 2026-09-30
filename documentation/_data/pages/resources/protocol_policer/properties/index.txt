---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 12364, "body_sha256": "sha256:5aa15e9262fb532a69e0e4ed2b35eb130d6e465e6087d860b3288bf6e855579f", "child_ids": ["xcsh-docs:resources:protocol_policer:properties:protocol_policer", "xcsh-docs:resources:protocol_policer:properties:timeouts"], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:reference", "parent_id": "xcsh-docs:resources:protocol_policer:fundamentals", "path": "documentation/resources/protocol_policer/properties/index.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
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

Name of the Protocol Policer. Must be unique within the namespace.

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

Namespace for the Protocol Policer. The F5 XC API restricts this resource to the system namespace;
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

- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/#schema-namespace) |
| `protocol_policer` | [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/#section) |
| `protocol_policer.policer` | [protocol_policer.policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/policer/#section) |
| `protocol_policer.policer.kind` | [protocol_policer.policer.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/policer/#schema-protocol_policer--policer--kind) |
| `protocol_policer.policer.name` | [protocol_policer.policer.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/policer/#schema-protocol_policer--policer--name) |
| `protocol_policer.policer.namespace` | [protocol_policer.policer.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/policer/#schema-protocol_policer--policer--namespace) |
| `protocol_policer.policer.tenant` | [protocol_policer.policer.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/policer/#schema-protocol_policer--policer--tenant) |
| `protocol_policer.policer.uid` | [protocol_policer.policer.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/policer/#schema-protocol_policer--policer--uid) |
| `protocol_policer.protocol` | [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/#section) |
| `protocol_policer.protocol.dns` | [protocol_policer.protocol.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/dns/#section) |
| `protocol_policer.protocol.icmp` | [protocol_policer.protocol.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/icmp/#section) |
| `protocol_policer.protocol.icmp.type` | [protocol_policer.protocol.icmp.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/icmp/#schema-protocol_policer--protocol--icmp--type) |
| `protocol_policer.protocol.tcp` | [protocol_policer.protocol.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/tcp/#section) |
| `protocol_policer.protocol.tcp.flags` | [protocol_policer.protocol.tcp.flags](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/tcp/#schema-protocol_policer--protocol--tcp--flags) |
| `protocol_policer.protocol.udp` | [protocol_policer.protocol.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/udp/#section) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/timeouts/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
