---
page_title: "internal_usb_device_rule"
subcategory: ""
description: "List of internal USB device rules for server."
xcsh_docs: {"aliases": ["internal usb device rule"], "body_bytes": 1511, "body_sha256": "sha256:b0c36f7b3f4f7366b329cce15526e904a2dc61c999414be1605f51a4d72ee8bb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/internal_usb_device_rule/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3001313320012012-1233202130023210-1212001121232030-1232021103332333-0212331312033303-1020321011122021-1210110331030331-3132113111230010", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["internal_usb_device_rule"], "schema_version": 1, "sections": [{"aliases": ["internal usb device rule b device class"], "anchor": "schema-internal_usb_device_rule--b_device_class", "description": "Class. The class of this device.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "b_device_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["internal usb device rule b device protocol"], "anchor": "schema-internal_usb_device_rule--b_device_protocol", "description": "The protocol (within the sub-class) of this device.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "b_device_protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["internal usb device rule b device sub class"], "anchor": "schema-internal_usb_device_rule--b_device_sub_class", "description": "The sub-class (within the class) of this device.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "b_device_sub_class"], "syntax": "attribute", "type": "string"}, {"aliases": ["internal usb device rule i serial"], "anchor": "schema-internal_usb_device_rule--i_serial", "description": "Index of Serial Number String Descriptor.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "i_serial"], "syntax": "attribute", "type": "string"}, {"aliases": ["internal usb device rule id product"], "anchor": "schema-internal_usb_device_rule--id_product", "description": "Product ID (Assigned by Manufacturer) in hex.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "id_product"], "syntax": "attribute", "type": "string"}, {"aliases": ["internal usb device rule id vendor"], "anchor": "schema-internal_usb_device_rule--id_vendor", "description": "Vendor ID. Vendor ID (Assigned by USB Org) in hex.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["internal_usb_device_rule", "id_vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/internal_usb_device_rule/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of internal USB device rules for server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
