---
page_title: "ethernet_interface.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config.router.dns_config for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 2260, "body_sha256": "sha256:f655f05ba539ccad6faca170afaebe2e06c41c52750a92a4518a70ffd5a5939c", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config.router.dns_config for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md)
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

- [configured_list](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md): complete subsection reference.

- [local_dns](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md): complete subsection reference.

## Next pages

- [ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md)
- [xcsh_network_interface](../resources/network_interface.md)
