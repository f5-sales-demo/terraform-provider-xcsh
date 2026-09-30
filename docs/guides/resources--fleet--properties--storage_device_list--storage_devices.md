---
page_title: "storage_device_list.storage_devices"
subcategory: ""
description: "storage_device_list.storage_devices for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 5494, "body_sha256": "sha256:111d9a1561e867fb32b5bd24fd075fe918c9d30e3e503b208fd48d72acb11750", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:custom_storage", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:pure_service_orchestrator"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list", "path": "docs/guides/resources--fleet--properties--storage_device_list--storage_devices.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list", "storage_devices"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list.storage_devices for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_device_list.storage_devices

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_device_list](resources--fleet--properties--storage_device_list.md)
- storage_device_list.storage_devices

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Devices. List of custom storage devices.

Upstream description:

List of custom storage devices.

Provider validators and defaults (from schema source):

```go
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

Upstream description:

Map of parameter name and string value.

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

- [custom_storage](resources--fleet--properties--storage_device_list--storage_devices--custom_storage.md): complete subsection reference.

- [hpe_storage](resources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md): complete subsection reference.

- [netapp_trident](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md): complete subsection reference.

- [pure_service_orchestrator](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--storage_device"></a>

### storage_device property

Type: `"string"`. Optional.

Storage Device. Storage device and device unit.

Upstream description:

Storage device and device unit.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [storage_device_list.storage_devices.custom_storage](resources--fleet--properties--storage_device_list--storage_devices--custom_storage.md)
- [storage_device_list.storage_devices.hpe_storage](resources--fleet--properties--storage_device_list--storage_devices--hpe_storage.md)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--properties--storage_device_list--storage_devices--netapp_trident.md)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--properties--storage_device_list--storage_devices--pure_service_orchestrator.md)
- [storage_device_list](resources--fleet--properties--storage_device_list.md)
- [xcsh_fleet](../resources/fleet.md)
