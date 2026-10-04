---
page_title: "aws_parameters.az_nodes.inside_subnet.subnet_param"
subcategory: ""
description: "Parameters for creating a new cloud subnet."
xcsh_docs: {"aliases": ["aws parameters az nodes inside subnet subnet param"], "body_bytes": 2800, "body_sha256": "sha256:fde46f3fb37640b70176b42b01f9b19497f674ca65f7a9bbd9c0989b20aba9ab", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet:subnet_param", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/subnet_param/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0212000320232300-0311222123313132-1320220103213130-1133122300233200-1321211111021310-2131223333130302-0203132033113202-3122213010132131", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "schema-aws_parameters--az_nodes--inside_subnet--subnet_param--ipv4", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes.inside_subnet.subnet_param:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet:subnet_param", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "az_nodes", "inside_subnet", "subnet_param"], "schema_version": 1, "sections": [{"aliases": ["aws parameters az nodes inside subnet subnet param ipv4"], "anchor": "schema-aws_parameters--az_nodes--inside_subnet--subnet_param--ipv4", "description": "IPv4 subnet prefix for this subnet.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet:subnet_param", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "az_nodes", "inside_subnet", "subnet_param", "ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Parameters for creating a new cloud subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.az_nodes.inside_subnet.subnet_param

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [aws_parameters.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/)
- [aws_parameters.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/)
- aws_parameters.az_nodes.inside_subnet.subnet_param

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_parameters--az_nodes--inside_subnet--subnet_param--ipv4"></a>

### ipv4 property

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

## Next pages

- [aws_parameters.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
