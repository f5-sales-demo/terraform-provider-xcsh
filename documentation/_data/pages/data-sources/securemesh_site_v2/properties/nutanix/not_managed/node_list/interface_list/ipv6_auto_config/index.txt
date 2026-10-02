---
page_title: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config"
subcategory: ""
description: "IPV6AutoConfigType."
xcsh_docs: {"aliases": ["nutanix not managed node list interface list ipv6 auto config"], "body_bytes": 2680, "body_sha256": "sha256:fae81fd4a845ec5ac496b45cb74da7a98b031fb58e8087637847238a2af306bf", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:host", "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config"], "schema_version": 1, "sections": [{"aliases": ["host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:host", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "host"], "syntax": "attribute", "type": "object"}, {"aliases": ["router"], "anchor": "section", "description": "IPV6AutoConfigRouterType.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IPV6AutoConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nutanix.not_managed.node_list.interface_list.ipv6_auto_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [nutanix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/)
- [nutanix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/)
- [nutanix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/)
- [nutanix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="section"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

## Direct properties

- [host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/host/): complete subsection reference.

- [router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/): complete subsection reference.

## Next pages

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/host/)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [nutanix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
