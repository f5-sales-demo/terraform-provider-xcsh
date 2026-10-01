---
page_title: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map"
subcategory: ""
description: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3160, "body_sha256": "sha256:d4613f57a11cbede9fccdc969147d0c35a63f03d04033c3acc5f1e4897f5de7c", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "path": "docs/guides/resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "interface_ip_map"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [nutanix](resources--securemesh_site_v2--properties--nutanix.md)
- [nutanix.not_managed](resources--securemesh_site_v2--properties--nutanix--not_managed.md)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list.md)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Map of Interface IPv6 assignments per node.

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

<a id="schema-nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

## Next pages

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
