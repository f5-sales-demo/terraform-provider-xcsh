---
page_title: "api_rate_limit"
subcategory: "Load Balancing"
description: "APIRateLimit."
xcsh_docs: {"aliases": ["api rate limit"], "body_bytes": 4591, "body_sha256": "sha256:295975f64fa994b7e59cb7442254773e0f0224e8991e8c010656669431ae3468", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/api_rate_limit/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,custom_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,custom_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit"], "schema_version": 1, "sections": [{"aliases": ["api rate limit api endpoint rules"], "anchor": "section", "description": "Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-api_rate_limit--api_endpoint_rules--specific_domain", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules:ConflictingListObjectAttributes:inline_rate_limiter,ref_rate_limiter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules:ConflictingListObjectAttributes:inline_rate_limiter,ref_rate_limiter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:ref_rate_limiter", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--api_endpoint_rules--api_endpoint_path", "enforcement": "provider-schema", "group": "api_rate_limit.api_endpoint_rules:RequiredListObjectAttributes:api_endpoint_path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "type": "requires"}], "schema_path": ["api_rate_limit", "api_endpoint_rules"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit bypass rate limiting rules"], "anchor": "section", "description": "This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit custom ip allowed list"], "anchor": "section", "description": "IP Allowed list using existing ip_prefix_set objects.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.custom_ip_allowed_list:RequiredObjectAttributes:rate_limiter_allowed_prefixes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes", "type": "requires"}], "schema_path": ["api_rate_limit", "custom_ip_allowed_list"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit ip allowed list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:ip_allowed_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "ip_allowed_list"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit no ip allowed list"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "no_ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules"], "anchor": "section", "description": "Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-api_rate_limit--server_url_rules--specific_domain", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:inline_rate_limiter,ref_rate_limiter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:inline_rate_limiter,ref_rate_limiter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:ref_rate_limiter", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--server_url_rules--base_path", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:RequiredListObjectAttributes:base_path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules", "type": "requires"}], "schema_path": ["api_rate_limit", "server_url_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "APIRateLimit.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- api_rate_limit

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\]
APIRateLimit.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "custom_ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list")}
```

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

- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/#section)
- [disable_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/disable_rate_limit/#section)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_rate_limit {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/): complete subsection reference.

- [bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/): complete subsection reference.

- [custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/): complete subsection reference.

- [ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/ip_allowed_list/): complete subsection reference.

- [no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/no_ip_allowed_list/): complete subsection reference.

- [server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [api_rate_limit.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/)
- [api_rate_limit.custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/)
- [api_rate_limit.ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/ip_allowed_list/)
- [api_rate_limit.no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/no_ip_allowed_list/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
