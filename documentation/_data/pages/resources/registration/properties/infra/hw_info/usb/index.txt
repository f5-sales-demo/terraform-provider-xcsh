---
page_title: "infra.hw_info.usb"
subcategory: ""
description: "List of USB devices in server."
xcsh_docs: {"aliases": ["infra hw info usb"], "body_bytes": 12256, "body_sha256": "sha256:378f6fb2e0d55107c6ad1e94343138a89151e1068c0b298258ac350b1691f8eb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "path": "documentation/resources/registration/properties/infra/hw_info/usb/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2133131333302012-2311223101000032-3011003103231233-1102312333003101-0202323211123222-0220232101330330-2012330211102002-3032003002021300", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "usb"], "schema_version": 1, "sections": [{"aliases": ["address"], "anchor": "schema-infra--hw_info--usb--address", "description": "Address of the device on the bus in decimal.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "address"], "syntax": "attribute", "type": "number"}, {"aliases": ["b device class"], "anchor": "schema-infra--hw_info--usb--b_device_class", "description": "The class of this device.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_device_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["b device protocol"], "anchor": "schema-infra--hw_info--usb--b_device_protocol", "description": "The protocol (within the sub-class) of this device.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_device_protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["b device sub class"], "anchor": "schema-infra--hw_info--usb--b_device_sub_class", "description": "The sub-class (within the class) of this device.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_device_sub_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["b max packet size"], "anchor": "schema-infra--hw_info--usb--b_max_packet_size", "description": "Maximum size of the control transfer.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "b_max_packet_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["bcd device"], "anchor": "schema-infra--hw_info--usb--bcd_device", "description": "The device version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "bcd_device"], "syntax": "attribute", "type": "string"}, {"aliases": ["bcd usb"], "anchor": "schema-infra--hw_info--usb--bcd_usb", "description": "USB Specification Release Number.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "bcd_usb"], "syntax": "attribute", "type": "string"}, {"aliases": ["bus"], "anchor": "schema-infra--hw_info--usb--bus", "description": "The bus on which the device was detected in decimal.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "bus"], "syntax": "attribute", "type": "number"}, {"aliases": ["description spec"], "anchor": "schema-infra--hw_info--usb--description_spec", "description": "Description. Device description.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["i manufacturer"], "anchor": "schema-infra--hw_info--usb--i_manufacturer", "description": "Manufacturer name.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "i_manufacturer"], "syntax": "attribute", "type": "string"}, {"aliases": ["i product"], "anchor": "schema-infra--hw_info--usb--i_product", "description": "Product name reported by device.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "i_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["i serial"], "anchor": "schema-infra--hw_info--usb--i_serial", "description": "Index of Serial Number String Descriptor.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "i_serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["id product"], "anchor": "schema-infra--hw_info--usb--id_product", "description": "Product ID (Assigned by Manufacturer) in hex.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "id_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["id vendor"], "anchor": "schema-infra--hw_info--usb--id_vendor", "description": "Vendor ID (Assigned by USB Org) in hex.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "id_vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-infra--hw_info--usb--port", "description": "Port on which the device was detected in decimal.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["product name"], "anchor": "schema-infra--hw_info--usb--product_name", "description": "Product ID translated to name (if available)", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "product_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["speed"], "anchor": "schema-infra--hw_info--usb--speed", "description": "The negotiated operating speed for the device.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "speed"], "syntax": "attribute", "type": "string"}, {"aliases": ["usb type"], "anchor": "schema-infra--hw_info--usb--usb_type", "description": "Type of USB device Unknown USB device type Internal USB present in Certified HW USB device present during node registration USB device that can be matched by USB rules.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "usb_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["vendor name"], "anchor": "schema-infra--hw_info--usb--vendor_name", "description": "Vendor ID translated to name (if available)", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:usb", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "usb", "vendor_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/usb/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of USB devices in server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.usb

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- infra.hw_info.usb

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
usb {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hw_info--usb--address"></a>

### address property

Type: `"number"`. Optional.

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

Type: `"number"`. Optional.

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

Type: `"string"`. Optional.

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

Type: `"string"`. Optional.

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

Type: `"number"`. Optional.

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

Type: `"string"`. Optional.

Description. Device description.

<a id="schema-infra--hw_info--usb--i_manufacturer"></a>

### i_manufacturer property

Type: `"string"`. Optional.

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

Type: `"string"`. Optional.

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

Type: `"number"`. Optional.

Port on which the device was detected in decimal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

Type: `"string"`. Optional.

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

Type: `"string"`. Optional.

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

Type: `"string"`. Optional.

\[Enum: UNKNOWN\_USB|INTERNAL|REGISTERED|CONFIGURABLE\] Type of USB device Unknown USB device type
Internal USB present in Certified HW USB device present during node registration USB device that can
be matched by USB rules. Possible values are \`UNKNOWN\_USB\`, \`INTERNAL\`, \`REGISTERED\`,
\`CONFIGURABLE\`. Defaults to \`UNKNOWN\_USB\`.

Upstream description:

Type of USB device

Unknown USB device type Internal USB present in Certified HW USB device present during node
registration USB device that can be matched by USB rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN_USB",
    "INTERNAL",
    "REGISTERED",
    "CONFIGURABLE"),
}
```

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

Type: `"string"`. Optional.

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

- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
