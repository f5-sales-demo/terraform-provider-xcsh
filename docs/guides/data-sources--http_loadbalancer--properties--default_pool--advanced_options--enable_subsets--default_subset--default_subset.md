---
page_title: "default_pool.advanced_options.enable_subsets.default_subset.default_subset"
subcategory: "Load Balancing"
description: "default_pool.advanced_options.enable_subsets.default_subset.default_subset for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1965, "body_sha256": "sha256:84f7dc8c3e640e2081221586696590e5c872b5ef3493de90d889550e59af2649", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset:default_subset", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset:default_subset", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "path": "docs/guides/data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--default_subset--default_subset.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets", "default_subset", "default_subset"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/default_subset/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options.enable_subsets.default_subset.default_subset for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets.default_subset.default_subset

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [default_pool](data-sources--http_loadbalancer--properties--default_pool.md)
- [default_pool.advanced_options](data-sources--http_loadbalancer--properties--default_pool--advanced_options.md)
- [default_pool.advanced_options.enable_subsets](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets.md)
- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--default_subset.md)
- default_pool.advanced_options.enable_subsets.default_subset.default_subset

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

- [default_pool.advanced_options.enable_subsets.default_subset](data-sources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets--default_subset.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
