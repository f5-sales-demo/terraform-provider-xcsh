---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults"
subcategory: ""
description: "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration."
xcsh_docs: {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults"], "body_bytes": 11946, "body_sha256": "sha256:57b2ea1def377665921cafea67e53c839737359f08aa298a5d1f0c269e55771c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults:no_qos"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/storage/volume_defaults/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1331203201033220-3033100023220013-2010010023203222-2210330212120313-2310202223112112-2033230330001120-3220221121000221-1210120221222223", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults:no_qos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults:no_qos", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults adaptive qos policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--adaptive_qos_policy", "description": "Exclusive with Enter Adaptive QoS Policy Name.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "adaptive_qos_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults encryption"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--encryption", "description": "Enable NetApp volume encryption.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "encryption"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults export policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--export_policy", "description": "Export policy to use.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "export_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults no qos"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults:no_qos", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "no_qos"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults qos policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--qos_policy", "description": "Exclusive with Enter QoS Policy Name.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "qos_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults security style"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--security_style", "description": "Security style for new volumes.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "security_style"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults snapshot dir"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_dir", "description": "Access to the .snapshot directory.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "snapshot_dir"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults snapshot policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_policy", "description": "Snapshot policy to use.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "snapshot_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults snapshot reserve"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_reserve", "description": "Percentage of volume reserved for snapshots. \"0\" if snapshot policy is \"none\", else \"\"", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "snapshot_reserve"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults space reserve"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--space_reserve", "description": "Space reservation mode; “none” (thin) or “volume” (thick)", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["none", "thick"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "space_reserve"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults split on clone"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--split_on_clone", "description": "Split a clone from its parent upon creation.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "split_on_clone"], "syntax": "attribute", "type": "bool"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults tiering policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--tiering_policy", "description": "Tiering policy to use. \"none\" is default.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "tiering_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap san storage volume defaults unix permissions"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--unix_permissions", "description": "Unix permission mode for new volumes. All allowed 777.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage", "volume_defaults", "unix_permissions"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/storage/volume_defaults/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fleetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/storage/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
```

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

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--adaptive_qos_policy"></a>

### adaptive_qos_policy property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--encryption"></a>

### encryption property

Type: `"bool"`. Optional.

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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--export_policy"></a>

### export_policy property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [no_qos](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/storage/volume_defaults/no_qos/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--qos_policy"></a>

### qos_policy property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--security_style"></a>

### security_style property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_dir"></a>

### snapshot_dir property

Type: `"bool"`. Optional.

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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_policy"></a>

### snapshot_policy property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--snapshot_reserve"></a>

### snapshot_reserve property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--space_reserve"></a>

### space_reserve property

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["none","thick"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--split_on_clone"></a>

### split_on_clone property

Type: `"bool"`. Optional.

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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--tiering_policy"></a>

### tiering_policy property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults--unix_permissions"></a>

### unix_permissions property

Type: `"number"`. Optional.

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
