---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret"
subcategory: ""
description: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 3172, "body_sha256": "sha256:48be1588af465f013ac5535534b9bb0606d777ce9e2b768c047f3c6b198bddc0", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret:blindfold_secret_info", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret:clear_secret_info"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "path": "docs/guides/resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_target_initiator_secret"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_target_initiator_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_device_list](resources--fleet--properties--storage_device_list.md)
- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

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
chap_target_initiator_secret {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--clear_secret_info.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--blindfold_secret_info.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret--clear_secret_info.md)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md)
- [xcsh_fleet](../resources/fleet.md)
