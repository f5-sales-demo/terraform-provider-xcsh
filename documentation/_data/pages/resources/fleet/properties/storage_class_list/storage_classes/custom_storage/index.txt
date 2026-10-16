---
page_title: "storage_class_list.storage_classes.custom_storage"
subcategory: ""
description: "Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site."
xcsh_docs: {"aliases": ["storage class list storage classes custom storage"], "body_bytes": 2420, "body_sha256": "sha256:2dbc0de6ea2d080005f557b476ce1c6787621426e16fa5db7f8ac1abc56ec0c5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:custom_storage", "parent_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "path": "documentation/resources/fleet/properties/storage_class_list/storage_classes/custom_storage/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3003121013331300-1300220222103221-2021323122133331-0011011023003132-0213101122200221-0222333220013210-1333213000113113-3202103202223221", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "custom_storage"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes custom storage yaml"], "anchor": "schema-storage_class_list--storage_classes--custom_storage--yaml", "description": "K8s YAML for StorageClass.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:custom_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "custom_storage", "yaml"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/storage_classes/custom_storage/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["fleetCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
