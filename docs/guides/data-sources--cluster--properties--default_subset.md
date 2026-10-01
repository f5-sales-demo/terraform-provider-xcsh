---
page_title: "default_subset"
subcategory: ""
description: "default_subset for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1261, "body_sha256": "sha256:721349badc73f4f6e27c75fd20438e360627865960af3b950b8020acbee8331b", "canonical_id": "xcsh-docs:data-sources:cluster:properties:default_subset", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:default_subset", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "docs/guides/data-sources--cluster--properties--default_subset.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_subset"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/default_subset/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_subset for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_subset

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- [Property reference](data-sources--cluster--reference.md)
- default_subset

<a id="section"></a>

Type: `"single"`. Computed.

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

Upstream description:

List of key-value pairs that define default subset. This subset can be referred in fallback\_policy
which gets used when route specifies no metadata or no subset matching the metadata exists.

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

- [Property reference](data-sources--cluster--reference.md)
- [xcsh_cluster](../data-sources/cluster.md)
