---
page_title: "ingress_egress_gw_ar.active_forward_proxy_policies"
subcategory: "Infrastructure"
description: "Ordered List of Forward Proxy Policies active."
xcsh_docs: {"aliases": ["ingress egress gw ar active forward proxy policies"], "body_bytes": 1647, "body_sha256": "sha256:0c0cbed4bd71116e76f20760cafd598c435a9f1fd511db90d96c0d083145f6f6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:active_forward_proxy_policies:forward_proxy_policies"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:active_forward_proxy_policies", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2301302013023033-2330332301010311-2322002330001233-2220010332300030-1220222100323300-0121231121232210-1012021230202220-1103301233313310", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw_ar", "active_forward_proxy_policies"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw ar active forward proxy policies forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:active_forward_proxy_policies:forward_proxy_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ingress_egress_gw_ar", "active_forward_proxy_policies", "forward_proxy_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Ordered List of Forward Proxy Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.active_forward_proxy_policies

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- ingress_egress_gw_ar.active_forward_proxy_policies

<a id="section"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

- [forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/forward_proxy_policies/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/forward_proxy_policies/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
