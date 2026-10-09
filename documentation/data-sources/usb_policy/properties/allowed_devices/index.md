---
page_title: "allowed_devices"
subcategory: ""
description: "List of allowed USB devices."
xcsh_docs: {"aliases": ["allowed devices"], "body_bytes": 4946, "body_sha256": "sha256:2c87256969dbbb4b08f8e33582c33e4ce3e7701865d8e9c6406647245046be06", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:usb_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "parent_id": "xcsh-docs:data-sources:usb_policy:reference", "path": "documentation/data-sources/usb_policy/properties/allowed_devices/index.md", "product": "distributed-cloud", "provider_name": "usb_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1222313110031021-0033310100131201-3003112111300200-1012320303311210-3122230300102333-2330103233013011-0311111233223203-0223210002011200", "registry_path": "docs/guides/data-sources--usb_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowed_devices"], "schema_version": 1, "sections": [{"aliases": ["allowed devices b device class"], "anchor": "schema-allowed_devices--b_device_class", "description": "The class of this device.", "document_id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "b_device_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices b device protocol"], "anchor": "schema-allowed_devices--b_device_protocol", "description": "The protocol (within the sub-class) of this device.", "document_id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "b_device_protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices b device sub class"], "anchor": "schema-allowed_devices--b_device_sub_class", "description": "The sub-class (within the class) of this device.", "document_id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "b_device_sub_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices i serial"], "anchor": "schema-allowed_devices--i_serial", "description": "Index of Serial Number String Descriptor.", "document_id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "i_serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices id product"], "anchor": "schema-allowed_devices--id_product", "description": "Product ID (Assigned by Manufacturer) in hex.", "document_id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "id_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices id vendor"], "anchor": "schema-allowed_devices--id_vendor", "description": "Vendor ID (Assigned by USB Org) in hex.", "document_id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "id_vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/usb_policy/properties/allowed_devices/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of allowed USB devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["usb_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowed_devices

Breadcrumbs:

- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/properties/)
- allowed_devices

<a id="section"></a>

Type: `"list"`. Computed.

Allowed USB devices. List of allowed USB devices.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

<a id="schema-allowed_devices--b_device_class"></a>

### b_device_class property

Type: `"string"`. Computed.

Class. The class of this device.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-allowed_devices--b_device_protocol"></a>

### b_device_protocol property

Type: `"string"`. Computed.

The protocol (within the sub-class) of this device.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-allowed_devices--b_device_sub_class"></a>

### b_device_sub_class property

Type: `"string"`. Computed.

The sub-class (within the class) of this device.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-allowed_devices--i_serial"></a>

### i_serial property

Type: `"string"`. Computed.

Index of Serial Number String Descriptor.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-allowed_devices--id_product"></a>

### id_product property

Type: `"string"`. Computed.

Product ID (Assigned by Manufacturer) in hex.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="schema-allowed_devices--id_vendor"></a>

### id_vendor property

Type: `"string"`. Computed.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
