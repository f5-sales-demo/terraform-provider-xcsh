---
page_title: "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router"
subcategory: ""
description: "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3533, "body_sha256": "sha256:898a2553ccd4f7de2ad9f12f70dcbc0492024f1d79fc173bef36af424c3682e7", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config:router", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:kvm:not_managed:node_list:interface_list:ipv6_auto_config", "path": "docs/guides/data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list--ipv6_auto_config--router.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kvm", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/kvm/not_managed/node_list/interface_list/ipv6_auto_config/router/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kvm.not_managed.node_list.interface_list.ipv6_auto_config.router for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kvm.not_managed.node_list.interface_list.ipv6_auto_config.router

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [kvm](data-sources--securemesh_site_v2--properties--kvm.md)
- [kvm.not_managed](data-sources--securemesh_site_v2--properties--kvm--not_managed.md)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list.md)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list.md)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list--ipv6_auto_config.md)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="section"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

## Direct properties

- [dns_config](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md): complete subsection reference.

<a id="schema-kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--network_prefix"></a>

### network_prefix property

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md): complete subsection reference.

## Next pages

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--kvm--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
