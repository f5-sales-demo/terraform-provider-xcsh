---
page_title: "origin_pool.use_tls.default_session_key_caching"
subcategory: "Load Balancing"
description: "origin_pool.use_tls.default_session_key_caching for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1116, "body_sha256": "sha256:64a691185b7a04ae4d1eb96d7ccd1b2483aa87c79a54aa484bd2036c7402a7e4", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:default_session_key_caching", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls:default_session_key_caching", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:use_tls", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--origin_pool--use_tls--default_session_key_caching.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "use_tls", "default_session_key_caching"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/use_tls/default_session_key_caching/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.use_tls.default_session_key_caching for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pool.use_tls.default_session_key_caching

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [origin_pool](data-sources--cdn_loadbalancer--properties--origin_pool.md)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--properties--origin_pool--use_tls.md)
- origin_pool.use_tls.default_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
