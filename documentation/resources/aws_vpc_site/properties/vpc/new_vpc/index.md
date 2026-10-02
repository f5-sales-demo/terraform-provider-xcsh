---
page_title: "vpc.new_vpc"
subcategory: "Infrastructure"
description: "Parameters to create new AWS VPC."
xcsh_docs: {"aliases": ["vpc new vpc"], "body_bytes": 4172, "body_sha256": "sha256:7ab0a11baa632f330bde14307101d52d6621ca8c89a0c8499aad18d43b7328d7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc:autogenerate"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc", "path": "documentation/resources/aws_vpc_site/properties/vpc/new_vpc/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3210012313332220-3200021210032131-1023033122320120-2203033212230103-3110133102212020-3300021333120300-0321300131112110-0222200222031002", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-005.md", "relationships": [{"anchor": "schema-vpc--new_vpc--name_tag", "enforcement": "provider-schema", "group": "vpc.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vpc.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc:autogenerate", "type": "conflicts"}, {"anchor": "schema-vpc--new_vpc--primary_ipv4", "enforcement": "provider-schema", "group": "vpc.new_vpc:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vpc", "new_vpc"], "schema_version": 1, "sections": [{"aliases": ["autogenerate"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc:autogenerate", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vpc", "new_vpc", "autogenerate"], "syntax": "attribute", "type": "object"}, {"aliases": ["name tag"], "anchor": "schema-vpc--new_vpc--name_tag", "description": "Exclusive with Specify the VPC Name.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vpc", "new_vpc", "name_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary ipv4"], "anchor": "schema-vpc--new_vpc--primary_ipv4", "description": "IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be modified. All subnets prefixes in this VPC must be part of this CIDR block.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vpc", "new_vpc", "primary_ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/vpc/new_vpc/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters to create new AWS VPC.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc.new_vpc

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/vpc/)
- vpc.new_vpc

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

- [autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/vpc/new_vpc/autogenerate/): complete subsection reference.

<a id="schema-vpc--new_vpc--name_tag"></a>

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

<a id="schema-vpc--new_vpc--primary_ipv4"></a>

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

- [vpc.new_vpc.autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/vpc/new_vpc/autogenerate/)
- [vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/vpc/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
