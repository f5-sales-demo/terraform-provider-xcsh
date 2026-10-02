---
page_title: "vpc"
subcategory: "Infrastructure"
description: "This defines choice about AWS VPC for a view."
xcsh_docs: {"aliases": ["vpc"], "body_bytes": 2746, "body_sha256": "sha256:031bacb12f3fe4e4e3e07dbf1eb83f928ee4c1b63b9d2847ff3b4b866054e4a8", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:vpc", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/vpc/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2003301201220030-2310111030110321-1023333120123110-3020131220223010-3230003013133033-1023220013310123-3332200230333020-2220002022320010", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-005.md", "relationships": [{"anchor": "schema-vpc--vpc_id", "enforcement": "provider-schema", "group": "vpc:ConflictingObjectAttributes:new_vpc,vpc_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vpc:ConflictingObjectAttributes:new_vpc,vpc_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vpc"], "schema_version": 1, "sections": [{"aliases": ["new vpc"], "anchor": "section", "description": "Parameters to create new AWS VPC.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-vpc--new_vpc--name_tag", "enforcement": "provider-schema", "group": "vpc.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vpc.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc:autogenerate", "type": "conflicts"}, {"anchor": "schema-vpc--new_vpc--primary_ipv4", "enforcement": "provider-schema", "group": "vpc.new_vpc:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "type": "requires"}], "schema_path": ["vpc", "new_vpc"], "syntax": "block", "type": "object"}, {"aliases": ["vpc id"], "anchor": "schema-vpc--vpc_id", "description": "Exclusive with Information about existing VPC ID.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vpc", "vpc_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/vpc/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines choice about AWS VPC for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- vpc

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about AWS VPC for a view.

Upstream description:

This defines choice about AWS VPC for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("new_vpc",
    "vpc_id")}
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
  "x-ves-oneof-field-choice": "[\"new_vpc\",\"vpc_id\"]"
}
```

Terraform syntax:

```terraform
vpc {
  # Configure direct properties listed below.
}
```

## Direct properties

- [new_vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/vpc/new_vpc/): complete subsection reference.

<a id="schema-vpc--vpc_id"></a>

### vpc_id property

Type: `"string"`. Optional.

Exclusive with \[new\_vpc\] Information about existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Information about existing VPC ID.

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
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [vpc.new_vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/vpc/new_vpc/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
