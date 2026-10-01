---
page_title: "storage_class_list.storage_classes.hpe_storage"
subcategory: ""
description: "storage_class_list.storage_classes.hpe_storage for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 11318, "body_sha256": "sha256:17fd0b0771bb4c459f164054f10271277cf76c85c56a5b6b816ab4af0ebef788", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "path": "docs/guides/data-sources--fleet--properties--storage_class_list--storage_classes--hpe_storage.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_class_list", "storage_classes", "hpe_storage"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_class_list/storage_classes/hpe_storage/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_class_list.storage_classes.hpe_storage for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes.hpe_storage

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_class_list](data-sources--fleet--properties--storage_class_list.md)
- [storage_class_list.storage_classes](data-sources--fleet--properties--storage_class_list--storage_classes.md)
- storage_class_list.storage_classes.hpe_storage

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--allow_mutations"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--allow_overrides"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--dedupe_enabled"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

<a id="schema-storage_class_list--storage_classes--hpe_storage--destroy_on_delete"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--encrypted"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--folder"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--limit_iops"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--limit_mbps"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--performance_policy"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--pool"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--protection_template"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--secret_name"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--secret_namespace"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--sync_on_detach"></a>

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

<a id="schema-storage_class_list--storage_classes--hpe_storage--thick"></a>

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

- [storage_class_list.storage_classes](data-sources--fleet--properties--storage_class_list--storage_classes.md)
- [xcsh_fleet](../data-sources/fleet.md)
