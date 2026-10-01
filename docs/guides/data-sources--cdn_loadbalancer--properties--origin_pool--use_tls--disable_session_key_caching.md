---
page_title: "origin_pool.use_tls.disable_session_key_caching"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.disable_session_key_caching for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1153, "body_sha256": "sha256:7dd988d8ddb91b7bcd2f965425e84d94a127c8bed50d3ffefe89a8de60c7a9e4", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_session_key_caching", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:disable_session_key_caching", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--disable_session_key_caching.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "disable_session_key_caching"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.disable_session_key_caching for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.disable_session_key_caching

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [origin_pool](data-sources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- origin_pool.use_tls.disable_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
