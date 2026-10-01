---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token"
subcategory: ""
description: "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 3162, "body_sha256": "sha256:f88ce1fd31535d8bf320dac7defe0007c044c4a9c394144f4b648595516f0add", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token:blindfold_secret_info", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "path": "docs/guides/data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "flash_blades", "api_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_device_list](data-sources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](data-sources--fleet--properties--storage_device_list--storage_devices.md)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

## Direct properties

- [blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--blindfold_secret_info.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--api_token--clear_secret_info.md)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md)
- [xcsh_fleet](../data-sources/fleet.md)
