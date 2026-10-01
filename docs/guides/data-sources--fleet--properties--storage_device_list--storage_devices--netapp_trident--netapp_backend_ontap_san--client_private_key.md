---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key"
subcategory: ""
description: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2595, "body_sha256": "sha256:d41feb98ebd4bdc63fd1c19f7b3699110074bd5e90f7426e30b826ac0aa229d6", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key:blindfold_secret_info", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "path": "docs/guides/data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "client_private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/client_private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_device_list](data-sources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](data-sources--fleet--properties--storage_device_list--storage_devices.md)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

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

- [blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--blindfold_secret_info.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_private_key--clear_secret_info.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [xcsh_fleet](../data-sources/fleet.md)
