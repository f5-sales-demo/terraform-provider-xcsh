---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap initiator secret"], "body_bytes": 2815, "body_sha256": "sha256:b6d02c07e579e6f6c2458dafa80799787837172a6a2bf2fc78d534da86b95ef7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:blindfold_secret_info", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_initiator_secret/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0033032203023221-3230223303120013-3120220030200030-0022302213110221-0011310133330111-1221021010221230-0021101120100321-0002031102332020", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_initiator_secret"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap initiator secret blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:blindfold_secret_info", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_initiator_secret", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap initiator secret clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info--url", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:clear_secret_info", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_initiator_secret", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_initiator_secret/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fleetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
chap_initiator_secret {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_initiator_secret/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_initiator_secret/clear_secret_info/): complete subsection reference.
