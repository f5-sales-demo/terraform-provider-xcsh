---
page_title: "device_list.devices.network_device"
subcategory: ""
description: "device_list.devices.network_device for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 3671, "body_sha256": "sha256:b234a583904a44a313086e18f8fab28c34f138a3d91df03d58b9c756a64bd6e8", "canonical_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "parent_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "path": "docs/guides/resources--fleet--properties--device_list--devices--network_device.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["device_list", "devices", "network_device"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/devices/network_device/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "device_list.devices.network_device for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list.devices.network_device

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [device_list](resources--fleet--properties--device_list.md)
- [device_list.devices](resources--fleet--properties--device_list--devices.md)
- device_list.devices.network_device

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Upstream description:

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Device mapping to nodes

A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An
interface in node inherits configuration from a device by matching, &#8203;- device\_name in Network
Interface for the device &#8203;- device name for physical-interface in the node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interface")}
```

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

Terraform syntax:

```terraform
network_device {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interface](resources--fleet--properties--device_list--devices--network_device--interface.md): complete subsection reference.

<a id="schema-device_list--devices--network_device--use"></a>

### use property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NETWORK_INTERFACE_USE_REGULAR",
    "NETWORK_INTERFACE_USE_OUTSIDE",
    "NETWORK_INTERFACE_USE_INSIDE"),
}
```

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

- [device_list.devices.network_device.interface](resources--fleet--properties--device_list--devices--network_device--interface.md)
- [device_list.devices](resources--fleet--properties--device_list--devices.md)
- [xcsh_fleet](../resources/fleet.md)
