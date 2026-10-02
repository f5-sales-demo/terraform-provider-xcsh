---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator"
subcategory: ""
description: "Device configuration for Pure Storage Service Orchestrator."
xcsh_docs: {"aliases": ["storage device list storage devices pure service orchestrator"], "body_bytes": 4524, "body_sha256": "sha256:da79c67e8d8c264a9dd069d68bb0120e1337db8508e552128a837a098187a640", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices", "path": "documentation/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022", "registry_path": "docs/guides/data-sources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator"], "schema_version": 1, "sections": [{"aliases": ["arrays"], "anchor": "section", "description": "Device configuration for PSO Arrays.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays"], "syntax": "attribute", "type": "object"}, {"aliases": ["cluster id"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--cluster_id", "description": "ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and underscores.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "cluster_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable storage topology"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_storage_topology", "description": "This option is to enable/disable the csi topology feature for pso-csi.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "enable_storage_topology"], "syntax": "attribute", "type": "bool"}, {"aliases": ["enable strict topology"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--enable_strict_topology", "description": "This option is to enable/disable the strict csi topology feature for pso-csi.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "enable_strict_topology"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Device configuration for Pure Storage Service Orchestrator.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- storage_device_list.storage_devices.pure_service_orchestrator

<a id="section"></a>

Type: `"single"`. Computed.

Device configuration for Pure Storage Service Orchestrator.

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

- [arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--cluster_id"></a>

### cluster_id property

Type: `"string"`. Computed.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays.

Upstream description:

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"bool"`. Computed.

Option is to enable/disable the csi topology feature for pso-csi.

Upstream description:

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

Type: `"bool"`. Computed.

Option is to enable/disable the strict csi topology feature for pso-csi.

Upstream description:

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

## Next pages

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
