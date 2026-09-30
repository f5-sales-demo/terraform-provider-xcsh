---
page_title: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 4765, "body_sha256": "sha256:d7dd80b64a57752d79a4594f2f30c8e53348ff0323dfff8264eb62f4bacb21a5", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token:blindfold_secret_info", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token:clear_secret_info"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "flash_blades", "api_token"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/)
- [custom_storage_config.storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/clear_secret_info/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/blindfold_secret_info/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/clear_secret_info/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
