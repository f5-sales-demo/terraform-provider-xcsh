---
page_title: "storage_device_list.storage_devices.custom_storage"
subcategory: ""
description: "storage_device_list.storage_devices.custom_storage for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1177, "body_sha256": "sha256:07bddac06009a21279d721e8bd54ac41065a39664b56dc502c896111e45819b2", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "path": "docs/guides/resources--fleet--properties--storage_device_list--storage_devices--custom_storage.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "custom_storage"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/custom_storage/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.custom_storage for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.custom_storage

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_device_list](resources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- storage_device_list.storage_devices.custom_storage

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for custom storage.

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
custom_storage = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- [xcsh_fleet](../resources/fleet.md)
