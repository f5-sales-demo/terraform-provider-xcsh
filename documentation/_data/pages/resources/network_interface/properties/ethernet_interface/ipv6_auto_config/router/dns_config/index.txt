---
page_title: "ethernet_interface.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "IPV6DnsConfig."
xcsh_docs: {"aliases": ["ethernet interface ipv6 auto config router dns config"], "body_bytes": 2802, "body_sha256": "sha256:0273d0b7b5b33ab1698b59ae818f306d5f58c760f7b3bc1c65984fd6191b0135", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "path": "documentation/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2111311031130302-0322033320230231-0000120233103131-3221023230032022-1220203201311031-0003011210112120-3212133012003330-3102212122201311", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "sections": [{"aliases": ["configured list"], "anchor": "section", "description": "IPV6DnsList.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.configured_list:RequiredObjectAttributes:dns_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "type": "requires"}], "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config", "configured_list"], "syntax": "block", "type": "object"}, {"aliases": ["local dns"], "anchor": "section", "description": "IPV6LocalDnsAddress.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,first_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}, {"anchor": "schema-ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,first_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:last_address", "type": "conflicts"}], "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config", "local_dns"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPV6DnsConfig.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/)
- [ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/)
- ethernet_interface.ipv6_auto_config.router.dns_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
```

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

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/): complete subsection reference.

- [local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/): complete subsection reference.

## Next pages

- [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/)
- [ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
