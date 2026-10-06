---
page_title: "api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api rate limit api endpoint rules request matcher cookie matchers check not present"], "body_bytes": 1754, "body_sha256": "sha256:ca15029776458013c5f7e5468e8c4a1d8a5cc674f7d0ab8f25ee7ac5d8e89586", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:cookie_matchers:check_not_present", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:cookie_matchers", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/cookie_matchers/check_not_present/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "request_matcher", "cookie_matchers", "check_not_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/cookie_matchers/check_not_present/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [api_rate_limit.api_endpoint_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/cookie_matchers/)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

This can be used for messages where no values are needed.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.
