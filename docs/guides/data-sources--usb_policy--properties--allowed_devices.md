---
page_title: "allowed_devices"
subcategory: ""
description: "allowed_devices for xcsh_usb_policy."
xcsh_docs: {"aliases": [], "body_bytes": 5040, "body_sha256": "sha256:f84b985585589773d84965953686a98551a33ee5222a9e132225e6a9a3f7a292", "canonical_id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "child_ids": [], "collection_id": "xcsh-docs:data-sources:usb_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:usb_policy:properties:allowed_devices", "parent_id": "xcsh-docs:data-sources:usb_policy:reference", "path": "docs/guides/data-sources--usb_policy--properties--allowed_devices.md", "provider_name": "usb_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allowed_devices"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/usb_policy/properties/allowed_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allowed_devices for xcsh_usb_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["usb_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allowed_devices

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md)
- [Property reference](data-sources--usb_policy--reference.md)
- allowed_devices

<a id="section"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-allowed_devices--id_vendor"></a>

### id_vendor property

Type: `"string"`. Computed.

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

- [Property reference](data-sources--usb_policy--reference.md)
- [xcsh_usb_policy](../data-sources/usb_policy.md)
