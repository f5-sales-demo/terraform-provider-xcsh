---
page_title: "infra.hw_info.usb"
subcategory: ""
description: "List of USB devices in server."
xcsh_docs: {"aliases": ["infra hw info usb"], "body_bytes": 11831, "body_sha256": "sha256:0ffde9b0d488b410a71eaa304ad6c96794bef28ce3eae4181e969b6a102c6441", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "documentation/data-sources/registration/properties/infra/hw_info/usb/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1110023331232012-0110323002300313-0131123232121130-0310333130302000-0321212033001212-2330232222300223-2313213031323201-0101213101122321", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "usb"], "schema_version": 1, "sections": [{"aliases": ["infra hw info usb address"], "anchor": "schema-infra--hw_info--usb--address", "description": "Address of the device on the bus in decimal.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "address"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info usb b device class"], "anchor": "schema-infra--hw_info--usb--b_device_class", "description": "The class of this device.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_device_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb b device protocol"], "anchor": "schema-infra--hw_info--usb--b_device_protocol", "description": "The protocol (within the sub-class) of this device.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_device_protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb b device sub class"], "anchor": "schema-infra--hw_info--usb--b_device_sub_class", "description": "The sub-class (within the class) of this device.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_device_sub_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb b max packet size"], "anchor": "schema-infra--hw_info--usb--b_max_packet_size", "description": "Maximum size of the control transfer.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_max_packet_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info usb bcd device"], "anchor": "schema-infra--hw_info--usb--bcd_device", "description": "The device version.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "bcd_device"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb bcd usb"], "anchor": "schema-infra--hw_info--usb--bcd_usb", "description": "USB Specification Release Number.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "bcd_usb"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb bus"], "anchor": "schema-infra--hw_info--usb--bus", "description": "The bus on which the device was detected in decimal.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "bus"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info usb description spec"], "anchor": "schema-infra--hw_info--usb--description_spec", "description": "Description. Device description.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb i manufacturer"], "anchor": "schema-infra--hw_info--usb--i_manufacturer", "description": "Manufacturer name.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "i_manufacturer"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb i product"], "anchor": "schema-infra--hw_info--usb--i_product", "description": "Product name reported by device.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "i_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb i serial"], "anchor": "schema-infra--hw_info--usb--i_serial", "description": "Index of Serial Number String Descriptor.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "i_serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb id product"], "anchor": "schema-infra--hw_info--usb--id_product", "description": "Product ID (Assigned by Manufacturer) in hex.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "id_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb id vendor"], "anchor": "schema-infra--hw_info--usb--id_vendor", "description": "Vendor ID (Assigned by USB Org) in hex.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "id_vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb port"], "anchor": "schema-infra--hw_info--usb--port", "description": "Port on which the device was detected in decimal.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["infra hw info usb product name"], "anchor": "schema-infra--hw_info--usb--product_name", "description": "Product ID translated to name (if available)", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "product_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb speed"], "anchor": "schema-infra--hw_info--usb--speed", "description": "The negotiated operating speed for the device.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "speed"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb usb type"], "anchor": "schema-infra--hw_info--usb--usb_type", "description": "Type of USB device Unknown USB device type Internal USB present in Certified HW USB device present during node registration USB device that can be matched by USB rules.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "usb_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info usb vendor name"], "anchor": "schema-infra--hw_info--usb--vendor_name", "description": "Vendor ID translated to name (if available)", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:usb", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "vendor_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/usb/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of USB devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.usb

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- infra.hw_info.usb

<a id="section"></a>

Type: `"list"`. Computed.

USB devices. List of USB devices in server.

Upstream description:

List of USB devices in server.

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

<a id="schema-infra--hw_info--usb--address"></a>

### address property

Type: `"number"`. Computed.

Address of the device on the bus in decimal.

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

<a id="schema-infra--hw_info--usb--b_device_class"></a>

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

<a id="schema-infra--hw_info--usb--b_device_protocol"></a>

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

<a id="schema-infra--hw_info--usb--b_device_sub_class"></a>

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

<a id="schema-infra--hw_info--usb--b_max_packet_size"></a>

### b_max_packet_size property

Type: `"number"`. Computed.

Max packet size. Maximum size of the control transfer.

Upstream description:

Maximum size of the control transfer.

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

<a id="schema-infra--hw_info--usb--bcd_device"></a>

### bcd_device property

Type: `"string"`. Computed.

BCD Device. The device version.

Upstream description:

The device version.

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

<a id="schema-infra--hw_info--usb--bcd_usb"></a>

### bcd_usb property

Type: `"string"`. Computed.

BCD Spec. USB Specification Release Number.

Upstream description:

USB Specification Release Number.

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

<a id="schema-infra--hw_info--usb--bus"></a>

### bus property

Type: `"number"`. Computed.

The bus on which the device was detected in decimal.

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

<a id="schema-infra--hw_info--usb--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Device description.

<a id="schema-infra--hw_info--usb--i_manufacturer"></a>

### i_manufacturer property

Type: `"string"`. Computed.

Manufacturer. Manufacturer name.

Upstream description:

Manufacturer name.

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

<a id="schema-infra--hw_info--usb--i_product"></a>

### i_product property

Type: `"string"`. Computed.

Device product. Product name reported by device.

Upstream description:

Product name reported by device.

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

<a id="schema-infra--hw_info--usb--i_serial"></a>

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

<a id="schema-infra--hw_info--usb--id_product"></a>

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

<a id="schema-infra--hw_info--usb--id_vendor"></a>

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

<a id="schema-infra--hw_info--usb--port"></a>

### port property

Type: `"number"`. Computed.

Port on which the device was detected in decimal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--usb--product_name"></a>

### product_name property

Type: `"string"`. Computed.

Product ID translated to name (if available).

Upstream description:

Product ID translated to name (if available)

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

<a id="schema-infra--hw_info--usb--speed"></a>

### speed property

Type: `"string"`. Computed.

The negotiated operating speed for the device.

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

<a id="schema-infra--hw_info--usb--usb_type"></a>

### usb_type property

Type: `"string"`. Computed.

\[Enum: UNKNOWN\_USB|INTERNAL|REGISTERED|CONFIGURABLE\] Type of USB device Unknown USB device type
Internal USB present in Certified HW USB device present during node registration USB device that can
be matched by USB rules. Possible values are \`UNKNOWN\_USB\`, \`INTERNAL\`, \`REGISTERED\`,
\`CONFIGURABLE\`. Defaults to \`UNKNOWN\_USB\`.

Upstream description:

Type of USB device

Unknown USB device type Internal USB present in Certified HW USB device present during node
registration USB device that can be matched by USB rules.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNKNOWN_USB",
  "enum": [
    "UNKNOWN_USB",
    "INTERNAL",
    "REGISTERED",
    "CONFIGURABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--usb--vendor_name"></a>

### vendor_name property

Type: `"string"`. Computed.

Vendor ID translated to name (if available).

Upstream description:

Vendor ID translated to name (if available)

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

- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
