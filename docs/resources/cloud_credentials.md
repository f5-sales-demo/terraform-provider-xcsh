---
page_title: "xcsh_cloud_credentials landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials landing."
---

# xcsh_cloud_credentials landing

<a id="canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210120202233232-1033332212033000-3202232122020232-2121013030010233-3113120222103321-3120123103003110-0110301200032132-0212031313232231"></a>

## xcsh_cloud_credentials — xcsh_cloud_credentials / 232012000010 / 2

Breadcrumbs:

- xcsh_cloud_credentials

Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud\_credentials
object. configuration.

<a id="canonical-2231103010121010-1013212011323220-1123122021121213-0031112300131200-2133300232311212-3133302012302032-3310302010112012-0121301111010233"></a>

## Prerequisites — xcsh_cloud_credentials / 232012000010 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1131103232233103-3002332211221230-2002200103303021-1310123011122323-3221023030003112-0120233222231101-1201223320110110-3003330011133011"></a>

## Minimal configuration — xcsh_cloud_credentials / 232012000010 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudCredentials Resource Example
# Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudCredentials configuration
resource "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}
```

<a id="canonical-1030230221303100-2001213230320230-3322010002112001-1123122003332000-1302103113221133-3332310100032213-1012200330021200-3212001333131330"></a>

## Root configuration — xcsh_cloud_credentials / 232012000010 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0121223001033113-3100322132221122-2133220023302310-1313212133022331-2210001233203102-2101030212133122-0230121030310232-0133230000032203"></a>

## Next pages — xcsh_cloud_credentials / 232012000010 / 6

- [Property reference](../guides/resources--cloud_credentials--reference--group-001.md#canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200)
- [Examples](../guides/resources--cloud_credentials--examples--group-001.md#canonical-3120020011033110-3030221321302133-3001201311303003-0200222331003321-3311303322032013-0001302102303201-3100011302032313-0333323311023021)
- [Import](../guides/resources--cloud_credentials--lifecycle--group-001.md#canonical-1100112101112012-0311100211012110-3221001133101211-2122001021022023-0232233203111021-3323211210312032-0320303120222020-0322212111313330)
- [Timeouts](../guides/resources--cloud_credentials--lifecycle--group-001.md#canonical-1011021231212103-3020302031233101-1131021320030301-0120112330301012-3303122231013323-0222301123003302-2321033022031131-0023313013030312)
