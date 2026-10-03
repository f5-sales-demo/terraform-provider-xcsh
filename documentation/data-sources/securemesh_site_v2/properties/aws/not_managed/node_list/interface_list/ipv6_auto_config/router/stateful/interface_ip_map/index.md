---
page_title: "aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map"
subcategory: ""
description: "Map of Interface IPv6 assignments per node."
xcsh_docs: {"aliases": ["aws not managed node list interface list ipv6 auto config router stateful interface ip map"], "body_bytes": 3528, "body_sha256": "sha256:af4ec7a15e5bf3d3f919b0a33f8577fdeef429a0cf70df37ef647be899ec1022", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "path": "documentation/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3033320221130332-2102030110033113-3021110323100200-1102113302220030-1113010313023020-3333031113102321-1123113100031321-3133223222002210", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "interface_ip_map"], "schema_version": 1, "sections": [{"aliases": ["aws not managed node list interface list ipv6 auto config router stateful interface ip map interface ip map"], "anchor": "schema-aws--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map", "description": "Map of Site:Node to IPv6 address.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "interface_ip_map", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Map of Interface IPv6 assignments per node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/)
- [aws.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/not_managed/)
- [aws.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/)
- [aws.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

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

<a id="schema-aws--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map"></a>

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

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
