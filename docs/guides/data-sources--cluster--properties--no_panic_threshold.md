---
page_title: "no_panic_threshold"
subcategory: ""
description: "no_panic_threshold for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1195, "body_sha256": "sha256:24c5b7acec5583de8f51073e5c2ee330fbc8e080bcd5edaa4432f0d1173c2570", "canonical_id": "xcsh-docs:data-sources:cluster:properties:no_panic_threshold", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:no_panic_threshold", "parent_id": "xcsh-docs:data-sources:cluster:reference", "path": "docs/guides/data-sources--cluster--properties--no_panic_threshold.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_panic_threshold"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/no_panic_threshold/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_panic_threshold for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_panic_threshold

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- [Property reference](data-sources--cluster--reference.md)
- no_panic_threshold

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_panic\_threshold, panic\_threshold; Default: no\_panic\_threshold\] Configuration
parameter for no panic threshold.

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

- [no_panic_threshold](data-sources--cluster--properties--no_panic_threshold.md#section)
- [panic_threshold](data-sources--cluster--reference.md#schema-panic_threshold)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--cluster--reference.md)
- [xcsh_cluster](../data-sources/cluster.md)
