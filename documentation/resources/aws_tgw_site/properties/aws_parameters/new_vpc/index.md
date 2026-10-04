---
page_title: "aws_parameters.new_vpc"
subcategory: ""
description: "Parameters to create new AWS VPC."
xcsh_docs: {"aliases": ["aws parameters new vpc"], "body_bytes": 4293, "body_sha256": "sha256:d2b2ada55bdfb44107a7b180178a89dff8de7411b662e07ea0af97a73b274485", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc:autogenerate"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/new_vpc/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0220020323123321-0313220113210230-0113310122202133-2213030113313113-2003203010103222-0032202230121212-2000112223330031-2120322213130322", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [{"anchor": "schema-aws_parameters--new_vpc--name_tag", "enforcement": "provider-schema", "group": "aws_parameters.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc:autogenerate", "type": "conflicts"}, {"anchor": "schema-aws_parameters--new_vpc--primary_ipv4", "enforcement": "provider-schema", "group": "aws_parameters.new_vpc:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "new_vpc"], "schema_version": 1, "sections": [{"aliases": ["aws parameters new vpc autogenerate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc:autogenerate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_vpc", "autogenerate"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters new vpc name tag"], "anchor": "schema-aws_parameters--new_vpc--name_tag", "description": "Exclusive with Specify the VPC Name.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_vpc", "name_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters new vpc primary ipv4"], "anchor": "schema-aws_parameters--new_vpc--primary_ipv4", "description": "IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be modified. All subnets prefixes in this VPC must be part of this CIDR block.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "new_vpc", "primary_ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/new_vpc/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Parameters to create new AWS VPC.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_vpc

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.new_vpc

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4"),
  validators.ConflictingObjectAttributes("autogenerate",
    "name_tag")}
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
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name_tag\"]"
}
```

Terraform syntax:

```terraform
new_vpc {
  # Configure direct properties listed below.
}
```

## Direct properties

- [autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/autogenerate/): complete subsection reference.

<a id="schema-aws_parameters--new_vpc--name_tag"></a>

### name_tag property

Type: `"string"`. Optional.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-aws_parameters--new_vpc--primary_ipv4"></a>

### primary_ipv4 property

Type: `"string"`. Optional.

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

Upstream description:

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  }
}
```

## Next pages

- [aws_parameters.new_vpc.autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/autogenerate/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
