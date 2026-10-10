---
page_title: "openstack.not_managed.node_list.interface_list.dhcp_server"
subcategory: ""
description: "DHCP server configuration for this interface."
xcsh_docs: {"aliases": ["openstack not managed node list interface list dhcp server"], "body_bytes": 4116, "body_sha256": "sha256:ea39b4827cf1be01fc18607a99b1a5e56a905b4a588ca9bb2b0cd5b8af3eee9c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:interface_ip_map"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0233013223130000-2322300320033203-1201323010010330-2313130013002230-1220010030002320-0331221200030012-2303311002213203-3202322310032312", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server"], "schema_version": 1, "sections": [{"aliases": ["openstack not managed node list interface list dhcp server automatic from end"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server", "automatic_from_end"], "syntax": "attribute", "type": "object"}, {"aliases": ["openstack not managed node list interface list dhcp server automatic from start"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server", "automatic_from_start"], "syntax": "attribute", "type": "object"}, {"aliases": ["openstack not managed node list interface list dhcp server dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_networks"], "syntax": "attribute", "type": "object"}, {"aliases": ["openstack not managed node list interface list dhcp server dhcp option82 tag"], "anchor": "schema-openstack--not_managed--node_list--interface_list--dhcp_server--dhcp_option82_tag", "description": "DHCP option 82 tag.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_option82_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["openstack not managed node list interface list dhcp server fixed ip map"], "anchor": "schema-openstack--not_managed--node_list--interface_list--dhcp_server--fixed_ip_map", "description": "Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server", "fixed_ip_map"], "syntax": "attribute", "type": "map"}, {"aliases": ["openstack not managed node list interface list dhcp server interface ip map"], "anchor": "section", "description": "Specify static IPv4 addresses per node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "dhcp_server", "interface_ip_map"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "DHCP server configuration for this interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed.node_list.interface_list.dhcp_server

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/)
- [openstack.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/)
- [openstack.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/)
- [openstack.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/)
- openstack.not_managed.node_list.interface_list.dhcp_server

<a id="section"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

## Direct properties

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/): complete subsection reference.

<a id="schema-openstack--not_managed--node_list--interface_list--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="schema-openstack--not_managed--node_list--interface_list--dhcp_server--fixed_ip_map"></a>

### fixed_ip_map property

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/): complete subsection reference.
