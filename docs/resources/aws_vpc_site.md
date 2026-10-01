---
page_title: "xcsh_aws_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site landing."
---

# xcsh_aws_vpc_site landing

<a id="canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322200302223000-1003103111313032-2032330110130211-0033210222033030-1203232111002020-0301322103030313-1230331013122013-3221010233032110"></a>

## xcsh_aws_vpc_site — xcsh_aws_vpc_site / 300321010123 / 2

Breadcrumbs:

- xcsh_aws_vpc_site

Manages a AWS VPC Site resource in F5 Distributed Cloud for deploying F5 sites within AWS VPC
environments.

<a id="canonical-2230303002031100-2332112220310133-1001011100301103-2031002011020103-0123133213200322-0321111021030211-2213111122013110-3212000211012213"></a>

## Prerequisites — xcsh_aws_vpc_site / 300321010123 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: AWS authentication for deployment

<a id="canonical-3110130232120310-0132300003020010-3130223320112200-1021222203121330-3220303102130103-2321333211000221-3231130300130020-0111000232113213"></a>

## Minimal configuration — xcsh_aws_vpc_site / 300321010123 / 4

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

<a id="canonical-2233112132032221-2111203223222312-0233311032211023-0222001013232302-0011102110321130-1331112221230122-3333023012231033-0020132221101201"></a>

## Root configuration — xcsh_aws_vpc_site / 300321010123 / 5

Required root properties: `aws_region`, `instance_type`, `name`, `namespace`, `ssh_key`. Full root flags and choices appear in the property reference.

<a id="canonical-3212113300032032-0113001012103330-2031203210301312-2101013121333012-0132200003020213-0313221002231133-2010210321331233-0100112322112011"></a>

## Next pages — xcsh_aws_vpc_site / 300321010123 / 6

- [Property reference](../guides/resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [Examples](../guides/resources--aws_vpc_site--examples--group-001.md#canonical-3021222013102130-2213322121010020-2103233032333010-0022211123132231-0000111302010333-2320100300121233-3002100010102022-0012131030023102)
- [Import](../guides/resources--aws_vpc_site--lifecycle--group-001.md#canonical-0332203321003012-2132203223333223-1211203112333223-0131320012011333-1301313111222113-2232300110111300-0203130020231103-0313101200330011)
- [Timeouts](../guides/resources--aws_vpc_site--lifecycle--group-001.md#canonical-3320220302321202-3320311311003031-1002120211012031-2313322022013021-1332123321020320-2233332101212213-0130012231333233-2302213332023311)
