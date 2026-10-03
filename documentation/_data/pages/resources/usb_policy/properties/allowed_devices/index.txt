---
page_title: "allowed_devices"
subcategory: ""
description: "List of allowed USB devices."
xcsh_docs: {"aliases": ["allowed devices"], "body_bytes": 5459, "body_sha256": "sha256:004983b3bd85d85d5a09a559ac542998872f2750d3acaeeeb08715063ac45284", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:usb_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:usb_policy:properties:allowed_devices", "parent_id": "xcsh-docs:resources:usb_policy:reference", "path": "documentation/resources/usb_policy/properties/allowed_devices/index.md", "product": "distributed-cloud", "provider_name": "usb_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3320231330031332-2111103103322211-1212323033212013-0110213130300113-0021012312002300-1222212221322113-0023133330222231-2322130331230220", "registry_path": "docs/guides/resources--usb_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["allowed_devices"], "schema_version": 1, "sections": [{"aliases": ["allowed devices b device class"], "anchor": "schema-allowed_devices--b_device_class", "description": "The class of this device.", "document_id": "xcsh-docs:resources:usb_policy:properties:allowed_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "b_device_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices b device protocol"], "anchor": "schema-allowed_devices--b_device_protocol", "description": "The protocol (within the sub-class) of this device.", "document_id": "xcsh-docs:resources:usb_policy:properties:allowed_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "b_device_protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices b device sub class"], "anchor": "schema-allowed_devices--b_device_sub_class", "description": "The sub-class (within the class) of this device.", "document_id": "xcsh-docs:resources:usb_policy:properties:allowed_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "b_device_sub_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices i serial"], "anchor": "schema-allowed_devices--i_serial", "description": "Index of Serial Number String Descriptor.", "document_id": "xcsh-docs:resources:usb_policy:properties:allowed_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "i_serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices id product"], "anchor": "schema-allowed_devices--id_product", "description": "Product ID (Assigned by Manufacturer) in hex.", "document_id": "xcsh-docs:resources:usb_policy:properties:allowed_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "id_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["allowed devices id vendor"], "anchor": "schema-allowed_devices--id_vendor", "description": "Vendor ID (Assigned by USB Org) in hex.", "document_id": "xcsh-docs:resources:usb_policy:properties:allowed_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["allowed_devices", "id_vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/usb_policy/properties/allowed_devices/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "List of allowed USB devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["usb_policyCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allowed_devices

Breadcrumbs:

- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/properties/)
- allowed_devices

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Allowed USB devices. List of allowed USB devices.

Upstream description:

List of allowed USB devices.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
allowed_devices {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-allowed_devices--b_device_class"></a>

### b_device_class property

Type: `"string"`. Optional.

Class. The class of this device.

Upstream description:

The class of this device.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Type: `"string"`. Optional.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

Upstream description:

Vendor ID (Assigned by USB Org) in hex.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/properties/)
- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/)
