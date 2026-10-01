---
page_title: "xcsh_forward_proxy_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy landing."
---

# xcsh_forward_proxy_policy landing

<a id="canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331201223030010-2312022111003221-1233230033101212-1020123332302133-3020120020002033-2012202201311103-1101030333101030-1032011311210220"></a>

## xcsh_forward_proxy_policy — xcsh_forward_proxy_policy / 002003121110 / 2

Breadcrumbs:

- xcsh_forward_proxy_policy

Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy
specification. configuration.

<a id="canonical-3131001231002213-1120023211223111-2102221103113032-3310020123000221-1323312221213000-2023323011220311-2000121102013123-0312300121323131"></a>

## Prerequisites — xcsh_forward_proxy_policy / 002003121110 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-1320230111103022-3031112320220223-2231220220021222-3013311223221311-1032303123302033-1202200300212020-1320010233332311-2210333011323202"></a>

## Minimal configuration — xcsh_forward_proxy_policy / 002003121110 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardProxyPolicy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardProxyPolicy configuration
resource "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}
```

<a id="canonical-2102323332323301-0332023322113110-0221320321121132-0113201222022311-0002320023210032-2100132133022120-1213303002202202-2102210211103332"></a>

## Root configuration — xcsh_forward_proxy_policy / 002003121110 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3123223302033320-0133230110231203-1102222031131202-1131300302212100-1230222122220012-1031032112233103-3111213230213200-1032331012233010"></a>

## Next pages — xcsh_forward_proxy_policy / 002003121110 / 6

- [Property reference](../guides/resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [Examples](../guides/resources--forward_proxy_policy--examples--group-001.md#canonical-2030102031112111-2133212023230213-0100322013131112-3131222303110302-1133313322002110-2320102233123123-2003020020003131-1100313123230120)
- [Import](../guides/resources--forward_proxy_policy--lifecycle--group-001.md#canonical-2020023102203003-2221121113230220-3010301330103023-0132312311022201-3310031220123101-1133300010033102-3001011102301121-1303200102312102)
- [Timeouts](../guides/resources--forward_proxy_policy--lifecycle--group-001.md#canonical-3202312002230100-2210013002033121-1213303313001213-1102030023010100-0122222313223230-3322033332102320-3130121032001021-0332300321013330)
