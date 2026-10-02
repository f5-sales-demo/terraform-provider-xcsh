---
page_title: "ingress_egress_gw.global_network_list.global_network_connections"
subcategory: "Infrastructure"
description: "Global network connections."
xcsh_docs: {"aliases": ["ingress egress gw global network list global network connections"], "body_bytes": 3146, "body_sha256": "sha256:90a63553d56011333b59dea1de93628ab16e0c50779326d47ea66e5b9f5eeeb7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections:slo_to_global_dr"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3133301113112332-3233003231333313-0020001320220300-0031000223132110-1021021122100302-2013200230133221-0330212013333302-2313022322211013", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections"], "schema_version": 1, "sections": [{"aliases": ["sli to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections", "sli_to_global_dr"], "syntax": "attribute", "type": "object"}, {"aliases": ["slo to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list:global_network_connections:slo_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections", "slo_to_global_dr"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Global network connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="section"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/): complete subsection reference.

- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/): complete subsection reference.

## Next pages

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/)
- [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
