---
page_title: "xcsh_ike_phase1_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile landing."
---

# xcsh_ike_phase1_profile landing

<a id="canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122013120200332-3230231111220011-3220232221123120-0312232111112101-1222131303222030-2012032111300113-2230132013221203-3100020131302232"></a>

## xcsh_ike_phase1_profile — xcsh_ike_phase1_profile / 233112112113 / 2

Breadcrumbs:

- xcsh_ike_phase1_profile

Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.
configuration.

<a id="canonical-0312031110333002-0001320003212010-3132103211123032-3201223112231031-0300023110311221-2333331330001213-3331110230200123-1002020202123331"></a>

## Prerequisites — xcsh_ike_phase1_profile / 233112112113 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3211012110100002-3233011110233303-1023010212111122-2313102022322222-3131310010210232-3021202132212332-3112331210212123-2130023211200113"></a>

## Minimal configuration — xcsh_ike_phase1_profile / 233112112113 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase1Profile Resource Example
# Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase1Profile configuration
resource "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  dh_group             = ["example-value"]
  encryption_algos     = ["example-value"]
  prf                  = ["example-value"]
}
```

<a id="canonical-1002130203300102-0331223020201223-0133132033222023-3023213112022100-3231303102010002-1330303210101032-1210303121323230-0002333013203311"></a>

## Root configuration — xcsh_ike_phase1_profile / 233112112113 / 5

Required root properties: `authentication_algos`, `dh_group`, `encryption_algos`, `name`, `namespace`, `prf`. Full root flags and choices appear in the property reference.

<a id="canonical-3102001200032212-0012220003110200-0311120321001023-1021333020211021-3123211000102001-1030131221123132-0012313000230102-3012210120313213"></a>

## Next pages — xcsh_ike_phase1_profile / 233112112113 / 6

- [Property reference](../guides/resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [Examples](../guides/resources--ike_phase1_profile--examples--group-001.md#canonical-1323012213312011-1122023312221233-1132121213232003-1210020322022232-2033112111301220-0323110113203212-2032122212220122-0211112132323330)
- [Import](../guides/resources--ike_phase1_profile--lifecycle--group-001.md#canonical-0322312002001201-1032331300011330-2313310000311320-1233221132030110-0033021131112120-0020032010303301-2000130222332033-1100112310212220)
- [Timeouts](../guides/resources--ike_phase1_profile--lifecycle--group-001.md#canonical-3112110222231100-0233032332133132-0333213211212212-2330122133133121-2012011122021112-2002130131103211-3121020021211300-3032021310322221)
