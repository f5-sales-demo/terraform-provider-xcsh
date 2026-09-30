---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_certified_hardware."
xcsh_docs: {"aliases": [], "body_bytes": 10601, "body_sha256": "sha256:facab76289024ed1396b860d90b7855f3618b4d2fd138955256827de41c7763c", "canonical_id": "xcsh-docs:data-sources:certified_hardware:reference", "child_ids": ["xcsh-docs:data-sources:certified_hardware:properties:devices", "xcsh-docs:data-sources:certified_hardware:properties:image_list", "xcsh-docs:data-sources:certified_hardware:properties:internal_usb_device_rule", "xcsh-docs:data-sources:certified_hardware:properties:numa_mem", "xcsh-docs:data-sources:certified_hardware:properties:vendor_model_list"], "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:reference", "parent_id": "xcsh-docs:data-sources:certified_hardware:fundamentals", "path": "docs/guides/data-sources--certified_hardware--reference.md", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_certified_hardware.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
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

- [devices](data-sources--certified_hardware--properties--devices.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

- [image_list](data-sources--certified_hardware--properties--image_list.md): complete subsection reference.

- [internal_usb_device_rule](data-sources--certified_hardware--properties--internal_usb_device_rule.md): complete subsection reference.

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

- [numa_mem](data-sources--certified_hardware--properties--numa_mem.md): complete subsection reference.

<a id="schema-numa_nodes"></a>

### numa_nodes property

Type: `"number"`. Computed.

The number of host NUMA nodes used in certified hardware.

- [vendor_model_list](data-sources--certified_hardware--properties--vendor_model_list.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--certified_hardware--reference.md#schema-annotations) |
| `certified_hardware_type` | [certified_hardware_type](data-sources--certified_hardware--reference.md#schema-certified_hardware_type) |
| `description` | [description](data-sources--certified_hardware--reference.md#schema-description) |
| `devices` | [devices](data-sources--certified_hardware--properties--devices.md#section) |
| `devices.device_list` | [devices.device_list](data-sources--certified_hardware--properties--devices.md#schema-devices--device_list) |
| `devices.max_unit` | [devices.max_unit](data-sources--certified_hardware--properties--devices.md#schema-devices--max_unit) |
| `devices.min_unit` | [devices.min_unit](data-sources--certified_hardware--properties--devices.md#schema-devices--min_unit) |
| `devices.name` | [devices.name](data-sources--certified_hardware--properties--devices.md#schema-devices--name) |
| `devices.type` | [devices.type](data-sources--certified_hardware--properties--devices.md#schema-devices--type) |
| `devices.use` | [devices.use](data-sources--certified_hardware--properties--devices.md#schema-devices--use) |
| `id` | [id](data-sources--certified_hardware--reference.md#schema-id) |
| `image_list` | [image_list](data-sources--certified_hardware--properties--image_list.md#section) |
| `image_list.aws` | [image_list.aws](data-sources--certified_hardware--properties--image_list--aws.md#section) |
| `image_list.aws.image_id` | [image_list.aws.image_id](data-sources--certified_hardware--properties--image_list--aws--image_id.md#section) |
| `image_list.aws.image_id.image_id` | [image_list.aws.image_id.image_id](data-sources--certified_hardware--properties--image_list--aws--image_id.md#schema-image_list--aws--image_id--image_id) |
| `image_list.aws.image_id.region` | [image_list.aws.image_id.region](data-sources--certified_hardware--properties--image_list--aws--image_id.md#schema-image_list--aws--image_id--region) |
| `image_list.azure` | [image_list.azure](data-sources--certified_hardware--properties--image_list--azure.md#section) |
| `image_list.azure.image_id` | [image_list.azure.image_id](data-sources--certified_hardware--properties--image_list--azure--image_id.md#section) |
| `image_list.azure.image_id.image_id` | [image_list.azure.image_id.image_id](data-sources--certified_hardware--properties--image_list--azure--image_id.md#schema-image_list--azure--image_id--image_id) |
| `image_list.azure.marketplace` | [image_list.azure.marketplace](data-sources--certified_hardware--properties--image_list--azure--marketplace.md#section) |
| `image_list.azure.marketplace.name` | [image_list.azure.marketplace.name](data-sources--certified_hardware--properties--image_list--azure--marketplace.md#schema-image_list--azure--marketplace--name) |
| `image_list.azure.marketplace.offer` | [image_list.azure.marketplace.offer](data-sources--certified_hardware--properties--image_list--azure--marketplace.md#schema-image_list--azure--marketplace--offer) |
| `image_list.azure.marketplace.publisher` | [image_list.azure.marketplace.publisher](data-sources--certified_hardware--properties--image_list--azure--marketplace.md#schema-image_list--azure--marketplace--publisher) |
| `image_list.azure.marketplace.sku` | [image_list.azure.marketplace.sku](data-sources--certified_hardware--properties--image_list--azure--marketplace.md#schema-image_list--azure--marketplace--sku) |
| `image_list.azure.marketplace.version` | [image_list.azure.marketplace.version](data-sources--certified_hardware--properties--image_list--azure--marketplace.md#schema-image_list--azure--marketplace--version) |
| `image_list.gcp` | [image_list.gcp](data-sources--certified_hardware--properties--image_list--gcp.md#section) |
| `image_list.gcp.image_id` | [image_list.gcp.image_id](data-sources--certified_hardware--properties--image_list--gcp--image_id.md#section) |
| `image_list.gcp.image_id.image_id` | [image_list.gcp.image_id.image_id](data-sources--certified_hardware--properties--image_list--gcp--image_id.md#schema-image_list--gcp--image_id--image_id) |
| `image_list.name` | [image_list.name](data-sources--certified_hardware--properties--image_list.md#schema-image_list--name) |
| `image_list.provider_ref` | [image_list.provider_ref](data-sources--certified_hardware--properties--image_list.md#schema-image_list--provider_ref) |
| `internal_usb_device_rule` | [internal_usb_device_rule](data-sources--certified_hardware--properties--internal_usb_device_rule.md#section) |
| `internal_usb_device_rule.b_device_class` | [internal_usb_device_rule.b_device_class](data-sources--certified_hardware--properties--internal_usb_device_rule.md#schema-internal_usb_device_rule--b_device_class) |
| `internal_usb_device_rule.b_device_protocol` | [internal_usb_device_rule.b_device_protocol](data-sources--certified_hardware--properties--internal_usb_device_rule.md#schema-internal_usb_device_rule--b_device_protocol) |
| `internal_usb_device_rule.b_device_sub_class` | [internal_usb_device_rule.b_device_sub_class](data-sources--certified_hardware--properties--internal_usb_device_rule.md#schema-internal_usb_device_rule--b_device_sub_class) |
| `internal_usb_device_rule.i_serial` | [internal_usb_device_rule.i_serial](data-sources--certified_hardware--properties--internal_usb_device_rule.md#schema-internal_usb_device_rule--i_serial) |
| `internal_usb_device_rule.id_product` | [internal_usb_device_rule.id_product](data-sources--certified_hardware--properties--internal_usb_device_rule.md#schema-internal_usb_device_rule--id_product) |
| `internal_usb_device_rule.id_vendor` | [internal_usb_device_rule.id_vendor](data-sources--certified_hardware--properties--internal_usb_device_rule.md#schema-internal_usb_device_rule--id_vendor) |
| `labels` | [labels](data-sources--certified_hardware--reference.md#schema-labels) |
| `mem_page_number` | [mem_page_number](data-sources--certified_hardware--reference.md#schema-mem_page_number) |
| `mem_page_size` | [mem_page_size](data-sources--certified_hardware--reference.md#schema-mem_page_size) |
| `name` | [name](data-sources--certified_hardware--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--certified_hardware--reference.md#schema-namespace) |
| `numa_mem` | [numa_mem](data-sources--certified_hardware--properties--numa_mem.md#section) |
| `numa_mem.memory` | [numa_mem.memory](data-sources--certified_hardware--properties--numa_mem.md#schema-numa_mem--memory) |
| `numa_mem.node` | [numa_mem.node](data-sources--certified_hardware--properties--numa_mem.md#schema-numa_mem--node) |
| `numa_nodes` | [numa_nodes](data-sources--certified_hardware--reference.md#schema-numa_nodes) |
| `vendor_model_list` | [vendor_model_list](data-sources--certified_hardware--properties--vendor_model_list.md#section) |
| `vendor_model_list.model` | [vendor_model_list.model](data-sources--certified_hardware--properties--vendor_model_list.md#schema-vendor_model_list--model) |
| `vendor_model_list.vendor` | [vendor_model_list.vendor](data-sources--certified_hardware--properties--vendor_model_list.md#schema-vendor_model_list--vendor) |

## Next pages

- [devices](data-sources--certified_hardware--properties--devices.md)
- [image_list](data-sources--certified_hardware--properties--image_list.md)
- [internal_usb_device_rule](data-sources--certified_hardware--properties--internal_usb_device_rule.md)
- [numa_mem](data-sources--certified_hardware--properties--numa_mem.md)
- [vendor_model_list](data-sources--certified_hardware--properties--vendor_model_list.md)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
