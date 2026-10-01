---
page_title: "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 3959, "body_sha256": "sha256:eef552f9a856d3dc7261f00a83c5665e4bfe967533679cb61c9862feae0f96d4", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:first_address", "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns:last_address"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:local_dns", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config", "local_dns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config.router.dns_config.local_dns for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

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

<a id="schema-ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address"></a>

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

- [first_address](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--first_address.md): complete subsection reference.

- [last_address](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--last_address.md): complete subsection reference.

## Next pages

- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--first_address.md)
- [ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--last_address.md)
- [ethernet_interface.ipv6_auto_config.router.dns_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- [xcsh_network_interface](../resources/network_interface.md)
