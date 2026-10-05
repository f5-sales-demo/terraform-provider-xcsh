---
page_title: "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list"
subcategory: ""
description: "IPV6DnsList."
xcsh_docs: {"aliases": ["oci not managed node list interface list ipv6 auto config router dns config configured list"], "body_bytes": 3857, "body_sha256": "sha256:a0acf8d3fdd08c409273509bf3be3cdf3dcd27650a7c026d6bef9af770d25dd4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config", "path": "documentation/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3313323101131111-0131211013201331-2323320220233101-0132223123303020-1203133011222221-3203303033323112-1222311023113113-3030312230130120", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list"], "schema_version": 1, "sections": [{"aliases": ["oci not managed node list interface list ipv6 auto config router dns config configured list dns list"], "anchor": "schema-oci--not_managed--node_list--interface_list--ipv6_auto_config--router--dns_config--configured_list--dns_list", "description": "List of IPv6 Addresses acting as DNS servers.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:dns_config:configured_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "dns_config", "configured_list", "dns_list"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/dns_config/configured_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "IPV6DnsList.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
