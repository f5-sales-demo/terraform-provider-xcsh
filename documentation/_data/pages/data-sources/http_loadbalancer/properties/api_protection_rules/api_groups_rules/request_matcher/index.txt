---
page_title: "api_protection_rules.api_groups_rules.request_matcher"
subcategory: "Load Balancing"
description: "Request conditions for matching a rule."
xcsh_docs: {"aliases": ["api protection rules api groups rules request matcher"], "body_bytes": 3265, "body_sha256": "sha256:d8d066eadbd6bfc47a0dc921d8559f51b1b60ba3e9a2288f6018cdb8f8fba00c", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:cookie_matchers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:jwt_claims", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:query_params"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "path": "documentation/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher"], "schema_version": 1, "sections": [{"aliases": ["cookie matchers"], "anchor": "section", "description": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must eva", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:cookie_matchers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher", "cookie_matchers"], "syntax": "attribute", "type": "object"}, {"aliases": ["headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header predicates", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt claims"], "anchor": "section", "description": "A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true. No", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:jwt_claims", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher", "jwt_claims"], "syntax": "attribute", "type": "object"}, {"aliases": ["query params"], "anchor": "section", "description": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that all", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:query_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher", "query_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Request conditions for matching a rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.request_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- api_protection_rules.api_groups_rules.request_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/cookie_matchers/): complete subsection reference.

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/headers/): complete subsection reference.

- [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/jwt_claims/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/query_params/): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/cookie_matchers/)
- [api_protection_rules.api_groups_rules.request_matcher.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/headers/)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/jwt_claims/)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/query_params/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
