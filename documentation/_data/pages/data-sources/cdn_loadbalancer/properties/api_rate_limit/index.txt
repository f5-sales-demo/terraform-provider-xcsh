---
page_title: "api_rate_limit"
subcategory: "Load Balancing"
description: "APIRateLimit."
xcsh_docs: {"aliases": ["api rate limit"], "body_bytes": 2570, "body_sha256": "sha256:bd0deecfeedd1a248ed82bc35cd7b903c7f7a64cb543db425cba319aa68cc197", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:ip_allowed_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_rate_limit/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit"], "schema_version": 1, "sections": [{"aliases": ["api rate limit api endpoint rules"], "anchor": "section", "description": "Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit bypass rate limiting rules"], "anchor": "section", "description": "This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit custom ip allowed list"], "anchor": "section", "description": "IP Allowed list using existing ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "custom_ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit ip allowed list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:ip_allowed_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit no ip allowed list"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "no_ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules"], "anchor": "section", "description": "Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "APIRateLimit.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- api_rate_limit

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\]
APIRateLimit.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/#section)
- [disable_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/disable_rate_limit/#section)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/): complete subsection reference.

- [bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/): complete subsection reference.

- [custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/): complete subsection reference.

- [ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/ip_allowed_list/): complete subsection reference.

- [no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/no_ip_allowed_list/): complete subsection reference.

- [server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/): complete subsection reference.
