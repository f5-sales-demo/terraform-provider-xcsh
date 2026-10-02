---
page_title: "rate_limit.rate_limiter"
subcategory: "Load Balancing"
description: "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period."
xcsh_docs: {"aliases": ["rate limit rate limiter"], "body_bytes": 6655, "body_sha256": "sha256:3431aa7e9095e74c87faa728c6480027065ed6f9753bb371b17fff84fcde6b95", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:disabled", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit", "path": "documentation/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter"], "schema_version": 1, "sections": [{"aliases": ["action block"], "anchor": "section", "description": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["burst multiplier"], "anchor": "schema-rate_limit--rate_limiter--burst_multiplier", "description": "The maximum burst of requests to accommodate, expressed as a multiple of the rate.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "burst_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["leaky bucket"], "anchor": "section", "description": "Leaky-Bucket is the default rate limiter algorithm for F5.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "leaky_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["period multiplier"], "anchor": "schema-rate_limit--rate_limiter--period_multiplier", "description": "This setting, combined with Per Period units, provides a duration.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "period_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["token bucket"], "anchor": "section", "description": "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "token_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["total number"], "anchor": "schema-rate_limit--rate_limiter--total_number", "description": "The total number of allowed requests per rate-limiting period.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "total_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["unit"], "anchor": "schema-rate_limit--rate_limiter--unit", "description": "Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "unit"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/)
- rate_limit.rate_limiter

<a id="section"></a>

Type: `"single"`. Computed.

Tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Upstream description:

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

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

## Direct properties

- [action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--burst_multiplier"></a>

### burst_multiplier property

Type: `"number"`. Computed.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/disabled/): complete subsection reference.

- [leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--period_multiplier"></a>

### period_multiplier property

Type: `"number"`. Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Upstream description:

This setting, combined with Per Period units, provides a duration.

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

- [token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--total_number"></a>

### total_number property

Type: `"number"`. Computed.

The total number of allowed requests per rate-limiting period.

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

Type: `"string"`. Computed.

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

- [rate_limit.rate_limiter.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/)
- [rate_limit.rate_limiter.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/disabled/)
- [rate_limit.rate_limiter.leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/)
- [rate_limit.rate_limiter.token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
