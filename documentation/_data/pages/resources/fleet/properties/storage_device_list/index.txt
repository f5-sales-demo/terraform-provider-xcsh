---
page_title: "storage_device_list"
subcategory: ""
description: "Add additional custom storage classes in Kubernetes for this fleet."
xcsh_docs: {"aliases": ["storage device list"], "body_bytes": 1348, "body_sha256": "sha256:20427725b642497deac87018b311f194daf3eb27a30f7cb4e8a72f315c2eb876", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/storage_device_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list"], "schema_version": 1, "sections": [{"aliases": ["storage devices"], "anchor": "section", "description": "List of custom storage devices.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--storage_device", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:RequiredListObjectAttributes:storage_device", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Add additional custom storage classes in Kubernetes for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- storage_device_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
storage_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
