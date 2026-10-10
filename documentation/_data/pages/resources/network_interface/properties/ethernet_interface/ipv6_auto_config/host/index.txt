---
page_title: "ethernet_interface.ipv6_auto_config.host"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ethernet interface ipv6 auto config host"], "body_bytes": 1231, "body_sha256": "sha256:7c79fcd92e590d56baadd66db9cd79010c006bdcb96789f41fc43e0b40d08e47", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config:host", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:ipv6_auto_config", "path": "documentation/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/host/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0210013010220110-2103131000100200-0020220111013130-0321133122222231-3001101223231232-0200220202022120-0000132001120013-2303013120031100", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "ipv6_auto_config", "host"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/host/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.ipv6_auto_config.host

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/ipv6_auto_config/)
- ethernet_interface.ipv6_auto_config.host

<a id="section"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

Additional upstream details:

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
host = {}
```

This is an empty object or choice marker. It has no direct properties.
