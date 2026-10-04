---
page_title: "devices"
subcategory: ""
description: "List of supported devices in this model."
xcsh_docs: {"aliases": ["devices"], "body_bytes": 3688, "body_sha256": "sha256:e5ff2fc3243296932ad44be532082e014434e5d51f7b0ca4424dd850f7911e38", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/devices/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1323223232320303-3210303020021333-0211300220322202-3033001020330232-1122111132321300-2001101311013022-0031101111011002-3103231103301231", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["devices"], "schema_version": 1, "sections": [{"aliases": ["devices device list"], "anchor": "schema-devices--device_list", "description": "In case of logical boot strap devices like LACP Link aggregation or RAID.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "device_list"], "syntax": "attribute", "type": "list"}, {"aliases": ["devices max unit"], "anchor": "schema-devices--max_unit", "description": "Last unit number of the device supported in this certified hardware.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "max_unit"], "syntax": "attribute", "type": "number"}, {"aliases": ["devices min unit"], "anchor": "schema-devices--min_unit", "description": "First unit number of the device supported in this certified hardware.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "min_unit"], "syntax": "attribute", "type": "number"}, {"aliases": ["devices name"], "anchor": "schema-devices--name", "description": "Device. Name of the device.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["devices type"], "anchor": "schema-devices--type", "description": "Different type of devices supported Invalid device or device that's not supported Ethernet device VIRTIO device TUNTAP device LACP based bond interface External iSCSI devices supported Nvidia GPU device used for machine learning. Possible values are `HARDWARE_DEVICE_INVALID`, `HARDWARE_DEVICE_ETHERNET`,", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "type"], "syntax": "attribute", "type": "string"}, {"aliases": ["devices use"], "anchor": "schema-devices--use", "description": "Defines how the device instance must be used If the device is owned by F5 Distributed Cloud software, it is available for users to configure as required Device reserved for internal use by F5 Distributed Cloud Node If the Network device is owned by VER, it is available for users to configure as.. Possible values are", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["devices", "use"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/devices/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of supported devices in this model.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# devices

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- devices

<a id="section"></a>

Type: `"list"`. Computed.

List of supported devices in this model.

## Direct properties

<a id="schema-devices--device_list"></a>

### device_list property

Type: `["list", "string"]`. Computed.

In case of logical boot strap devices like LACP Link aggregation or RAID.

<a id="schema-devices--max_unit"></a>

### max_unit property

Type: `"number"`. Computed.

Last unit number of the device supported in this certified hardware.

<a id="schema-devices--min_unit"></a>

### min_unit property

Type: `"number"`. Computed.

First unit number of the device supported in this certified hardware.

<a id="schema-devices--name"></a>

### name property

Type: `"string"`. Computed.

Device. Name of the device.

<a id="schema-devices--type"></a>

### type property

Type: `"string"`. Computed.

\[Enum:
HARDWARE\_DEVICE\_INVALID|HARDWARE\_DEVICE\_ETHERNET|HARDWARE\_DEVICE\_VIRTIO|HARDWARE\_DEVICE\_TUNTAP|HARDWARE\_DEVICE\_BOND|HARDWARE\_DEVICE\_EXTERNAL\_ISCSI\_STORTAGE|HARDWARE\_DEVICE\_NVIDIA\_GPU\]
Different type of devices supported Invalid device or device that's not supported Ethernet device
VIRTIO device TUNTAP device LACP based bond interface External iSCSI devices supported Nvidia GPU
device used for machine learning. Possible values are \`HARDWARE\_DEVICE\_INVALID\`,
\`HARDWARE\_DEVICE\_ETHERNET\`, \`HARDWARE\_DEVICE\_VIRTIO\`, \`HARDWARE\_DEVICE\_TUNTAP\`,
\`HARDWARE\_DEVICE\_BOND\`, \`HARDWARE\_DEVICE\_EXTERNAL\_ISCSI\_STORTAGE\`,
\`HARDWARE\_DEVICE\_NVIDIA\_GPU\`. Defaults to \`HARDWARE\_DEVICE\_INVALID\`.

<a id="schema-devices--use"></a>

### use property

Type: `"string"`. Computed.

\[Enum:
HARDWARE\_DEVICE\_USE\_REGULAR|HARDWARE\_DEVICE\_USE\_INTERNAL|HARDWARE\_NETWORK\_DEVICE\_USE\_REGULAR|HARDWARE\_NETWORK\_DEVICE\_USE\_INTERNAL|HARDWARE\_NETWORK\_DEVICE\_USE\_MANAGEMENT|HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE|HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE|HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE\_LAG|HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE\_LAG|HARDWARE\_NETWORK\_DEVICE\_USE\_LAG\_MEMBER|HARDWARE\_NETWORK\_DEVICE\_USE\_STORAGE|HARDWARE\_NETWORK\_DEVICE\_USE\_FALLBACK\_MANAGEMENT\]
Defines how the device instance must be used If the device is owned by F5 Distributed Cloud
software, it is available for users to configure as required Device reserved for internal use by F5
Distributed Cloud Node If the Network device is owned by VER, it is available for users to configure
as.. Possible values are \`HARDWARE\_DEVICE\_USE\_REGULAR\`, \`HARDWARE\_DEVICE\_USE\_INTERNAL\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_REGULAR\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_INTERNAL\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_MANAGEMENT\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE\_LAG\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE\_LAG\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_LAG\_MEMBER\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_STORAGE\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_FALLBACK\_MANAGEMENT\`. Defaults to
\`HARDWARE\_DEVICE\_USE\_REGULAR\`.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
