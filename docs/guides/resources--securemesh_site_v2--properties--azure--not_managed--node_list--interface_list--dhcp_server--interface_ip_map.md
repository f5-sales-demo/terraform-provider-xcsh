---
page_title: "azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map"
subcategory: ""
description: "azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2519, "body_sha256": "sha256:36efb5aaa2e2120a10e90fe0ab756f20d50d3c8d296370855124252c62860d3c", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:dhcp_server:interface_ip_map", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:dhcp_server", "path": "docs/guides/resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--dhcp_server--interface_ip_map.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure", "not_managed", "node_list", "interface_list", "dhcp_server", "interface_ip_map"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/dhcp_server/interface_ip_map/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [azure](resources--securemesh_site_v2--properties--azure.md)
- [azure.not_managed](resources--securemesh_site_v2--properties--azure--not_managed.md)
- [azure.not_managed.node_list](resources--securemesh_site_v2--properties--azure--not_managed--node_list.md)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list.md)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--dhcp_server.md)
- azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-azure--not_managed--node_list--interface_list--dhcp_server--interface_ip_map--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Optional.

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

- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--dhcp_server.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
