---
page_title: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims"
subcategory: "Load Balancing"
description: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6264, "body_sha256": "sha256:6aac1c4232c37e468ce89113102a83e626ced32d25ebfd38f2930b86f02e58f7", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:jwt_claims", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:jwt_claims:check_not_present", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:jwt_claims:check_present", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:jwt_claims:item"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher:jwt_claims", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:request_matcher", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "request_matcher", "jwt_claims"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/request_matcher/jwt_claims/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher.md)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="section"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

- [check_not_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--check_not_present.md): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--check_present.md): complete subsection reference.

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--item.md): complete subsection reference.

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--name"></a>

### name property

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--check_not_present.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--check_present.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher--jwt_claims--item.md)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--request_matcher.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
