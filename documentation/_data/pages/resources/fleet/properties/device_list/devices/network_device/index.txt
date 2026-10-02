---
page_title: "device_list.devices.network_device"
subcategory: ""
description: "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in"
xcsh_docs: {"aliases": ["device list devices network device"], "body_bytes": 4073, "body_sha256": "sha256:d08154948f6c7833942c668595e52ede86f63e40856918b9f5c17da9ebb56cb7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "parent_id": "xcsh-docs:resources:fleet:properties:device_list:devices", "path": "documentation/resources/fleet/properties/device_list/devices/network_device/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "device_list.devices.network_device:RequiredObjectAttributes:interface", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["device_list", "devices", "network_device"], "schema_version": 1, "sections": [{"aliases": ["interface"], "anchor": "section", "description": "Network Interface attributes for the device. User network interface configuration for this network device. Attributes like labels, MTU from the 'interface' are applied to corresponding interface in VER node If network interface refers to a virtual-network, the virtual-netowrk type must be consistent with use attribute", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device:interface", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["device_list", "devices", "network_device", "interface"], "syntax": "block", "type": "object"}, {"aliases": ["use"], "anchor": "schema-device_list--devices--network_device--use", "description": "Defines how the device is used If networking device is owned by VER, it is available for users to configure as required If networking device is owned by VER, it is included in bootstrap config and member of outside network. If networking device is owned by VER, it is included in bootstrap config and member of inside", "document_id": "xcsh-docs:resources:fleet:properties:device_list:devices:network_device", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["device_list", "devices", "network_device", "use"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/devices/network_device/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Represents physical network interface. The 'interface' reference points to a Network Interface object. Attributes such as Labels, MTU from Network Interface must be applied to the device. Device mapping to nodes A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An interface in", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [device_list.devices.network_device.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/network_device/interface/)
- [device_list.devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/device_list/devices/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
