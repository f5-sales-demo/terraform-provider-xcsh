---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip"
subcategory: ""
description: "Configure Static IP parameters for cluster."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface static ipv6 address cluster static ip"], "body_bytes": 3245, "body_sha256": "sha256:87fa9e664eac82c39d4f9938abd598b505ca64a3b1110384f5e1c33d93bdeb55", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address:cluster_static_ip", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address", "path": "documentation/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/cluster_static_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1000111233132313-3131130113320001-2321010330322313-2302033130112012-3023313121021323-3310022011301000-3322332200132033-1301000123300103", "registry_path": "docs/guides/resources--securemesh_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ipv6_address", "cluster_static_ip"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces ethernet interface static ipv6 address cluster static ip interface ip map"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip--interface_ip_map", "description": "Map of Node to Static IP configuration value, Key:Node, Value:IP Address.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address:cluster_static_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ipv6_address", "cluster_static_ip", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters for cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
