---
page_title: "xcsh_trusted_ca_list landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list landing."
---

# xcsh_trusted_ca_list landing

<a id="canonical-0222000113113203-2212232323102203-2101230202220010-2030322030013021-2133220023123121-2113200100110031-3003002313103330-2200302323113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220210000213111-2021032300333210-1232130110303312-2330023303112222-0320122301113233-3232101020122003-3222002213212001-2221011210020100"></a>

## xcsh_trusted_ca_list — xcsh_trusted_ca_list / 333212322031 / 2

Breadcrumbs:

- xcsh_trusted_ca_list

Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list
management.

<a id="canonical-2113010021322301-3022100302021221-1133120032101312-3320001002123112-3003011030203212-0111133321001230-2023103101002102-2103201113231310"></a>

## Prerequisites — xcsh_trusted_ca_list / 333212322031 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1310213330322223-3323111132123111-1311001011101012-1120120033103102-3310330213321232-3031320023331101-3100102220002233-1130331230020121"></a>

## Minimal configuration — xcsh_trusted_ca_list / 333212322031 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```

<a id="canonical-2302203031332011-1000210203113321-0302023313110022-1213303122203323-1113030001013021-2200230122300020-2122232202100221-1330320013223212"></a>

## Root configuration — xcsh_trusted_ca_list / 333212322031 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0212212303312001-3210001222300303-1121103111022102-1103013203332100-1202323222001032-2320302311230300-0200023023202012-0022033323211002"></a>

## Next pages — xcsh_trusted_ca_list / 333212322031 / 6

- [Property reference](../guides/resources--trusted_ca_list--reference--group-001.md#canonical-2011100021221213-2030030021232013-1113200112210011-3323103123123311-0202132011000133-0020212102203021-2031323023203030-2231003231000112)
- [Examples](../guides/resources--trusted_ca_list--examples--group-001.md#canonical-1313012232330110-3010303033130202-2023002102131213-1311200033113301-3030011313223302-0010020201221102-3110312100122103-0133111300230220)
- [Import](../guides/resources--trusted_ca_list--lifecycle--group-001.md#canonical-2100031021303333-2312102102232223-1130221210320233-0033003021231131-1111033121121320-1122131302001000-3121332113222202-3222212211031113)
- [Timeouts](../guides/resources--trusted_ca_list--lifecycle--group-001.md#canonical-2133211020202300-3231300233020200-3023311303321120-1231013100223101-2322103231000131-0330022131222113-3120203022101332-3210022211131300)
