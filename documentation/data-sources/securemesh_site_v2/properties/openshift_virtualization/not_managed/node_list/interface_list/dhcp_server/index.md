---
page_title: "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server"
subcategory: ""
description: "DHCP server configuration for this interface."
xcsh_docs: {"aliases": ["openshift virtualization not managed node list interface list dhcp server"], "body_bytes": 5246, "body_sha256": "sha256:51beb2d8ec78d9679d5ddc0da8e6b46bd7d0822bcd465ab0fd21eca5c9f226bd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:interface_ip_map"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1303223332121112-3103203011131220-0203232011211120-1312002032330003-2012303323201203-2322113230013031-1132311310100300-0110332321320231", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server"], "schema_version": 1, "sections": [{"aliases": ["automatic from end"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "automatic_from_end"], "syntax": "attribute", "type": "object"}, {"aliases": ["automatic from start"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "automatic_from_start"], "syntax": "attribute", "type": "object"}, {"aliases": ["dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_networks"], "syntax": "attribute", "type": "object"}, {"aliases": ["dhcp option82 tag"], "anchor": "schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--dhcp_option82_tag", "description": "DHCP option 82 tag.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_option82_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["fixed ip map"], "anchor": "schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--fixed_ip_map", "description": "Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "fixed_ip_map"], "syntax": "attribute", "type": "map"}, {"aliases": ["interface ip map"], "anchor": "section", "description": "Specify static IPv4 addresses per node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "interface_ip_map"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DHCP server configuration for this interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization.not_managed.node_list.interface_list.dhcp_server

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [openshift_virtualization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/)
- [openshift_virtualization.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/)
- [openshift_virtualization.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/)
- [openshift_virtualization.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server

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

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/): complete subsection reference.

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--fixed_ip_map"></a>

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

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/): complete subsection reference.

## Next pages

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_end/)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_start/)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/)
- [openshift_virtualization.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
