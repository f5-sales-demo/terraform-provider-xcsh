---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap"
subcategory: ""
description: "Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP."
xcsh_docs: {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap"], "body_bytes": 4056, "body_sha256": "sha256:fa76208ea95d755cdd6a33f0972c9cb5f9e2a46ba5b68b9f7079bf736dc77ca1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "path": "documentation/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223", "registry_path": "docs/guides/data-sources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap initiator secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_initiator_secret"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap target initiator secret"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_target_initiator_secret"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap target username"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_username", "description": "Target username. Required if useCHAP=true.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_target_username"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san use chap chap username"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_username", "description": "Inbound username. Required if useCHAP=true.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["fleetCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="section"></a>

Type: `"single"`. Computed.

Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.

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

- [chap_initiator_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_initiator_secret/): complete subsection reference.

- [chap_target_initiator_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_target_initiator_secret/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_username"></a>

### chap_target_username property

Type: `"string"`. Computed.

Target username. Required if useCHAP=true.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_username"></a>

### chap_username property

Type: `"string"`. Computed.

Inbound username. Required if useCHAP=true.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```
