---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "IPV6DnsConfig."
xcsh_docs: {"aliases": ["custom storage config storage interface list storage interfaces storage interface ipv6 auto config router dns config"], "body_bytes": 4327, "body_sha256": "sha256:bde36e44843d7fce5f3a91c0e1661a32ef0a987f0911d323a6407db99103c91b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/dns_config/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321", "registry_path": "docs/guides/resources--voltstack_site--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:configured_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "sections": [{"aliases": ["configured list"], "anchor": "section", "description": "IPV6DnsList.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:configured_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list:RequiredObjectAttributes:dns_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:configured_list", "type": "requires"}], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "ipv6_auto_config", "router", "dns_config", "configured_list"], "syntax": "block", "type": "object"}, {"aliases": ["local dns"], "anchor": "section", "description": "IPV6LocalDnsAddress.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,first_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,first_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router:dns_config:local_dns:last_address", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "ipv6_auto_config", "router", "dns_config", "local_dns"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPV6DnsConfig.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [custom_storage_config.storage_interface_list.storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config

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

- [configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/dns_config/configured_list/): complete subsection reference.

- [local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/dns_config/local_dns/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/dns_config/configured_list/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/dns_config/local_dns/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/router/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
