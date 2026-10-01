---
page_title: "api_rate_limit.api_endpoint_rules.request_matcher"
subcategory: "Load Balancing"
description: "api_rate_limit.api_endpoint_rules.request_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2491, "body_sha256": "sha256:102b6f7eb5df518e69b532746d4c78452b21123d6c4d5adfc17f57c3e894d815", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:cookie_matchers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:jwt_claims", "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:query_params"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "request_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.api_endpoint_rules.request_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.request_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- api_rate_limit.api_endpoint_rules.request_matcher

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

- [cookie_matchers](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--cookie_matchers.md): complete subsection reference.

- [headers](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers.md): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--jwt_claims.md): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--query_params.md): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--cookie_matchers.md)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers.md)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--jwt_claims.md)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--query_params.md)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
