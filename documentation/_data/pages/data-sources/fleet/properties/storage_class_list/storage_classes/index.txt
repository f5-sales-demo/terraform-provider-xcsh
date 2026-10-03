---
page_title: "storage_class_list.storage_classes"
subcategory: ""
description: "List of custom storage classes."
xcsh_docs: {"aliases": ["storage class list storage classes"], "body_bytes": 8309, "body_sha256": "sha256:9319c8465e982cc1e15c7d211a09a5932d54fb05ab1346b416054e119ed864d3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:custom_storage", "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:pure_service_orchestrator"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list", "path": "documentation/data-sources/fleet/properties/storage_class_list/storage_classes/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3332030213120331-3133123310112000-0030211332231133-1313131033100011-1002001323111021-1322321023132033-3130112202202101-3033013311323132", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_class_list", "storage_classes"], "schema_version": 1, "sections": [{"aliases": ["storage class list storage classes advanced storage parameters"], "anchor": "schema-storage_class_list--storage_classes--advanced_storage_parameters", "description": "Map of parameter name and string value.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "advanced_storage_parameters"], "syntax": "attribute", "type": "map"}, {"aliases": ["storage class list storage classes allow volume expansion"], "anchor": "schema-storage_class_list--storage_classes--allow_volume_expansion", "description": "Allow volume expansion.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "allow_volume_expansion"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class list storage classes custom storage"], "anchor": "section", "description": "Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into given site.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:custom_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "custom_storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage class list storage classes default storage class"], "anchor": "schema-storage_class_list--storage_classes--default_storage_class", "description": "Make this storage class default storage class for the K8s cluster.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "default_storage_class"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage class list storage classes description spec"], "anchor": "schema-storage_class_list--storage_classes--description_spec", "description": "Storage Class Description. Description for this storage class.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes hpe storage"], "anchor": "section", "description": "Storage class Device configuration for HPE Storage.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:hpe_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "hpe_storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage class list storage classes netapp trident"], "anchor": "section", "description": "Storage class Device configuration for NetApp Trident.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:netapp_trident", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "netapp_trident"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage class list storage classes pure service orchestrator"], "anchor": "section", "description": "Storage class Device configuration for Pure Service Orchestrator.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes:pure_service_orchestrator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "pure_service_orchestrator"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage class list storage classes reclaim policy"], "anchor": "schema-storage_class_list--storage_classes--reclaim_policy", "description": "Reclaim Policy.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "reclaim_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes storage class name"], "anchor": "schema-storage_class_list--storage_classes--storage_class_name", "description": "Name of the storage class as it will appear in K8s.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "storage_class_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage class list storage classes storage device"], "anchor": "schema-storage_class_list--storage_classes--storage_device", "description": "Storage device that this class will use. The Device name defined at previous step.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_class_list:storage_classes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_class_list", "storage_classes", "storage_device"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_class_list/storage_classes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of custom storage classes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_class_list.storage_classes

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/)
- storage_class_list.storage_classes

<a id="section"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-storage_class_list--storage_classes--advanced_storage_parameters"></a>

### advanced_storage_parameters property

Type: `["map", "string"]`. Computed.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-storage_class_list--storage_classes--allow_volume_expansion"></a>

### allow_volume_expansion property

Type: `"bool"`. Computed.

Allow Volume Expansion. Allow volume expansion.

Upstream description:

Allow volume expansion.

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

- [custom_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/custom_storage/): complete subsection reference.

<a id="schema-storage_class_list--storage_classes--default_storage_class"></a>

### default_storage_class property

Type: `"bool"`. Computed.

Make this storage class default storage class for the K8s cluster.

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

<a id="schema-storage_class_list--storage_classes--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Storage Class Description. Description for this storage class.

- [hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/hpe_storage/): complete subsection reference.

- [netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/): complete subsection reference.

- [pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/pure_service_orchestrator/): complete subsection reference.

<a id="schema-storage_class_list--storage_classes--reclaim_policy"></a>

### reclaim_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Reclaim Policy.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="schema-storage_class_list--storage_classes--storage_class_name"></a>

### storage_class_name property

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="schema-storage_class_list--storage_classes--storage_device"></a>

### storage_device property

Type: `"string"`. Computed.

Storage device that this class will use. The Device name defined at previous step.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

## Next pages

- [storage_class_list.storage_classes.custom_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/custom_storage/)
- [storage_class_list.storage_classes.hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/hpe_storage/)
- [storage_class_list.storage_classes.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/netapp_trident/)
- [storage_class_list.storage_classes.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/storage_classes/pure_service_orchestrator/)
- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_class_list/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
