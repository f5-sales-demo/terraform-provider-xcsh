---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_forwarding_class."
xcsh_docs: {"aliases": ["forwarding class"], "body_bytes": 13514, "body_sha256": "sha256:f861cbfe342908dfed1fff50b784f886f0b69f76e8b70a01ce0b0ef01918c1d4", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forwarding_class:properties:dscp", "xcsh-docs:data-sources:forwarding_class:properties:dscp_based_queue", "xcsh-docs:data-sources:forwarding_class:properties:no_marking", "xcsh-docs:data-sources:forwarding_class:properties:no_policer", "xcsh-docs:data-sources:forwarding_class:properties:policer"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forwarding_class:reference", "parent_id": "xcsh-docs:data-sources:forwarding_class:fundamentals", "path": "documentation/data-sources/forwarding_class/properties/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213", "registry_path": "docs/guides/data-sources--forwarding_class--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["dscp"], "anchor": "section", "description": "DSCP marking setting as per RFC 2475.", "document_id": "xcsh-docs:data-sources:forwarding_class:properties:dscp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dscp"], "syntax": "attribute", "type": "object"}, {"aliases": ["dscp based queue"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forwarding_class:properties:dscp_based_queue", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dscp_based_queue"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interface group"], "anchor": "schema-interface_group", "description": "Interface group, group membership by adding group label to interface Choose any of the available interfaces Choose all interfaces with label group1 Choose all interfaces with label group2 Choose all interfaces with label group3.", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["interface_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["no marking"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forwarding_class:properties:no_marking", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_marking"], "syntax": "attribute", "type": "object"}, {"aliases": ["no policer"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:forwarding_class:properties:no_policer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_policer"], "syntax": "attribute", "type": "object"}, {"aliases": ["policer"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:forwarding_class:properties:policer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["policer"], "syntax": "attribute", "type": "object"}, {"aliases": ["queue id to use"], "anchor": "schema-queue_id_to_use", "description": "DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["queue_id_to_use"], "syntax": "attribute", "type": "string"}, {"aliases": ["tos value"], "anchor": "schema-tos_value", "description": "Exclusive with Decimal value of raw 8 bit TOS. In above example DSCP 10 = Precedence Class 1 and drop precedence low.", "document_id": "xcsh-docs:data-sources:forwarding_class:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tos_value"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/properties/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Property reference for xcsh_forwarding_class.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_forwarding_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

Description of the ForwardingClass.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [dscp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp/): complete subsection reference.

- [dscp_based_queue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp_based_queue/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-interface_group"></a>

### interface_group property

Type: `"string"`. Computed.

\[Enum: ANY\_AVAILABLE\_INTERFACE|INTERFACE\_GROUP1|INTERFACE\_GROUP2|INTERFACE\_GROUP3\] Interface
group, group membership by adding group label to interface Choose any of the available interfaces
Choose all interfaces with label group1 Choose all interfaces with label group2 Choose all
interfaces with label group3. Possible values are \`ANY\_AVAILABLE\_INTERFACE\`,
\`INTERFACE\_GROUP1\`, \`INTERFACE\_GROUP2\`, \`INTERFACE\_GROUP3\`. Defaults to
\`ANY\_AVAILABLE\_INTERFACE\`.

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

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

Name of the ForwardingClass.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Namespace where the ForwardingClass exists.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [no_marking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/no_marking/): complete subsection reference.

- [no_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/no_policer/): complete subsection reference.

- [policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/policer/): complete subsection reference.

<a id="schema-queue_id_to_use"></a>

### queue_id_to_use property

Type: `"string"`. Computed.

\[Enum:
DSCP\_BEST\_EFFORT|DSCP\_CLASS1|DSCP\_CLASS2|DSCP\_CLASS3|DSCP\_CLASS4|DSCP\_EXPRESS\_FORWARDING|DSCP\_CONTROL\_L3|DSCP\_CONTROL\_L2\]
DSCP Precedence Level Values Best Effort service will GET any available bandwidth DSCP Class 1
service DSCP Class 2 service DSCP Class 3 service DSCP Class 4 service Express Forwarding is used
for low latency traffic Control is used for routing traffic, not recommended Link Layer traffic
like.. Possible values are \`DSCP\_BEST\_EFFORT\`, \`DSCP\_CLASS1\`, \`DSCP\_CLASS2\`,
\`DSCP\_CLASS3\`, \`DSCP\_CLASS4\`, \`DSCP\_EXPRESS\_FORWARDING\`, \`DSCP\_CONTROL\_L3\`,
\`DSCP\_CONTROL\_L2\`. Defaults to \`DSCP\_BEST\_EFFORT\`.

Additional upstream details:

DSCP Precedence Level Values

Best Effort service will GET any available bandwidth DSCP Class 1 service DSCP Class 2 service DSCP
Class 3 service DSCP Class 4 service Express Forwarding is used for low latency traffic Control is
used for routing traffic, not recommended Link Layer traffic like LACP or keepalive, not
recommended.

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

<a id="schema-tos_value"></a>

### tos_value property

Type: `"number"`. Computed.

Exclusive with \[dscp no\_marking\] Decimal value of raw 8 bit TOS. In above example DSCP 10 =
Precedence Class 1 and drop precedence low.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-description) |
| `dscp` | [dscp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp/#section) |
| `dscp.drop_precedence` | [dscp.drop_precedence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp/#schema-dscp--drop_precedence) |
| `dscp.dscp_class` | [dscp.dscp_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp/#schema-dscp--dscp_class) |
| `dscp_based_queue` | [dscp_based_queue](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/dscp_based_queue/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-id) |
| `interface_group` | [interface_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-interface_group) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-namespace) |
| `no_marking` | [no_marking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/no_marking/#section) |
| `no_policer` | [no_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/no_policer/#section) |
| `policer` | [policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/policer/#section) |
| `policer.name` | [policer.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/policer/#schema-policer--name) |
| `policer.namespace` | [policer.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/policer/#schema-policer--namespace) |
| `policer.tenant` | [policer.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/policer/#schema-policer--tenant) |
| `queue_id_to_use` | [queue_id_to_use](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-queue_id_to_use) |
| `tos_value` | [tos_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/#schema-tos_value) |
