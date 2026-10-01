---
page_title: "xcsh_nat_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy landing."
---

# xcsh_nat_policy landing

<a id="canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322133320110031-3201011000003130-0023331322201102-3013322220213321-3223333021002211-0130112102121202-1332013013112201-1103113201013212"></a>

## xcsh_nat_policy — xcsh_nat_policy / 102111121212 / 2

Breadcrumbs:

- xcsh_nat_policy

Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures
nat policy with multiple rules,. configuration.

<a id="canonical-0112120221002332-2002210230133101-1233211213013122-0112233320002210-3133200020303111-1210200312031120-1131331200010010-1321122013130311"></a>

## Prerequisites — xcsh_nat_policy / 102111121212 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2221022312110302-3210301120323200-3323231220101003-1321022103121022-0123001310300122-0220110121322220-2211231300030322-1313013033000221"></a>

## Minimal configuration — xcsh_nat_policy / 102111121212 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NATPolicy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NATPolicy configuration
resource "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}
```

<a id="canonical-1133131003101200-3000033220110321-0202301201031020-0030233133112113-2132301333330212-1220000221100002-2330201222003131-0023033002202100"></a>

## Root configuration — xcsh_nat_policy / 102111121212 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3103003332012012-2210213333132031-1003332320121030-0032132331111200-2003123322211223-2003011223321201-2000022122312013-3302000312022322"></a>

## Next pages — xcsh_nat_policy / 102111121212 / 6

- [Property reference](../guides/resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [Examples](../guides/resources--nat_policy--examples--group-001.md#canonical-2133002012221220-1223212000131030-1323323013232213-1331030310331231-3031032213110301-1030200010103130-1213033032023003-3130112123011213)
- [Import](../guides/resources--nat_policy--lifecycle--group-001.md#canonical-1123202213313131-0010333323011313-3030022210010321-1312021003132132-0312003201012133-2322333120210231-0330001112212200-1212302033221113)
- [Timeouts](../guides/resources--nat_policy--lifecycle--group-001.md#canonical-2221120221333101-0031032101111121-0112031211102231-3210231223003322-0133302130213012-3001010031023300-2101202130130200-3333120220130223)
