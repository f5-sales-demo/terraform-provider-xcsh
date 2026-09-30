---
page_title: "coalescing_options.default_coalescing"
subcategory: ""
description: "coalescing_options.default_coalescing for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 961, "body_sha256": "sha256:e7fc0219ccf52d3bb0ca486971fbf97a6b5ca05d2acf9fa4b8db1323a61f91e7", "canonical_id": "xcsh-docs:resources:virtual_host:properties:coalescing_options:default_coalescing", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:coalescing_options:default_coalescing", "parent_id": "xcsh-docs:resources:virtual_host:properties:coalescing_options", "path": "docs/guides/resources--virtual_host--properties--coalescing_options--default_coalescing.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["coalescing_options", "default_coalescing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/coalescing_options/default_coalescing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "coalescing_options.default_coalescing for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# coalescing_options.default_coalescing

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [coalescing_options](resources--virtual_host--properties--coalescing_options.md)
- coalescing_options.default_coalescing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [coalescing_options](resources--virtual_host--properties--coalescing_options.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
