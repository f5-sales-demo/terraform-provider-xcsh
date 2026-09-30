---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful"
subcategory: ""
description: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 5734, "body_sha256": "sha256:788fa82d0b295880d47429a45ae55f3c85435b51f2d5773722ee3842aed1007b", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:stateful", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:stateful:automatic_from_end", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:stateful:automatic_from_start", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:stateful:dhcp_networks", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:stateful:interface_ip_map"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:stateful", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "ipv6_auto_config", "router", "stateful"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/stateful/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router.md)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful

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

- [automatic_from_end](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--dhcp_networks.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--fixed_ip_map"></a>

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

- [interface_ip_map](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--interface_ip_map.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--automatic_from_end.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--automatic_from_start.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--dhcp_networks.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--stateful--interface_ip_map.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
