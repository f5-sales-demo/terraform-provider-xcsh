---
page_title: "equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "IPV6DnsConfig."
xcsh_docs: {"aliases": ["equinix not managed node list interface list ipv6 auto config router dns config"], "body_bytes": 2443, "body_sha256": "sha256:7db11cc8c6cb516ed619f7849dfe1145abd6945b8914ea08b66babad806d9ae8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router", "path": "documentation/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0020032302230111-0001313312320112-2321112130303111-3033230322332223-1312002032003321-1130332110330211-1031320023121001-0120131303003032", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "sections": [{"aliases": ["equinix not managed node list interface list ipv6 auto config router dns config configured list"], "anchor": "section", "description": "IPV6DnsList.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["equinix not managed node list interface list ipv6 auto config router dns config local dns"], "anchor": "section", "description": "IPV6LocalDnsAddress.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "IPV6DnsConfig.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
