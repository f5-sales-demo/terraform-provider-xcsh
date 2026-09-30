---
page_title: "api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item"
subcategory: "Load Balancing"
description: "api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6067, "body_sha256": "sha256:36605b7c5c0022e521f9c671d8d9b0e519f33c83bda490ddedaf36fc05c9f18d", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:cookie_matchers:item", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:cookie_matchers:item", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:cookie_matchers", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--request_matcher--cookie_matchers--item.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "cookie_matchers", "item"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/cookie_matchers/item/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_rate_limit](data-sources--cdn_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules.md)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--request_matcher.md)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--request_matcher--cookie_matchers.md)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="section"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

## Direct properties

<a id="schema-api_rate_limit--server_url_rules--request_matcher--cookie_matchers--item--exact_values"></a>

### exact_values property

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-api_rate_limit--server_url_rules--request_matcher--cookie_matchers--item--regex_values"></a>

### regex_values property

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-api_rate_limit--server_url_rules--request_matcher--cookie_matchers--item--transformers"></a>

### transformers property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--properties--api_rate_limit--server_url_rules--request_matcher--cookie_matchers.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
