---
page_title: "baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful"
subcategory: ""
description: "baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5459, "body_sha256": "sha256:c4a64ca43a4cd64e55b3dfb4deee6ad1b3aa175fd3c4b1d6fdc09de5c7d23dbd", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "docs/guides/resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [baremetal](resources--securemesh_site_v2--properties--baremetal.md)
- [baremetal.not_managed](resources--securemesh_site_v2--properties--baremetal--not_managed.md)
- [baremetal.not_managed.node_list](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list.md)
- [baremetal.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list.md)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

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
stateful {
  # Configure direct properties listed below.
}
```

## Direct properties

- [automatic_from_end](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md): complete subsection reference.

<a id="schema-baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--fixed_ip_map"></a>

### fixed_ip_map property

Type: `["map", "string"]`. Optional.

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

- [interface_ip_map](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map.md): complete subsection reference.

## Next pages

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_end.md)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_start.md)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map.md)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
