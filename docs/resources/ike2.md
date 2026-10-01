---
page_title: "xcsh_ike2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 landing."
---

# xcsh_ike2 landing

<a id="canonical-3103220130122111-0023320210302011-3122102311032022-0223033112130223-1201210032310333-1323122321230203-0200202200211032-1120020123001232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011103032113112-2000133323311031-2210230311203033-1221221312013223-2002013213231312-2112130023301212-1120030100203112-2200003030020000"></a>

## xcsh_ike2 — xcsh_ike2 / 111112220111 / 2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

<a id="canonical-2221301002222033-1303201013302103-1201323122113003-3132122310001113-3020100312322013-0032011310332210-0001203021022203-2031033013333120"></a>

## Prerequisites — xcsh_ike2 / 111112220111 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3211330002331303-3011200112131223-3332320123100223-0303230031300301-3300211221011232-3103331111211230-0201333010302131-3202301301332220"></a>

## Minimal configuration — xcsh_ike2 / 111112220111 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike2 configuration
resource "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}
```

<a id="canonical-1313030010120202-1133113200033323-0123022303233123-0032220012310010-0233031112321311-1321130022021001-0011120113113020-0222131022001033"></a>

## Root configuration — xcsh_ike2 / 111112220111 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3000330133222122-1033211303323311-3033131020333113-0231132133200223-2332230322130111-2200133321131212-0312032213102132-0221330103013023"></a>

## Next pages — xcsh_ike2 / 111112220111 / 6

- [Property reference](../guides/resources--ike2--reference--group-001.md#canonical-0121100022303232-1202012003120020-0122303122032001-2001323012112121-3102220103221003-0032210001132223-1221211112013122-1321230021230203)
- [Examples](../guides/resources--ike2--examples--group-001.md#canonical-0111201230203323-2121321302031333-0133222302222032-3320123331001032-3023312131303101-0131233020020330-3233122333320210-0213113213030103)
- [Import](../guides/resources--ike2--lifecycle--group-001.md#canonical-3103333320220320-3210330102222123-1313013231312013-2023232031020333-3010223230203201-1023203012130221-0032332213032112-0120213001112222)
- [Timeouts](../guides/resources--ike2--lifecycle--group-001.md#canonical-2020120020322003-3220331301011200-2112200331122133-0300311223302123-2203132302000300-2213333113200221-3030120212202202-1311321103111330)
