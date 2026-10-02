---
page_title: "aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map"
subcategory: ""
description: "Map of Interface IPv6 assignments per node."
xcsh_docs: {"aliases": ["aws not managed node list interface list ipv6 auto config router stateful interface ip map"], "body_bytes": 3620, "body_sha256": "sha256:8996c4a7729dd8e3f381d6a1f40e20aef7e6ce74a244ec7468bc6cbadec396e7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "path": "documentation/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1330001013210123-1120220010020113-3202320223302023-2231201331210030-0101122001201110-3230010032333130-0003121112232331-0031022010210210", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "interface_ip_map"], "schema_version": 1, "sections": [{"aliases": ["interface ip map"], "anchor": "schema-aws--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map", "description": "Map of Site:Node to IPv6 address.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:aws:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:interface_ip_map", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "interface_ip_map", "interface_ip_map"], "syntax": "attribute", "type": "map"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/interface_ip_map/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Map of Interface IPv6 assignments per node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/)
- [aws.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/)
- [aws.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/)
- [aws.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map"></a>

### interface_ip_map property

Type: `["map", "string"]`. Optional.

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

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/aws/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
