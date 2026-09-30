---
page_title: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3895, "body_sha256": "sha256:f30e7c99c0011e7d06c0336e910ab6785d20b3e102f60237ec7a08c78575f1b3", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator.md)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays.md)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

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

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--enable_snapshot_directory"></a>

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

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--export_rules"></a>

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

- [flash_blades](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades.md)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
