---
page_title: "rate_limit.rate_limiter.leaky_bucket"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter.leaky_bucket for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 974, "body_sha256": "sha256:54994b5a2601dbd0ded9a3c0faff56c22d1df8401edf90fa11bb97118db139ab", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter", "path": "docs/guides/data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--leaky_bucket.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "leaky_bucket"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter.leaky_bucket for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rate_limit.rate_limiter.leaky_bucket

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter.md)
- rate_limit.rate_limiter.leaky_bucket

<a id="section"></a>

Type: `["object", {}]`. Computed.

Leaky-Bucket is the default rate limiter algorithm for F5.

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

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rate_limit.rate_limiter](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
