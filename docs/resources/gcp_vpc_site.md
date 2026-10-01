---
page_title: "xcsh_gcp_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site landing."
---

# xcsh_gcp_vpc_site landing

<a id="canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003202101322120-3022002331201312-0331210230032230-1021130333022233-0001103333111111-1330013303000001-0321030303303210-3320133013300033"></a>

## xcsh_gcp_vpc_site — xcsh_gcp_vpc_site / 222333122320 / 2

Breadcrumbs:

- xcsh_gcp_vpc_site

Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud
VPC environments.

<a id="canonical-0231303232231120-1331022320113230-0331302332322303-3221012320110023-2333322320211320-0110102133022222-3332213303301102-3110022323310130"></a>

## Prerequisites — xcsh_gcp_vpc_site / 222333122320 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: GCP authentication for deployment

<a id="canonical-1312100333321032-2233110113300133-3331002220220311-1200302213132230-0200113011002010-3211030030100110-0230103130233132-3001212311121021"></a>

## Minimal configuration — xcsh_gcp_vpc_site / 222333122320 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GCPVPCSite Resource Example
# Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GCPVPCSite configuration
resource "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"

  gcp_region    = "example-value"
  instance_type = "example-value"
  ssh_key       = "example-value"
}
```

<a id="canonical-3021201310313320-1230102213131032-2020310213321120-0320223133120330-3310020130133123-2330332312113121-3102023032110323-3120123223303233"></a>

## Root configuration — xcsh_gcp_vpc_site / 222333122320 / 5

Required root properties: `gcp_region`, `instance_type`, `name`, `namespace`, `ssh_key`. Full root flags and choices appear in the property reference.

<a id="canonical-0111010323101122-1203202002131303-1301130332130001-1031132001223210-0021200203302002-1111102323102203-0330323102111011-1103002012313110"></a>

## Next pages — xcsh_gcp_vpc_site / 222333122320 / 6

- [Property reference](../guides/resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [Examples](../guides/resources--gcp_vpc_site--examples--group-001.md#canonical-2130032311011020-2313232302331000-1323131223333302-2021301313123311-3212333323210202-3210003231323310-2331000223320022-1131312223331202)
- [Import](../guides/resources--gcp_vpc_site--lifecycle--group-001.md#canonical-1013101201001231-3023300112221101-3303121320311033-1130112222231020-1130202201320022-0311221223222322-1130012013132020-3120221200333112)
- [Timeouts](../guides/resources--gcp_vpc_site--lifecycle--group-001.md#canonical-1230020002210331-2002123033023113-3211202330330122-3200320123101132-1232001211101333-1312232201021131-3131120332103021-2210220223232222)
