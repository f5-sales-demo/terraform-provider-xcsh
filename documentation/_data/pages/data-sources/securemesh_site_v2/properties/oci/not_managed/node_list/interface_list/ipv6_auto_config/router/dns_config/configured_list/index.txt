---
page_title: "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list"
subcategory: ""
description: "IPV6DnsList."
xcsh_docs: {"aliases": ["oci not managed node list interface list ipv6 auto config router dns config configured list"], "body_bytes": 3857, "body_sha256": "sha256:52c2f9314542567387386f9ea0c6c3737bcaae50b60ddefb5f75e3f06a3b700d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "documentation/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3313323101131111-0131211013201331-2323320220233101-0132223123303020-1203133011222221-3203303033323112-1222311023113113-3030312230130120", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list"], "schema_version": 1, "sections": [{"aliases": ["dns list"], "anchor": "schema-oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list--dns_list", "description": "List of IPv6 Addresses acting as DNS servers.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list", "dns_list"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPV6DnsList.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [oci](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/)
- [oci.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/)
- [oci.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/)
- [oci.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

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

<a id="schema-oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list--dns_list"></a>

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

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
