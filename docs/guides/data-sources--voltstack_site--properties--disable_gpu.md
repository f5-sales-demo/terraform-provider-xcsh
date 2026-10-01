---
page_title: "disable_gpu"
subcategory: ""
description: "disable_gpu for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1284, "body_sha256": "sha256:19016f20675ead224352fdb747e8965b433bd4fffb032bb4daaefdbfd3cf8c58", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:disable_gpu", "child_ids": [], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:disable_gpu", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--disable_gpu.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_gpu"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/disable_gpu/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_gpu for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_gpu

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- disable_gpu

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_gpu, enable\_gpu, enable\_vgpu; Default: disable\_gpu\] Configuration parameter
for disable gpu.

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

- [disable_gpu](data-sources--voltstack_site--properties--disable_gpu.md#section)
- [enable_gpu](data-sources--voltstack_site--properties--enable_gpu.md#section)
- [enable_vgpu](data-sources--voltstack_site--properties--enable_vgpu.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
