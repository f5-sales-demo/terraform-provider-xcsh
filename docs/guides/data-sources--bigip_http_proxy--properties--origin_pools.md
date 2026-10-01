---
page_title: "origin_pools"
subcategory: ""
description: "origin_pools for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 963, "body_sha256": "sha256:b2b671884c048884c4f7f7a0cd6e932db7cbde4afb9d451218667c35d99df45b", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "path": "docs/guides/data-sources--bigip_http_proxy--properties--origin_pools.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- origin_pools

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for origin pools.

Upstream description:

List of Origin Pools.

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

- [pools](data-sources--bigip_http_proxy--properties--origin_pools--pools.md): complete subsection reference.

## Next pages

- [origin_pools.pools](data-sources--bigip_http_proxy--properties--origin_pools--pools.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
