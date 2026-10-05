---
page_title: "api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api protection rules api endpoint rules request matcher jwt claims check present"], "body_bytes": 2157, "body_sha256": "sha256:7a10c84e76bb5869942ab4741b22bb1ba9a620f2e22f3e6659a730b626daf7ed", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:jwt_claims:check_present", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules:request_matcher:jwt_claims", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/check_present/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1013231211303123-1220030002202332-1310300330221320-0313113300201011-2221331103333213-1100201111213201-0312333232202311-0113333203010301", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_endpoint_rules", "request_matcher", "jwt_claims", "check_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/check_present/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/)
- [api_protection_rules.api_endpoint_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_endpoint_rules/request_matcher/jwt_claims/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
