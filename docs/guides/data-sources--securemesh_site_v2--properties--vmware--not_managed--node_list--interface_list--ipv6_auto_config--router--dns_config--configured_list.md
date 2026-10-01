---
page_title: "vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list"
subcategory: ""
description: "vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3378, "body_sha256": "sha256:e58ddcb430196b45ec7d4448fb97bebbacd8bcc5d2fcd1fdf41300509caf2583", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "docs/guides/data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [vmware](data-sources--securemesh_site_v2--properties--vmware.md)
- [vmware.not_managed](data-sources--securemesh_site_v2--properties--vmware--not_managed.md)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list.md)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list.md)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="section"></a>

Type: `"single"`. Computed.

IPV6DnsList.

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

<a id="schema-vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list--dns_list"></a>

### dns_list property

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--properties--vmware--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
