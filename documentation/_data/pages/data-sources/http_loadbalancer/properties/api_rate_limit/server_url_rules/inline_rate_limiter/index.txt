---
page_title: "api_rate_limit.server_url_rules.inline_rate_limiter"
subcategory: "Load Balancing"
description: "Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the required rate_limiter_choice when no stored rate-limiter object is used."
xcsh_docs: {"aliases": ["api rate limit server url rules inline rate limiter"], "body_bytes": 3520, "body_sha256": "sha256:c4bb9a71cc6d249e846634f5c40da26c1c251f56c99eb236d985a3553588bb8c", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter:ref_user_id", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter:use_http_lb_user_id"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "path": "documentation/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/inline_rate_limiter/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0322033310030223-3022000203103231-3203020110202213-2221301013202210-3110312310203011-0023332220310100-1311220130202000-3201021231210030", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "inline_rate_limiter"], "schema_version": 1, "sections": [{"aliases": ["api rate limit server url rules inline rate limiter ref user id"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter:ref_user_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "inline_rate_limiter", "ref_user_id"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules inline rate limiter threshold"], "anchor": "schema-api_rate_limit--server_url_rules--inline_rate_limiter--threshold", "description": "The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified period.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "inline_rate_limiter", "threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["api rate limit server url rules inline rate limiter unit"], "anchor": "schema-api_rate_limit--server_url_rules--inline_rate_limiter--unit", "description": "Unit for the period per which the rate limit is applied. - SECOND: Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR: Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "inline_rate_limiter", "unit"], "syntax": "attribute", "type": "string"}, {"aliases": ["api rate limit server url rules inline rate limiter use http lb user id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter:use_http_lb_user_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "inline_rate_limiter", "use_http_lb_user_id"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/inline_rate_limiter/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the required rate_limiter_choice when no stored rate-limiter object is used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.inline_rate_limiter

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="section"></a>

Type: `"single"`. Computed.

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

- [ref_user_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/inline_rate_limiter/ref_user_id/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--inline_rate_limiter--threshold"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-api_rate_limit--server_url_rules--inline_rate_limiter--unit"></a>

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

- [use_http_lb_user_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/server_url_rules/inline_rate_limiter/use_http_lb_user_id/): complete subsection reference.
