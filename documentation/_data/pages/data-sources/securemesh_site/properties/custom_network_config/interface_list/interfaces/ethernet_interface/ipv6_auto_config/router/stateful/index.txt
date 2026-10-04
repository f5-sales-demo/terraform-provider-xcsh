---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful"
subcategory: ""
description: "DHCPIPV6 Stateful Server."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful"], "body_bytes": 6924, "body_sha256": "sha256:0e11c3f9bb7639c8eef699117ddd85fc0f61cf457e4b5122a81dc47d78887e1e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_end", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_start", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:dhcp_networks", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:interface_ip_map"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3313303221102132-3232330322221311-0130032001103331-3310303022103310-2122003200103232-3201230202102213-3333003302022330-3001223333031321", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful automatic from end"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_end", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "automatic_from_end"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful automatic from start"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_start", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "automatic_from_start"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP server can allocate IP addresses.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "dhcp_networks"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful fixed ip map"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--fixed_ip_map", "description": "Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6 addresses based on the MAC Address of the DHCP Client.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "fixed_ip_map"], "syntax": "attribute", "type": "map"}, {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful interface ip map"], "anchor": "section", "description": "Map of Interface IPv6 assignments per node.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:interface_ip_map", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "interface_ip_map"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "DHCPIPV6 Stateful Server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful

<a id="section"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

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

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--fixed_ip_map"></a>

### fixed_ip_map property

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
      "type": "string"
    }
  },
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
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/interface_ip_map/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_end/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_start/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/interface_ip_map/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
