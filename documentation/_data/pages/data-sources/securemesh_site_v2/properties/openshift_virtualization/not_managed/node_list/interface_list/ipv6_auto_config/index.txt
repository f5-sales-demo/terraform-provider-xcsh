---
page_title: "openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config"
subcategory: ""
description: "IPV6AutoConfigType."
xcsh_docs: {"aliases": ["openshift virtualization not managed node list interface list ipv6 auto config"], "body_bytes": 2986, "body_sha256": "sha256:8daba4363c5f807d2313765414e9d405a58fa502366fae838859dc83d125591e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:host", "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list", "path": "documentation/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2231300321021130-3311032001203032-0233302102210321-2100102010022003-3302001313000302-0222213232322203-1311120323222210-2332101101121320", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "ipv6_auto_config"], "schema_version": 1, "sections": [{"aliases": ["openshift virtualization not managed node list interface list ipv6 auto config host"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:host", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "host"], "syntax": "attribute", "type": "object"}, {"aliases": ["openshift virtualization not managed node list interface list ipv6 auto config router"], "anchor": "section", "description": "IPV6AutoConfigRouterType.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:openshift_virtualization:not_managed:node_list:interface_list:ipv6_auto_config:router", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["openshift_virtualization", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "IPV6AutoConfigType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [openshift_virtualization](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/)
- [openshift_virtualization.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/)
- [openshift_virtualization.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/)
- [openshift_virtualization.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config

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

- [host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/host/): complete subsection reference.

- [router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/): complete subsection reference.

## Next pages

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/host/)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/ipv6_auto_config/router/)
- [openshift_virtualization.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/openshift_virtualization/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
