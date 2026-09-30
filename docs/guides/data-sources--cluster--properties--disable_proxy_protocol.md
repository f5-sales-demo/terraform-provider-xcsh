---
page_title: "disable_proxy_protocol"
subcategory: ""
description: "disable_proxy_protocol for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1242, "body_sha256": "sha256:51470f7bdca75689a564d9e57284fea1c1c5984ac90ef374c40b686ad288c8f2", "canonical_id": "xcsh-docs:data-sources:cluster:properties:disable_proxy_protocol", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:disable_proxy_protocol", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "docs/guides/data-sources--cluster--properties--disable_proxy_protocol.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_proxy_protocol"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/disable_proxy_protocol/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_proxy_protocol for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_proxy_protocol

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- [Property reference](data-sources--cluster--reference.md)
- disable_proxy_protocol

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_proxy\_protocol, proxy\_protocol\_v1, proxy\_protocol\_v2; Default:
disable\_proxy\_protocol\] Configuration parameter for disable proxy protocol.

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

OneOf alternatives in this subsection:

- [disable_proxy_protocol](data-sources--cluster--properties--disable_proxy_protocol.md#section)
- [proxy_protocol_v1](data-sources--cluster--properties--proxy_protocol_v1.md#section)
- [proxy_protocol_v2](data-sources--cluster--properties--proxy_protocol_v2.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--cluster--reference.md)
- [xcsh_cluster](../data-sources/cluster.md)
