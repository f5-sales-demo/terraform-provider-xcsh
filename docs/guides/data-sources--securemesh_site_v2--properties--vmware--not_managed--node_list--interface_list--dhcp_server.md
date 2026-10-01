---
page_title: "vmware.not_managed.node_list.interface_list.dhcp_server"
subcategory: ""
description: "vmware.not_managed.node_list.interface_list.dhcp_server for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 4009, "body_sha256": "sha256:b76362640fe32161777ed61493a0ccaee9086ad14979fd73e96784705c74b2fb", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server:interface_ip_map"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "path": "docs/guides/data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "dhcp_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vmware.not_managed.node_list.interface_list.dhcp_server for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware.not_managed.node_list.interface_list.dhcp_server

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [vmware](data-sources--securemesh_site_v2--properties--vmware.md)
- [vmware.not_managed](data-sources--securemesh_site_v2--properties--vmware--not_managed.md)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list.md)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list.md)
- vmware.not_managed.node_list.interface_list.dhcp_server

<a id="section"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

## Direct properties

- [automatic_from_end](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--dhcp_networks.md): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--interface_list--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="schema-vmware--not_managed--node_list--interface_list--dhcp_server--fixed_ip_map"></a>

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

- [interface_ip_map](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--interface_ip_map.md): complete subsection reference.

## Next pages

- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--automatic_from_end.md)
- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--automatic_from_start.md)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--dhcp_networks.md)
- [vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--dhcp_server--interface_ip_map.md)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
