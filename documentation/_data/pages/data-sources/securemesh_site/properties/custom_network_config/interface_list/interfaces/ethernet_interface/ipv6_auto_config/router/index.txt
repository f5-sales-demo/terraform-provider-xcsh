---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router"
subcategory: ""
description: "IPV6AutoConfigRouterType."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router"], "body_bytes": 4627, "body_sha256": "sha256:0eb86858d1bb2eaf88b61db98923c8dbdd5b988907dbad3150b0585923342745", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1001033021011103-2110300330002020-1210301200322110-2321120320011300-1331320231221231-0103323010021201-1010030010312310-3302310201012010", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router"], "schema_version": 1, "sections": [{"aliases": ["dns config"], "anchor": "section", "description": "IPV6DnsConfig.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "dns_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["network prefix"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix", "description": "Exclusive with Network prefix that is used as Prefix information Allowed only /64 prefix length as per RFC 4862.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful"], "anchor": "section", "description": "DHCPIPV6 Stateful Server.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IPV6AutoConfigRouterType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router

<a id="section"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

## Direct properties

- [dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix"></a>

### network_prefix property

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
