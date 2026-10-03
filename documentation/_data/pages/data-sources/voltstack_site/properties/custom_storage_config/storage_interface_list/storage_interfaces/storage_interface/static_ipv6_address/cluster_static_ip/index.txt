---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.cluster_static_ip"
subcategory: ""
description: "Configure Static IP parameters for cluster."
xcsh_docs: {"aliases": ["custom storage config storage interface list storage interfaces storage interface static ipv6 address cluster static ip"], "body_bytes": 3318, "body_sha256": "sha256:2fed76301416f1d0da9414b6a3e212a29ab8f42fd6619c284d1bf51802e3eb80", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address:cluster_static_ip", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ipv6_address/cluster_static_ip/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0330003113311232-3031101302333123-2323020103310230-3213203211130213-0303021320113113-0132130232332301-0221310311113021-1110202332200323", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ipv6_address", "cluster_static_ip"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage interface list storage interfaces storage interface static ipv6 address cluster static ip interface ip map"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ipv6_address--cluster_static_ip--interface_ip_map", "description": "Map of Node to Static IP configuration value, Key:Node, Value:IP Address.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address:cluster_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ipv6_address", "cluster_static_ip", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ipv6_address/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters for cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.cluster_static_ip

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [custom_storage_config.storage_interface_list.storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ipv6_address/)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.cluster_static_ip

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

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ipv6_address--cluster_static_ip--interface_ip_map"></a>

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

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ipv6_address/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
