---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage"
subcategory: ""
description: "List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection."
xcsh_docs: {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas storage"], "body_bytes": 4901, "body_sha256": "sha256:64447d1bee909a3f7e54fea3e9a49a4516890bc62bd325da39b44111033420d6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage"], "schema_version": 1, "sections": [{"aliases": ["labels"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--labels", "description": "List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match selection.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["volume defaults"], "anchor": "section", "description": "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults:no_qos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage:volume_defaults:no_qos", "type": "conflicts"}], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage", "volume_defaults"], "syntax": "block", "type": "object"}, {"aliases": ["zone"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--zone", "description": "Virtual Storage Pool zone definition.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage", "zone"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

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
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/volume_defaults/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage--zone"></a>

### zone property

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
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

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/volume_defaults/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
