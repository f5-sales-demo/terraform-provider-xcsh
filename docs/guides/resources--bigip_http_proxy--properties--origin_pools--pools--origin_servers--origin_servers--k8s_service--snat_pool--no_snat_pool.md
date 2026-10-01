---
page_title: "origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1993, "body_sha256": "sha256:c745f65197c56558536fa7e85ef05222ee18898267acd3d22f1d32764940f368", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:snat_pool:no_snat_pool", "child_ids": [], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:snat_pool", "path": "docs/guides/resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--no_snat_pool.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "k8s_service", "snat_pool", "no_snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](resources--bigip_http_proxy--properties--origin_pools--pools.md)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool.md)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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

Terraform syntax:

```terraform
no_snat_pool = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
