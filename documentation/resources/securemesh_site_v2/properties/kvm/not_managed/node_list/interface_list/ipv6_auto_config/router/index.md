---
page_title: "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router"
subcategory: ""
description: "IPV6AutoConfigRouterType."
xcsh_docs: {"aliases": ["kvm not managed node list interface list ipv6 auto config router"], "body_bytes": 3603, "body_sha256": "sha256:fcdae594a83da9130fc3a44b9fc78b098e7a7f0a4d7efda4c9d0edfc51133445", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config", "path": "documentation/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1310323013123101-0320203111002303-2002231332302132-3311333233122010-0122001311022212-0211001231113131-0132231210001011-1300221211233301", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-010.md", "relationships": [{"anchor": "schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "schema_version": 1, "sections": [{"aliases": ["kvm not managed node list interface list ipv6 auto config router dns config"], "anchor": "section", "description": "IPV6DnsConfig.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "syntax": "block", "type": "object"}, {"aliases": ["kvm not managed node list interface list ipv6 auto config router network prefix"], "anchor": "schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix", "description": "Exclusive with Network prefix that is used as Prefix information Allowed only /64 prefix length as per RFC 4862.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["kvm not managed node list interface list ipv6 auto config router stateful"], "anchor": "section", "description": "DHCPIPV6 Stateful Server.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:RequiredObjectAttributes:dhcp_networks", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "type": "requires"}], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "IPV6AutoConfigRouterType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kvm.not_managed.node_list.interface_list.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [kvm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/)
- [kvm.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/)
- [kvm.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/): complete subsection reference.

<a id="schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix"></a>

### network_prefix property

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/): complete subsection reference.
