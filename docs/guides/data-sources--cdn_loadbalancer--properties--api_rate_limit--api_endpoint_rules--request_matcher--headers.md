---
page_title: "api_rate_limit.api_endpoint_rules.request_matcher.headers"
subcategory: "Load Balancing"
description: "api_rate_limit.api_endpoint_rules.request_matcher.headers for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5584, "body_sha256": "sha256:d54a0b0b04b045e84fe2011a3479c38f2283efda48b16b84f440b5f598f4f8fb", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers:check_not_present", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers:check_present", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers:item"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "request_matcher", "headers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.api_endpoint_rules.request_matcher.headers for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.request_matcher.headers

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_rate_limit](data-sources--cdn_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.api_endpoint_rules](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="section"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Direct properties

- [check_not_present](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers--check_not_present.md): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers--check_present.md): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--request_matcher--headers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers--item.md): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--request_matcher--headers--name"></a>

### name property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers--check_not_present.md)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers--check_present.md)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.item](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers--item.md)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
