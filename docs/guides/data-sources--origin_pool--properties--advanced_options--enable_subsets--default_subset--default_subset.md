---
page_title: "advanced_options.enable_subsets.default_subset.default_subset"
subcategory: "Load Balancing"
description: "advanced_options.enable_subsets.default_subset.default_subset for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1698, "body_sha256": "sha256:64f826125c8438fafa36bb4b252e196661a877d97add81d5dec6e126e1a0e751", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:default_subset:default_subset", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:default_subset:default_subset", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "path": "docs/guides/data-sources--origin_pool--properties--advanced_options--enable_subsets--default_subset--default_subset.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "enable_subsets", "default_subset", "default_subset"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/enable_subsets/default_subset/default_subset/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.enable_subsets.default_subset.default_subset for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.enable_subsets.default_subset.default_subset

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- [advanced_options.enable_subsets](data-sources--origin_pool--properties--advanced_options--enable_subsets.md)
- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--properties--advanced_options--enable_subsets--default_subset.md)
- advanced_options.enable_subsets.default_subset.default_subset

<a id="section"></a>

Type: `"single"`. Computed.

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. Which gets used when route specifies no metadata
or no subset matching the metadata exists.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "32"
  }
}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_options.enable_subsets.default_subset](data-sources--origin_pool--properties--advanced_options--enable_subsets--default_subset.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
