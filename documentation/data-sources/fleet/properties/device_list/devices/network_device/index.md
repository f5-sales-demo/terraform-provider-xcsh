---
page_title: "device_list.devices.network_device"
subcategory: ""
description: "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in "
xcsh_docs: {"aliases": ["device list devices network device"], "body_bytes": 3594, "body_sha256": "sha256:e4a958c4af0f1be067f0fe62090534e0b0e83737717a0e3a941c72eb833c4028", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device:interface"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device", "parent_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices", "path": "documentation/data-sources/fleet/properties/device_list/devices/network_device/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2131311101003331-0110112131012232-1121310012222213-1000103201021121-0000131232300330-1023322300032223-3331100203312202-0321013001002303", "registry_path": "docs/guides/data-sources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list", "devices", "network_device"], "schema_version": 1, "sections": [{"aliases": ["interface"], "anchor": "section", "description": "Network Interface attributes for the device. User network interface configuration for this network device. Attributes like labels, MTU from the 'interface' are applied to corresponding interface in VER node If network interface refers to a virtual-network, the virtual-netowrk type must be consistent with use attribute ", "document_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["device_list", "devices", "network_device", "interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["use"], "anchor": "schema-device_list--devices--network_device--use", "description": "Defines how the device is used If networking device is owned by VER, it is available for users to configure as required If networking device is owned by VER, it is included in bootstrap config and member of outside network. If networking device is owned by VER, it is included in bootstrap config and member of inside ne", "document_id": "xcsh-docs:data-sources:fleet:properties:device_list:devices:network_device", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "network_device", "use"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/device_list/devices/network_device/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in ", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
