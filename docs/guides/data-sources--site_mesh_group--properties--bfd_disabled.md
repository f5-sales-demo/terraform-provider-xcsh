---
page_title: "bfd_disabled"
subcategory: "Infrastructure"
description: "bfd_disabled for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1159, "body_sha256": "sha256:1e0e3984a0ba06fc9b35f6bcfb91063cde4aa67c0d0d2b18a6c5dc991cd7b405", "canonical_id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_disabled", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:bfd_disabled", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "docs/guides/data-sources--site_mesh_group--properties--bfd_disabled.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bfd_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/bfd_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bfd_disabled for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bfd_disabled

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
- [Property reference](data-sources--site_mesh_group--reference.md)
- bfd_disabled

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: bfd\_disabled, bfd\_enabled\] Enable this option

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

- [bfd_disabled](data-sources--site_mesh_group--properties--bfd_disabled.md#section)
- [bfd_enabled](data-sources--site_mesh_group--properties--bfd_enabled.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--site_mesh_group--reference.md)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
