---
page_title: "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful"
subcategory: ""
description: "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 6253, "body_sha256": "sha256:971ac09e01ad594a98582cc95143036c023236d922e533a16a5edd8cbc88b793", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "documentation/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [openshift_virtualization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/)
- [openshift_virtualization.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/)
- [openshift_virtualization.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/)
- [openshift_virtualization.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

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

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/): complete subsection reference.

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--fixed_ip_map"></a>

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

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/): complete subsection reference.

## Next pages

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_end/)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_start/)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
