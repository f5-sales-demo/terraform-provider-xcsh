---
page_title: "custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["custom storage config storage device list storage devices hpe storage iscsi chap password"], "body_bytes": 3552, "body_sha256": "sha256:cf9cf948365ab0a8fbda01861a140b038f2d1466f99f5a02c9dfce5eb475178c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3002023122213001-1201323231321233-1331331032312113-0112303002001310-0211221330312123-1033033230333000-3330013010023110-2320220313320220", "registry_path": "docs/guides/resources--voltstack_site--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage device list storage devices hpe storage iscsi chap password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "type": "requires"}], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["custom storage config storage device list storage devices hpe storage iscsi chap password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--clear_secret_info--url", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info", "type": "requires"}], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/)
- [custom_storage_config.storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/clear_secret_info/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/blindfold_secret_info/)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/clear_secret_info/)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/hpe_storage/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
