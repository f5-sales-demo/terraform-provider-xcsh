---
page_title: "gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip"
subcategory: ""
description: "gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2376, "body_sha256": "sha256:8fe453ec8932d69ec2abe980b2f3149adeeeb2631201266f8d09e550e49bf9c2", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:static_ipv6_address", "path": "docs/guides/data-sources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [gcp](data-sources--securemesh_site_v2--properties--gcp.md)
- [gcp.not_managed](data-sources--securemesh_site_v2--properties--gcp--not_managed.md)
- [gcp.not_managed.node_list](data-sources--securemesh_site_v2--properties--gcp--not_managed--node_list.md)
- [gcp.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list.md)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--static_ipv6_address.md)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

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

<a id="schema-gcp--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip--interface_ip_map"></a>

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

- [gcp.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--properties--gcp--not_managed--node_list--interface_list--static_ipv6_address.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
