---
page_title: "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful"
subcategory: ""
description: "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5894, "body_sha256": "sha256:307dfd93a8c6a6174c03cd7b16ec2491db58ee03f92839accfa16bafaf9784b1", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "docs/guides/resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [openshift_virtualization](resources--securemesh_site_v2--properties--openshift_virtualization.md)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed.md)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list.md)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

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

- [automatic_from_end](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md): complete subsection reference.

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--fixed_ip_map"></a>

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

- [interface_ip_map](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map.md): complete subsection reference.

## Next pages

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_end.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--automatic_from_start.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map.md)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
