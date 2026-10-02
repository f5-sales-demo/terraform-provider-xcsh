---
page_title: "storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade"
subcategory: ""
description: "Specify what storage flash blades should be managed the plugin."
xcsh_docs: {"aliases": ["storage device list storage devices pure service orchestrator arrays flash blade"], "body_bytes": 3902, "body_sha256": "sha256:be80e2b22f8d57f3c2eb891b4d9710ffd1fec27d0aecc167e150c97105e74e8b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "path": "documentation/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303", "registry_path": "docs/guides/data-sources--fleet--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade"], "schema_version": 1, "sections": [{"aliases": ["enable snapshot directory"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--enable_snapshot_directory", "description": "Enable/Disable FlashBlade snapshots.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "enable_snapshot_directory"], "syntax": "attribute", "type": "bool"}, {"aliases": ["export rules"], "anchor": "schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--export_rules", "description": "NFS Export rules.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "export_rules"], "syntax": "attribute", "type": "string"}, {"aliases": ["flash blades"], "anchor": "section", "description": "For FlashBlades you must set the \"mgmt_endpoint\", \"api_token\" and nfs_endpoint.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "flash_blades"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specify what storage flash blades should be managed the plugin.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="section"></a>

Type: `"single"`. Computed.

Specify what storage flash blades should be managed the plugin.

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

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--enable_snapshot_directory"></a>

### enable_snapshot_directory property

Type: `"bool"`. Computed.

Enable Snapshot Directory. Enable/Disable FlashBlade snapshots.

Upstream description:

Enable/Disable FlashBlade snapshots.

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

<a id="schema-storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--export_rules"></a>

### export_rules property

Type: `"string"`. Computed.

NFS Export Rules. NFS Export rules.

Upstream description:

NFS Export rules.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 250,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 250,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [flash_blades](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
