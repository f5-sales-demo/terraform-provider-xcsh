---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server"
subcategory: ""
description: "Configuration parameter for dhcp server."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server"], "body_bytes": 7079, "body_sha256": "sha256:5cc463392d389e8ec5f6b7b0d49db41762fc1fb98dfc5f7bf1c77fd48064c5c6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_end", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_start", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:interface_ip_map"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "path": "documentation/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3313213021322122-1010212200031300-3231310301132213-3032330000302101-2002123121031231-2330221130022000-3311311112303222-3230203002321032", "registry_path": "docs/guides/resources--voltstack_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server:RequiredObjectAttributes:dhcp_networks", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces ethernet interface dhcp server automatic from end"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_end", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "automatic_from_end"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server automatic from start"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:automatic_from_start", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "automatic_from_start"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dns_address", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:same_as_dgw", "type": "conflicts"}], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp option82 tag"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_option82_tag", "description": "DHCP option 82 tag.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_option82_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server fixed ip map"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--fixed_ip_map", "description": "Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "fixed_ip_map"], "syntax": "attribute", "type": "map"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server interface ip map"], "anchor": "section", "description": "Specify static IPv4 addresses per node.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:interface_ip_map", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "interface_ip_map"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration parameter for dhcp server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
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

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--fixed_ip_map"></a>

### fixed_ip_map property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

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
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/interface_ip_map/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/automatic_from_end/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/automatic_from_start/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/interface_ip_map/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
