---
page_title: "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns"
subcategory: ""
description: "IPV6LocalDnsAddress."
xcsh_docs: {"aliases": ["openstack not managed node list interface list ipv6 auto config router dns config local dns"], "body_bytes": 5527, "body_sha256": "sha256:b664380c6617ca7d79c1e86334f1a0042d36cb7279126f3761c8eddc9e9dc0e3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "documentation/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0010123202020021-3232021123110202-3012220233303110-1301320101300200-1120320103000333-1123112210030331-2212202103301313-1012312131301132", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-016.md", "relationships": [{"anchor": "schema-openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address", "enforcement": "provider-schema", "group": "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,first_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}, {"anchor": "schema-openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address", "enforcement": "provider-schema", "group": "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,first_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:configured_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns:ConflictingObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "schema_version": 1, "sections": [{"aliases": ["configured address"], "anchor": "schema-openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address", "description": "Exclusive with Configured address from the network prefix is chosen as DNS server.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns", "configured_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["first address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns", "first_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["last address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:openstack:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openstack", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns", "last_address"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IPV6LocalDnsAddress.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [openstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/)
- [openstack.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/)
- [openstack.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/)
- [openstack.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-openstack--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address"></a>

### configured_address property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/first_address/): complete subsection reference.

- [last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/last_address/): complete subsection reference.

## Next pages

- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/first_address/)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/last_address/)
- [openstack.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/openstack/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
