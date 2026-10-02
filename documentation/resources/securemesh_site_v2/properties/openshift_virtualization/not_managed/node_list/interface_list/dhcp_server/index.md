---
page_title: "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server"
subcategory: ""
description: "DHCP server configuration for this interface."
xcsh_docs: {"aliases": ["openshift virtualization not managed node list interface list dhcp server"], "body_bytes": 5740, "body_sha256": "sha256:166fceb1dfe2264730070d0003fde601702230d80042b7eb30078bbfe372f9d8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:interface_ip_map"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-015.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server:RequiredObjectAttributes:dhcp_networks", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server"], "schema_version": 1, "sections": [{"aliases": ["automatic from end"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_end", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "automatic_from_end"], "syntax": "attribute", "type": "object"}, {"aliases": ["automatic from start"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:automatic_from_start", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "automatic_from_start"], "syntax": "attribute", "type": "object"}, {"aliases": ["dhcp networks"], "anchor": "section", "description": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--dns_address", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:same_as_dgw", "type": "conflicts"}], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_networks"], "syntax": "block", "type": "object"}, {"aliases": ["dhcp option82 tag"], "anchor": "schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--dhcp_option82_tag", "description": "DHCP option 82 tag.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_option82_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["fixed ip map"], "anchor": "schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--fixed_ip_map", "description": "Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "fixed_ip_map"], "syntax": "attribute", "type": "map"}, {"aliases": ["interface ip map"], "anchor": "section", "description": "Specify static IPv4 addresses per node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "dhcp_server", "interface_ip_map"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DHCP server configuration for this interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization.not_managed.node_list.interface_list.dhcp_server

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [openshift_virtualization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/)
- [openshift_virtualization.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/)
- [openshift_virtualization.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/)
- [openshift_virtualization.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

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

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/): complete subsection reference.

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--dhcp_server--fixed_ip_map"></a>

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

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/): complete subsection reference.

## Next pages

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_end/)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/automatic_from_start/)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/)
- [openshift_virtualization.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
