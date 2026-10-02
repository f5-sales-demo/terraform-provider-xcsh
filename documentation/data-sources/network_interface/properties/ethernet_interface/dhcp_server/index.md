---
page_title: "ethernet_interface.dhcp_server"
subcategory: ""
description: "Configuration parameter for dhcp server."
xcsh_docs: {"aliases": ["ethernet interface dhcp server"], "body_bytes": 3898, "body_sha256": "sha256:318cae670e9feb49678dae9784dcc7c37b7586663bc208507b81b7b09689d2e7", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_end", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_start", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:interface_ip_map"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "path": "documentation/data-sources/network_interface/properties/ethernet_interface/dhcp_server/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2201300003132010-2001113313300220-3310301103130123-0223210020331113-3222121212002301-1002313312212312-2323031230022222-1220213121000033", "registry_path": "docs/guides/data-sources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "dhcp_server"], "schema_version": 1, "sections": [{"aliases": ["automatic from end"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_end", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "automatic_from_end"], "syntax": "attribute", "type": "object"}, {"aliases": ["automatic from start"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:automatic_from_start", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "automatic_from_start"], "syntax": "attribute", "type": "object"}, {"aliases": ["dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks"], "syntax": "attribute", "type": "object"}, {"aliases": ["dhcp option82 tag"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_option82_tag", "description": "DHCP option 82 tag.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_option82_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["fixed ip map"], "anchor": "schema-ethernet_interface--dhcp_server--fixed_ip_map", "description": "Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "fixed_ip_map"], "syntax": "attribute", "type": "map"}, {"aliases": ["interface ip map"], "anchor": "section", "description": "Specify static IPv4 addresses per node.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:dhcp_server:interface_ip_map", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "interface_ip_map"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for dhcp server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.dhcp_server

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/)
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

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/): complete subsection reference.

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

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/interface_ip_map/): complete subsection reference.

## Next pages

- [ethernet_interface.dhcp_server.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/automatic_from_end/)
- [ethernet_interface.dhcp_server.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/automatic_from_start/)
- [ethernet_interface.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/)
- [ethernet_interface.dhcp_server.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/dhcp_server/interface_ip_map/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
