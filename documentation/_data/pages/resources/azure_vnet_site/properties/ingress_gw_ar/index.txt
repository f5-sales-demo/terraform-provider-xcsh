---
page_title: "ingress_gw_ar"
subcategory: "Infrastructure"
description: "Single interface Azure ingress site."
xcsh_docs: {"aliases": ["ingress gw ar"], "body_bytes": 3707, "body_sha256": "sha256:5348d441ee4683a1f2f937a0340718df873997d72d62ac5cff61582bdd917697", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "documentation/resources/azure_vnet_site/properties/ingress_gw_ar/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3100320201112313-1011012113210322-1000233330130201-0100100022211103-0200311001112230-3202222013102010-0332322302323212-1211221101203010", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "schema-ingress_gw_ar--azure_certified_hw", "enforcement": "provider-schema", "group": "ingress_gw_ar:RequiredObjectAttributes:azure_certified_hw", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar"], "schema_version": 1, "sections": [{"aliases": ["accelerated networking"], "anchor": "section", "description": "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.accelerated_networking:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:accelerated_networking:enable", "type": "conflicts"}], "schema_path": ["ingress_gw_ar", "accelerated_networking"], "syntax": "block", "type": "object"}, {"aliases": ["azure certified hw"], "anchor": "schema-ingress_gw_ar--azure_certified_hw", "description": "Name for Azure certified hardware.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "azure_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["node"], "anchor": "section", "description": "Parameters for creating Single interface Node for Alternate Region.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_gw_ar--node--fault_domain", "enforcement": "provider-schema", "group": "ingress_gw_ar.node:RequiredObjectAttributes:fault_domain,node_number,update_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node", "type": "requires"}, {"anchor": "schema-ingress_gw_ar--node--node_number", "enforcement": "provider-schema", "group": "ingress_gw_ar.node:RequiredObjectAttributes:fault_domain,node_number,update_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node", "type": "requires"}, {"anchor": "schema-ingress_gw_ar--node--update_domain", "enforcement": "provider-schema", "group": "ingress_gw_ar.node:RequiredObjectAttributes:fault_domain,node_number,update_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:node", "type": "requires"}], "schema_path": ["ingress_gw_ar", "node"], "syntax": "block", "type": "object"}, {"aliases": ["performance enhancement mode"], "anchor": "section", "description": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced", "type": "conflicts"}], "schema_path": ["ingress_gw_ar", "performance_enhancement_mode"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw_ar/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Single interface Azure ingress site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- ingress_gw_ar

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ingress gw ar.

Upstream description:

Single interface Azure ingress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("azure_certified_hw")}
```

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

Terraform syntax:

```terraform
ingress_gw_ar {
  # Configure direct properties listed below.
}
```

## Direct properties

- [accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/): complete subsection reference.

<a id="schema-ingress_gw_ar--azure_certified_hw"></a>

### azure_certified_hw property

Type: `"string"`. Optional.

\[Enum: azure-byol-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware. The only
possible value is \`azure-byol-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltmesh"),
}
```

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

- [node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/node/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/): complete subsection reference.

## Next pages

- [ingress_gw_ar.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/)
- [ingress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/node/)
- [ingress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
