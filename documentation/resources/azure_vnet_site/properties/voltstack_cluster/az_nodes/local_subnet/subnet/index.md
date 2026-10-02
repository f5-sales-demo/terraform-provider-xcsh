---
page_title: "voltstack_cluster.az_nodes.local_subnet.subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["voltstack cluster az nodes local subnet subnet"], "body_bytes": 4639, "body_sha256": "sha256:ccb62aa8323e23e4752b99d57feadf299bd19025ea139606fd43417d2196579a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet:vnet_resource_group"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2321021200010031-2303131232032002-0211311332133223-3032302213231102-2003332231102312-0320003032310310-1332021002130023-2103133233230020", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_resource_grp", "enforcement": "provider-schema", "group": "voltstack_cluster.az_nodes.local_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.az_nodes.local_subnet.subnet:ConflictingObjectAttributes:subnet_resource_grp,vnet_resource_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet:vnet_resource_group", "type": "conflicts"}, {"anchor": "schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_name", "enforcement": "provider-schema", "group": "voltstack_cluster.az_nodes.local_subnet.subnet:RequiredObjectAttributes:subnet_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet"], "schema_version": 1, "sections": [{"aliases": ["subnet name"], "anchor": "schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_name", "description": "Name of existing subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet", "subnet_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["subnet resource grp"], "anchor": "schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_resource_grp", "description": "Exclusive with Specify name of Resource Group.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet", "subnet_resource_grp"], "syntax": "attribute", "type": "string"}, {"aliases": ["vnet resource group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet:vnet_resource_group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet", "subnet", "vnet_resource_group"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.az_nodes.local_subnet.subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/)
- [voltstack_cluster.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/)
- voltstack_cluster.az_nodes.local_subnet.subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_name"></a>

### subnet_name property

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_resource_grp"></a>

### subnet_resource_grp property

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/vnet_resource_group/): complete subsection reference.

## Next pages

- [voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/vnet_resource_group/)
- [voltstack_cluster.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
