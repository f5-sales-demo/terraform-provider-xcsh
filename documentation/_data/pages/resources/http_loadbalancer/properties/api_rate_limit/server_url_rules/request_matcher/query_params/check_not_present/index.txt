---
page_title: "api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api rate limit server url rules request matcher query params check not present"], "body_bytes": 2093, "body_sha256": "sha256:0ca2575ab1d66cd8722406da4451347f2c610cfef1c55452604ea0c39f61f0df", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params:check_not_present", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:request_matcher:query_params", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/check_not_present/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3021120003000002-2122133132200033-2111100220122321-1130122312232121-1123322213333012-2222213213222122-0331233023132103-2030020303303221", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "request_matcher", "query_params", "check_not_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/check_not_present/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [api_rate_limit.server_url_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/)
- [api_rate_limit.server_url_rules.request_matcher.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_rate_limit.server_url_rules.request_matcher.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/request_matcher/query_params/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
