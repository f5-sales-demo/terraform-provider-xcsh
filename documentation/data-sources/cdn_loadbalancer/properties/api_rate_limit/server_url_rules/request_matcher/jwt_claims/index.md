---
page_title: "api_rate_limit.server_url_rules.request_matcher.jwt_claims"
subcategory: "Load Balancing"
description: "A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true."
xcsh_docs: {"aliases": ["api rate limit server url rules request matcher jwt claims"], "body_bytes": 4606, "body_sha256": "sha256:64510dfb2936c99339dba9232497aa834720402cc4cda33a387c26f24e87e427", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims:check_not_present", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims:check_present", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims:item"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/jwt_claims/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "jwt_claims"], "schema_version": 1, "sections": [{"aliases": ["api rate limit server url rules request matcher jwt claims check not present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims:check_not_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "jwt_claims", "check_not_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules request matcher jwt claims check present"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims:check_present", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "jwt_claims", "check_present"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules request matcher jwt claims invert matcher"], "anchor": "schema-api_rate_limit--server_url_rules--request_matcher--jwt_claims--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "jwt_claims", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["api rate limit server url rules request matcher jwt claims item", "succeeded", "success", "successful"], "anchor": "section", "description": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims:item", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "jwt_claims", "item"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules request matcher jwt claims name"], "anchor": "schema-api_rate_limit--server_url_rules--request_matcher--jwt_claims--name", "description": "JWT claim name.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:jwt_claims", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "jwt_claims", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/jwt_claims/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.request_matcher.jwt_claims

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [api_rate_limit.server_url_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="section"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/jwt_claims/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/jwt_claims/check_present/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--request_matcher--jwt_claims--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/jwt_claims/item/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--request_matcher--jwt_claims--name"></a>

### name property

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
