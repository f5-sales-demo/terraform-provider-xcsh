---
page_title: "xcsh_aws_vpc_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site examples."
---

# xcsh_aws_vpc_site examples

<a id="canonical-5c3437aa50cde77e2eedddc34f6ed350d27b75623455e69714667fa60cf9f156"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e481df96f32639f4005ea8a01ef6e8a9c119b4630249683347b2245c152055a"></a>

## Examples — Examples / 41aa4ae5f6c5 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- Examples

<a id="canonical-797df2a6036a9f234d281255b2f7a111837f20d8d9d9e4c26f149287d9112b25"></a>

## Complete configurations — Examples / 41aa4ae5f6c5 / 3

- [Data source](data-sources--aws_vpc_site--examples--group-001.md#canonical-86f032310541025fc60b528e4ed6bd96b47bbc7da63161ca1d8f501f25927e3d): valid configuration.

<a id="canonical-c5a4f581f66c3a334224ded4df76fb5a9c53e6c1158bd712ead5e35347ced8e7"></a>

## Next pages — Examples / 41aa4ae5f6c5 / 4

- [Data source](data-sources--aws_vpc_site--examples--group-001.md#canonical-86f032310541025fc60b528e4ed6bd96b47bbc7da63161ca1d8f501f25927e3d)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-86f032310541025fc60b528e4ed6bd96b47bbc7da63161ca1d8f501f25927e3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bca0546ee7cb50a016e97a9265ce335cd2d24a2517c675a0a8c8b10c9be0c1b9"></a>

## Data source — Data source / 0090ce5468ae / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Examples](data-sources--aws_vpc_site--examples--group-001.md#canonical-5c3437aa50cde77e2eedddc34f6ed350d27b75623455e69714667fa60cf9f156)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_aws_vpc_site/data-source.tf`; digest `sha256:255a62b4362db5d970e6f776cb0dcab3af764d3d9fda6ded76847f22ef47dd92`.

```terraform
# AWSVPCSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSVPCSite by name
data "xcsh_aws_vpc_site" "example" {
  name      = "example-aws-vpc-site"
  namespace = "staging"
}

output "aws_vpc_site_id" {
  value = data.xcsh_aws_vpc_site.example.id
}
```

<a id="canonical-1c59ec6752b5bf4b6e2ec4e13ffc973525be562779982053619e09ce516a9cba"></a>

## Next pages — Data source / 0090ce5468ae / 3

- [Examples](data-sources--aws_vpc_site--examples--group-001.md#canonical-5c3437aa50cde77e2eedddc34f6ed350d27b75623455e69714667fa60cf9f156)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
