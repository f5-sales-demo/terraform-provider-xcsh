---
page_title: "equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns"
subcategory: ""
description: "equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 4084, "body_sha256": "sha256:85a0d787321ab709594407843cb0fc98f54ba664fb3a343cde9deeeb145bba8f", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:first_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns:last_address"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:local_dns", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "docs/guides/data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "local_dns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/local_dns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [equinix](data-sources--securemesh_site_v2--properties--equinix.md)
- [equinix.not_managed](data-sources--securemesh_site_v2--properties--equinix--not_managed.md)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list.md)
- [equinix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list.md)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

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

<a id="schema-equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--configured_address"></a>

### configured_address property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

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

- [first_address](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--first_address.md): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--last_address.md): complete subsection reference.

## Next pages

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--first_address.md)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--local_dns--last_address.md)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
