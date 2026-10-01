---
page_title: "origin_pool.more_origin_options"
subcategory: "Load Balancing"
description: "origin_pool.more_origin_options for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1605, "body_sha256": "sha256:193f5d056c0a514791fdb67dffc462aa8b89783e51a78cc79bd0c3bdf10eb13d", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--origin_pool--more_origin_options.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "more_origin_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/more_origin_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.more_origin_options for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.more_origin_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [origin_pool](data-sources--cdn_loadbalancer--properties--origin_pool.md)
- origin_pool.more_origin_options

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for more origin options.

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

<a id="schema-origin_pool--more_origin_options--enable_byte_range_request"></a>

### enable_byte_range_request property

Type: `"bool"`. Computed.

Choice to enable/disable byte range requests towards origin.

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

<a id="schema-origin_pool--more_origin_options--websocket_proxy"></a>

### websocket_proxy property

Type: `"bool"`. Computed.

Option to enable proxying of websocket connections to the origin server.

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

## Next pages

- [origin_pool](data-sources--cdn_loadbalancer--properties--origin_pool.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
