---
page_title: "ethernet_interface.dhcp_server"
subcategory: ""
description: "ethernet_interface.dhcp_server for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 3150, "body_sha256": "sha256:7cfdf0ec5be80eee006bd7a579d423f70e9b946646266869627a1af61da22508", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_end", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_start", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:interface_ip_map"], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "path": "docs/guides/data-sources--network_interface--properties--ethernet_interface--dhcp_server.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "dhcp_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.dhcp_server for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ethernet_interface.dhcp_server

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- ethernet_interface.dhcp_server

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dhcp server.

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

## Direct properties

- [automatic_from_end](data-sources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](data-sources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md): complete subsection reference.

<a id="schema-ethernet_interface--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="schema-ethernet_interface--dhcp_server--fixed_ip_map"></a>

### fixed_ip_map property

Type: `["map", "string"]`. Computed.

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

- [interface_ip_map](data-sources--network_interface--properties--ethernet_interface--dhcp_server--interface_ip_map.md): complete subsection reference.

## Next pages

- [ethernet_interface.dhcp_server.automatic_from_end](data-sources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_end.md)
- [ethernet_interface.dhcp_server.automatic_from_start](data-sources--network_interface--properties--ethernet_interface--dhcp_server--automatic_from_start.md)
- [ethernet_interface.dhcp_server.dhcp_networks](data-sources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md)
- [ethernet_interface.dhcp_server.interface_ip_map](data-sources--network_interface--properties--ethernet_interface--dhcp_server--interface_ip_map.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
