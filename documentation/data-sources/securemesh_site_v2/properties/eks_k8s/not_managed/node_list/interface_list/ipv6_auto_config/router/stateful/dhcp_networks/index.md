---
page_title: "eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks"
subcategory: ""
description: "List of networks from which DHCP server can allocate IP addresses."
xcsh_docs: {"aliases": ["eks k8s not managed node list interface list ipv6 auto config router stateful dhcp networks"], "body_bytes": 6151, "body_sha256": "sha256:b9d050cd3c9229d1e323ffb9ab689b7911766ca0b8f15ad3807706f4a58763ed", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "path": "documentation/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2010132122000111-0200232110303303-0330111320313312-2330201212200103-0203310320332233-0121310313012122-2322233210313002-2203200230202212", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks"], "schema_version": 1, "sections": [{"aliases": ["eks k8s not managed node list interface list ipv6 auto config router stateful dhcp networks network prefix"], "anchor": "schema-eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix", "description": "Exclusive with Network Prefix to be used for IPv6 address auto configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["eks k8s not managed node list interface list ipv6 auto config router stateful dhcp networks pool settings"], "anchor": "schema-eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings", "description": "Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pool_settings"], "syntax": "attribute", "type": "string"}, {"aliases": ["eks k8s not managed node list interface list ipv6 auto config router stateful dhcp networks pools"], "anchor": "section", "description": "List of non overlapping IP address ranges.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of networks from which DHCP server can allocate IP addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [eks_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/)
- [eks_k8s.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/)
- [eks_k8s.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/)
- [eks_k8s.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="section"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Computed.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="schema-eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings"></a>

### pool_settings property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/): complete subsection reference.

## Next pages

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
