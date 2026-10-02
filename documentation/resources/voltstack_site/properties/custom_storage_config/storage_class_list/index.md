---
page_title: "custom_storage_config.storage_class_list"
subcategory: ""
description: "Add additional custom storage classes in Kubernetes for this fleet."
xcsh_docs: {"aliases": ["custom storage config storage class list"], "body_bytes": 1691, "body_sha256": "sha256:a06312037428bc047d9427e6d19975d815ea14a826f05d55c60b6a7eebb1cd12", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_class_list/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3200223201220132-0311220310300003-2103003133003311-1323333023323123-3011330010121132-2203203012313212-2030201103311030-0213120200223032", "registry_path": "docs/guides/resources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list"], "schema_version": 1, "sections": [{"aliases": ["storage classes"], "anchor": "section", "description": "List of custom storage classes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_class_list--storage_classes--storage_class_name", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name,storage_device", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "type": "requires"}, {"anchor": "schema-custom_storage_config--storage_class_list--storage_classes--storage_device", "enforcement": "provider-schema", "group": "custom_storage_config.storage_class_list.storage_classes:RequiredListObjectAttributes:storage_class_name,storage_device", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "type": "requires"}], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_class_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Add additional custom storage classes in Kubernetes for this fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- custom_storage_config.storage_class_list

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

- [storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
