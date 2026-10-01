---
page_title: "vpc.new_vpc"
subcategory: "Infrastructure"
description: "vpc.new_vpc for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 3817, "body_sha256": "sha256:d8aabd062c18ec04cf1a8a5d49761ab47b3d4c649748ba65eebf0a511380c285", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc:autogenerate"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:vpc:new_vpc", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:vpc", "path": "docs/guides/resources--aws_vpc_site--properties--vpc--new_vpc.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vpc", "new_vpc"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/vpc/new_vpc/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vpc.new_vpc for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc.new_vpc

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [vpc](resources--aws_vpc_site--properties--vpc.md)
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

- [autogenerate](resources--aws_vpc_site--properties--vpc--new_vpc--autogenerate.md): complete subsection reference.

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

- [vpc.new_vpc.autogenerate](resources--aws_vpc_site--properties--vpc--new_vpc--autogenerate.md)
- [vpc](resources--aws_vpc_site--properties--vpc.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
