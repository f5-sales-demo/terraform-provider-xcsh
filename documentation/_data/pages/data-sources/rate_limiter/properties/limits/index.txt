---
page_title: "limits"
subcategory: "Security"
description: "A list of RateLimitValues that specifies the total number of allowed requests for each specified period."
xcsh_docs: {"aliases": ["limits"], "body_bytes": 5223, "body_sha256": "sha256:903cbe8c6148d6af24d14687efb932d12a7aa2bda75236539464b037c361a7be", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter:properties:limits:action_block", "xcsh-docs:data-sources:rate_limiter:properties:limits:disabled", "xcsh-docs:data-sources:rate_limiter:properties:limits:leaky_bucket", "xcsh-docs:data-sources:rate_limiter:properties:limits:token_bucket"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "parent_id": "xcsh-docs:data-sources:rate_limiter:reference", "path": "documentation/data-sources/rate_limiter/properties/limits/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1210221011020222-1230201312322022-0013012231131330-1221200323202101-0311331002301321-2120032131302113-1312313103113202-2030231023103011", "registry_path": "docs/guides/data-sources--rate_limiter--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["limits"], "schema_version": 1, "sections": [{"aliases": ["limits action block"], "anchor": "section", "description": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["limits burst multiplier"], "anchor": "schema-limits--burst_multiplier", "description": "The maximum burst of requests to accommodate, expressed as a multiple of the rate.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "burst_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["limits disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["limits leaky bucket"], "anchor": "section", "description": "Leaky-Bucket is the default rate limiter algorithm for F5.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:leaky_bucket", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "leaky_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["limits period multiplier"], "anchor": "schema-limits--period_multiplier", "description": "This setting, combined with Per Period units, provides a duration.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "period_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["limits token bucket"], "anchor": "section", "description": "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:token_bucket", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "token_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["limits total number"], "anchor": "schema-limits--total_number", "description": "The total number of allowed requests per rate-limiting period.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "total_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["limits unit"], "anchor": "schema-limits--unit", "description": "Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["limits", "unit"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/properties/limits/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "A list of RateLimitValues that specifies the total number of allowed requests for each specified period.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/)
- limits

<a id="section"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/disabled/): complete subsection reference.

- [leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/leaky_bucket/): complete subsection reference.

<a id="schema-limits--period_multiplier"></a>

### period_multiplier property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/token_bucket/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
