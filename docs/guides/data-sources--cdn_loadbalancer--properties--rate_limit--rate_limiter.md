---
page_title: "rate_limit.rate_limiter"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6006, "body_sha256": "sha256:b50b8ec24853d942c65cd2177de491b53ae0f60a32c8145befcc55649b57f37b", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:disabled", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "rate_limiter"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [rate_limit](data-sources--cdn_loadbalancer--properties--rate_limit.md)
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

- [action_block](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block.md): complete subsection reference.

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

- [disabled](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--disabled.md): complete subsection reference.

- [leaky_bucket](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--leaky_bucket.md): complete subsection reference.

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

- [token_bucket](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--token_bucket.md): complete subsection reference.

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

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block.md)
- [rate_limit.rate_limiter.disabled](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--disabled.md)
- [rate_limit.rate_limiter.leaky_bucket](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--leaky_bucket.md)
- [rate_limit.rate_limiter.token_bucket](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--token_bucket.md)
- [rate_limit](data-sources--cdn_loadbalancer--properties--rate_limit.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
