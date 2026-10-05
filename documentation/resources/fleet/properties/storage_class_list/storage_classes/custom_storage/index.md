---
page_title: "storage_class_list.storage_classes.custom_storage"
subcategory: ""
description: "Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site."
xcsh_docs: {"aliases": ["storage class list storage classes custom storage"], "body_bytes": 2706, "body_sha256": "sha256:990f0212d440e04f0ec80b10a32efa9fb48a869a1f33329222b3c2751b4f674d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:custom_storage", "parent_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "path": "documentation/resources/fleet/properties/storage_class_list/storage_classes/custom_storage/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3003121013331300-1300220222103221-2021323122133331-0011011023003132-0213101122200221-0222333220013210-1333213000113113-3202103202223221", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "custom_storage"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes custom storage yaml"], "anchor": "schema-storage_class_list--storage_classes--custom_storage--yaml", "description": "K8s YAML for StorageClass.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:custom_storage", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "custom_storage", "yaml"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/storage_classes/custom_storage/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.custom_storage

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- storage_class_list.storage_classes.custom_storage

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

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
custom_storage {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_class_list--storage_classes--custom_storage--yaml"></a>

### yaml property

Type: `"string"`. Optional.

Storage Class YAML. K8s YAML for StorageClass.

Upstream description:

K8s YAML for StorageClass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
