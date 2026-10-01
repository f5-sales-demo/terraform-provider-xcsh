---
page_title: "custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2600, "body_sha256": "sha256:466c5a3f76af210c6f7e30b061a5c98344bf1a3bc0a22de722ac674a18bb9c65", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:cluster_static_ip", "child_ids": [], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:cluster_static_ip", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip", "path": "docs/guides/data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--static_ip--cluster_static_ip.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface", "static_ip", "cluster_static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](data-sources--voltstack_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface.md)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--static_ip.md)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-custom_network_config--interface_list--interfaces--tunnel_interface--static_ip--cluster_static_ip--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

## Next pages

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--properties--custom_network_config--interface_list--interfaces--tunnel_interface--static_ip.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
