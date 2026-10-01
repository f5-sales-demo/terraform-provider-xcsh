---
page_title: "aws_parameters.new_vpc"
subcategory: ""
description: "aws_parameters.new_vpc for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 3481, "body_sha256": "sha256:8fc001154ceaa9e21b5e726aec21c551825859a5275dbd525384e0cf18a9b92d", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_vpc", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_vpc:autogenerate"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:new_vpc", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters", "path": "docs/guides/data-sources--aws_tgw_site--properties--aws_parameters--new_vpc.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_parameters", "new_vpc"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/aws_parameters/new_vpc/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_parameters.new_vpc for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.new_vpc

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [aws_parameters](data-sources--aws_tgw_site--properties--aws_parameters.md)
- aws_parameters.new_vpc

<a id="section"></a>

Type: `"single"`. Computed.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

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

## Direct properties

- [autogenerate](data-sources--aws_tgw_site--properties--aws_parameters--new_vpc--autogenerate.md): complete subsection reference.

<a id="schema-aws_parameters--new_vpc--name_tag"></a>

### name_tag property

Type: `"string"`. Computed.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

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

<a id="schema-aws_parameters--new_vpc--primary_ipv4"></a>

### primary_ipv4 property

Type: `"string"`. Computed.

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

- [aws_parameters.new_vpc.autogenerate](data-sources--aws_tgw_site--properties--aws_parameters--new_vpc--autogenerate.md)
- [aws_parameters](data-sources--aws_tgw_site--properties--aws_parameters.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
