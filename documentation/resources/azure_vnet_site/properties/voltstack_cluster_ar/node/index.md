---
page_title: "voltstack_cluster_ar.node"
subcategory: "Infrastructure"
description: "Parameters for creating Single interface Node for Alternate Region."
xcsh_docs: {"aliases": ["voltstack cluster ar node"], "body_bytes": 4424, "body_sha256": "sha256:568d3e1c13919bfb767e10dd1675e71e54b48618e5c46e737eb3e5c75873930a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-010.md", "relationships": [{"anchor": "schema-voltstack_cluster_ar--node--fault_domain", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.node:RequiredObjectAttributes:fault_domain,node_number,update_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "type": "requires"}, {"anchor": "schema-voltstack_cluster_ar--node--node_number", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.node:RequiredObjectAttributes:fault_domain,node_number,update_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "type": "requires"}, {"anchor": "schema-voltstack_cluster_ar--node--update_domain", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.node:RequiredObjectAttributes:fault_domain,node_number,update_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster_ar", "node"], "schema_version": 1, "sections": [{"aliases": ["fault domain"], "anchor": "schema-voltstack_cluster_ar--node--fault_domain", "description": "Namuber of fault domains to be used while creating the availability set.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "node", "fault_domain"], "syntax": "attribute", "type": "number"}, {"aliases": ["local subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.node.local_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster_ar.node.local_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node:local_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["voltstack_cluster_ar", "node", "local_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["node number"], "anchor": "schema-voltstack_cluster_ar--node--node_number", "description": "Number of main nodes to create, either 1 or 3.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "node", "node_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["update domain"], "anchor": "schema-voltstack_cluster_ar--node--update_domain", "description": "Namuber of update domains to be used while creating the availability set.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:node", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster_ar", "node", "update_domain"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for creating Single interface Node for Alternate Region.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster_ar.node

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- voltstack_cluster_ar.node

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating Single interface Node for Alternate Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fault_domain",
    "node_number",
    "update_domain")}
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
node {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster_ar--node--fault_domain"></a>

### fault_domain property

Type: `"number"`. Optional.

Namuber of fault domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 3),
}
```

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

- [local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/): complete subsection reference.

<a id="schema-voltstack_cluster_ar--node--node_number"></a>

### node_number property

Type: `"number"`. Optional.

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

<a id="schema-voltstack_cluster_ar--node--update_domain"></a>

### update_domain property

Type: `"number"`. Optional.

Namuber of update domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

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

- [voltstack_cluster_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
