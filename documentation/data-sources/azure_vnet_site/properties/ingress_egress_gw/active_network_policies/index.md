---
page_title: "ingress_egress_gw.active_network_policies"
subcategory: "Infrastructure"
description: "List of firewall policy views."
xcsh_docs: {"aliases": ["ingress egress gw active network policies"], "body_bytes": 1627, "body_sha256": "sha256:238297aabdef759f0926d45d84fa2c97e7d48c5202fe02bf41a2a9d985a833f3", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies:network_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0321022000222022-3131202022210023-2322210213232131-0013202311322003-0111112011313133-1112011131000311-2201003120201203-0321303301130222", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "active_network_policies"], "schema_version": 1, "sections": [{"aliases": ["network policies"], "anchor": "section", "description": "Ordered List of Firewall Policies active for this network firewall.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies:network_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ingress_egress_gw", "active_network_policies", "network_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of firewall policy views.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.active_network_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- ingress_egress_gw.active_network_policies

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

- [network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/network_policies/): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/network_policies/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
