---
page_title: "xcsh_cloud_link landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link landing."
---

# xcsh_cloud_link landing

<a id="canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200232332100232-2223323332331332-3032203322002203-3322322303313333-3320320131020303-2011000100030201-2310330123101220-3030002332122030"></a>

## xcsh_cloud_link — xcsh_cloud_link / 131131023011 / 2

Breadcrumbs:

- xcsh_cloud_link

Manages new CloudLink with configured parameters in F5 Distributed Cloud.

<a id="canonical-1233021331101003-0231003113010023-2220021303203231-3123233221020103-1323332103203320-3102313112222212-2232013112232102-0122321000233331"></a>

## Prerequisites — xcsh_cloud_link / 131131023011 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1030132331110032-3232312103030222-0103013200320213-1113211201231000-3103113321223020-0130332111003021-0201001201300133-2223310222010100"></a>

## Minimal configuration — xcsh_cloud_link / 131131023011 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudLink Resource Example
# Manages new CloudLink with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudLink configuration
resource "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}
```

<a id="canonical-1111202101322231-1232021000203232-1202300201130231-3213320332303011-2133023030213333-0003200213312222-3312030002320301-3233031302132333"></a>

## Root configuration — xcsh_cloud_link / 131131023011 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2333023303211231-1133323301123121-1331002313310302-0300333031200200-3301121211231010-1132120011110313-2210312332233200-3301311213321102"></a>

## Next pages — xcsh_cloud_link / 131131023011 / 6

- [Property reference](../guides/resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [Examples](../guides/resources--cloud_link--examples--group-001.md#canonical-3223203122101022-3131311230130300-2110032021232020-1231210112230010-1102111103020220-0113113113231101-1320200100312123-1100101312003100)
- [Import](../guides/resources--cloud_link--lifecycle--group-001.md#canonical-2300023222031022-0032122313133302-2300303231023300-0101221030033112-2223123212030203-1200232132030122-0012022101020313-0011133123203011)
- [Timeouts](../guides/resources--cloud_link--lifecycle--group-001.md#canonical-3333201131131133-1203003313232001-3133120120113201-2010023122100220-0211010021321122-3031311311211023-2222221221103320-0331102333323220)
