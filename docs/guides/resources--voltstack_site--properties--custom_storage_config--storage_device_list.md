---
page_title: "custom_storage_config.storage_device_list"
subcategory: ""
description: "custom_storage_config.storage_device_list for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1342, "body_sha256": "sha256:bc72c22ce45fe5b4c76b7bedde9dc73069f1a58b6e9da2ecf2ed70703340e54b", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_device_list.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- custom_storage_config.storage_device_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

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
storage_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_devices](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
