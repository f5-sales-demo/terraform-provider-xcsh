---
page_title: "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools"
subcategory: ""
description: "List of non overlapping IP address ranges."
xcsh_docs: {"aliases": ["kvm not managed node list interface list ipv6 auto config router stateful dhcp networks pools"], "body_bytes": 5659, "body_sha256": "sha256:5935a2f74ead0a3e5162ed8b45b239f0c304f7189a5d1e80196f804d8aa7a066", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "path": "documentation/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3032031320322210-0322031100313310-3211130211232131-3301030321102023-1220112100023201-3112313120232333-2321112233022223-3122202310323212", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools"], "schema_version": 1, "sections": [{"aliases": ["kvm not managed node list interface list ipv6 auto config router stateful dhcp networks pools end ip"], "anchor": "schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip", "description": "Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on network prefix.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools", "end_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["kvm not managed node list interface list ipv6 auto config router stateful dhcp networks pools start ip"], "anchor": "schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip", "description": "Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on network prefix. 2001::1 with prefix length of 64, start offset is 5.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools", "start_ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of non overlapping IP address ranges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [kvm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/)
- [kvm.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/)
- [kvm.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/)
- [kvm.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="section"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip"></a>

### end_ip property

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip"></a>

### start_ip property

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

## Next pages

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
