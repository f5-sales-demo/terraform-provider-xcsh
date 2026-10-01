---
page_title: "xcsh_nfv_service landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service landing."
---

# xcsh_nfv_service landing

<a id="canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332102132223213-0211031113011202-1311300133012020-0222121310230001-3212213333122313-3010210002000311-3200331133020221-3132022200022122"></a>

## xcsh_nfv_service — xcsh_nfv_service / 220000122333 / 2

Breadcrumbs:

- xcsh_nfv_service

Manages new NFV service with configured parameters in F5 Distributed Cloud.

<a id="canonical-1103012211100133-2033113112001100-0312020103130222-0312130231113122-0123003011222210-3313013010020212-0213230100123320-1331031212200000"></a>

## Prerequisites — xcsh_nfv_service / 220000122333 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2320300013122032-2321000020113003-0333031230101212-3031301102031303-2110100331221003-3320233113102312-2210311021013101-1211331100111210"></a>

## Minimal configuration — xcsh_nfv_service / 220000122333 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NfvService Resource Example
# Manages new NFV service with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NfvService configuration
resource "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}
```

<a id="canonical-3232332131113211-1111322223001121-1321313033330303-2212211221133011-2122330322133312-2322013231121320-3302220121000301-0000231013203301"></a>

## Root configuration — xcsh_nfv_service / 220000122333 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0021032203233310-0111201313032201-1003002220002333-2302020323031112-2212133003332102-2002210332313033-0333002330033110-2232221021020023"></a>

## Next pages — xcsh_nfv_service / 220000122333 / 6

- [Property reference](../guides/resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [Examples](../guides/resources--nfv_service--examples--group-001.md#canonical-3323120001301022-1300222011110323-3230233332130232-2013313320213001-1331011000332003-3011300111102130-2323011313120301-3001310223031303)
- [Import](../guides/resources--nfv_service--lifecycle--group-001.md#canonical-2323020230322330-0003322023110322-1303021130122331-1203333202001210-1120123332203332-1332031031321101-1200331223123102-2321212200001223)
- [Timeouts](../guides/resources--nfv_service--lifecycle--group-001.md#canonical-2022222132300301-1033310210110013-0102133302110111-1222222201010222-1003113322323222-0102301223131222-1310312100112200-2203030121313002)
