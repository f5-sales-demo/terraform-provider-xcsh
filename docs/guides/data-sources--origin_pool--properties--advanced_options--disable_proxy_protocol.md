---
page_title: "advanced_options.disable_proxy_protocol"
subcategory: "Load Balancing"
description: "advanced_options.disable_proxy_protocol for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1007, "body_sha256": "sha256:7f8db2d5e4573f9ba38162af737ec6fb1d44ac4a65f7fa8d65edc12589b9b5fe", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_proxy_protocol", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_proxy_protocol", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "docs/guides/data-sources--origin_pool--properties--advanced_options--disable_proxy_protocol.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "disable_proxy_protocol"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/disable_proxy_protocol/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.disable_proxy_protocol for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.disable_proxy_protocol

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- advanced_options.disable_proxy_protocol

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable proxy protocol.

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

- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
