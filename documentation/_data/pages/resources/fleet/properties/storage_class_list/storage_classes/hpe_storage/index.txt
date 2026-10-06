---
page_title: "storage_class_list.storage_classes.hpe_storage"
subcategory: ""
description: "Storage class Device configuration for HPE Storage."
xcsh_docs: {"aliases": ["storage class list storage classes hpe storage"], "body_bytes": 12801, "body_sha256": "sha256:2f42ee43dfa77f19a23ba47b1a6bb6252f47fdffa8162fc9d814cdb7f66faaa9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "parent_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes", "path": "documentation/resources/fleet/properties/storage_class_list/storage_classes/hpe_storage/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-2032302011222201-1310233322022302-0123100203223233-1210011333013002-2322333333201313-1212222011311012-2021103000021222-3002302213031013", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "hpe_storage"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes hpe storage allow mutations"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--allow_mutations", "description": "Mutation can override specified parameters.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "allow_mutations"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage allow overrides"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--allow_overrides", "description": "PVC can override specified parameters.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "allow_overrides"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage dedupe enabled"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--dedupe_enabled", "description": "Indicates that the volume should enable deduplication.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "dedupe_enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class list storage classes hpe storage description spec"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--description_spec", "description": "The SecretName parameter is used to identify name of secret to identify backend storage's auth information.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage destroy on delete"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--destroy_on_delete", "description": "Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is deleted.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "destroy_on_delete"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class list storage classes hpe storage encrypted"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--encrypted", "description": "Indicates that the volume should be encrypted.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "encrypted"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class list storage classes hpe storage folder"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--folder", "description": "The name of the folder in which to place the volume.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "folder"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage limit iops"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--limit_iops", "description": "The IOPS limit of the volume.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "limit_iops"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage limit mbps"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--limit_mbps", "description": "The IOPS limit of the volume.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "limit_mbps"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage performance policy"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--performance_policy", "description": "The name of the performance policy to assign to the volume.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "performance_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage pool"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--pool", "description": "The name of the pool in which to place the volume.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "pool"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage protection template"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--protection_template", "description": "The name of the performance policy to assign to the volume.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "protection_template"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage secret name"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--secret_name", "description": "The SecretName parameter is used to identify name of secret to identify backend storage's auth information.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "secret_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage secret namespace"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--secret_namespace", "description": "The SecretNamespace parameter is used to identify name of namespace where secret resides.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "secret_namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage sync on detach"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--sync_on_detach", "description": "Indicates that a snapshot of the volume should be synced to the replication partner each time it is detached from a node.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "sync_on_detach"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class list storage classes hpe storage thick"], "anchor": "schema-storage_class_list--storage_classes--hpe_storage--thick", "description": "Indicates that the volume should be thick provisioned.", "document_id": "xcsh-docs:resources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage", "thick"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_class_list/storage_classes/hpe_storage/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Storage class Device configuration for HPE Storage.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.hpe_storage

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/)
- [storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_class_list/storage_classes/)
- storage_class_list.storage_classes.hpe_storage

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for HPE Storage.

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
hpe_storage {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_class_list--storage_classes--hpe_storage--allow_mutations"></a>

### allow_mutations property

Type: `"string"`. Optional.

Mutation can override specified parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--allow_overrides"></a>

### allow_overrides property

Type: `"string"`. Optional.

AllowOverrides. PVC can override specified parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--dedupe_enabled"></a>

### dedupe_enabled property

Type: `"bool"`. Optional.

Indicates that the volume should enable deduplication.

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--destroy_on_delete"></a>

### destroy_on_delete property

Type: `"bool"`. Optional.

Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is
deleted.

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--encrypted"></a>

### encrypted property

Type: `"bool"`. Optional.

Indicates that the volume should be encrypted.

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--folder"></a>

### folder property

Type: `"string"`. Optional.

The name of the folder in which to place the volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--limit_iops"></a>

### limit_iops property

Type: `"string"`. Optional.

LimitIops. The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--limit_mbps"></a>

### limit_mbps property

Type: `"string"`. Optional.

LimitMbps. The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--performance_policy"></a>

### performance_policy property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--pool"></a>

### pool property

Type: `"string"`. Optional.

The name of the pool in which to place the volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--protection_template"></a>

### protection_template property

Type: `"string"`. Optional.

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--secret_name"></a>

### secret_name property

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--secret_namespace"></a>

### secret_namespace property

Type: `"string"`. Optional.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-storage_class_list--storage_classes--hpe_storage--sync_on_detach"></a>

### sync_on_detach property

Type: `"bool"`. Optional.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--thick"></a>

### thick property

Type: `"bool"`. Optional.

Indicates that the volume should be thick provisioned.

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
