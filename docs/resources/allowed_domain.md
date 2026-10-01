---
page_title: "xcsh_allowed_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain landing."
---

# xcsh_allowed_domain landing

<a id="canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202301223332013-0102331003223302-0213010231330111-2313311012120122-2012300133021222-0333212232320000-3020132303133100-1013212112323120"></a>

## xcsh_allowed_domain — xcsh_allowed_domain / 012132321012 / 2

Breadcrumbs:

- xcsh_allowed_domain

Manages allowed domain in F5 Distributed Cloud.

<a id="canonical-2221130003332010-2200030002030303-3223033011331013-2023033110200000-2200200102013020-0202113001120033-2301210210110002-0322232201302100"></a>

## Prerequisites — xcsh_allowed_domain / 012132321012 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0333212213220012-0112333030113221-0121300202000000-0201120001221133-0201332312010123-2310112131222113-0230032321322231-3110231133231001"></a>

## Minimal configuration — xcsh_allowed_domain / 012132321012 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```

<a id="canonical-0103022201103030-0300201330122121-3202113201230132-1300220303032320-2012320132321011-3211003303301013-0111313123233102-0012232102120021"></a>

## Root configuration — xcsh_allowed_domain / 012132321012 / 5

Required root properties: `allowed_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2232120313002303-0003200211033110-0131132011032130-1023123312102301-2103313012232013-3103132201223312-1211332312301021-0213313322213330"></a>

## Next pages — xcsh_allowed_domain / 012132321012 / 6

- [Property reference](../guides/resources--allowed_domain--reference--group-001.md#canonical-2310331131020212-1322032001003131-0111302203012211-3201023230322111-0103212210213030-2323300331323331-1322110232101022-3321121103131313)
- [Examples](../guides/resources--allowed_domain--examples--group-001.md#canonical-0100103011021332-2023322133032001-2333102002332112-0211100100330332-3023211020232111-2132102123233230-3312033032113322-3030130003223213)
- [Import](../guides/resources--allowed_domain--lifecycle--group-001.md#canonical-2312212133331322-0022331202123303-2200333212310221-0221320020030101-1112000122330103-3121213033203022-3022300322203033-2121132322131320)
- [Timeouts](../guides/resources--allowed_domain--lifecycle--group-001.md#canonical-2123103123222230-2123332130230132-3032013023131032-0121211211321223-2020121021201233-1113103310323310-2231321030232313-3230102223020312)
