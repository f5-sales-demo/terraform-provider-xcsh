---
page_title: "ingress_egress_gw.global_network_list.global_network_connections"
subcategory: "Infrastructure"
description: "Global network connections."
xcsh_docs: {"aliases": ["ingress egress gw global network list global network connections"], "body_bytes": 3110, "body_sha256": "sha256:7fa1e29559ff1e5fe5a4815eebe37c2e37f918fbb1a927d5cb46cf86a963603a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections:slo_to_global_dr"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list", "path": "documentation/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw global network list global network connections sli to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections:sli_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections", "sli_to_global_dr"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw global network list global network connections slo to global dr"], "anchor": "section", "description": "Global network reference for direct connection.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections:slo_to_global_dr", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "global_network_list", "global_network_connections", "slo_to_global_dr"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Global network connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/)
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/): complete subsection reference.

- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/): complete subsection reference.

## Next pages

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/)
- [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
