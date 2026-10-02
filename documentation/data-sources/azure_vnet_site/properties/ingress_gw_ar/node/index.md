---
page_title: "ingress_gw_ar.node"
subcategory: "Infrastructure"
description: "Parameters for creating Single interface Node for Alternate Region."
xcsh_docs: {"aliases": ["ingress gw ar node"], "body_bytes": 3792, "body_sha256": "sha256:16270fa7769c4796b64645470fb811be1c0f1b2395b8f89ab23342f235ef1c07", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0010210313032011-1011113033113211-3123112133302321-2302321012213301-3310100230200303-0330123211122331-1003000213311010-3020103101121002", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar", "node"], "schema_version": 1, "sections": [{"aliases": ["fault domain"], "anchor": "schema-ingress_gw_ar--node--fault_domain", "description": "Namuber of fault domains to be used while creating the availability set.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "node", "fault_domain"], "syntax": "attribute", "type": "number"}, {"aliases": ["local subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node:local_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw_ar", "node", "local_subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["node number"], "anchor": "schema-ingress_gw_ar--node--node_number", "description": "Number of main nodes to create, either 1 or 3.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "node", "node_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["update domain"], "anchor": "schema-ingress_gw_ar--node--update_domain", "description": "Namuber of update domains to be used while creating the availability set.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar:node", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw_ar", "node", "update_domain"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for creating Single interface Node for Alternate Region.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.node

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/)
- ingress_gw_ar.node

<a id="section"></a>

Type: `"single"`. Computed.

Parameters for creating Single interface Node for Alternate Region.

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

<a id="schema-ingress_gw_ar--node--fault_domain"></a>

### fault_domain property

Type: `"number"`. Computed.

Namuber of fault domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/): complete subsection reference.

<a id="schema-ingress_gw_ar--node--node_number"></a>

### node_number property

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

<a id="schema-ingress_gw_ar--node--update_domain"></a>

### update_domain property

Type: `"number"`. Computed.

Namuber of update domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

## Next pages

- [ingress_gw_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
