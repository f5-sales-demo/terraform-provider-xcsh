---
page_title: "storage_class_list.storage_classes.netapp_trident"
subcategory: ""
description: "Storage class Device configuration for NetApp Trident."
xcsh_docs: {"aliases": ["storage class list storage classes netapp trident"], "body_bytes": 2091, "body_sha256": "sha256:a9e4ce938ff36ae901d05b6c4b828fe6342f7a6fc44fedfc903a5b7bb6eb6f5e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "path": "documentation/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2203313120110100-3230202300332001-2221002111100231-0300101000123331-3101321212203303-2333320001101321-0131102322311203-0002311222210022", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "netapp_trident"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes netapp trident selector"], "anchor": "section", "description": "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "netapp_trident", "selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage class list storage classes netapp trident storage pools"], "anchor": "schema-storage_class_list--storage_classes--netapp_trident--storage_pools", "description": "The storagePools parameter is used to further restrict the set of pools that match any specified attributes.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "netapp_trident", "storage_pools"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Storage class Device configuration for NetApp Trident.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["fleetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.netapp_trident

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/)
- storage_class_list.storage_classes.netapp_trident

<a id="section"></a>

Type: `"single"`. Computed.

Storage class Device configuration for NetApp Trident.

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

## Direct properties

- [selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/): complete subsection reference.

<a id="schema-storage_class_list--storage_classes--netapp_trident--storage_pools"></a>

### storage_pools property

Type: `"string"`. Computed.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-20T00:10:01+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```
