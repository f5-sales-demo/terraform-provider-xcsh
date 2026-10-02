---
page_title: "api_rate_limit"
subcategory: "Load Balancing"
description: "Path- or API-group-scoped rate limiting. Define server_url_rules or api_endpoint_rules and choose inline_rate_limiter for an inline limit, or ref_rate_limiter for a stored rate-limiter reference."
xcsh_docs: {"aliases": ["api rate limit"], "body_bytes": 5030, "body_sha256": "sha256:475dbe6eac308fd301a76f4606d99ca8aa7f7da6bd29db646b6c4e04818a0405", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:ip_allowed_list", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,custom_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,custom_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:bypass_rate_limiting_rules,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:custom_ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit:ConflictingObjectAttributes:ip_allowed_list,no_ip_allowed_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit"], "schema_version": 1, "sections": [{"aliases": ["api endpoint rules"], "anchor": "section", "description": "Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules"], "syntax": "block", "type": "object"}, {"aliases": ["bypass rate limiting rules"], "anchor": "section", "description": "This category defines rules per URL or API group. If request matches any of these rules, skip Rate Limiting.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules"], "syntax": "block", "type": "object"}, {"aliases": ["custom ip allowed list"], "anchor": "section", "description": "IP Allowed list using existing ip_prefix_set objects.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.custom_ip_allowed_list:RequiredObjectAttributes:rate_limiter_allowed_prefixes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes", "type": "requires"}], "schema_path": ["api_rate_limit", "custom_ip_allowed_list"], "syntax": "block", "type": "object"}, {"aliases": ["ip allowed list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:ip_allowed_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "ip_allowed_list"], "syntax": "block", "type": "object"}, {"aliases": ["no ip allowed list"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:no_ip_allowed_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "no_ip_allowed_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["server url rules"], "anchor": "section", "description": "Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Path- or API-group-scoped rate limiting. Define server_url_rules or api_endpoint_rules and choose inline_rate_limiter for an inline limit, or ref_rate_limiter for a stored rate-limiter reference.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- api_rate_limit

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\] Path-
or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and choose
inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Upstream description:

Path- or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and
choose inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

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

- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/#section)
- [disable_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/disable_rate_limit/#section)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_rate_limit {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/): complete subsection reference.

- [bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/): complete subsection reference.

- [custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/): complete subsection reference.

- [ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/ip_allowed_list/): complete subsection reference.

- [no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/no_ip_allowed_list/): complete subsection reference.

- [server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [api_rate_limit.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/)
- [api_rate_limit.custom_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/)
- [api_rate_limit.ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/ip_allowed_list/)
- [api_rate_limit.no_ip_allowed_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/no_ip_allowed_list/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
