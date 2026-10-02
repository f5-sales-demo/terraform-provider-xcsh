---
page_title: "api_protection_rules.api_endpoint_rules.request_matcher"
subcategory: "Load Balancing"
description: "Request conditions for matching a rule."
xcsh_docs: {"aliases": ["api protection rules api endpoint rules request matcher"], "body_bytes": 3301, "body_sha256": "sha256:d203c4e8ba6ae185e94c92ee1215ce3a530f42dae4a8c3ae0996c9f5a1133486", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:jwt_claims", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:query_params"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "path": "documentation/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher"], "schema_version": 1, "sections": [{"aliases": ["cookie matchers"], "anchor": "section", "description": "A list of predicates for all cookies that need to be matched. The criteria for matching each cookie is described in individual instances of CookieMatcherType. The actual cookie values are extracted from the request API as a list of strings for each cookie name. Note that all specified cookie matcher predicates must", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "cookie_matchers"], "syntax": "attribute", "type": "object"}, {"aliases": ["headers"], "anchor": "section", "description": "A list of predicates for various HTTP headers that need to match. The criteria for matching each HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values are extracted from the request API as a list of strings for each HTTP header type. Note that all specified header", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt claims"], "anchor": "section", "description": "A list of predicates for various JWT claims that need to match. The criteria for matching each JWT claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates must evaluate to true.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:jwt_claims", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "jwt_claims"], "syntax": "attribute", "type": "object"}, {"aliases": ["query params"], "anchor": "section", "description": "A list of predicates for all query parameters that need to be matched. The criteria for matching each query parameter are described in individual instances of QueryParameterMatcherType. The actual query parameter values are extracted from the request API as a list of strings for each query parameter name. Note that", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:query_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "query_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Request conditions for matching a rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.request_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- api_protection_rules.api_endpoint_rules.request_matcher

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

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/): complete subsection reference.

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/headers/): complete subsection reference.

- [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/query_params/): complete subsection reference.

## Next pages

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/headers/)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/query_params/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
