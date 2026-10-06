---
page_title: "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns"
subcategory: ""
description: "IPV6LocalDnsAddress."
xcsh_docs: {"aliases": ["kvm not managed node list interface list ipv6 auto config router dns config local dns"], "body_bytes": 3652, "body_sha256": "sha256:1163ae6abfd631e64feea05c7662f729a57e1bfb205694fbb06dda7a611df0a5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "documentation/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3013321020113100-2323323031022311-3022022223211201-2213033231100323-2120230310013232-1310102122000002-2302230121313031-1121233020312221", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "schema_version": 1, "sections": [{"aliases": ["kvm not managed node list interface list ipv6 auto config router dns config local dns configured address"], "anchor": "schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address", "description": "Exclusive with Configured address from the network prefix is chosen as DNS server.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns", "configured_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["kvm not managed node list interface list ipv6 auto config router dns config local dns first address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns", "first_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["kvm not managed node list interface list ipv6 auto config router dns config local dns last address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns", "last_address"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "IPV6LocalDnsAddress.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [kvm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/)
- [kvm.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/)
- [kvm.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="section"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

## Direct properties

<a id="schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address"></a>

### configured_address property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/first_address/): complete subsection reference.

- [last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/last_address/): complete subsection reference.
