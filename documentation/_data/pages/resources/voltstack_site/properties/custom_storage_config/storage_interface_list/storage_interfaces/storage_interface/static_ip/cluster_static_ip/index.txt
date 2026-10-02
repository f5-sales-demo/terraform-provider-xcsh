---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip"
subcategory: ""
description: "Configure Static IP parameters for cluster."
xcsh_docs: {"aliases": ["custom storage config storage interface list storage interfaces storage interface static ip cluster static ip"], "body_bytes": 3347, "body_sha256": "sha256:546fb9990f35171b1d22de0bbbb9ffb3aeece45a6bdf014646777c10775abee4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:cluster_static_ip", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/cluster_static_ip/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3310011330130321-1101330002000001-2002102232030320-1112201101230231-2011313322002022-1212322011112133-3131333102011230-1012332122233321", "registry_path": "docs/guides/resources--voltstack_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ip", "cluster_static_ip"], "schema_version": 1, "sections": [{"aliases": ["interface ip map"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--cluster_static_ip--interface_ip_map", "description": "Map of Node to Static IP configuration value, Key:Node, Value:IP Address.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:cluster_static_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ip", "cluster_static_ip", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure Static IP parameters for cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [custom_storage_config.storage_interface_list.storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip

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

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip--cluster_static_ip--interface_ip_map"></a>

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

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
