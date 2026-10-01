---
page_title: "ethernet_interface.dhcp_server"
subcategory: ""
description: "ethernet_interface.dhcp_server for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 3752, "body_sha256": "sha256:65743044a748baafbf53d0f0b0c14cdfad1360f7eac0c5f969243719fe723b07", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_end", "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_start", "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:interface_ip_map"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--dhcp_server.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "dhcp_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.dhcp_server for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.dhcp_server

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- ethernet_interface.dhcp_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dhcp server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [automatic_from_end](resources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](resources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](resources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md): complete subsection reference.

<a id="schema-ethernet_interface--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="schema-ethernet_interface--dhcp_server--fixed_ip_map"></a>

### fixed_ip_map property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--network_interface--properties--ethernet_interface--dhcp_server--interface_ip_map.md): complete subsection reference.

## Next pages

- [ethernet_interface.dhcp_server.automatic_from_end](resources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_end.md)
- [ethernet_interface.dhcp_server.automatic_from_start](resources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_start.md)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md)
- [ethernet_interface.dhcp_server.interface_ip_map](resources--network_interface--properties--ethernet_interface--dhcp_server--interface_ip_map.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [xcsh_network_interface](../resources/network_interface.md)
