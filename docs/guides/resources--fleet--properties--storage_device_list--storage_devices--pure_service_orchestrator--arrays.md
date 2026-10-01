---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator.arrays"
subcategory: ""
description: "storage_device_list.storage_devices.pure_service_orchestrator.arrays for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2148, "body_sha256": "sha256:6f0e75a45655d630e9ea59ee0a94d5c5f775874bf6ad5479f812fb07c9fbe536", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "path": "docs/guides/resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.pure_service_orchestrator.arrays for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator.arrays

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_device_list](resources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Arrays Configuration. Device configuration for PSO Arrays.

Upstream description:

Device configuration for PSO Arrays.

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
arrays {
  # Configure direct properties listed below.
}
```

## Direct properties

- [flash_array](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md): complete subsection reference.

- [flash_blade](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md)
- [xcsh_fleet](../resources/fleet.md)
