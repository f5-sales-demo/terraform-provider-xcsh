---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator.arrays"
subcategory: ""
description: "Device configuration for PSO Arrays."
xcsh_docs: {"aliases": ["storage device list storage devices pure service orchestrator arrays"], "body_bytes": 1809, "body_sha256": "sha256:0688bd655d60143cd8c385b6da4fd5359d6d296e82d31153939ae37d1ac7cdde", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices pure service orchestrator arrays flash array"], "anchor": "section", "description": "Specify what storage flash arrays should be managed the plugin.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_type", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "type": "requires"}, {"anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--iscsi_login_timeout", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "type": "requires"}, {"anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--san_type", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices pure service orchestrator arrays flash blade"], "anchor": "section", "description": "Specify what storage flash blades should be managed the plugin.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade:RequiredObjectAttributes:flash_blades", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Device configuration for PSO Arrays.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["fleetCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator.arrays

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Arrays Configuration. Device configuration for PSO Arrays.

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

- [flash_array](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/): complete subsection reference.

- [flash_blade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/): complete subsection reference.
