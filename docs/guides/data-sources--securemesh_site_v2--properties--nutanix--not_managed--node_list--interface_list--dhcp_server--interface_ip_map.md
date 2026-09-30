---
page_title: "nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map"
subcategory: ""
description: "nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2451, "body_sha256": "sha256:1c6ad0765fe7b64c40337d0d2cce6c335fad998eb9f8813b25b0917a1ce12164", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:dhcp_server", "path": "docs/guides/data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--dhcp_server--interface_ip_map.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "dhcp_server", "interface_ip_map"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [nutanix](data-sources--securemesh_site_v2--properties--nutanix.md)
- [nutanix.not_managed](data-sources--securemesh_site_v2--properties--nutanix--not_managed.md)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list.md)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list.md)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--dhcp_server.md)
- nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

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

<a id="schema-nutanix--not_managed--node_list--interface_list--dhcp_server--interface_ip_map--interface_ip_map"></a>

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

- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--dhcp_server.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
