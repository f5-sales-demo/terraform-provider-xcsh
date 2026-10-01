---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token"
subcategory: ""
description: "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 3426, "body_sha256": "sha256:20ce9f979ec056ddfd235eb915416705a0cd239824b6dad9af6b87c6795aa2fa", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token:blindfold_secret_info", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token:clear_secret_info"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "path": "docs/guides/resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "flash_blades", "api_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_device_list](resources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
api_token {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md)
- [xcsh_fleet](../resources/fleet.md)
