---
page_title: "storage_class_list"
subcategory: ""
description: "Add additional custom storage classes in Kubernetes for this fleet."
xcsh_docs: {"aliases": ["storage class list"], "body_bytes": 1342, "body_sha256": "sha256:48bbd3e9ff420c9b706cc8791dbfb971f83865cfd78f57311d3bf09909dec802", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/storage_class_list/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes"], "anchor": "section", "description": "List of custom storage classes.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "schema-storage_class_list--storage_classes--storage_class_name", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name,storage_device", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "type": "requires"}, {"anchor": "schema-storage_class_list--storage_classes--storage_device", "enforcement": "provider-schema", "group": "storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name,storage_device", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "type": "requires"}], "schema_path": ["storage_class_list", "storage_classes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Add additional custom storage classes in Kubernetes for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- storage_class_list

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
storage_class_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/): complete subsection reference.

## Next pages

- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
