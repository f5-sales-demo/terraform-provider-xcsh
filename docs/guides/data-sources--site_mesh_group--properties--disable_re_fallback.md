---
page_title: "disable_re_fallback"
subcategory: "Infrastructure"
description: "disable_re_fallback for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1279, "body_sha256": "sha256:b9021653134ac9e95116aa2668d175798b7aded293f7e91208c499927bf425cf", "canonical_id": "xcsh-docs:data-sources:site_mesh_group:properties:disable_re_fallback", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:disable_re_fallback", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "docs/guides/data-sources--site_mesh_group--properties--disable_re_fallback.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_re_fallback"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/disable_re_fallback/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_re_fallback for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_re_fallback

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
- [Property reference](data-sources--site_mesh_group--reference.md)
- disable_re_fallback

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_re\_fallback, enable\_re\_fallback; Default: disable\_re\_fallback\] Configuration
parameter for disable re fallback.

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

- [disable_re_fallback](data-sources--site_mesh_group--properties--disable_re_fallback.md#section)
- [enable_re_fallback](data-sources--site_mesh_group--properties--enable_re_fallback.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--site_mesh_group--reference.md)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
