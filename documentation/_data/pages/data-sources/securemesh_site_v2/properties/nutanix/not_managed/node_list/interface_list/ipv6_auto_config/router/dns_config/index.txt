---
page_title: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "IPV6DnsConfig."
xcsh_docs: {"aliases": ["nutanix not managed node list interface list ipv6 auto config router dns config"], "body_bytes": 3402, "body_sha256": "sha256:7c09c8d8e579034a2b23dcf4e54827ab48e6da2fa0c710ce95b24934151d38d0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "documentation/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "sections": [{"aliases": ["configured list"], "anchor": "section", "description": "IPV6DnsList.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["local dns"], "anchor": "section", "description": "IPV6LocalDnsAddress.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPV6DnsConfig.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [nutanix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/)
- [nutanix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/)
- [nutanix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/)
- [nutanix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

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

- [configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/): complete subsection reference.

- [local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/): complete subsection reference.

## Next pages

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
