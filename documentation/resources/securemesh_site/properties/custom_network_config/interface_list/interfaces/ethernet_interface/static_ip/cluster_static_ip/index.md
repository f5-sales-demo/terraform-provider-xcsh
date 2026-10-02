---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip"
subcategory: ""
description: "Configure Static IP parameters for cluster."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface static ip cluster static ip"], "body_bytes": 3175, "body_sha256": "sha256:791e20fb6c4d448ff151bf989b1625934287eb751c53b656eabf4a85cb111025", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:cluster_static_ip", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip", "path": "documentation/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3032300233030002-3113011321302333-0320321211303111-3022212002201020-1332202031223020-0221120113311023-2220303003123223-2030200202320031", "registry_path": "docs/guides/resources--securemesh_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip", "cluster_static_ip"], "schema_version": 1, "sections": [{"aliases": ["interface ip map"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip--interface_ip_map", "description": "Map of Node to Static IP configuration value, Key:Node, Value:IP Address.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip:cluster_static_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ip", "cluster_static_ip", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters for cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip

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

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip--interface_ip_map"></a>

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

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
