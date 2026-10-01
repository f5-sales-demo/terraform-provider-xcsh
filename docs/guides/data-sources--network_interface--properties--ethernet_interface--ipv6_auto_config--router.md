---
page_title: "ethernet_interface.ipv6_auto_config.router"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config.router for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 2910, "body_sha256": "sha256:7118602bf0e1b121cc72487884e02dc3a34718e5a0f3a773b76266c28c610b90", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:stateful"], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config", "path": "docs/guides/data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config.router for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- ethernet_interface.ipv6_auto_config.router

<a id="section"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

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

## Direct properties

- [dns_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md): complete subsection reference.

<a id="schema-ethernet_interface--ipv6_auto_config--router--network_prefix"></a>

### network_prefix property

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

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

- [stateful](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md): complete subsection reference.

## Next pages

- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- [ethernet_interface.ipv6_auto_config.router.stateful](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
