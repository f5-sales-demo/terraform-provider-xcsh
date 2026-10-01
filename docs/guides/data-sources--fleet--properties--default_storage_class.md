---
page_title: "default_storage_class"
subcategory: ""
description: "default_storage_class for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1213, "body_sha256": "sha256:c25692a8a4efae6f33fd65b2f69c42542d459f58b595f8c24f8ee1757918df11", "canonical_id": "xcsh-docs:data-sources:fleet:properties:default_storage_class", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:default_storage_class", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "docs/guides/data-sources--fleet--properties--default_storage_class.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_storage_class"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/default_storage_class/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_storage_class for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_storage_class

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- default_storage_class

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_storage\_class, storage\_class\_list; Default: default\_storage\_class\]
Configuration parameter for default storage class.

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

- [default_storage_class](data-sources--fleet--properties--default_storage_class.md#section)
- [storage_class_list](data-sources--fleet--properties--storage_class_list.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--fleet--reference.md)
- [xcsh_fleet](../data-sources/fleet.md)
