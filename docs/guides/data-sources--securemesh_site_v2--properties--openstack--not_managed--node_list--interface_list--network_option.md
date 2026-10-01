---
page_title: "openstack.not_managed.node_list.interface_list.network_option"
subcategory: ""
description: "openstack.not_managed.node_list.interface_list.network_option for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3017, "body_sha256": "sha256:25df7629983c0764d66201afe5c5645b41c7acc2ded965bd7800d450b8afc804", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:network_option", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:network_option:site_local_inside_network", "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:network_option:site_local_network"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:network_option", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list", "path": "docs/guides/data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--network_option.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "network_option"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/network_option/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openstack.not_managed.node_list.interface_list.network_option for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed.node_list.interface_list.network_option

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [openstack](data-sources--securemesh_site_v2--properties--openstack.md)
- [openstack.not_managed](data-sources--securemesh_site_v2--properties--openstack--not_managed.md)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list.md)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list.md)
- openstack.not_managed.node_list.interface_list.network_option

<a id="section"></a>

Type: `"single"`. Computed.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

## Direct properties

- [site_local_inside_network](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--network_option--site_local_inside_network.md): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--network_option--site_local_network.md): complete subsection reference.

## Next pages

- [openstack.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--network_option--site_local_inside_network.md)
- [openstack.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--network_option--site_local_network.md)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
