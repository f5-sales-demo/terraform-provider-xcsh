---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2485, "body_sha256": "sha256:f257c7ccc77ab0bc80359fe59c217914a9016474f4bf4571075768ff367d047f", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "path": "docs/guides/data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [origin_pools](data-sources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](data-sources--bigip_http_proxy--properties--origin_pools--pools.md)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

<a id="section"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

## Direct properties

- [no_snat_pool](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--no_snat_pool.md): complete subsection reference.

- [snat_pool](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool.md): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--no_snat_pool.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool.md)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
