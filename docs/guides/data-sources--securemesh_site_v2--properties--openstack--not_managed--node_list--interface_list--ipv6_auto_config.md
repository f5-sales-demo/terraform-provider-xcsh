---
page_title: "openstack.not_managed.node_list.interface_list.ipv6_auto_config"
subcategory: ""
description: "openstack.not_managed.node_list.interface_list.ipv6_auto_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2032, "body_sha256": "sha256:cd98c1edea03b32b5a6d242bd17b17684c09b85fc404fee14e4e3935c97059da", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:host", "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list", "path": "docs/guides/data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ipv6_auto_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "openstack.not_managed.node_list.interface_list.ipv6_auto_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# openstack.not_managed.node_list.interface_list.ipv6_auto_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [openstack](data-sources--securemesh_site_v2--properties--openstack.md)
- [openstack.not_managed](data-sources--securemesh_site_v2--properties--openstack--not_managed.md)
- [openstack.not_managed.node_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list.md)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list.md)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config

<a id="section"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

## Direct properties

- [host](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--host.md): complete subsection reference.

- [router](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--router.md): complete subsection reference.

## Next pages

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--host.md)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [openstack.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--openstack--not_managed--node_list--interface_list.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
