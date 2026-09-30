---
page_title: "default_subset"
subcategory: ""
description: "default_subset for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1162, "body_sha256": "sha256:fb6aa43ffef537d3a26c71f5f51dcf256cac6ea603a5762cdb09f4172636da9d", "canonical_id": "xcsh-docs:data-sources:cluster:properties:default_subset", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:default_subset", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "docs/guides/data-sources--cluster--properties--default_subset.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_subset"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/default_subset/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_subset for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
