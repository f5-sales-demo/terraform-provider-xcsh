---
page_title: "internal_usb_device_rule"
subcategory: ""
description: "List of internal USB device rules for server."
xcsh_docs: {"aliases": ["internal usb device rule"], "body_bytes": 1769, "body_sha256": "sha256:a091993a3916c0ebc6c4d61077606be53b69ef2d2428eca711c5647b3ace9f45", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/internal_usb_device_rule/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3001313320012012-1233202130023210-1212001121232030-1232021103332333-0212331312033303-1020321011122021-1210110331030331-3132113111230010", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["internal_usb_device_rule"], "schema_version": 1, "sections": [{"aliases": ["b device class"], "anchor": "schema-internal_usb_device_rule--b_device_class", "description": "Class. The class of this device.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "b_device_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["b device protocol"], "anchor": "schema-internal_usb_device_rule--b_device_protocol", "description": "The protocol (within the sub-class) of this device.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "b_device_protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["b device sub class"], "anchor": "schema-internal_usb_device_rule--b_device_sub_class", "description": "The sub-class (within the class) of this device.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "b_device_sub_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["i serial"], "anchor": "schema-internal_usb_device_rule--i_serial", "description": "Index of Serial Number String Descriptor.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "i_serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["id product"], "anchor": "schema-internal_usb_device_rule--id_product", "description": "Product ID (Assigned by Manufacturer) in hex.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "id_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["id vendor"], "anchor": "schema-internal_usb_device_rule--id_vendor", "description": "Vendor ID. Vendor ID (Assigned by USB Org) in hex.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "id_vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/internal_usb_device_rule/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of internal USB device rules for server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# internal_usb_device_rule

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- internal_usb_device_rule

<a id="section"></a>

Type: `"list"`. Computed.

List of internal USB device rules for server.

## Direct properties

<a id="schema-internal_usb_device_rule--b_device_class"></a>

### b_device_class property

Type: `"string"`. Computed.

Class. The class of this device.

<a id="schema-internal_usb_device_rule--b_device_protocol"></a>

### b_device_protocol property

Type: `"string"`. Computed.

The protocol (within the sub-class) of this device.

<a id="schema-internal_usb_device_rule--b_device_sub_class"></a>

### b_device_sub_class property

Type: `"string"`. Computed.

The sub-class (within the class) of this device.

<a id="schema-internal_usb_device_rule--i_serial"></a>

### i_serial property

Type: `"string"`. Computed.

Index of Serial Number String Descriptor.

<a id="schema-internal_usb_device_rule--id_product"></a>

### id_product property

Type: `"string"`. Computed.

Product ID (Assigned by Manufacturer) in hex.

<a id="schema-internal_usb_device_rule--id_vendor"></a>

### id_vendor property

Type: `"string"`. Computed.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
