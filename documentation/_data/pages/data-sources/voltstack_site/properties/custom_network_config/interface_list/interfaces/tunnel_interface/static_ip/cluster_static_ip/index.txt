---
page_title: "custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip"
subcategory: ""
description: "Configure Static IP parameters for cluster."
xcsh_docs: {"aliases": ["custom network config interface list interfaces tunnel interface static ip cluster static ip"], "body_bytes": 3047, "body_sha256": "sha256:756e27552c9ab1320c4151e93c86657329ea4c9f77fd628ce483b46e27b60d54", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:cluster_static_ip", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/cluster_static_ip/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0002001312201211-0031010301102132-0000302231021303-2322302120120103-2302002303230320-0010030113002120-0332120221203022-1212230013012130", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface", "static_ip", "cluster_static_ip"], "schema_version": 1, "sections": [{"aliases": ["interface ip map"], "anchor": "schema-custom_network_config--interface_list--interfaces--tunnel_interface--static_ip--cluster_static_ip--interface_ip_map", "description": "Map of Node to Static IP configuration value, Key:Node, Value:IP Address.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list:interfaces:tunnel_interface:static_ip:cluster_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "tunnel_interface", "static_ip", "cluster_static_ip", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters for cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-custom_network_config--interface_list--interfaces--tunnel_interface--static_ip--cluster_static_ip--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Computed.

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

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/interface_list/interfaces/tunnel_interface/static_ip/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
