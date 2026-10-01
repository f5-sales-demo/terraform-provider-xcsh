---
page_title: "api_protection_rules.api_groups_rules.request_matcher"
subcategory: "Load Balancing"
description: "api_protection_rules.api_groups_rules.request_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2575, "body_sha256": "sha256:256a042b91f6b28ea30080ee7212f0734fb2fa6ece39fd9b9a7ef6c4114459bd", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:cookie_matchers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:jwt_claims", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher:query_params"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "request_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/request_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_groups_rules.request_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.request_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_protection_rules](data-sources--http_loadbalancer--properties--api_protection_rules.md)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md)
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

- [cookie_matchers](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--cookie_matchers.md): complete subsection reference.

- [headers](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--headers.md): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--jwt_claims.md): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--query_params.md): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--cookie_matchers.md)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--headers.md)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--jwt_claims.md)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher--query_params.md)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
