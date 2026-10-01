---
page_title: "xcsh_aws_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site landing."
---

# xcsh_aws_vpc_site landing

<a id="canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303233013320312-2222013303211101-3301313011230012-0122333031232030-3313002012001301-2321321100033033-2110222202113320-3323102202002012"></a>

## xcsh_aws_vpc_site — xcsh_aws_vpc_site / 121232012012 / 2

Breadcrumbs:

- xcsh_aws_vpc_site

Manages a AWS VPC Site resource in F5 Distributed Cloud for deploying F5 sites within AWS VPC
environments.

<a id="canonical-0302010103312231-3331102111100223-3023211122223220-1302021311323220-1220323311103320-0302112102202032-0300230123201222-0312213133023312"></a>

## Prerequisites — xcsh_aws_vpc_site / 121232012012 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: AWS authentication for deployment

<a id="canonical-2322121101210030-1232223302210211-3200311200022033-1303112333303312-2332221002212311-1210211121221031-1332210130021202-3222332102000020"></a>

## Minimal configuration — xcsh_aws_vpc_site / 121232012012 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1111313121110331-2010221203132311-2302002300111203-2201123020310201-0101011112130300-3102213132012121-2110211131330302-0000313200312331"></a>

## Root configuration — xcsh_aws_vpc_site / 121232012012 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0021021110223000-2133120200322232-2133111123322011-3002201112111321-1300230333230233-2203301320301223-0321101113330310-3030302300032032"></a>

## Next pages — xcsh_aws_vpc_site / 121232012012 / 6

- [Property reference](../guides/data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [Examples](../guides/data-sources--aws_vpc_site--examples--group-001.md#canonical-1130031003132222-1100303132131332-0232323131313003-1033123231031100-3102132313111202-0310111132122113-0110121213332212-0030332133011112)
