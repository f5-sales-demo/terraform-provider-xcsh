---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_certified_hardware."
xcsh_docs: {"aliases": ["certified hardware"], "body_bytes": 13802, "body_sha256": "sha256:a7c54201f63bb5dfd009afd582524ce05c0e30d93309574b0ef45c2820d78180", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:certified_hardware:properties:devices", "xcsh-docs:data-sources:certified_hardware:properties:image_list", "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "xcsh-docs:data-sources:certified_hardware:properties:vendor_model_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:reference", "parent_id": "xcsh-docs:data-sources:certified_hardware:fundamentals", "path": "documentation/data-sources/certified_hardware/properties/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0302130231202120-2203110233320123-1120332022311120-3011032230213311-2103323001103203-0212122123333022-1013300331312111-0101311012103201", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["certified hardware type"], "anchor": "schema-certified_hardware_type", "description": "Different type of certified HW for billing rate. Possible values are `VOLTMESH`, `VOLTSTACK_COMBO`, `CLOUD_MARKET_PLACE`. Defaults to `VOLTMESH`.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["certified_hardware_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Description.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["devices"], "anchor": "section", "description": "List of supported devices in this model.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["devices"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["image list"], "anchor": "section", "description": "List of image names with providers for this certified hardware, e.g. AWS ami-0f99d090261d2acd5.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["image_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["internal usb device rule"], "anchor": "section", "description": "List of internal USB device rules for server.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["internal_usb_device_rule"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Labels.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["mem page number"], "anchor": "schema-mem_page_number", "description": "Number of pages allocated in this certified hardware for Hugepages. Each page size is defined above in 'mem_page_size' Total memory reserved for Hugepages is 'mem_page_size * mem_page_number'.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mem_page_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["mem page size"], "anchor": "schema-mem_page_size", "description": "Memory for packets buffers etc are allocated in blocks of pages. Size of each memory page is defined here. Invalid Page size Page size of 4KB Page size of 2MB Page size of 1GB. Possible values are `HARDWARE_MEM_PAGE_SIZE_INVALID`, `HARDWARE_MEM_PAGE_SIZE_4KB`, `HARDWARE_MEM_PAGE_SIZE_2MB`,", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["mem_page_size"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Name of the CertifiedHardware to look up.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace of the CertifiedHardware.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["numa mem"], "anchor": "section", "description": "List of Numa nodes with the number of MB of instance memory to map to node instance If not specified, memory is evenly divided among available NUMA nodes.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["numa_mem"], "syntax": "attribute", "type": "object"}, {"aliases": ["numa nodes"], "anchor": "schema-numa_nodes", "description": "The number of host NUMA nodes used in certified hardware.", "document_id": "xcsh-docs:data-sources:certified_hardware:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["numa_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["vendor model list"], "anchor": "section", "description": "List of supported hardware vendor and model for this certified hardware.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:vendor_model_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["vendor_model_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_certified_hardware.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

<a id="schema-certified_hardware_type"></a>

### certified_hardware_type property

Type: `"string"`. Computed.

\[Enum: VOLTMESH|VOLTSTACK\_COMBO|CLOUD\_MARKET\_PLACE\] Different type of certified HW for billing
rate. Possible values are \`VOLTMESH\`, \`VOLTSTACK\_COMBO\`, \`CLOUD\_MARKET\_PLACE\`. Defaults to
\`VOLTMESH\`.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

- [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

- [image_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/): complete subsection reference.

- [internal_usb_device_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-mem_page_number"></a>

### mem_page_number property

Type: `"number"`. Computed.

Number of pages allocated in this certified hardware for Hugepages. Each page size is defined above
in 'mem\_page\_size' Total memory reserved for Hugepages is 'mem\_page\_size \* mem\_page\_number'.

<a id="schema-mem_page_size"></a>

### mem_page_size property

Type: `"string"`. Computed.

\[Enum:
HARDWARE\_MEM\_PAGE\_SIZE\_INVALID|HARDWARE\_MEM\_PAGE\_SIZE\_4KB|HARDWARE\_MEM\_PAGE\_SIZE\_2MB|HARDWARE\_MEM\_PAGE\_SIZE\_1GB\]
Memory for packets buffers etc are allocated in blocks of pages. Size of each memory page is defined
here. Invalid Page size Page size of 4KB Page size of 2MB Page size of 1GB. Possible values are
\`HARDWARE\_MEM\_PAGE\_SIZE\_INVALID\`, \`HARDWARE\_MEM\_PAGE\_SIZE\_4KB\`,
\`HARDWARE\_MEM\_PAGE\_SIZE\_2MB\`, \`HARDWARE\_MEM\_PAGE\_SIZE\_1GB\`. Defaults to
\`HARDWARE\_MEM\_PAGE\_SIZE\_INVALID\`.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the CertifiedHardware to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the CertifiedHardware.

- [numa_mem](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/numa_mem/): complete subsection reference.

<a id="schema-numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Computed.

The number of host NUMA nodes used in certified hardware.

- [vendor_model_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/vendor_model_list/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-annotations) |
| `certified_hardware_type` | [certified_hardware_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-certified_hardware_type) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-description) |
| `devices` | [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/#section) |
| `devices.device_list` | [devices.device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/#schema-devices--device_list) |
| `devices.max_unit` | [devices.max_unit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/#schema-devices--max_unit) |
| `devices.min_unit` | [devices.min_unit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/#schema-devices--min_unit) |
| `devices.name` | [devices.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/#schema-devices--name) |
| `devices.type` | [devices.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/#schema-devices--type) |
| `devices.use` | [devices.use](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/#schema-devices--use) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-id) |
| `image_list` | [image_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/#section) |
| `image_list.aws` | [image_list.aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/aws/#section) |
| `image_list.aws.image_id` | [image_list.aws.image_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/aws/image_id/#section) |
| `image_list.aws.image_id.image_id` | [image_list.aws.image_id.image_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/aws/image_id/#schema-image_list--aws--image_id--image_id) |
| `image_list.aws.image_id.region` | [image_list.aws.image_id.region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/aws/image_id/#schema-image_list--aws--image_id--region) |
| `image_list.azure` | [image_list.azure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/#section) |
| `image_list.azure.image_id` | [image_list.azure.image_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/image_id/#section) |
| `image_list.azure.image_id.image_id` | [image_list.azure.image_id.image_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/image_id/#schema-image_list--azure--image_id--image_id) |
| `image_list.azure.marketplace` | [image_list.azure.marketplace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/marketplace/#section) |
| `image_list.azure.marketplace.name` | [image_list.azure.marketplace.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/marketplace/#schema-image_list--azure--marketplace--name) |
| `image_list.azure.marketplace.offer` | [image_list.azure.marketplace.offer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/marketplace/#schema-image_list--azure--marketplace--offer) |
| `image_list.azure.marketplace.publisher` | [image_list.azure.marketplace.publisher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/marketplace/#schema-image_list--azure--marketplace--publisher) |
| `image_list.azure.marketplace.sku` | [image_list.azure.marketplace.sku](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/marketplace/#schema-image_list--azure--marketplace--sku) |
| `image_list.azure.marketplace.version` | [image_list.azure.marketplace.version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/marketplace/#schema-image_list--azure--marketplace--version) |
| `image_list.gcp` | [image_list.gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/gcp/#section) |
| `image_list.gcp.image_id` | [image_list.gcp.image_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/gcp/image_id/#section) |
| `image_list.gcp.image_id.image_id` | [image_list.gcp.image_id.image_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/gcp/image_id/#schema-image_list--gcp--image_id--image_id) |
| `image_list.name` | [image_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/#schema-image_list--name) |
| `image_list.provider_ref` | [image_list.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/#schema-image_list--provider_ref) |
| `internal_usb_device_rule` | [internal_usb_device_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/#section) |
| `internal_usb_device_rule.b_device_class` | [internal_usb_device_rule.b_device_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/#schema-internal_usb_device_rule--b_device_class) |
| `internal_usb_device_rule.b_device_protocol` | [internal_usb_device_rule.b_device_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/#schema-internal_usb_device_rule--b_device_protocol) |
| `internal_usb_device_rule.b_device_sub_class` | [internal_usb_device_rule.b_device_sub_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/#schema-internal_usb_device_rule--b_device_sub_class) |
| `internal_usb_device_rule.i_serial` | [internal_usb_device_rule.i_serial](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/#schema-internal_usb_device_rule--i_serial) |
| `internal_usb_device_rule.id_product` | [internal_usb_device_rule.id_product](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/#schema-internal_usb_device_rule--id_product) |
| `internal_usb_device_rule.id_vendor` | [internal_usb_device_rule.id_vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/#schema-internal_usb_device_rule--id_vendor) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-labels) |
| `mem_page_number` | [mem_page_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-mem_page_number) |
| `mem_page_size` | [mem_page_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-mem_page_size) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-namespace) |
| `numa_mem` | [numa_mem](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/numa_mem/#section) |
| `numa_mem.memory` | [numa_mem.memory](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/numa_mem/#schema-numa_mem--memory) |
| `numa_mem.node` | [numa_mem.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/numa_mem/#schema-numa_mem--node) |
| `numa_nodes` | [numa_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/#schema-numa_nodes) |
| `vendor_model_list` | [vendor_model_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/vendor_model_list/#section) |
| `vendor_model_list.model` | [vendor_model_list.model](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/vendor_model_list/#schema-vendor_model_list--model) |
| `vendor_model_list.vendor` | [vendor_model_list.vendor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/vendor_model_list/#schema-vendor_model_list--vendor) |

## Next pages

- [devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/devices/)
- [image_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/)
- [internal_usb_device_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/internal_usb_device_rule/)
- [numa_mem](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/numa_mem/)
- [vendor_model_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/vendor_model_list/)
- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
