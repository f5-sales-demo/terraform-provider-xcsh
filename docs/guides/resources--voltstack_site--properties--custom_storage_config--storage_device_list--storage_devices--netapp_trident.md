---
page_title: "custom_storage_config.storage_device_list.storage_devices.netapp_trident"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.netapp_trident for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2538, "body_sha256": "sha256:f8e536231a7048a7d41e9edc5f8885894123a3398ed35700b0f86a683c9499d2", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.netapp_trident for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_device_list.storage_devices.netapp_trident

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_device_list](resources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident

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

- [netapp_backend_ontap_nas](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md): complete subsection reference.

- [netapp_backend_ontap_san](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
