---
page_title: "ethernet_interface.ipv6_auto_config"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1473, "body_sha256": "sha256:6146872d2228246ba39d0c278c4713309216897658a4c2365978a995f1d112b0", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config:router"], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:ipv6_auto_config", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "path": "docs/guides/data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- ethernet_interface.ipv6_auto_config

<a id="section"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

## Direct properties

- [host](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--host.md): complete subsection reference.

- [router](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md): complete subsection reference.

## Next pages

- [ethernet_interface.ipv6_auto_config.host](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--host.md)
- [ethernet_interface.ipv6_auto_config.router](data-sources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
