---
page_title: "origin_servers.k8s_service.inside_network"
subcategory: "Load Balancing"
description: "origin_servers.k8s_service.inside_network for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1122, "body_sha256": "sha256:4b5d31ea0b54cc7951e519deecb632e83b92f05e40284a6f39b7e601062077e1", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service:inside_network", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service:inside_network", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:origin_servers:k8s_service", "path": "docs/guides/data-sources--origin_pool--properties--origin_servers--k8s_service--inside_network.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "k8s_service", "inside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/origin_servers/k8s_service/inside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.k8s_service.inside_network for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.k8s_service.inside_network

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [origin_servers](data-sources--origin_pool--properties--origin_servers.md)
- [origin_servers.k8s_service](data-sources--origin_pool--properties--origin_servers--k8s_service.md)
- origin_servers.k8s_service.inside_network

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

- [origin_servers.k8s_service](data-sources--origin_pool--properties--origin_servers--k8s_service.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
