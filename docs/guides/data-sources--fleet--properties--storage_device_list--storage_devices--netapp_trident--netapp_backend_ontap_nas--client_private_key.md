---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key"
subcategory: ""
description: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2595, "body_sha256": "sha256:afb2836454b2eab91a6d48411eadb83fa536e2e37b1ae04e8ea148c791f8e754", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:blindfold_secret_info", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "path": "docs/guides/data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_device_list](data-sources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](data-sources--fleet--properties--storage_device_list--storage_devices.md)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

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

- [blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--blindfold_secret_info.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_private_key--clear_secret_info.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md)
- [xcsh_fleet](../data-sources/fleet.md)
