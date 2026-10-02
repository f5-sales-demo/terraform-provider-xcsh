---
page_title: "limits"
subcategory: "Security"
description: "A list of RateLimitValues that specifies the total number of allowed requests for each specified period."
xcsh_docs: {"aliases": ["limits"], "body_bytes": 7480, "body_sha256": "sha256:43115e3863ec7f18349bb1bf6f72c961ee4e88a6fcd24583cbef10bd179f6214", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:rate_limiter:properties:limits:action_block", "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "xcsh-docs:resources:rate_limiter:properties:limits:leaky_bucket", "xcsh-docs:resources:rate_limiter:properties:limits:token_bucket"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits", "parent_id": "xcsh-docs:resources:rate_limiter:reference", "path": "documentation/resources/rate_limiter/properties/limits/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3310131000101102-2223023103233001-2320233313020222-3232300220331020-0300101132023133-1111032010320111-3101031132110011-3311313100200222", "registry_path": "docs/guides/resources--rate_limiter--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:leaky_bucket", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits:ConflictingListObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:token_bucket", "type": "conflicts"}, {"anchor": "schema-limits--total_number", "enforcement": "provider-schema", "group": "limits:RequiredListObjectAttributes:total_number", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["limits"], "schema_version": 1, "sections": [{"aliases": ["action block"], "anchor": "section", "description": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "limits.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:rate_limiter:properties:limits:action_block:seconds", "type": "conflicts"}], "schema_path": ["limits", "action_block"], "syntax": "block", "type": "object"}, {"aliases": ["burst multiplier"], "anchor": "schema-limits--burst_multiplier", "description": "The maximum burst of requests to accommodate, expressed as a multiple of the rate.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "burst_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["leaky bucket"], "anchor": "section", "description": "Leaky-Bucket is the default rate limiter algorithm for F5.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:leaky_bucket", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "leaky_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["period multiplier"], "anchor": "schema-limits--period_multiplier", "description": "This setting, combined with Per Period units, provides a duration.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "period_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["token bucket"], "anchor": "section", "description": "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits:token_bucket", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "token_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["total number"], "anchor": "schema-limits--total_number", "description": "The total number of allowed requests per rate-limiting period.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "total_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["unit"], "anchor": "schema-limits--unit", "description": "Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days.", "document_id": "xcsh-docs:resources:rate_limiter:properties:limits", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "unit"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of RateLimitValues that specifies the total number of allowed requests for each specified period.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/)
- limits

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Upstream description:

A list of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("total_number"),
  validators.ConflictingListObjectAttributes("action_block",
    "disabled"),
  validators.ConflictingListObjectAttributes("leaky_bucket",
    "token_bucket")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
limits {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/): complete subsection reference.

<a id="schema-limits--burst_multiplier"></a>

### burst_multiplier property

Type: `"number"`. Optional.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/disabled/): complete subsection reference.

- [leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/leaky_bucket/): complete subsection reference.

<a id="schema-limits--period_multiplier"></a>

### period_multiplier property

Type: `"number"`. Optional.

Setting, combined with Per Period units, provides a duration.

Upstream description:

This setting, combined with Per Period units, provides a duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/token_bucket/): complete subsection reference.

<a id="schema-limits--total_number"></a>

### total_number property

Type: `"number"`. Optional.

The total number of allowed requests per rate-limiting period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="schema-limits--unit"></a>

### unit property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [limits.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/action_block/)
- [limits.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/disabled/)
- [limits.leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/leaky_bucket/)
- [limits.token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/limits/token_bucket/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
