---
page_title: "rate_limit.rate_limiter.token_bucket"
subcategory: "Load Balancing"
description: "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits."
xcsh_docs: {"aliases": ["rate limit rate limiter token bucket"], "body_bytes": 1133, "body_sha256": "sha256:53649ed50d744e4f68c51cac1b3dc7e4659a1e3db7f0dbce759c5421e451b2c2", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "path": "documentation/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3211020022010202-1101201222211102-3010213311221001-2102010330320122-0120301001332101-0332320102011122-0321200121122313-3032003110020000", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "token_bucket"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.token_bucket

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/)
- rate_limit.rate_limiter.token_bucket

<a id="section"></a>

Type: `["object", {}]`. Optional.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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
token_bucket = {}
```

This is an empty object or choice marker. It has no direct properties.
