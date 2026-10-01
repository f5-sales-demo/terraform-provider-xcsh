---
page_title: "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns"
subcategory: ""
description: "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 4719, "body_sha256": "sha256:578902ed9df88ebacc93994ceb573368ebccf40865d5d1f85835b17bba5d30df", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "docs/guides/resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["azure", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [azure](resources--securemesh_site_v2--properties--azure.md)
- [azure.not_managed](resources--securemesh_site_v2--properties--azure--not_managed.md)
- [azure.not_managed.node_list](resources--securemesh_site_v2--properties--azure--not_managed--node_list.md)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list.md)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

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

<a id="schema-azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address"></a>

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

- [first_address](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--first_address.md): complete subsection reference.

- [last_address](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--last_address.md): complete subsection reference.

## Next pages

- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--first_address.md)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--last_address.md)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--properties--azure--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
