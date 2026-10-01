---
page_title: "vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2712, "body_sha256": "sha256:4931ca645fc51712f631a3021d09446c55623eaa32f10b5b2961fba755ef901e", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "docs/guides/data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [vmware](data-sources--securemesh_site_v2--properties--vmware.md)
- [vmware.not_managed](data-sources--securemesh_site_v2--properties--vmware--not_managed.md)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list.md)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list.md)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="section"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

## Direct properties

- [configured_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md): complete subsection reference.

## Next pages

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
