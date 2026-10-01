---
page_title: "xcsh_aws_tgw_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site landing."
---

# xcsh_aws_tgw_site landing

<a id="canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1653dee1a8e793e930c7b70a8fdae5d08097b3254866bbb55aa4147842b2709a"></a>

## xcsh_aws_tgw_site — xcsh_aws_tgw_site / f9cce303d3c3 / 2

Breadcrumbs:

- xcsh_aws_tgw_site

Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS
Transit Gateway.

<a id="canonical-33113d79e0291c271573dacd64c0e14022f6044c28ef95299ab2b926ca3bbfc1"></a>

## Prerequisites — xcsh_aws_tgw_site / f9cce303d3c3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-928c8e0a06e37eaee7da0380a0550717dab9361b53364953d23a46c4a3e1f232"></a>

## Minimal configuration — xcsh_aws_tgw_site / f9cce303d3c3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AWSTGWSite Resource Example
# Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS Transit Gateway.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AWSTGWSite configuration
resource "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}
```

<a id="canonical-2823d4776de37ce82ab1a974f59f5cab980e77cf5dd1c881561d6bd29fe2c8f3"></a>

## Root configuration — xcsh_aws_tgw_site / f9cce303d3c3 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0319d613d013cff993dd2a919e28f1887c02c41a22d29b9f1237650751a38e66"></a>

## Next pages — xcsh_aws_tgw_site / f9cce303d3c3 / 6

- [Property reference](../guides/resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [Examples](../guides/resources--aws_tgw_site--examples--group-001.md#canonical-42c06f13d48332ed5021e5a105e3d74e9187d47b541c18dbff2f2277fff7eadb)
- [Import](../guides/resources--aws_tgw_site--lifecycle--group-001.md#canonical-076a116e8d168cba4a4cc59ed6f317f9fb7fb28e95b33ba227cb1d98a16cd004)
- [Timeouts](../guides/resources--aws_tgw_site--lifecycle--group-001.md#canonical-3468a7efe44f9c635fa1b18471f8644156f31532d003999d1e3bb4f1b53ae64c)
