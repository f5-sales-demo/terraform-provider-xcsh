---
page_title: "xcsh_advertise_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy landing."
---

# xcsh_advertise_policy landing

<a id="canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231201013023103-2201210033300221-3032002131102120-0103111132131323-2202132022323000-0013330323210023-2300032110331210-0102023122311120"></a>

## xcsh_advertise_policy — xcsh_advertise_policy / 112032010100 / 2

Breadcrumbs:

- xcsh_advertise_policy

Manages a Advertise Policy resource in F5 Distributed Cloud for advertise\_policy object controls
how and where a service represented by a given virtual\_host object is advertised to consumers.
configuration.

<a id="canonical-1220021302200200-0101311231112103-0132032112100022-2132022132222111-0232332223210022-3312010223333332-1020011112322313-2203200100133230"></a>

## Prerequisites — xcsh_advertise_policy / 112032010100 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2003130022202300-3112002122212000-1011023023101102-0202120232200032-0011030200233122-2322131210311132-2111021000212021-3203221013021033"></a>

## Minimal configuration — xcsh_advertise_policy / 112032010100 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AdvertisePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AdvertisePolicy by name
data "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}

output "advertise_policy_id" {
  value = data.xcsh_advertise_policy.example.id
}
```

<a id="canonical-3002012201221233-3012201203211332-3033031320020321-1031102323310300-1331201212203320-0200121330132023-0111110022203002-0030030221123133"></a>

## Root configuration — xcsh_advertise_policy / 112032010100 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3213130301031332-0022332003313002-0200110330320213-3113312213213020-1212101122230222-3103311230000000-1022031201031030-2230303230232010"></a>

## Next pages — xcsh_advertise_policy / 112032010100 / 6

- [Property reference](../guides/data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [Examples](../guides/data-sources--advertise_policy--examples--group-001.md#canonical-3001020230333331-2030223330301300-0330111210101131-2021213112120012-0220120323212011-1312001132020230-2100222333330312-3132321031200111)
