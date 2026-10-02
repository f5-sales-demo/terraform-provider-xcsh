---
page_title: "rate_limit.rate_limiter"
subcategory: "Load Balancing"
description: "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period."
xcsh_docs: {"aliases": ["rate limit rate limiter"], "body_bytes": 7610, "body_sha256": "sha256:f56f6d220d17ceae5d4a83d5e9551998b2fd845221ed5abb3309292f2f92bf2c", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:disabled", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit", "path": "documentation/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2221301003332310-3123302231002320-3221021022221231-2302132120132200-3303010010220200-2112200111213023-0322311302303232-3302113001113211", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-014.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:action_block,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:ConflictingObjectAttributes:leaky_bucket,token_bucket", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "type": "conflicts"}, {"anchor": "schema-rate_limit--rate_limiter--total_number", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter:RequiredObjectAttributes:total_number", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter"], "schema_version": 1, "sections": [{"aliases": ["action block"], "anchor": "section", "description": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:hours,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rate_limit.rate_limiter.action_block:ConflictingObjectAttributes:minutes,seconds", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "type": "conflicts"}], "schema_path": ["rate_limit", "rate_limiter", "action_block"], "syntax": "block", "type": "object"}, {"aliases": ["burst multiplier"], "anchor": "schema-rate_limit--rate_limiter--burst_multiplier", "description": "The maximum burst of requests to accommodate, expressed as a multiple of the rate.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "burst_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["leaky bucket"], "anchor": "section", "description": "Leaky-Bucket is the default rate limiter algorithm for F5.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "leaky_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["period multiplier"], "anchor": "schema-rate_limit--rate_limiter--period_multiplier", "description": "This setting, combined with Per Period units, provides a duration.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "period_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["token bucket"], "anchor": "section", "description": "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "token_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["total number"], "anchor": "schema-rate_limit--rate_limiter--total_number", "description": "The total number of allowed requests per rate-limiting period.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "total_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["unit"], "anchor": "schema-rate_limit--rate_limiter--unit", "description": "Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "unit"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/)
- rate_limit.rate_limiter

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Upstream description:

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("total_number"),
  validators.ConflictingObjectAttributes("action_block",
    "disabled"),
  validators.ConflictingObjectAttributes("leaky_bucket",
    "token_bucket")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

Terraform syntax:

```terraform
rate_limiter {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--burst_multiplier"></a>

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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/disabled/): complete subsection reference.

- [leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--period_multiplier"></a>

### period_multiplier property

Type: `"number"`. Optional, Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

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

- [token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--total_number"></a>

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

<a id="schema-rate_limit--rate_limiter--unit"></a>

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

- [rate_limit.rate_limiter.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/)
- [rate_limit.rate_limiter.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/disabled/)
- [rate_limit.rate_limiter.leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/)
- [rate_limit.rate_limiter.token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
