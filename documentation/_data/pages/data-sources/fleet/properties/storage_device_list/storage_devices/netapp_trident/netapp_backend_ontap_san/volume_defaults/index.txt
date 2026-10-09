---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults"
subcategory: ""
description: "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration."
xcsh_docs: {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults"], "body_bytes": 10356, "body_sha256": "sha256:687d4234bda14572eaf0063bd348ffda4e2aaa352f770c28c6c3f0583e559e4a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults:no_qos"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "path": "documentation/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/volume_defaults/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130", "registry_path": "docs/guides/data-sources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults adaptive qos policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--adaptive_qos_policy", "description": "Exclusive with Enter Adaptive QoS Policy Name.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "adaptive_qos_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults encryption"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--encryption", "description": "Enable NetApp volume encryption.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "encryption"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults export policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--export_policy", "description": "Export policy to use.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "export_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults no qos"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults:no_qos", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "no_qos"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults qos policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--qos_policy", "description": "Exclusive with Enter QoS Policy Name.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "qos_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults security style"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--security_style", "description": "Security style for new volumes.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "security_style"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults snapshot dir"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_dir", "description": "Access to the .snapshot directory.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "snapshot_dir"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults snapshot policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_policy", "description": "Snapshot policy to use.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "snapshot_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults snapshot reserve"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_reserve", "description": "Percentage of volume reserved for snapshots. \"0\" if snapshot policy is \"none\", else \"\"", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "snapshot_reserve"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults space reserve"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--space_reserve", "description": "Space reservation mode; “none” (thin) or “volume” (thick)", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "space_reserve"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults split on clone"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--split_on_clone", "description": "Split a clone from its parent upon creation.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "split_on_clone"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults tiering policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--tiering_policy", "description": "Tiering policy to use. \"none\" is default.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "tiering_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san volume defaults unix permissions"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--unix_permissions", "description": "Unix permission mode for new volumes. All allowed 777.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults", "unix_permissions"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/volume_defaults/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["fleetCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="section"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

## Direct properties

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--adaptive_qos_policy"></a>

### adaptive_qos_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--encryption"></a>

### encryption property

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--export_policy"></a>

### export_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Export policy to use.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [no_qos](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/volume_defaults/no_qos/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--qos_policy"></a>

### qos_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--security_style"></a>

### security_style property

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_dir"></a>

### snapshot_dir property

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_policy"></a>

### snapshot_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Snapshot policy to use.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--snapshot_reserve"></a>

### snapshot_reserve property

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--space_reserve"></a>

### space_reserve property

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--split_on_clone"></a>

### split_on_clone property

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--tiering_policy"></a>

### tiering_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Tiering policy to use. "none" is default.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--unix_permissions"></a>

### unix_permissions property

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

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
