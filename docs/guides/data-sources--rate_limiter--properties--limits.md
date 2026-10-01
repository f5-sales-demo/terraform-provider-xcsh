---
page_title: "limits"
subcategory: "Security"
description: "limits for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 5920, "body_sha256": "sha256:078b71f6c4efeff8885f07f004df3fe1409e8b77cfcdd38f7cc374c7fc8096af", "canonical_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "child_ids": ["xcsh-docs:data-sources:rate_limiter:properties:limits:action_block", "xcsh-docs:data-sources:rate_limiter:properties:limits:disabled", "xcsh-docs:data-sources:rate_limiter:properties:limits:leaky_bucket", "xcsh-docs:data-sources:rate_limiter:properties:limits:token_bucket"], "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "parent_id": "xcsh-docs:data-sources:rate_limiter:reference", "path": "docs/guides/data-sources--rate_limiter--properties--limits.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["limits"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/properties/limits/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "limits for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md)
- [Property reference](data-sources--rate_limiter--reference.md)
- limits

<a id="section"></a>

Type: `"list"`. Computed.

List of RateLimitValues that specifies the total number of allowed requests for each specified
period.

Upstream description:

A list of RateLimitValues that specifies the total number of allowed requests for each specified
period.

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

## Direct properties

- [action_block](data-sources--rate_limiter--properties--limits--action_block.md): complete subsection reference.

<a id="schema-limits--burst_multiplier"></a>

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

- [disabled](data-sources--rate_limiter--properties--limits--disabled.md): complete subsection reference.

- [leaky_bucket](data-sources--rate_limiter--properties--limits--leaky_bucket.md): complete subsection reference.

<a id="schema-limits--period_multiplier"></a>

### period_multiplier property

Type: `"number"`. Computed.

Setting, combined with Per Period units, provides a duration.

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

- [token_bucket](data-sources--rate_limiter--properties--limits--token_bucket.md): complete subsection reference.

<a id="schema-limits--total_number"></a>

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

<a id="schema-limits--unit"></a>

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

- [limits.action_block](data-sources--rate_limiter--properties--limits--action_block.md)
- [limits.disabled](data-sources--rate_limiter--properties--limits--disabled.md)
- [limits.leaky_bucket](data-sources--rate_limiter--properties--limits--leaky_bucket.md)
- [limits.token_bucket](data-sources--rate_limiter--properties--limits--token_bucket.md)
- [Property reference](data-sources--rate_limiter--reference.md)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md)
