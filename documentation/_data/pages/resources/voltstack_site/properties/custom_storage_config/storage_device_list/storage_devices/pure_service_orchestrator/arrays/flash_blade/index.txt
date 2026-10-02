---
page_title: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade"
subcategory: ""
description: "Specify what storage flash blades should be managed the plugin."
xcsh_docs: {"aliases": ["custom storage config storage device list storage devices pure service orchestrator arrays flash blade"], "body_bytes": 4911, "body_sha256": "sha256:26fe21bad8dd2230746717e526ad26395e4cd90c113c05ca71ce674a3f2afa7c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113", "registry_path": "docs/guides/resources--voltstack_site--reference--group-007.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade:RequiredObjectAttributes:flash_blades", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade"], "schema_version": 1, "sections": [{"aliases": ["enable snapshot directory"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--enable_snapshot_directory", "description": "Enable/Disable FlashBlade snapshots.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "enable_snapshot_directory"], "syntax": "attribute", "type": "bool"}, {"aliases": ["export rules"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--export_rules", "description": "NFS Export rules.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "export_rules"], "syntax": "attribute", "type": "string"}, {"aliases": ["flash blades"], "anchor": "section", "description": "For FlashBlades you must set the \"mgmt_endpoint\", \"api_token\" and nfs_endpoint.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--mgmt_dns_name", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades:ConflictingListObjectAttributes:mgmt_dns_name,mgmt_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--mgmt_ip", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades:ConflictingListObjectAttributes:mgmt_dns_name,mgmt_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--nfs_endpoint_dns_name", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades:ConflictingListObjectAttributes:nfs_endpoint_dns_name,nfs_endpoint_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--nfs_endpoint_ip", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades:ConflictingListObjectAttributes:nfs_endpoint_dns_name,nfs_endpoint_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "flash_blades"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specify what storage flash blades should be managed the plugin.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/)
- [custom_storage_config.storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash blades should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("flash_blades")}
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
flash_blade {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--enable_snapshot_directory"></a>

### enable_snapshot_directory property

Type: `"bool"`. Optional.

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

Type: `"string"`. Optional.

NFS Export Rules. NFS Export rules.

Upstream description:

NFS Export rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 250),
}
```

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

- [flash_blades](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
