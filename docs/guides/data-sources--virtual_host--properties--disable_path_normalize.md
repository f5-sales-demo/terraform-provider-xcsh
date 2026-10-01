---
page_title: "disable_path_normalize"
subcategory: ""
description: "disable_path_normalize for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1252, "body_sha256": "sha256:8de4cbe143946010f492698794625ea5e821205d4ce01083a9c9889aa333c082", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:disable_path_normalize", "child_ids": [], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:disable_path_normalize", "parent_id": "xcsh-docs:data-sources:virtual_host:reference", "path": "docs/guides/data-sources--virtual_host--properties--disable_path_normalize.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_path_normalize"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/disable_path_normalize/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_path_normalize for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_path_normalize

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- disable_path_normalize

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_path\_normalize, enable\_path\_normalize; Default: disable\_path\_normalize\]
Enable this option

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

- [disable_path_normalize](data-sources--virtual_host--properties--disable_path_normalize.md#section)
- [enable_path_normalize](data-sources--virtual_host--properties--enable_path_normalize.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--virtual_host--reference.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
