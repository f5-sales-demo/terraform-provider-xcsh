---
page_title: "rate_limit.rate_limiter"
subcategory: "Load Balancing"
description: "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period."
xcsh_docs: {"aliases": ["rate limit rate limiter"], "body_bytes": 5394, "body_sha256": "sha256:5d85fc3fe4cd0cde576749b354f5c65698fa59325db6154baad47ff9c9fa4da9", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:disabled", "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:token_bucket"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit", "path": "documentation/resources/http_loadbalancer/properties/rate_limit/rate_limiter/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter"], "schema_version": 1, "sections": [{"aliases": ["rate limit rate limiter action block"], "anchor": "section", "description": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block"], "syntax": "block", "type": "object"}, {"aliases": ["rate limit rate limiter burst multiplier"], "anchor": "schema-rate_limit--rate_limiter--burst_multiplier", "description": "The maximum burst of requests to accommodate, expressed as a multiple of the rate.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "burst_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["rate limit rate limiter disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["rate limit rate limiter leaky bucket"], "anchor": "section", "description": "Leaky-Bucket is the default rate limiter algorithm for F5.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "leaky_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["rate limit rate limiter period multiplier"], "anchor": "schema-rate_limit--rate_limiter--period_multiplier", "description": "This setting, combined with Per Period units, provides a duration.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "period_multiplier"], "syntax": "attribute", "type": "number"}, {"aliases": ["rate limit rate limiter token bucket"], "anchor": "section", "description": "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "token_bucket"], "syntax": "attribute", "type": "object"}, {"aliases": ["rate limit rate limiter total number"], "anchor": "schema-rate_limit--rate_limiter--total_number", "description": "The total number of allowed requests per rate-limiting period.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "total_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["rate limit rate limiter unit"], "anchor": "schema-rate_limit--rate_limiter--unit", "description": "Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "unit"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/rate_limit/rate_limiter/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "A tuple consisting of a rate limit period unit and the total number of allowed requests for that period.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/)
- rate_limit.rate_limiter

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
rate_limiter {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--burst_multiplier"></a>

### burst_multiplier property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/disabled/): complete subsection reference.

- [leaky_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--period_multiplier"></a>

### period_multiplier property

Type: `"number"`. Optional, Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [token_bucket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/): complete subsection reference.

<a id="schema-rate_limit--rate_limiter--total_number"></a>

### total_number property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
