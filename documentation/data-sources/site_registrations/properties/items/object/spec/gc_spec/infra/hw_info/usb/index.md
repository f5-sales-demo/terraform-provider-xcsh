---
page_title: "items.object.spec.gc_spec.infra.hw_info.usb"
subcategory: ""
description: "items.object.spec.gc_spec.infra.hw_info.usb for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 5801, "body_sha256": "sha256:1019e1f80a02dea73b34785c41c452aa38123ff35ddd905af7d602a6a29a1c9f", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info:usb", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/usb/index.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "hw_info", "usb"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/usb/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.hw_info.usb for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items.object.spec.gc_spec.infra.hw_info.usb

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- items.object.spec.gc_spec.infra.hw_info.usb

<a id="section"></a>

Type: `"list"`. Computed.

USB devices. List of USB devices in server.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--address"></a>

### address property

Type: `"number"`. Computed.

Address of the device on the bus in decimal.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--b_device_class"></a>

### b_device_class property

Type: `"string"`. Computed.

Class. The class of this device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--b_device_protocol"></a>

### b_device_protocol property

Type: `"string"`. Computed.

The protocol (within the sub-class) of this device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--b_device_sub_class"></a>

### b_device_sub_class property

Type: `"string"`. Computed.

The sub-class (within the class) of this device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--b_max_packet_size"></a>

### b_max_packet_size property

Type: `"number"`. Computed.

Max packet size. Maximum size of the control transfer.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--bcd_device"></a>

### bcd_device property

Type: `"string"`. Computed.

BCD Device. The device version.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--bcd_usb"></a>

### bcd_usb property

Type: `"string"`. Computed.

BCD Spec. USB Specification Release Number.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--bus"></a>

### bus property

Type: `"number"`. Computed.

The bus on which the device was detected in decimal.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Device description.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--i_manufacturer"></a>

### i_manufacturer property

Type: `"string"`. Computed.

Manufacturer. Manufacturer name.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--i_product"></a>

### i_product property

Type: `"string"`. Computed.

Device product. Product name reported by device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--i_serial"></a>

### i_serial property

Type: `"string"`. Computed.

Index of Serial Number String Descriptor.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--id_product"></a>

### id_product property

Type: `"string"`. Computed.

Product ID (Assigned by Manufacturer) in hex.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--id_vendor"></a>

### id_vendor property

Type: `"string"`. Computed.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--port"></a>

### port property

Type: `"number"`. Computed.

Port on which the device was detected in decimal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--product_name"></a>

### product_name property

Type: `"string"`. Computed.

Product ID translated to name (if available).

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--speed"></a>

### speed property

Type: `"string"`. Computed.

The negotiated operating speed for the device.

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--usb_type"></a>

### usb_type property

Type: `"string"`. Computed.

\[Enum: UNKNOWN\_USB|INTERNAL|REGISTERED|CONFIGURABLE\] Type of USB device Unknown USB device type
Internal USB present in Certified HW USB device present during node registration USB device that can
be matched by USB rules. Possible values are \`UNKNOWN\_USB\`, \`INTERNAL\`, \`REGISTERED\`,
\`CONFIGURABLE\`. Defaults to \`UNKNOWN\_USB\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN_USB",
    "INTERNAL",
    "REGISTERED",
    "CONFIGURABLE"),
}
```

<a id="schema-items--object--spec--gc_spec--infra--hw_info--usb--vendor_name"></a>

### vendor_name property

Type: `"string"`. Computed.

Vendor ID translated to name (if available).

## Next pages

- [items.object.spec.gc_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
