---
page_title: "api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present"
subcategory: "Load Balancing"
description: "api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1656, "body_sha256": "sha256:a36100b917e588310691feff5ef1a3b3c0115b41c3f1b2e1763c2238e7d6b525", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers:check_present", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers:check_present", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:request_matcher:headers", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers--check_present.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "request_matcher", "headers", "check_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/request_matcher/headers/check_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_rate_limit](resources--cdn_loadbalancer--properties--api_rate_limit.md)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules.md)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher.md)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers.md)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present

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

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--cdn_loadbalancer--properties--api_rate_limit--api_endpoint_rules--request_matcher--headers.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
