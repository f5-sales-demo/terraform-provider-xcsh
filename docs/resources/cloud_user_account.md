---
page_title: "xcsh_cloud_user_account landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account landing."
---

# xcsh_cloud_user_account landing

<a id="canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320303033110132-3020203122213330-1333032022303123-2031220321230113-3133313212131121-0231321212230323-1203101021201111-0001310021302112"></a>

## xcsh_cloud_user_account — xcsh_cloud_user_account / 021122021132 / 2

Breadcrumbs:

- xcsh_cloud_user_account

Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create
specifications. configuration.

<a id="canonical-0003132103030202-1322032113331331-0311102000110012-0001030332123110-1201001032032132-1330110220220000-0312223321322222-3131101021301323"></a>

## Prerequisites — xcsh_cloud_user_account / 021122021132 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1230303211321231-2022203200312301-3320302200121112-0022323312211010-0333130012130223-2302233311210302-3210203212200033-1222033213331230"></a>

## Minimal configuration — xcsh_cloud_user_account / 021122021132 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudUserAccount Resource Example
# Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create specifications.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudUserAccount configuration
resource "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}
```

<a id="canonical-2312321021033010-3100121232132213-1023112110230102-2001202102321121-3312301233233032-0003221120101230-2302112003011033-1312112310003013"></a>

## Root configuration — xcsh_cloud_user_account / 021122021132 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2113002012222301-3101132003223003-2131100210213302-3001300302131132-0232031011210211-1211332303311301-3132201003223103-1223132223111210"></a>

## Next pages — xcsh_cloud_user_account / 021122021132 / 6

- [Property reference](../guides/resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [Examples](../guides/resources--cloud_user_account--examples--group-001.md#canonical-1201010100032201-3121233112123000-2101330232023112-1331201022311202-1331232332031200-3132300021332320-3012231020101202-0111122311221301)
- [Import](../guides/resources--cloud_user_account--lifecycle--group-001.md#canonical-3201120320023322-0133231032231320-1222233130132021-2211030001331102-3101023332001003-2022030322301001-2200000132022010-1133230210131223)
- [Timeouts](../guides/resources--cloud_user_account--lifecycle--group-001.md#canonical-3021311123021021-1232011023331221-1133112020230020-0320122303201100-1102130212021200-3300003133022100-0301203212200310-2320223000311203)
