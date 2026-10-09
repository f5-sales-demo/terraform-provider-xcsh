---
page_title: "storage_device_list.storage_devices"
subcategory: ""
description: "List of custom storage devices."
xcsh_docs: {"aliases": ["storage device list storage devices"], "body_bytes": 6483, "body_sha256": "sha256:4621b4b0557ba9398899773c6bce675ab0296c61968d31a7fb4c0670a4d4ee00", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,hpe_storage", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,netapp_trident", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:custom_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:hpe_storage,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:ConflictingListObjectAttributes:netapp_trident,pure_service_orchestrator", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--storage_device", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices:RequiredListObjectAttributes:storage_device", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices advanced advanced parameters"], "anchor": "schema-storage_device_list--storage_devices--advanced_advanced_parameters", "description": "Map of parameter name and string value.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "advanced_advanced_parameters"], "syntax": "attribute", "type": "map"}, {"aliases": ["storage device list storage devices custom storage"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "custom_storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices hpe storage"], "anchor": "section", "description": "Device configuration for HPE Storage.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--hpe_storage--api_server_port", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage:RequiredObjectAttributes:api_server_port,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "requires"}, {"anchor": "schema-storage_device_list--storage_devices--hpe_storage--username", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage:RequiredObjectAttributes:api_server_port,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident"], "anchor": "section", "description": "Device configuration for NetApp Trident Storage.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident:ConflictingObjectAttributes:netapp_backend_ontap_nas,netapp_backend_ontap_san", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident:ConflictingObjectAttributes:netapp_backend_ontap_nas,netapp_backend_ontap_san", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "type": "conflicts"}], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices pure service orchestrator"], "anchor": "section", "description": "Device configuration for Pure Storage Service Orchestrator.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--cluster_id", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator:RequiredObjectAttributes:cluster_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "requires"}], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices storage device"], "anchor": "schema-storage_device_list--storage_devices--storage_device", "description": "Storage device and device unit.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "storage_device"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of custom storage devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["fleetCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- storage_device_list.storage_devices

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Devices. List of custom storage devices.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
storage_devices {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_device_list--storage_devices--advanced_advanced_parameters"></a>

### advanced_advanced_parameters property

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
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

- [custom_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/custom_storage/): complete subsection reference.

- [hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/): complete subsection reference.

- [netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/): complete subsection reference.

- [pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--storage_device"></a>

### storage_device property

Type: `"string"`. Optional.

Storage Device. Storage device and device unit.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```
