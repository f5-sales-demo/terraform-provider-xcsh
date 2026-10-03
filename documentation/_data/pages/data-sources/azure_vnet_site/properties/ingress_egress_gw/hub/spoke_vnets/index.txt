---
page_title: "ingress_egress_gw.hub.spoke_vnets"
subcategory: "Infrastructure"
description: "Spoke VNet Peering."
xcsh_docs: {"aliases": ["ingress egress gw hub spoke vnets"], "body_bytes": 3352, "body_sha256": "sha256:1f64b6131d44999d0d80b3ee60baf039446518265d56ef4bba0d5fc383c680d1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:auto", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:labels", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2011313111320333-2003012011313212-1220020212101211-0113121032211100-3220331213131000-2310303231122302-2313230021031003-0331211133311100", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw hub spoke vnets auto"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:auto", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "auto"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw hub spoke vnets labels"], "anchor": "section", "description": "Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall policy These labels used must be from known key and label defined in shared namespace.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:labels", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "labels"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw hub spoke vnets manual"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "manual"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw hub spoke vnets vnet"], "anchor": "section", "description": "Resource group and name of existing Azure VNet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "vnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Spoke VNet Peering.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.spoke_vnets

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- ingress_egress_gw.hub.spoke_vnets

<a id="section"></a>

Type: `"list"`. Computed.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/auto/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/labels/): complete subsection reference.

- [manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/manual/): complete subsection reference.

- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/): complete subsection reference.

## Next pages

- [ingress_egress_gw.hub.spoke_vnets.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/auto/)
- [ingress_egress_gw.hub.spoke_vnets.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/labels/)
- [ingress_egress_gw.hub.spoke_vnets.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/manual/)
- [ingress_egress_gw.hub.spoke_vnets.vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
