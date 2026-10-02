---
page_title: "limits.token_bucket"
subcategory: "Security"
description: "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits."
xcsh_docs: {"aliases": ["limits token bucket"], "body_bytes": 1138, "body_sha256": "sha256:8123610f35841fc6f22b5812f0a0936018f7d331771ea6d1905d43cbcdeae415", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter:properties:limits:token_bucket", "parent_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "path": "documentation/data-sources/rate_limiter/properties/limits/token_bucket/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2212113022203132-2321331333203133-3102222103200021-0310003111032310-3320200113203333-1303022001101013-2122211002201100-3302033133201321", "registry_path": "docs/guides/data-sources--rate_limiter--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["limits", "token_bucket"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/properties/limits/token_bucket/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.token_bucket

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/)
- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/)
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

- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
