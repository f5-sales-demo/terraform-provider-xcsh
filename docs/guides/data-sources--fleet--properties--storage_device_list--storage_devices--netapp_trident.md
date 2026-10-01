---
page_title: "storage_device_list.storage_devices.netapp_trident"
subcategory: ""
description: "storage_device_list.storage_devices.netapp_trident for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1832, "body_sha256": "sha256:7c341997f93120d9bc65bc0f5a3aa11d4fd3ceee63f2562d7f6cc8dc9c3afcab", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices", "path": "docs/guides/data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.netapp_trident for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_device_list](data-sources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](data-sources--fleet--properties--storage_device_list--storage_devices.md)
- storage_device_list.storage_devices.netapp_trident

<a id="section"></a>

Type: `"single"`. Computed.

Device configuration for NetApp Trident Storage.

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

## Direct properties

- [netapp_backend_ontap_nas](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md): complete subsection reference.

- [netapp_backend_ontap_san](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [storage_device_list.storage_devices](data-sources--fleet--properties--storage_device_list--storage_devices.md)
- [xcsh_fleet](../data-sources/fleet.md)
