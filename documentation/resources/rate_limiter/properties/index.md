---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["rate limiter"], "body_bytes": 13455, "body_sha256": "sha256:1a99d19ae1dc393ed145ef5b5fcd8668bfd6f4f65d50bb698bf06cc7517f0e31", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:rate_limiter:properties:limits", "xcsh-docs:resources:rate_limiter:properties:timeouts", "xcsh-docs:resources:rate_limiter:properties:user_identification"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:reference", "parent_id": "xcsh-docs:resources:rate_limiter:fundamentals", "path": "documentation/resources/rate_limiter/properties/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2123321020010221-3010212320031032-2010230230101221-1113323033011311-0030013332100232-3330302220301023-0133012200323100-2102212210333022", "registry_path": "docs/guides/resources--rate_limiter--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:rate_limiter:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:rate_limiter:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:rate_limiter:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:rate_limiter:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:rate_limiter:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["limits"], "anchor": "section", "description": "A list of RateLimitValues that specifies the total number of allowed requests for each specified period.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:leaky_bucket", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:token_bucket", "type": "conflicts"}, {"anchor": "schema-limits--total_number", "enforcement": "provider-schema", "group": "limits:RequiredListObjectAttributes:total_number", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits", "type": "requires"}], "schema_path": ["limits"], "syntax": "block", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:rate_limiter:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:rate_limiter:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:rate_limiter:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["user identification"], "anchor": "section", "description": "A reference to user_identification object. The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.", "document_id": "xcsh-docs:resources:rate_limiter:properties:user_identification", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["user_identification"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
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

- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Rate Limiter. Must be unique within the namespace.

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

Namespace where the Rate Limiter is created.

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

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/timeouts/): complete subsection reference.

- [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/#schema-labels) |
| `limits` | [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/#section) |
| `limits.action_block` | [limits.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/#section) |
| `limits.action_block.hours` | [limits.action_block.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/hours/#section) |
| `limits.action_block.hours.duration` | [limits.action_block.hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/hours/#schema-limits--action_block--hours--duration) |
| `limits.action_block.minutes` | [limits.action_block.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/minutes/#section) |
| `limits.action_block.minutes.duration` | [limits.action_block.minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/minutes/#schema-limits--action_block--minutes--duration) |
| `limits.action_block.seconds` | [limits.action_block.seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/seconds/#section) |
| `limits.action_block.seconds.duration` | [limits.action_block.seconds.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/seconds/#schema-limits--action_block--seconds--duration) |
| `limits.burst_multiplier` | [limits.burst_multiplier](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/#schema-limits--burst_multiplier) |
| `limits.disabled` | [limits.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/disabled/#section) |
| `limits.leaky_bucket` | [limits.leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/leaky_bucket/#section) |
| `limits.period_multiplier` | [limits.period_multiplier](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/#schema-limits--period_multiplier) |
| `limits.token_bucket` | [limits.token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/token_bucket/#section) |
| `limits.total_number` | [limits.total_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/#schema-limits--total_number) |
| `limits.unit` | [limits.unit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/#schema-limits--unit) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/timeouts/#schema-timeouts--update) |
| `user_identification` | [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/#section) |
| `user_identification.kind` | [user_identification.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/#schema-user_identification--kind) |
| `user_identification.name` | [user_identification.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/#schema-user_identification--name) |
| `user_identification.namespace` | [user_identification.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/#schema-user_identification--namespace) |
| `user_identification.tenant` | [user_identification.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/#schema-user_identification--tenant) |
| `user_identification.uid` | [user_identification.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/#schema-user_identification--uid) |

## Next pages

- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/timeouts/)
- [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/user_identification/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
