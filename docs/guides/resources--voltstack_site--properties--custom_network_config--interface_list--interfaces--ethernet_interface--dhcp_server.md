---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 5061, "body_sha256": "sha256:49d99e5e9a8e42a86dd5231c7fd49681e606a4d6198cce16f0e284077f85ec4c", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_end", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_start", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:interface_ip_map"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](resources--voltstack_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server

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

- [automatic_from_end](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--fixed_ip_map"></a>

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

- [interface_ip_map](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map.md): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_end.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--automatic_from_start.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
