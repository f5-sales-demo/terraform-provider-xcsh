---
page_title: "rate_limit.rate_limiter.token_bucket"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter.token_bucket for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1025, "body_sha256": "sha256:6526a638ef8992180b70e6c11c17e7afcfcb9fc3a524be9c63df5bca78ffb7ed", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:token_bucket", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter", "path": "docs/guides/resources--cdn_loadbalancer--properties--rate_limit--rate_limiter--token_bucket.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "token_bucket"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/token_bucket/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter.token_bucket for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rate_limit.rate_limiter.token_bucket

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md)
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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rate_limit.rate_limiter](resources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
