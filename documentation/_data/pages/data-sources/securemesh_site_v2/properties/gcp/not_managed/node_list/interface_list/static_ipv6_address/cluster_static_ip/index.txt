---
page_title: "gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip"
subcategory: ""
description: "Configure Static IP parameters for cluster."
xcsh_docs: {"aliases": ["gcp not managed node list interface list static ipv6 address cluster static ip"], "body_bytes": 2823, "body_sha256": "sha256:eac3de7bd909fd2cf58a125233ce1f2e1b83911c5a02f7188247648975bc093b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:static_ipv6_address", "path": "documentation/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0331332131031021-3020331333322311-0120103103321133-1323002020232333-2233103203203200-3220210210002011-1000330020321112-2222030002310231", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip"], "schema_version": 1, "sections": [{"aliases": ["interface ip map"], "anchor": "schema-gcp--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip--interface_ip_map", "description": "Map of Node to Static IP configuration value, Key:Node, Value:IP Address.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure Static IP parameters for cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/)
- [gcp.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/)
- [gcp.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/)
- [gcp.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/static_ipv6_address/)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

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

<a id="schema-gcp--not_managed--node_list--interface_list--static_ipv6_address--cluster_static_ip--interface_ip_map"></a>

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

- [gcp.not_managed.node_list.interface_list.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/static_ipv6_address/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
