---
page_title: "device_list.devices.network_device"
subcategory: ""
description: "device_list.devices.network_device for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 3594, "body_sha256": "sha256:e4a958c4af0f1be067f0fe62090534e0b0e83737717a0e3a941c72eb833c4028", "child_ids": ["xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device:interface"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device", "parent_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices", "path": "documentation/data-sources/fleet/properties/device_list/devices/network_device/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["device_list", "devices", "network_device"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/device_list/devices/network_device/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "device_list.devices.network_device for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list.devices.network_device

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/)
- [device_list.devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/devices/)
- device_list.devices.network_device

<a id="section"></a>

Type: `"single"`. Computed.

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Upstream description:

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Device mapping to nodes

A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An
interface in node inherits configuration from a device by matching, &#8203;- device\_name in Network
Interface for the device &#8203;- device name for physical-interface in the node.

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

## Direct properties

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/devices/network_device/interface/): complete subsection reference.

<a id="schema-device_list--devices--network_device--use"></a>

### use property

Type: `"string"`. Computed.

\[Enum:
NETWORK\_INTERFACE\_USE\_REGULAR|NETWORK\_INTERFACE\_USE\_OUTSIDE|NETWORK\_INTERFACE\_USE\_INSIDE\]
Defines how the device is used If networking device is owned by VER, it is available for users to
configure as required If networking device is owned by VER, it is included in bootstrap config and
member of outside network. If networking device is owned by VER, it is included in bootstrap
config.. Possible values are \`NETWORK\_INTERFACE\_USE\_REGULAR\`,
\`NETWORK\_INTERFACE\_USE\_OUTSIDE\`, \`NETWORK\_INTERFACE\_USE\_INSIDE\`. Defaults to
\`NETWORK\_INTERFACE\_USE\_REGULAR\`.

Upstream description:

Defines how the device is used

If networking device is owned by VER, it is available for users to configure as required If
networking device is owned by VER, it is included in bootstrap config and member of outside network.
If networking device is owned by VER, it is included in bootstrap config and member of inside
network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NETWORK_INTERFACE_USE_REGULAR",
  "enum": [
    "NETWORK_INTERFACE_USE_REGULAR",
    "NETWORK_INTERFACE_USE_OUTSIDE",
    "NETWORK_INTERFACE_USE_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [device_list.devices.network_device.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/devices/network_device/interface/)
- [device_list.devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/device_list/devices/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
