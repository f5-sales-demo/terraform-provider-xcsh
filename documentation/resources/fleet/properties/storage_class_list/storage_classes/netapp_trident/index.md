---
page_title: "storage_class_list.storage_classes.netapp_trident"
subcategory: ""
description: "Storage class Device configuration for NetApp Trident."
xcsh_docs: {"aliases": ["storage class list storage classes netapp trident"], "body_bytes": 2808, "body_sha256": "sha256:b7a3764975ad671367b9174c5a5fd542488f156a594ec658167797f3a8b7cccb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "parent_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "path": "documentation/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "netapp_trident"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes netapp trident selector"], "anchor": "section", "description": "Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a volume. The volume will have the aspects defined in the chosen virtual pool.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident:selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "netapp_trident", "selector"], "syntax": "block", "type": "object"}, {"aliases": ["storage class list storage classes netapp trident storage pools"], "anchor": "schema-storage_class_list--storage_classes--netapp_trident--storage_pools", "description": "The storagePools parameter is used to further restrict the set of pools that match any specified attributes.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "netapp_trident", "storage_pools"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Storage class Device configuration for NetApp Trident.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.netapp_trident

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- storage_class_list.storage_classes.netapp_trident

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
netapp_trident {
  # Configure direct properties listed below.
}
```

## Direct properties

- [selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/): complete subsection reference.

<a id="schema-storage_class_list--storage_classes--netapp_trident--storage_pools"></a>

### storage_pools property

Type: `"string"`. Optional.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [storage_class_list.storage_classes.netapp_trident.selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/netapp_trident/selector/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
