---
page_title: "device_list.devices.network_device"
subcategory: ""
description: "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in"
xcsh_docs: {"aliases": ["device list devices network device"], "body_bytes": 3816, "body_sha256": "sha256:043fbb5c05df14a61cff45d79fda605d8448fe2edad8577933c1acc15b2b46e2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "parent_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "path": "documentation/resources/fleet/properties/device_list/devices/network_device/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "device_list.devices.network_device:RequiredObjectAttributes:interface", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list", "devices", "network_device"], "schema_version": 1, "sections": [{"aliases": ["device list devices network device interface"], "anchor": "section", "description": "Network Interface attributes for the device. User network interface configuration for this network device. Attributes like labels, MTU from the 'interface' are applied to corresponding interface in VER node If network interface refers to a virtual-network, the virtual-netowrk type must be consistent with use attribute", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["device_list", "devices", "network_device", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["device list devices network device use"], "anchor": "schema-device_list--devices--network_device--use", "description": "Defines how the device is used If networking device is owned by VER, it is available for users to configure as required If networking device is owned by VER, it is included in bootstrap config and member of outside network. If networking device is owned by VER, it is included in bootstrap config and member of inside", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["NETWORK_INTERFACE_USE_INSIDE", "NETWORK_INTERFACE_USE_OUTSIDE", "NETWORK_INTERFACE_USE_REGULAR"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "network_device", "use"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/devices/network_device/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fleetCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list.devices.network_device

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/)
- [device_list.devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/)
- device_list.devices.network_device

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Device mapping to nodes

A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An
interface in node inherits configuration from a device by matching, &#8203;- device\_name in Network
Interface for the device &#8203;- device name for physical-interface in the node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/network_device/interface/): complete subsection reference.

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

Additional upstream details:

Defines how the device is used

If networking device is owned by VER, it is available for users to configure as required If
networking device is owned by VER, it is included in bootstrap config and member of outside network.
If networking device is owned by VER, it is included in bootstrap config and member of inside
network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NETWORK_INTERFACE_USE_INSIDE","NETWORK_INTERFACE_USE_OUTSIDE","NETWORK_INTERFACE_USE_REGULAR"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
