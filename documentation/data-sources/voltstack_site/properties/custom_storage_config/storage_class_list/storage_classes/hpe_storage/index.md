---
page_title: "custom_storage_config.storage_class_list.storage_classes.hpe_storage"
subcategory: ""
description: "Storage class Device configuration for HPE Storage."
xcsh_docs: {"aliases": ["custom storage config storage class list storage classes hpe storage"], "body_bytes": 12385, "body_sha256": "sha256:a08d4f30b4759635c477b9460b2c2949a74fade9c5bc133226e811752b000e03", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/hpe_storage/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1001010133313201-1021212203222132-1023231101320002-0011003123210100-2131032021320320-3303023111020310-2312213010202322-3210102210222120", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage class list storage classes hpe storage allow mutations"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--allow_mutations", "description": "Mutation can override specified parameters.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "allow_mutations"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage allow overrides"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--allow_overrides", "description": "PVC can override specified parameters.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "allow_overrides"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage dedupe enabled"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--dedupe_enabled", "description": "Indicates that the volume should enable deduplication.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "dedupe_enabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["custom storage config storage class list storage classes hpe storage description spec"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--description_spec", "description": "The SecretName parameter is used to identify name of secret to identify backend storage's auth information.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage destroy on delete"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--destroy_on_delete", "description": "Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is deleted.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "destroy_on_delete"], "syntax": "attribute", "type": "bool"}, {"aliases": ["custom storage config storage class list storage classes hpe storage encrypted"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--encrypted", "description": "Indicates that the volume should be encrypted.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "encrypted"], "syntax": "attribute", "type": "bool"}, {"aliases": ["custom storage config storage class list storage classes hpe storage folder"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--folder", "description": "The name of the folder in which to place the volume.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "folder"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage limit iops"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--limit_iops", "description": "The IOPS limit of the volume.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "limit_iops"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage limit mbps"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--limit_mbps", "description": "The IOPS limit of the volume.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "limit_mbps"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage performance policy"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--performance_policy", "description": "The name of the performance policy to assign to the volume.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "performance_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage pool"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--pool", "description": "The name of the pool in which to place the volume.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "pool"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage protection template"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--protection_template", "description": "The name of the performance policy to assign to the volume.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "protection_template"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage secret name"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--secret_name", "description": "The SecretName parameter is used to identify name of secret to identify backend storage's auth information.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "secret_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage secret namespace"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--secret_namespace", "description": "The SecretNamespace parameter is used to identify name of namespace where secret resides.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "secret_namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage class list storage classes hpe storage sync on detach"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--sync_on_detach", "description": "Indicates that a snapshot of the volume should be synced to the replication partner each time it is detached from a node.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "sync_on_detach"], "syntax": "attribute", "type": "bool"}, {"aliases": ["custom storage config storage class list storage classes hpe storage thick"], "anchor": "schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--thick", "description": "Indicates that the volume should be thick provisioned.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes", "hpe_storage", "thick"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/hpe_storage/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Storage class Device configuration for HPE Storage.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes.hpe_storage

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- custom_storage_config.storage_class_list.storage_classes.hpe_storage

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--allow_mutations"></a>

### allow_mutations property

Type: `"string"`. Computed.

Mutation can override specified parameters.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--allow_overrides"></a>

### allow_overrides property

Type: `"string"`. Computed.

AllowOverrides. PVC can override specified parameters.

Upstream description:

PVC can override specified parameters.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--dedupe_enabled"></a>

### dedupe_enabled property

Type: `"bool"`. Computed.

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--destroy_on_delete"></a>

### destroy_on_delete property

Type: `"bool"`. Computed.

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--encrypted"></a>

### encrypted property

Type: `"bool"`. Computed.

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--folder"></a>

### folder property

Type: `"string"`. Computed.

The name of the folder in which to place the volume.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--limit_iops"></a>

### limit_iops property

Type: `"string"`. Computed.

LimitIops. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--limit_mbps"></a>

### limit_mbps property

Type: `"string"`. Computed.

LimitMbps. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--performance_policy"></a>

### performance_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

The name of the performance policy to assign to the volume.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--pool"></a>

### pool property

Type: `"string"`. Computed.

The name of the pool in which to place the volume.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--protection_template"></a>

### protection_template property

Type: `"string"`. Computed.

The name of the performance policy to assign to the volume.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--secret_name"></a>

### secret_name property

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--secret_namespace"></a>

### secret_namespace property

Type: `"string"`. Computed.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--sync_on_detach"></a>

### sync_on_detach property

Type: `"bool"`. Computed.

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--hpe_storage--thick"></a>

### thick property

Type: `"bool"`. Computed.

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

## Next pages

- [custom_storage_config.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
