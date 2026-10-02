---
page_title: "ingress_egress_gw.az_nodes"
subcategory: "Infrastructure"
description: "Only Single AZ or Three AZ(s) nodes are supported currently."
xcsh_docs: {"aliases": ["ingress egress gw az nodes"], "body_bytes": 3284, "body_sha256": "sha256:c49f9de4a41b763fa4f7ab5b714fcd9d76fb353ac21d00eb74c410483ea61d47", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2320003002102002-3031230002210222-2200333233212032-0122112133200323-0211203203202110-2310220113320323-3021301221013322-1130210221221230", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes"], "schema_version": 1, "sections": [{"aliases": ["azure az"], "anchor": "schema-ingress_egress_gw--az_nodes--azure_az", "description": "A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "azure_az"], "syntax": "attribute", "type": "string"}, {"aliases": ["inside subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:inside_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "inside_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Only Single AZ or Three AZ(s) nodes are supported currently.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- ingress_egress_gw.az_nodes

<a id="section"></a>

Type: `"list"`. Computed.

Only Single AZ or Three AZ(s) nodes are supported currently.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

## Direct properties

<a id="schema-ingress_egress_gw--az_nodes--azure_az"></a>

### azure_az property

Type: `"string"`. Computed.

\[Enum: 1|2|3\] Zone depicting a grouping of datacenters within an Azure region. Expecting numeric
input. Possible values are \`1\`, \`2\`, \`3\`.

Upstream description:

A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "1",
    "2",
    "3"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  }
}
```

- [inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/): complete subsection reference.

- [outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/)
- [ingress_egress_gw.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
