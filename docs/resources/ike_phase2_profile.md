---
page_title: "xcsh_ike_phase2_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile landing."
---

# xcsh_ike_phase2_profile landing

<a id="canonical-1201310231032303-0210221131023113-3010000231013011-2100332100031002-2123212020212030-3030213230201210-3322203100220221-3230221201202231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123310112122321-0123133133030212-3220022311201012-0222021112013312-2200320001222313-0113213001111230-1231030201302113-1313330210313111"></a>

## xcsh_ike_phase2_profile — xcsh_ike_phase2_profile / 012301230233 / 2

Breadcrumbs:

- xcsh_ike_phase2_profile

Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.
configuration.

<a id="canonical-2231222333133031-0121120232213112-1233332121132221-3301113131203011-2033032013000333-1131330331202321-0303303212232322-1203303213200313"></a>

## Prerequisites — xcsh_ike_phase2_profile / 012301230233 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1220301102232330-3110002000313131-0231102300113121-0120301030121120-1300102121200313-2311111033013022-2123212313210101-2122010210311320"></a>

## Minimal configuration — xcsh_ike_phase2_profile / 012301230233 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase2Profile Resource Example
# Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase2Profile configuration
resource "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  encryption_algos     = ["example-value"]
}
```

<a id="canonical-3300133110110301-2001002030222202-2113331302032203-1322033133132103-1232311131322220-2311100002133130-0220230230101300-0232110120233122"></a>

## Root configuration — xcsh_ike_phase2_profile / 012301230233 / 5

Required root properties: `authentication_algos`, `encryption_algos`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2300323123101233-1320223220011001-1131130133123231-0320011122011122-0320013330300123-1032303003331323-2033132301212331-0022020102312331"></a>

## Next pages — xcsh_ike_phase2_profile / 012301230233 / 6

- [Property reference](../guides/resources--ike_phase2_profile--reference--group-001.md#canonical-2020002230212010-1313022100100310-3220332211132032-0022322013313122-2101323310333120-1102230003321002-1003312302120232-3300001312010212)
- [Examples](../guides/resources--ike_phase2_profile--examples--group-001.md#canonical-3103222302032203-1330320330012010-0213312233122220-1021012121002330-3303103000313230-1233321202101112-0230001312120213-0030313130000110)
- [Import](../guides/resources--ike_phase2_profile--lifecycle--group-001.md#canonical-2300110122312323-3231330033101312-3121322201131323-0030230000230002-2320230022010111-2030230103013132-0102222020322131-3031130303213331)
- [Timeouts](../guides/resources--ike_phase2_profile--lifecycle--group-001.md#canonical-3303011203233323-2200311123101230-3222321311220312-3130000200110130-0301033121112332-0023233020213230-0032102213223210-1302311202233210)
