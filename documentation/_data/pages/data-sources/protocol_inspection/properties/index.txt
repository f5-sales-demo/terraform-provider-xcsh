---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_protocol_inspection."
xcsh_docs: {"aliases": [], "body_bytes": 11266, "body_sha256": "sha256:75c9e78832f614a0d42596e1d172e7c35669f01983b9dce0aa522201df192571", "child_ids": ["xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_compliance_checks", "xcsh-docs:data-sources:protocol_inspection:properties:enable_disable_signatures"], "collection_id": "xcsh-docs:data-sources:protocol_inspection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_inspection:reference", "parent_id": "xcsh-docs:data-sources:protocol_inspection:fundamentals", "path": "documentation/data-sources/protocol_inspection/properties/index.md", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_inspection/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_protocol_inspection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
- Property reference

## Direct properties

<a id="schema-action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: ALLOW|DENY|DROP\] Action after inspection - ALLOW: Allow Allow traffic - DENY: Deny Throw
RST error for TCP and ICMP error for UDP - DROP: DROP Silently drop traffic. Possible values are
\`ALLOW\`, \`DENY\`, \`DROP\`. Defaults to \`ALLOW\`. Server applies default when omitted.

Upstream description:

Action after inspection

&#8203;- ALLOW: Allow

Allow traffic &#8203;- DENY: Deny

Throw RST error for TCP and ICMP error for UDP &#8203;- DROP: DROP

Silently drop traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "ALLOW",
  "enum": [
    "ALLOW",
    "DENY",
    "DROP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

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

Description of the ProtocolInspection.

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

- [enable_disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/): complete subsection reference.

- [enable_disable_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/): complete subsection reference.

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

Name of the ProtocolInspection.

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

Namespace where the ProtocolInspection exists.

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

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/#schema-action) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/#schema-description) |
| `enable_disable_compliance_checks` | [enable_disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/#section) |
| `enable_disable_compliance_checks.disable_compliance_checks` | [enable_disable_compliance_checks.disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/disable_compliance_checks/#section) |
| `enable_disable_compliance_checks.enable_compliance_checks` | [enable_disable_compliance_checks.enable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#section) |
| `enable_disable_compliance_checks.enable_compliance_checks.name` | [enable_disable_compliance_checks.enable_compliance_checks.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#schema-enable_disable_compliance_checks--enable_compliance_checks--name) |
| `enable_disable_compliance_checks.enable_compliance_checks.namespace` | [enable_disable_compliance_checks.enable_compliance_checks.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#schema-enable_disable_compliance_checks--enable_compliance_checks--namespace) |
| `enable_disable_compliance_checks.enable_compliance_checks.tenant` | [enable_disable_compliance_checks.enable_compliance_checks.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/enable_compliance_checks/#schema-enable_disable_compliance_checks--enable_compliance_checks--tenant) |
| `enable_disable_signatures` | [enable_disable_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/#section) |
| `enable_disable_signatures.disable_signature` | [enable_disable_signatures.disable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/disable_signature/#section) |
| `enable_disable_signatures.enable_signature` | [enable_disable_signatures.enable_signature](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/enable_signature/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/#schema-namespace) |

## Next pages

- [enable_disable_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_compliance_checks/)
- [enable_disable_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/properties/enable_disable_signatures/)
- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_inspection/)
