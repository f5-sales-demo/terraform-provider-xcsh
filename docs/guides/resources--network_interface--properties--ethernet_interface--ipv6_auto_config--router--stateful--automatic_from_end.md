---
page_title: "ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end"
subcategory: ""
description: "ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1648, "body_sha256": "sha256:3427387b5c2382e7ce77ba74f99dd055f7ccaa0dfdb6297896d3fbc130c1e191", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_end", "child_ids": [], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:stateful:automatic_from_end", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:router:stateful", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful--automatic_from_end.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "router", "stateful", "automatic_from_end"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_end/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [ethernet_interface.ipv6_auto_config](resources--network_interface--properties--ethernet_interface--ipv6_auto_config.md)
- [ethernet_interface.ipv6_auto_config.router](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router.md)
- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md)
- ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ethernet_interface.ipv6_auto_config.router.stateful](resources--network_interface--properties--ethernet_interface--ipv6_auto_config--router--stateful.md)
- [xcsh_network_interface](../resources/network_interface.md)
