---
page_title: "storage_device_list.storage_devices.netapp_trident"
subcategory: ""
description: "storage_device_list.storage_devices.netapp_trident for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2123, "body_sha256": "sha256:96f9326242b26f1c61c85e13708bb247a3ffe4531d5e91237ea872481333f315", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "path": "docs/guides/resources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.netapp_trident for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_device_list](resources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- storage_device_list.storage_devices.netapp_trident

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for NetApp Trident Storage.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("netapp_backend_ontap_nas",
    "netapp_backend_ontap_san")}
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
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

## Direct properties

- [netapp_backend_ontap_nas](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md): complete subsection reference.

- [netapp_backend_ontap_san](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- [xcsh_fleet](../resources/fleet.md)
