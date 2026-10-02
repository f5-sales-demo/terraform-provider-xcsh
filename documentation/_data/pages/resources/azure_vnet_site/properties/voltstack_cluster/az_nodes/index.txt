---
page_title: "voltstack_cluster.az_nodes"
subcategory: "Infrastructure"
description: "Only Single AZ or Three AZ(s) nodes are supported currently."
xcsh_docs: {"aliases": ["voltstack cluster az nodes"], "body_bytes": 3296, "body_sha256": "sha256:a0746ddc025fb7c53510672be0f3124312769199c2c9297fe278ddb5179c4865", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1110030203112331-1013331332022220-1131111111202200-1313201113200032-0100300023212210-1230202123102322-2322310110221331-2322301103133021", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "schema-voltstack_cluster--az_nodes--azure_az", "enforcement": "provider-schema", "group": "voltstack_cluster.az_nodes:RequiredListObjectAttributes:azure_az", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "az_nodes"], "schema_version": 1, "sections": [{"aliases": ["azure az"], "anchor": "schema-voltstack_cluster--az_nodes--azure_az", "description": "A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes", "azure_az"], "syntax": "attribute", "type": "string"}, {"aliases": ["local subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.az_nodes.local_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.az_nodes.local_subnet:ConflictingObjectAttributes:subnet,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:az_nodes:local_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["voltstack_cluster", "az_nodes", "local_subnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Only Single AZ or Three AZ(s) nodes are supported currently.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.az_nodes

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- voltstack_cluster.az_nodes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("azure_az")}
```

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

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-voltstack_cluster--az_nodes--azure_az"></a>

### azure_az property

Type: `"string"`. Optional.

\[Enum: 1|2|3\] Zone depicting a grouping of datacenters within an Azure region. Expecting numeric
input. Possible values are \`1\`, \`2\`, \`3\`.

Upstream description:

A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("1",
    "2",
    "3"),
}
```

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

- [local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/): complete subsection reference.

## Next pages

- [voltstack_cluster.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
