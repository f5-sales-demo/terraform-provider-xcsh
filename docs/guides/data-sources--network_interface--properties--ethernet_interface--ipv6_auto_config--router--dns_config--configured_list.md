---
page_title: "ethernet_interface.ipv6_auto_config.router.dns_config.configured_list"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config.router.dns_config.configured_list for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 2651, "body_sha256": "sha256:d5804d0783a2c62c2f2b011d8d102c9b5bebcad9cc0b27e95dcf10f525ac4afe", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config:configured_list", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:dns_config", "path": "docs/guides/data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "dns_config", "configured_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config.router.dns_config.configured_list for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- [ethernet_interface.ipv6_auto_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md)
- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="section"></a>

Type: `"single"`. Computed.

IPV6DnsList.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list"></a>

### dns_list property

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--dns_config.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
