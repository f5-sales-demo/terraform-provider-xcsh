---
page_title: "limits.token_bucket"
subcategory: "Security"
description: "limits.token_bucket for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 881, "body_sha256": "sha256:8f349f6fcc65ea95d1e40bf18163d061920504b2759e52d4fac4bfe2052b6a84", "canonical_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:token_bucket", "child_ids": [], "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter:properties:limits:token_bucket", "parent_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "path": "docs/guides/data-sources--rate_limiter--properties--limits--token_bucket.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["limits", "token_bucket"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/properties/limits/token_bucket/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "limits.token_bucket for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.token_bucket

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md)
- [Property reference](data-sources--rate_limiter--reference.md)
- [limits](data-sources--rate_limiter--properties--limits.md)
- limits.token_bucket

<a id="section"></a>

Type: `["object", {}]`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [limits](data-sources--rate_limiter--properties--limits.md)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md)
