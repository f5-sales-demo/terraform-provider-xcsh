---
page_title: "api_rate_limit.api_endpoint_rules"
subcategory: "Load Balancing"
description: "api_rate_limit.api_endpoint_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5357, "body_sha256": "sha256:1191ef6663e04ebad6f71283b206c4b6559020022e4982651fe9c2603b0f1ecf", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:api_endpoint_method", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:ref_rate_limiter", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.api_endpoint_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- api_rate_limit.api_endpoint_rules

<a id="section"></a>

Type: `"list"`. Computed.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

## Direct properties

- [any_domain](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--any_domain.md): complete subsection reference.

- [api_endpoint_method](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--api_endpoint_method.md): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--api_endpoint_path"></a>

### api_endpoint_path property

Type: `"string"`. Computed.

API Endpoint. The endpoint (path) of the request.

Upstream description:

The endpoint (path) of the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher.md): complete subsection reference.

- [inline_rate_limiter](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter.md): complete subsection reference.

- [ref_rate_limiter](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--ref_rate_limiter.md): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

## Next pages

- [api_rate_limit.api_endpoint_rules.any_domain](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--any_domain.md)
- [api_rate_limit.api_endpoint_rules.api_endpoint_method](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--api_endpoint_method.md)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--client_matcher.md)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--inline_rate_limiter.md)
- [api_rate_limit.api_endpoint_rules.ref_rate_limiter](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--ref_rate_limiter.md)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
