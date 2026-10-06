---
page_title: "Property reference"
subcategory: "Security"
description: "Property reference for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["rate limiter"], "body_bytes": 11766, "body_sha256": "sha256:e0cc90842ae1d57f78c7f944fee19a0586a416fa8973196279b7d74ed0b5d512", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter:properties:limits", "xcsh-docs:data-sources:rate_limiter:properties:user_identification"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter:reference", "parent_id": "xcsh-docs:data-sources:rate_limiter:fundamentals", "path": "documentation/data-sources/rate_limiter/properties/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1213113300212133-2333323311332223-0113322313302233-1120122211232010-0333010000311333-1323210023210012-1203111201003021-1223213101213223", "registry_path": "docs/guides/data-sources--rate_limiter--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:rate_limiter:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:rate_limiter:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:rate_limiter:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:rate_limiter:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["limits"], "anchor": "section", "description": "A list of RateLimitValues that specifies the total number of allowed requests for each specified period.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["limits"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:rate_limiter:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:rate_limiter:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["user identification"], "anchor": "section", "description": "A reference to user_identification object. The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:user_identification", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["user_identification"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
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

Description of the RateLimiter.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

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

- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the RateLimiter.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Namespace where the RateLimiter exists.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/user_identification/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/#schema-labels) |
| `limits` | [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/#section) |
| `limits.action_block` | [limits.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/#section) |
| `limits.action_block.hours` | [limits.action_block.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/hours/#section) |
| `limits.action_block.hours.duration` | [limits.action_block.hours.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/hours/#schema-limits--action_block--hours--duration) |
| `limits.action_block.minutes` | [limits.action_block.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/minutes/#section) |
| `limits.action_block.minutes.duration` | [limits.action_block.minutes.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/minutes/#schema-limits--action_block--minutes--duration) |
| `limits.action_block.seconds` | [limits.action_block.seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/seconds/#section) |
| `limits.action_block.seconds.duration` | [limits.action_block.seconds.duration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/seconds/#schema-limits--action_block--seconds--duration) |
| `limits.burst_multiplier` | [limits.burst_multiplier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/#schema-limits--burst_multiplier) |
| `limits.disabled` | [limits.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/disabled/#section) |
| `limits.leaky_bucket` | [limits.leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/leaky_bucket/#section) |
| `limits.period_multiplier` | [limits.period_multiplier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/#schema-limits--period_multiplier) |
| `limits.token_bucket` | [limits.token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/token_bucket/#section) |
| `limits.total_number` | [limits.total_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/#schema-limits--total_number) |
| `limits.unit` | [limits.unit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/#schema-limits--unit) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/#schema-namespace) |
| `user_identification` | [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/user_identification/#section) |
| `user_identification.kind` | [user_identification.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/user_identification/#schema-user_identification--kind) |
| `user_identification.name` | [user_identification.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/user_identification/#schema-user_identification--name) |
| `user_identification.namespace` | [user_identification.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/user_identification/#schema-user_identification--namespace) |
| `user_identification.tenant` | [user_identification.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/user_identification/#schema-user_identification--tenant) |
| `user_identification.uid` | [user_identification.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/user_identification/#schema-user_identification--uid) |
