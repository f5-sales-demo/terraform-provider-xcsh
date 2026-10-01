---
page_title: "limits"
subcategory: "Security"
description: "limits for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 7480, "body_sha256": "sha256:43115e3863ec7f18349bb1bf6f72c961ee4e88a6fcd24583cbef10bd179f6214", "child_ids": ["xcsh-docs:resources:rate_limiter:properties:limits:action_block", "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "xcsh-docs:resources:rate_limiter:properties:limits:leaky_bucket", "xcsh-docs:resources:rate_limiter:properties:limits:token_bucket"], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits", "parent_id": "xcsh-docs:resources:rate_limiter:reference", "path": "documentation/resources/rate_limiter/properties/limits/index.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["limits"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "limits for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
