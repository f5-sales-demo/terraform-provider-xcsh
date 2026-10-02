---
page_title: "baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map"
subcategory: ""
description: "Map of Interface IPv6 assignments per node."
xcsh_docs: {"aliases": ["baremetal not managed node list interface list ipv6 auto config router stateful interface ip map"], "body_bytes": 3642, "body_sha256": "sha256:74d49264d4c346d1815b811646a304642d18e5a89424661ad29a9d103aa86af6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "path": "documentation/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2101030003110223-3122331020132001-2122001333121022-0021333132221333-1300001030021301-0101323013232233-0200020332112311-2220220111020311", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "interface_ip_map"], "schema_version": 1, "sections": [{"aliases": ["interface ip map"], "anchor": "schema-baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map", "description": "Map of Site:Node to IPv6 address.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "interface_ip_map", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Map of Interface IPv6 assignments per node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [baremetal](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/)
- [baremetal.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/)
- [baremetal.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/)
- [baremetal.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="section"></a>

Type: `"single"`. Computed.

Map of Interface IPv6 assignments per node.

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

<a id="schema-baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

## Next pages

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
