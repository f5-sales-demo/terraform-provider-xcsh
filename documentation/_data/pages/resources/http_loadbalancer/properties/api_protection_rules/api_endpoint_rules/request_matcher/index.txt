---
page_title: "api_protection_rules.api_endpoint_rules.request_matcher"
subcategory: "Load Balancing"
description: "api_protection_rules.api_endpoint_rules.request_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3383, "body_sha256": "sha256:df8a99a5a6a61c4128efd8c8fee6299052bb26708c6493b0786dab9c4bf50165", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:cookie_matchers", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:headers", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:jwt_claims", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:query_params"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_endpoint_rules.request_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.request_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- api_protection_rules.api_endpoint_rules.request_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/): complete subsection reference.

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/headers/): complete subsection reference.

- [jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/query_params/): complete subsection reference.

## Next pages

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/cookie_matchers/)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/headers/)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/query_params/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
