---
page_title: "ingress_gw.az_nodes"
subcategory: "Infrastructure"
description: "Only Single AZ or Three AZ(s) nodes are supported currently."
xcsh_docs: {"aliases": ["ingress gw az nodes"], "body_bytes": 2681, "body_sha256": "sha256:ede8b7702e6f872058132fc72fd78e12932c7799013a57d03720f97e8fc4f0fb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes:local_subnet"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw", "path": "documentation/resources/aws_vpc_site/properties/ingress_gw/az_nodes/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2330123030021300-0333222331002031-2122221322133023-1323322003320020-1210003321202232-0102331330021201-0022103012213020-0103331133211210", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-004.md", "relationships": [{"anchor": "schema-ingress_gw--az_nodes--aws_az_name", "enforcement": "provider-schema", "group": "ingress_gw.az_nodes:RequiredListObjectAttributes:aws_az_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "az_nodes"], "schema_version": 1, "sections": [{"aliases": ["ingress gw az nodes aws az name"], "anchor": "schema-ingress_gw--az_nodes--aws_az_name", "description": "AWS availability zone, must be consistent with the selected AWS region.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "az_nodes", "aws_az_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress gw az nodes local subnet"], "anchor": "section", "description": "Parameters for AWS subnet.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes:local_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_gw--az_nodes--local_subnet--existing_subnet_id", "enforcement": "provider-schema", "group": "ingress_gw.az_nodes.local_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes:local_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw.az_nodes.local_subnet:ConflictingObjectAttributes:existing_subnet_id,subnet_param", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:az_nodes:local_subnet:subnet_param", "type": "conflicts"}], "schema_path": ["ingress_gw", "az_nodes", "local_subnet"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_gw/az_nodes/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Only Single AZ or Three AZ(s) nodes are supported currently.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.az_nodes

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/)
- ingress_gw.az_nodes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name")}
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

<a id="schema-ingress_gw--az_nodes--aws_az_name"></a>

### aws_az_name property

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/): complete subsection reference.

## Next pages

- [ingress_gw.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/az_nodes/local_subnet/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
