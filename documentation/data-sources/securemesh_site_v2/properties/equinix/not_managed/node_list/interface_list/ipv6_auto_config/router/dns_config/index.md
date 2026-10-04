---
page_title: "equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "IPV6DnsConfig."
xcsh_docs: {"aliases": ["equinix not managed node list interface list ipv6 auto config router dns config"], "body_bytes": 3402, "body_sha256": "sha256:de84ff262096b86c73390eb51ed5d566d9a1860abe26c06c4fea0d5ac996e1db", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "documentation/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "sections": [{"aliases": ["equinix not managed node list interface list ipv6 auto config router dns config configured list"], "anchor": "section", "description": "IPV6DnsList.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["equinix not managed node list interface list ipv6 auto config router dns config local dns"], "anchor": "section", "description": "IPV6LocalDnsAddress.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "IPV6DnsConfig.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [equinix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/)
- [equinix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/)
- [equinix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/)
- [equinix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

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

- [configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/): complete subsection reference.

- [local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/): complete subsection reference.

## Next pages

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
