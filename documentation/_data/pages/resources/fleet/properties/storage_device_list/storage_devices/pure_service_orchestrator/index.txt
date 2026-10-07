---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator"
subcategory: ""
description: "Device configuration for Pure Storage Service Orchestrator."
xcsh_docs: {"aliases": ["storage device list storage devices pure service orchestrator"], "body_bytes": 4026, "body_sha256": "sha256:1006b900a55eb8670d7d371f4925a7fcb0816c65bf650aad41a97093a9dfadba", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--cluster_id", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.pure_service_orchestrator:RequiredObjectAttributes:cluster_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices pure service orchestrator arrays"], "anchor": "section", "description": "Device configuration for PSO Arrays.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices pure service orchestrator cluster id"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--cluster_id", "description": "ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and underscores.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "cluster_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices pure service orchestrator enable storage topology"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_storage_topology", "description": "This option is to enable/disable the csi topology feature for pso-csi.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "enable_storage_topology"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage device list storage devices pure service orchestrator enable strict topology"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_strict_topology", "description": "This option is to enable/disable the strict csi topology feature for pso-csi.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "enable_strict_topology"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Device configuration for Pure Storage Service Orchestrator.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["fleetCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- storage_device_list.storage_devices.pure_service_orchestrator

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for Pure Storage Service Orchestrator.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_id")}
```

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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

## Direct properties

- [arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--cluster_id"></a>

### cluster_id property

Type: `"string"`. Optional.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 22),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 22,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9_]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  }
}
```

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_storage_topology"></a>

### enable_storage_topology property

Type: `"bool"`. Optional.

This option is to enable/disable the csi topology feature for pso-csi.

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

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_strict_topology"></a>

### enable_strict_topology property

Type: `"bool"`. Optional.

This option is to enable/disable the strict csi topology feature for pso-csi.

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
