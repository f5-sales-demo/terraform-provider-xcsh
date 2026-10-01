---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 4648, "body_sha256": "sha256:cb1cb9eee28e158c8c38d8dfa6216690a267eb5e88fe1c7d9834d4bb1f66b651", "canonical_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:first_address", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:last_address"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config", "path": "docs/guides/data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "dns_config", "local_dns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
- [Property reference](data-sources--securemesh_site--reference.md)
- [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](data-sources--securemesh_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

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

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address"></a>

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

- [first_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--first_address.md): complete subsection reference.

- [last_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--last_address.md): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--first_address.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--last_address.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
