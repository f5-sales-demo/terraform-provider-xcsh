---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": ["fast acl rule"], "body_bytes": 15539, "body_sha256": "sha256:14e90705e7e1aefb48851d72724f7cdae06eeab1105702def1718ffe386ab66d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl_rule:properties:action", "xcsh-docs:resources:fast_acl_rule:properties:ip_prefix_set", "xcsh-docs:resources:fast_acl_rule:properties:port", "xcsh-docs:resources:fast_acl_rule:properties:prefix", "xcsh-docs:resources:fast_acl_rule:properties:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:reference", "parent_id": "xcsh-docs:resources:fast_acl_rule:fundamentals", "path": "documentation/resources/fast_acl_rule/properties/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232", "registry_path": "docs/guides/resources--fast_acl_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "section", "description": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-action--simple_action", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "type": "conflicts"}, {"anchor": "schema-action--simple_action", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:protocol_policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:action:protocol_policer_action", "type": "conflicts"}], "schema_path": ["action"], "syntax": "block", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:fast_acl_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:fast_acl_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:fast_acl_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:fast_acl_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:ip_prefix_set", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:fast_acl_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:fast_acl_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:fast_acl_rule:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "section", "description": "L4 port numbers to match.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-port--user_defined", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port", "type": "conflicts"}, {"anchor": "schema-port--user_defined", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:dns", "type": "conflicts"}], "schema_path": ["port"], "syntax": "block", "type": "object"}, {"aliases": ["prefix"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:prefix", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["prefix"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_fast_acl_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
- Property reference

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Additional upstream details:

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

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

Name of the Fast ACL Rule. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

Namespace where the Fast ACL Rule is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/): complete subsection reference.

- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/prefix/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/#section) |
| `action.policer_action` | [action.policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/#section) |
| `action.policer_action.ref` | [action.policer_action.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/#section) |
| `action.policer_action.ref.kind` | [action.policer_action.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/#schema-action--policer_action--ref--kind) |
| `action.policer_action.ref.name` | [action.policer_action.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/#schema-action--policer_action--ref--name) |
| `action.policer_action.ref.namespace` | [action.policer_action.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/#schema-action--policer_action--ref--namespace) |
| `action.policer_action.ref.tenant` | [action.policer_action.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/#schema-action--policer_action--ref--tenant) |
| `action.policer_action.ref.uid` | [action.policer_action.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/#schema-action--policer_action--ref--uid) |
| `action.protocol_policer_action` | [action.protocol_policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/#section) |
| `action.protocol_policer_action.ref` | [action.protocol_policer_action.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/ref/#section) |
| `action.protocol_policer_action.ref.kind` | [action.protocol_policer_action.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/ref/#schema-action--protocol_policer_action--ref--kind) |
| `action.protocol_policer_action.ref.name` | [action.protocol_policer_action.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/ref/#schema-action--protocol_policer_action--ref--name) |
| `action.protocol_policer_action.ref.namespace` | [action.protocol_policer_action.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/ref/#schema-action--protocol_policer_action--ref--namespace) |
| `action.protocol_policer_action.ref.tenant` | [action.protocol_policer_action.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/ref/#schema-action--protocol_policer_action--ref--tenant) |
| `action.protocol_policer_action.ref.uid` | [action.protocol_policer_action.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/protocol_policer_action/ref/#schema-action--protocol_policer_action--ref--uid) |
| `action.simple_action` | [action.simple_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/#schema-action--simple_action) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/#schema-id) |
| `ip_prefix_set` | [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/#section) |
| `ip_prefix_set.ref` | [ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/ref/#section) |
| `ip_prefix_set.ref.kind` | [ip_prefix_set.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--kind) |
| `ip_prefix_set.ref.name` | [ip_prefix_set.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--name) |
| `ip_prefix_set.ref.namespace` | [ip_prefix_set.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--namespace) |
| `ip_prefix_set.ref.tenant` | [ip_prefix_set.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--tenant) |
| `ip_prefix_set.ref.uid` | [ip_prefix_set.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/ip_prefix_set/ref/#schema-ip_prefix_set--ref--uid) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/#schema-namespace) |
| `port` | [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/#section) |
| `port.all` | [port.all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/all/#section) |
| `port.dns` | [port.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/dns/#section) |
| `port.user_defined` | [port.user_defined](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/#schema-port--user_defined) |
| `prefix` | [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/prefix/#section) |
| `prefix.prefix` | [prefix.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/prefix/#schema-prefix--prefix) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/timeouts/#schema-timeouts--update) |
