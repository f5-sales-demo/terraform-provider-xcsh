---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map"
subcategory: ""
description: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2886, "body_sha256": "sha256:80122538212d9887300e1013df323d1f3b2841a45eda76376e444fc634e822b6", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:interface_ip_map", "child_ids": [], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:interface_ip_map", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--interface_ip_map.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "interface_ip_map"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/interface_ip_map/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server.md)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map

<a id="section"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--interface_ip_map--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
