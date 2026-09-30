---
page_title: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers"
subcategory: "Load Balancing"
description: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6236, "body_sha256": "sha256:83e3ecdaeb3d9e19b89dd534c25b80dbd4a961882cc42f24f7648a6cccf31444", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:cookie_matchers", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:cookie_matchers:check_not_present", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:cookie_matchers:check_present", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:cookie_matchers:item"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:cookie_matchers", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "request_matcher", "cookie_matchers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/request_matcher/cookie_matchers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher.md)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="section"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
    }
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

- [check_not_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--check_not_present.md): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--check_present.md): complete subsection reference.

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--item.md): complete subsection reference.

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--name"></a>

### name property

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--check_not_present.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--check_present.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--cookie_matchers--item.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
