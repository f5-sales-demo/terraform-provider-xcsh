---
page_title: "advanced_options.no_panic_threshold"
subcategory: "Load Balancing"
description: "advanced_options.no_panic_threshold for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 958, "body_sha256": "sha256:2913f8b4d9208213df4119e121e08af467068a3d350c78b63e4689fb5680699a", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:no_panic_threshold", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:no_panic_threshold", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "docs/guides/data-sources--origin_pool--properties--advanced_options--no_panic_threshold.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "no_panic_threshold"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/no_panic_threshold/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.no_panic_threshold for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advanced_options.no_panic_threshold

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- advanced_options.no_panic_threshold

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no panic threshold. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
