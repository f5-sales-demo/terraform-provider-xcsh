---
page_title: "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["storage device list storage devices hpe storage iscsi chap password"], "body_bytes": 2941, "body_sha256": "sha256:1980c1756f13fa1737cec3ca51b9b9af5f0827dde2264579ccfe7beb01c60bf1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0220213210310012-3110112220202121-0102302122303013-3211121021300211-2332032122311201-0303302300223213-2311000121321111-3230211321012132", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices hpe storage iscsi chap password blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices hpe storage iscsi chap password clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_password--clear_secret_info--url", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
