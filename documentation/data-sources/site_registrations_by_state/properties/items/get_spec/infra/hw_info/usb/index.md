---
page_title: "items.get_spec.infra.hw_info.usb"
subcategory: ""
description: "items.get_spec.infra.hw_info.usb for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 5361, "body_sha256": "sha256:61ff8be8c9617189d0c26634996a33703b82f13218c1f22e94411afbbe232ed3", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info:usb", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/usb/index.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "usb"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/usb/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.infra.hw_info.usb for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.usb

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.usb

<a id="section"></a>

Type: `"list"`. Computed.

USB devices. List of USB devices in server.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--usb--address"></a>

### address property

Type: `"number"`. Computed.

Address of the device on the bus in decimal.

<a id="schema-items--get_spec--infra--hw_info--usb--b_device_class"></a>

### b_device_class property

Type: `"string"`. Computed.

Class. The class of this device.

<a id="schema-items--get_spec--infra--hw_info--usb--b_device_protocol"></a>

### b_device_protocol property

Type: `"string"`. Computed.

The protocol (within the sub-class) of this device.

<a id="schema-items--get_spec--infra--hw_info--usb--b_device_sub_class"></a>

### b_device_sub_class property

Type: `"string"`. Computed.

The sub-class (within the class) of this device.

<a id="schema-items--get_spec--infra--hw_info--usb--b_max_packet_size"></a>

### b_max_packet_size property

Type: `"number"`. Computed.

Max packet size. Maximum size of the control transfer.

<a id="schema-items--get_spec--infra--hw_info--usb--bcd_device"></a>

### bcd_device property

Type: `"string"`. Computed.

BCD Device. The device version.

<a id="schema-items--get_spec--infra--hw_info--usb--bcd_usb"></a>

### bcd_usb property

Type: `"string"`. Computed.

BCD Spec. USB Specification Release Number.

<a id="schema-items--get_spec--infra--hw_info--usb--bus"></a>

### bus property

Type: `"number"`. Computed.

The bus on which the device was detected in decimal.

<a id="schema-items--get_spec--infra--hw_info--usb--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Device description.

<a id="schema-items--get_spec--infra--hw_info--usb--i_manufacturer"></a>

### i_manufacturer property

Type: `"string"`. Computed.

Manufacturer. Manufacturer name.

<a id="schema-items--get_spec--infra--hw_info--usb--i_product"></a>

### i_product property

Type: `"string"`. Computed.

Device product. Product name reported by device.

<a id="schema-items--get_spec--infra--hw_info--usb--i_serial"></a>

### i_serial property

Type: `"string"`. Computed.

Index of Serial Number String Descriptor.

<a id="schema-items--get_spec--infra--hw_info--usb--id_product"></a>

### id_product property

Type: `"string"`. Computed.

Product ID (Assigned by Manufacturer) in hex.

<a id="schema-items--get_spec--infra--hw_info--usb--id_vendor"></a>

### id_vendor property

Type: `"string"`. Computed.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

<a id="schema-items--get_spec--infra--hw_info--usb--port"></a>

### port property

Type: `"number"`. Computed.

Port on which the device was detected in decimal.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

<a id="schema-items--get_spec--infra--hw_info--usb--product_name"></a>

### product_name property

Type: `"string"`. Computed.

Product ID translated to name (if available).

<a id="schema-items--get_spec--infra--hw_info--usb--speed"></a>

### speed property

Type: `"string"`. Computed.

The negotiated operating speed for the device.

<a id="schema-items--get_spec--infra--hw_info--usb--usb_type"></a>

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

<a id="schema-items--get_spec--infra--hw_info--usb--vendor_name"></a>

### vendor_name property

Type: `"string"`. Computed.

Vendor ID translated to name (if available).

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
