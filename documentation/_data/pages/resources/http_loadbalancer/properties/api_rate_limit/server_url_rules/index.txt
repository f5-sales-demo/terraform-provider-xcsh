---
page_title: "api_rate_limit.server_url_rules"
subcategory: "Load Balancing"
description: "Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter."
xcsh_docs: {"aliases": ["api rate limit server url rules"], "body_bytes": 7915, "body_sha256": "sha256:d44079215c63e349423093c9497b5d6241105a3c3e1bd39235a5e56724b654cc", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:any_domain", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:ref_rate_limiter", "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-008.md", "relationships": [{"anchor": "schema-api_rate_limit--server_url_rules--specific_domain", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:any_domain,specific_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:inline_rate_limiter,ref_rate_limiter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:ConflictingListObjectAttributes:inline_rate_limiter,ref_rate_limiter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:ref_rate_limiter", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--server_url_rules--base_path", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules:RequiredListObjectAttributes:base_path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules"], "schema_version": 1, "sections": [{"aliases": ["api rate limit server url rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:any_domain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["api rate limit server url rules api group"], "anchor": "schema-api_rate_limit--server_url_rules--api_group", "description": "API groups derived from API Definition swaggers. For example oas-all-operations including all paths and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the swaggers. Custom groups can be created if user tags paths or operations with \"x-F5 Distributed Cloud-API-group\" extensions", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "api_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["api rate limit server url rules base path"], "anchor": "schema-api_rate_limit--server_url_rules--base_path", "description": "Prefix of the request path.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "base_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["api rate limit server url rules client matcher"], "anchor": "section", "description": "Client conditions for matching a rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_client,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:client_selector,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_list,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:asn_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:any_client,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_threat_category_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher:ConflictingObjectAttributes:client_selector,ip_threat_category_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_threat_category_list", "type": "conflicts"}], "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit server url rules inline rate limiter"], "anchor": "section", "description": "Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the required rate_limiter_choice when no stored rate-limiter object is used.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.inline_rate_limiter:ConflictingObjectAttributes:ref_user_id,use_http_lb_user_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter:ref_user_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.inline_rate_limiter:ConflictingObjectAttributes:ref_user_id,use_http_lb_user_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter:use_http_lb_user_id", "type": "conflicts"}, {"anchor": "schema-api_rate_limit--server_url_rules--inline_rate_limiter--threshold", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.inline_rate_limiter:RequiredObjectAttributes:threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:inline_rate_limiter", "type": "requires"}], "schema_path": ["api_rate_limit", "server_url_rules", "inline_rate_limiter"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit server url rules ref rate limiter"], "anchor": "section", "description": "Reference to a stored rate-limiter object for this scoped rule. Select exactly one of ref_rate_limiter and inline_rate_limiter.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:ref_rate_limiter", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_rate_limit--server_url_rules--ref_rate_limiter--name", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.ref_rate_limiter:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:ref_rate_limiter", "type": "requires"}], "schema_path": ["api_rate_limit", "server_url_rules", "ref_rate_limiter"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit server url rules request matcher"], "anchor": "section", "description": "Request conditions for matching a rule.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["api rate limit server url rules specific domain"], "anchor": "schema-api_rate_limit--server_url_rules--specific_domain", "description": "Exclusive with The rule will apply for a specific domain.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "specific_domain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one rate_limiter_choice: inline_rate_limiter or ref_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- api_rate_limit.server_url_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
server_url_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/any_domain/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--api_group"></a>

### api_group property

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-api_rate_limit--server_url_rules--base_path"></a>

### base_path property

Type: `"string"`. Optional.

Base Path. Prefix of the request path.

Upstream description:

Prefix of the request path.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/): complete subsection reference.

- [inline_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/inline_rate_limiter/): complete subsection reference.

- [ref_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/ref_rate_limiter/): complete subsection reference.

- [request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/): complete subsection reference.

<a id="schema-api_rate_limit--server_url_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [api_rate_limit.server_url_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/any_domain/)
- [api_rate_limit.server_url_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/)
- [api_rate_limit.server_url_rules.inline_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/inline_rate_limiter/)
- [api_rate_limit.server_url_rules.ref_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/ref_rate_limiter/)
- [api_rate_limit.server_url_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
