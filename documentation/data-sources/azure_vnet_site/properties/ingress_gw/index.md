---
page_title: "ingress_gw"
subcategory: "Infrastructure"
description: "Single interface Azure ingress site on on Recommended Region."
xcsh_docs: {"aliases": ["ingress gw"], "body_bytes": 3211, "body_sha256": "sha256:88109c7b6a5ed38181a37a92199ade7cd5995e0286f85fe3653727d15e7bf932", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:accelerated_networking", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:az_nodes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_gw/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3131002031302201-3131200302211201-3321302320022222-0121110331203131-0213023331312031-1121133023012302-0231001231103032-0203312022210001", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw"], "schema_version": 1, "sections": [{"aliases": ["accelerated networking"], "anchor": "section", "description": "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:accelerated_networking", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw", "accelerated_networking"], "syntax": "attribute", "type": "object"}, {"aliases": ["az nodes"], "anchor": "section", "description": "Only Single AZ or Three AZ(s) nodes are supported currently.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:az_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ingress_gw", "az_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure certified hw"], "anchor": "schema-ingress_gw--azure_certified_hw", "description": "Name for Azure certified hardware.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "azure_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["performance enhancement mode"], "anchor": "section", "description": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw", "performance_enhancement_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Single interface Azure ingress site on on Recommended Region.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- ingress_gw

<a id="section"></a>

Type: `"single"`. Computed.

Single interface Azure ingress site on on Recommended Region.

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

- [accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/accelerated_networking/): complete subsection reference.

- [az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/): complete subsection reference.

<a id="schema-ingress_gw--azure_certified_hw"></a>

### azure_certified_hw property

Type: `"string"`. Computed.

\[Enum: azure-byol-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware. The only
possible value is \`azure-byol-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/): complete subsection reference.

## Next pages

- [ingress_gw.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/accelerated_networking/)
- [ingress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/)
- [ingress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
