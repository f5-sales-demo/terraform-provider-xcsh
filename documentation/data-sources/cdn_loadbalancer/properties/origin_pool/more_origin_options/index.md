---
page_title: "origin_pool.more_origin_options"
subcategory: "Load Balancing"
description: "Configuration parameter for more origin options."
xcsh_docs: {"aliases": ["backend servers", "origin pool more origin options", "origin servers", "upstream servers"], "body_bytes": 1862, "body_sha256": "sha256:442a9bfef4dddd9514e6ccbb7b7c46bd3f664396c4eaf6461ad82432480350dd", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool", "path": "documentation/data-sources/cdn_loadbalancer/properties/origin_pool/more_origin_options/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2100100112333133-0231133201332010-3330303123020131-2331222323022102-0213320032321313-1122011332212031-0323101302201211-0202110320102031", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "more_origin_options"], "schema_version": 1, "sections": [{"aliases": ["origin pool more origin options enable byte range request"], "anchor": "schema-origin_pool--more_origin_options--enable_byte_range_request", "description": "Choice to enable/disable byte range requests towards origin.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "more_origin_options", "enable_byte_range_request"], "syntax": "attribute", "type": "bool"}, {"aliases": ["origin pool more origin options websocket proxy"], "anchor": "schema-origin_pool--more_origin_options--websocket_proxy", "description": "Option to enable proxying of websocket connections to the origin server.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:origin_pool:more_origin_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "more_origin_options", "websocket_proxy"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/origin_pool/more_origin_options/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration parameter for more origin options.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.more_origin_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/)
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

- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/origin_pool/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
