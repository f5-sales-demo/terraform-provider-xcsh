---
page_title: "xcsh_aws_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site landing."
---

# xcsh_aws_vpc_site landing

<a id="canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a832ac0434d5dce8ef147250f92a3cc63b9508831e933376cf47687e912f394"></a>

## xcsh_aws_vpc_site — xcsh_aws_vpc_site / 97c11fc3911b / 2

Breadcrumbs:

- xcsh_aws_vpc_site

Manages a AWS VPC Site resource in F5 Distributed Cloud for deploying F5 sites within AWS VPC
environments.

<a id="canonical-accc2350be5a8d1f41150c538d0852131b7e783a39549325a755a1d4e60251a7"></a>

## Prerequisites — xcsh_aws_vpc_site / 97c11fc3911b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: AWS authentication for deployment

<a id="canonical-d472e6341ec03204dcaf85a049aa367ce8cd2713b9fe5029ed7307081502e5e7"></a>

## Minimal configuration — xcsh_aws_vpc_site / 97c11fc3911b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AWSVPCSite Resource Example
# Manages a AWS VPC Site resource in F5 Distributed Cloud for deploying F5 sites within AWS VPC environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AWSVPCSite configuration
resource "xcsh_aws_vpc_site" "example" {
  name      = "example-aws-vpc-site"
  namespace = "staging"

  aws_region    = "example-value"
  instance_type = "example-value"
  ssh_key       = "example-value"
}
```

<a id="canonical-af59e3a9958ebab62fd4e94b2a047bb205494e5c7d5a9b1aff2c6b4f087a9461"></a>

## Root configuration — xcsh_aws_vpc_site / 97c11fc3911b / 5

Required root properties: `aws_region`, `instance_type`, `name`, `namespace`, `ssh_key`. Full root flags and choices appear in the property reference.

<a id="canonical-e65f038e170464fc8d8e4c76911d9fc61e80322737a42b5f84939f6f105ba585"></a>

## Next pages — xcsh_aws_vpc_site / 97c11fc3911b / 6

- [Property reference](../guides/resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [Examples](../guides/resources--aws_vpc_site--examples--group-001.md#canonical-c9a8749ca7e9910893bcefc40a95b7ad0057213fb843066fc240448a0674c2d2)
- [Import](../guides/resources--aws_vpc_site--lifecycle--group-001.md#canonical-3e8f90c69e8ebfeb658d6feb1de0617f71dd5a97aec1457023708b5337460f05)
- [Timeouts](../guides/resources--aws_vpc_site--lifecycle--group-001.md#canonical-f8a32e62f8d750cd4262518db7e8a1c97e6f9238aff919a71c1adfefb29fe2f5)
