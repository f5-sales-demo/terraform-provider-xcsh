---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3272, "body_sha256": "sha256:98e186914ce7cd405a056aeb9e5014574425df92cdbf93aa1d9533c7baaa6fbc", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:dns_config", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "dns_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](resources--voltstack_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router.md)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config

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

- [configured_list](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md): complete subsection reference.

- [local_dns](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
