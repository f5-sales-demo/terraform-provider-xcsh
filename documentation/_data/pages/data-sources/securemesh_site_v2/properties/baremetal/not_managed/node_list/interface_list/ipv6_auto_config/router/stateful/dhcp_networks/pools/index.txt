---
page_title: "baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools"
subcategory: ""
description: "List of non overlapping IP address ranges."
xcsh_docs: {"aliases": ["baremetal not managed node list interface list ipv6 auto config router stateful dhcp networks pools"], "body_bytes": 5791, "body_sha256": "sha256:b0e4c433c8743c09c1800146a4a83bc83bab30097faee565f38e1d807e0e3aa8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "path": "documentation/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0122223303102223-1111323101132330-1203300222021330-1133200123303113-3131023021220233-2300202300012213-3233111231100022-2313122020301220", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools"], "schema_version": 1, "sections": [{"aliases": ["end ip"], "anchor": "schema-baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip", "description": "Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on network prefix.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools", "end_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["start ip"], "anchor": "schema-baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip", "description": "Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on network prefix. 2001::1 with prefix length of 64, start offset is 5.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools", "start_ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of non overlapping IP address ranges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

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
- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/)
- baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-baremetal--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [baremetal.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
