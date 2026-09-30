---
page_title: "openstack.not_managed.node_list.interface_list.static_ipv6_address"
subcategory: ""
description: "openstack.not_managed.node_list.interface_list.static_ipv6_address for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2255, "body_sha256": "sha256:c868efed03aa2dca874384cfd77b1778e637f7bcde28b83fe2a9b43698179de2", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:static_ipv6_address", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list", "path": "docs/guides/data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--static_ipv6_address.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "static_ipv6_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openstack.not_managed.node_list.interface_list.static_ipv6_address for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# openstack.not_managed.node_list.interface_list.static_ipv6_address

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [openstack](data-sources--securemesh_site_v2--properties--openstack.md)
- [openstack.not_managed](data-sources--securemesh_site_v2--properties--openstack--not_managed.md)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list.md)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list.md)
- openstack.not_managed.node_list.interface_list.static_ipv6_address

<a id="section"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

## Direct properties

- [cluster_static_ip](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip.md): complete subsection reference.

## Next pages

- [openstack.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip.md)
- [openstack.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip.md)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
