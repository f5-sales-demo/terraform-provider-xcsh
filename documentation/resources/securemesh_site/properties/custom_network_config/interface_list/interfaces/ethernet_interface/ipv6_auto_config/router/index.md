---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router"
subcategory: ""
description: "IPV6AutoConfigRouterType."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router"], "body_bytes": 5014, "body_sha256": "sha256:22693a3a687017dd60bdf28e4f4eea7be193b7f9d65b776f3044e463f1ce5267", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config", "path": "documentation/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3110013313013003-0203321132202300-1031303111312132-2113230221131231-2222133030013203-2232022303100221-3110100300231300-1130331201233201", "registry_path": "docs/guides/resources--securemesh_site--reference--group-002.md", "relationships": [{"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router"], "schema_version": 1, "sections": [{"aliases": ["dns config"], "anchor": "section", "description": "IPV6DnsConfig.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "dns_config"], "syntax": "block", "type": "object"}, {"aliases": ["network prefix"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix", "description": "Exclusive with Network prefix that is used as Prefix information Allowed only /64 prefix length as per RFC 4862.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful"], "anchor": "section", "description": "DHCPIPV6 Stateful Server.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful:RequiredObjectAttributes:dhcp_networks", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:dhcp_networks", "type": "requires"}], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPV6AutoConfigRouterType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix"></a>

### network_prefix property

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

- [stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
