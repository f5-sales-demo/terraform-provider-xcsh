---
page_title: "api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api protection rules api groups rules request matcher jwt claims check not present"], "body_bytes": 2153, "body_sha256": "sha256:75bd17d81eb7c71cd7f2b03368cb87716065cb1d74331ead220b5e734d9e86aa", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:jwt_claims:check_not_present", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:jwt_claims", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/jwt_claims/check_not_present/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1221022321103011-2323212110312021-3301313311210313-2313023113311122-0220130312322000-3020003110301201-1211210030123330-1113023111131212", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher", "jwt_claims", "check_not_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/jwt_claims/check_not_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [api_protection_rules.api_groups_rules.request_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/jwt_claims/)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

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

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/jwt_claims/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
