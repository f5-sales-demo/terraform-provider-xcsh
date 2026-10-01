---
page_title: "xcsh_ike1 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 landing."
---

# xcsh_ike1 landing

<a id="canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232110012311323-0332110310101000-1003221122001321-2212103200203203-3213022101122200-2212201000130222-2011003031303022-2113102003231001"></a>

## xcsh_ike1 — xcsh_ike1 / 233232002001 / 2

Breadcrumbs:

- xcsh_ike1

Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.

<a id="canonical-0101112111213332-3322000100331032-3032200213211201-3100200232330000-1210021000020321-2320111201111111-1323023122201302-0000102102031212"></a>

## Prerequisites — xcsh_ike1 / 233232002001 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2130003213133001-3110003012020233-2310103332112333-1032031120020021-1103120302012232-2103223012311302-1200310101330001-1232130203200130"></a>

## Minimal configuration — xcsh_ike1 / 233232002001 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```

<a id="canonical-0032122322323222-0110101131233121-0232110301112101-0333100311303031-1131230112200310-2232130102333233-2020223133312031-0331332322032000"></a>

## Root configuration — xcsh_ike1 / 233232002001 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3033331202210031-0303113033313103-2300232020213112-2121310333012021-3331200301010331-0101113313101310-1132201120131323-3021033310121323"></a>

## Next pages — xcsh_ike1 / 233232002001 / 6

- [Property reference](../guides/resources--ike1--reference--group-001.md#canonical-0231322323021312-2223311002332013-3200002221122100-2200120310330232-3230312122213012-1111030312321102-2331002032223020-1222113231322110)
- [Examples](../guides/resources--ike1--examples--group-001.md#canonical-2221002230320112-3330322031333012-1003220211230121-0013321303220131-0320313301233213-2101130002222202-3222220033130000-1222323203032201)
- [Import](../guides/resources--ike1--lifecycle--group-001.md#canonical-0002211210321023-0122000321130313-0022200202111332-1322032202323010-0301322000002202-0111113320133100-1100033011001013-2330132103302121)
- [Timeouts](../guides/resources--ike1--lifecycle--group-001.md#canonical-0010113102323002-3310312132111312-2000310000132222-0031201123032302-3201011012302233-3123113232301030-1322213203033001-2020321310233130)
