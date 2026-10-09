---
page_title: "api_rate_limit.api_endpoint_rules"
subcategory: "Load Balancing"
description: "Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter."
xcsh_docs: {"aliases": ["api rate limit api endpoint rules"], "body_bytes": 4538, "body_sha256": "sha256:03c8d7edead11b289d518910595c71535959b0251479d49fe77fdb0668f303a7", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:any_domain", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:api_endpoint_method", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:ref_rate_limiter", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3311331301112212-2122231110001213-0022021101212332-2101323001120013-3121233021100121-0012133103000222-0012121003200302-1333312110021222", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules"], "schema_version": 1, "sections": [{"aliases": ["api rate limit api endpoint rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:any_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules api endpoint method", "succeeded", "success", "successful"], "anchor": "section", "description": "A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is considered successful if the input method is a member of the list. The result of the match based on the method list is inverted if invert_matcher is true.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:api_endpoint_method", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "api_endpoint_method"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules api endpoint path"], "anchor": "schema-api_rate_limit--api_endpoint_rules--api_endpoint_path", "description": "The endpoint (path) of the request.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "api_endpoint_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["api rate limit api endpoint rules client matcher"], "anchor": "section", "description": "Client conditions for matching a rule.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules inline rate limiter"], "anchor": "section", "description": "Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the required rate_limiter_choice when no stored rate-limiter object is used.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules ref rate limiter"], "anchor": "section", "description": "Reference to a stored rate-limiter object for this scoped rule. Select exactly one of ref_rate_limiter and inline_rate_limiter.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:ref_rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "ref_rate_limiter"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules request matcher"], "anchor": "section", "description": "Request conditions for matching a rule.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "request_matcher"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit api endpoint rules specific domain"], "anchor": "schema-api_rate_limit--api_endpoint_rules--specific_domain", "description": "Exclusive with The rule will apply for a specific domain.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "specific_domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/)
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/any_domain/): complete subsection reference.

- [api_endpoint_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/api_endpoint_method/): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--api_endpoint_path"></a>

### api_endpoint_path property

Type: `"string"`. Computed.

API Endpoint. The endpoint (path) of the request.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/): complete subsection reference.

- [inline_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/): complete subsection reference.

- [ref_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/ref_rate_limiter/): complete subsection reference.

- [request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/): complete subsection reference.

<a id="schema-api_rate_limit--api_endpoint_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
