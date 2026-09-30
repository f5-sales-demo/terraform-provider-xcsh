---
page_title: "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password"
subcategory: ""
description: "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2842, "body_sha256": "sha256:c7227be8de409a167fcad5adf495433e7d690bdbbd3c800ec7dc23d0f6f320f2", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

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
iscsi_chap_password {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/clear_secret_info/): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/blindfold_secret_info/)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/clear_secret_info/)
- [storage_device_list.storage_devices.hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
