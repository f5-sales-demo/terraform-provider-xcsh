---
page_title: "openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip"
subcategory: ""
description: "openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2790, "body_sha256": "sha256:c9d7a99185e6ab1bfbe2a77acf0504d12aacba01170799ee19e58b8be6472dba", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:static_ipv6_address", "path": "docs/guides/resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [openshift_virtualization](resources--securemesh_site_v2--properties--openshift_virtualization.md)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed.md)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list.md)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list.md)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--static_ipv6_address.md)
- openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
cluster_static_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-openshift_virtualization--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Optional.

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

- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--properties--openshift_virtualization--not_managed--node_list--interface_list--static_ipv6_address.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
