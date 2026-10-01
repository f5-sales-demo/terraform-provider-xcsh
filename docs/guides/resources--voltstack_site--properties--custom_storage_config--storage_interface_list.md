---
page_title: "custom_storage_config.storage_interface_list"
subcategory: ""
description: "custom_storage_config.storage_interface_list for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1515, "body_sha256": "sha256:1e5ec9284286d914e8823ca739e765d1bc17357f1494f87ff6a1f2d9fca33359", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_interface_list.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- custom_storage_config.storage_interface_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure storage interfaces for this App Stack site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_interfaces")}
```

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
storage_interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_interfaces](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
