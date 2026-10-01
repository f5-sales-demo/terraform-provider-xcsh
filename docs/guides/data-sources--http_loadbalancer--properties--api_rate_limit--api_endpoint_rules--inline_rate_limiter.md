---
page_title: "api_rate_limit.api_endpoint_rules.inline_rate_limiter"
subcategory: "Load Balancing"
description: "api_rate_limit.api_endpoint_rules.inline_rate_limiter for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4207, "body_sha256": "sha256:55c73d2ac980c1a61439d4778aee662e790479b8e4addd84b2f95c8b66799f58", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:ref_user_id", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:use_http_lb_user_id"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.api_endpoint_rules.inline_rate_limiter for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.inline_rate_limiter

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for inline rate limiter.

Upstream description:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

## Direct properties

- [ref_user_id](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter--ref_user_id.md): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--threshold"></a>

### threshold property

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

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
    "minimum": 1,
    "multipleOf": 1
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

<a id="schema-api_rate_limit--api_endpoint_rules--inline_rate_limiter--unit"></a>

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

- [use_http_lb_user_id](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter--use_http_lb_user_id.md): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter--ref_user_id.md)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter--use_http_lb_user_id.md)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
