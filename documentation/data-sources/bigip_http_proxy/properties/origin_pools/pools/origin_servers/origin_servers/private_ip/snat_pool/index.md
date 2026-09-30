---
page_title: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3013, "body_sha256": "sha256:3fb79fae11e4d6fd601b860c996edf6c3c917ebc641170d517acbb32888e6509", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:no_snat_pool", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool:snat_pool"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip:snat_pool", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:private_ip", "path": "documentation/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "private_ip", "snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/)
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

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/snat_pool/): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/no_snat_pool/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/snat_pool/)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
