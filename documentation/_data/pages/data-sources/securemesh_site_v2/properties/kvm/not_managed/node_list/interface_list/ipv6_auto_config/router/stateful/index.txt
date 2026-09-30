---
page_title: "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful"
subcategory: ""
description: "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5545, "body_sha256": "sha256:2b4814e052f872bae8f067e902c4f94d3faadd5e079d781bbc57b2a0df463277", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "documentation/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [kvm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/)
- [kvm.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/)
- [kvm.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

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

- [automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_end/): complete subsection reference.

- [automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_start/): complete subsection reference.

- [dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/): complete subsection reference.

<a id="schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--fixed_ip_map"></a>

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

- [interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/): complete subsection reference.

## Next pages

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_end/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/automatic_from_start/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
