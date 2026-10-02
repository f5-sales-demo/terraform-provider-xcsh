---
page_title: "azure.not_managed.node_list.interface_list.ipv6_auto_config.router"
subcategory: ""
description: "IPV6AutoConfigRouterType."
xcsh_docs: {"aliases": ["azure not managed node list interface list ipv6 auto config router"], "body_bytes": 4589, "body_sha256": "sha256:19adf4f6a0dcb101c24b3ea86cf31b65a99a50c2328dd0eb03f7b940fd6f3957", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config", "path": "documentation/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/router/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3121312031011313-3300231331310200-1100200320002310-3030100321120121-0030110021331111-2100012231002202-1003132031130311-2010201110230021", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-005.md", "relationships": [{"anchor": "schema-azure--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router:ConflictingObjectAttributes:network_prefix,stateful", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["azure", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "schema_version": 1, "sections": [{"aliases": ["dns config"], "anchor": "section", "description": "IPV6DnsConfig.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config:ConflictingObjectAttributes:configured_list,local_dns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}], "schema_path": ["azure", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config"], "syntax": "block", "type": "object"}, {"aliases": ["network prefix"], "anchor": "schema-azure--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix", "description": "Exclusive with Network prefix that is used as Prefix information Allowed only /64 prefix length as per RFC 4862.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful"], "anchor": "section", "description": "DHCPIPV6 Stateful Server.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful:RequiredObjectAttributes:dhcp_networks", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "type": "requires"}], "schema_path": ["azure", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IPV6AutoConfigRouterType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure.not_managed.node_list.interface_list.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [azure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/)
- [azure.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/)
- [azure.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/)
- [azure.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router

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

- [dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/): complete subsection reference.

<a id="schema-azure--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix"></a>

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

- [stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/): complete subsection reference.

## Next pages

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
