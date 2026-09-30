---
page_title: "origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator"
subcategory: ""
description: "origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 2455, "body_sha256": "sha256:f369c6487cbc16275761909948bc7b469bd5f9eb36b67044ed55c02ac5621950", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:site_locator", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:site_locator:site", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:site_locator:virtual_site"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service:site_locator", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:origin_servers:k8s_service", "path": "docs/guides/data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "origin_servers", "k8s_service", "site_locator"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [origin_pools](data-sources--bigip_http_proxy--properties--origin_pools.md)
- [origin_pools.pools](data-sources--bigip_http_proxy--properties--origin_pools--pools.md)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator

<a id="section"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

## Direct properties

- [site](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site.md): complete subsection reference.

- [virtual_site](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md): complete subsection reference.

## Next pages

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site.md)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
