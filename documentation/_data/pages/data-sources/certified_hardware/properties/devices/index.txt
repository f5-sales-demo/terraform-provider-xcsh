---
page_title: "devices"
subcategory: ""
description: "devices for xcsh_certified_hardware."
xcsh_docs: {"aliases": [], "body_bytes": 3688, "body_sha256": "sha256:e5ff2fc3243296932ad44be532082e014434e5d51f7b0ca4424dd850f7911e38", "child_ids": [], "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:devices", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/devices/index.md", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["devices"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "devices for xcsh_certified_hardware.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
