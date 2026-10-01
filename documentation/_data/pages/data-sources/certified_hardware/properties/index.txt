---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_certified_hardware."
xcsh_docs: {"aliases": [], "body_bytes": 13802, "body_sha256": "sha256:a7c54201f63bb5dfd009afd582524ce05c0e30d93309574b0ef45c2820d78180", "child_ids": ["xcsh-docs:data-sources:certified_hardware:properties:devices", "xcsh-docs:data-sources:certified_hardware:properties:image_list", "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "xcsh-docs:data-sources:certified_hardware:properties:vendor_model_list"], "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:reference", "parent_id": "xcsh-docs:data-sources:certified_hardware:fundamentals", "path": "documentation/data-sources/certified_hardware/properties/index.md", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_certified_hardware.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
