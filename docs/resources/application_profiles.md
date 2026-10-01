---
page_title: "xcsh_application_profiles landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles landing."
---

# xcsh_application_profiles landing

<a id="canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012211133003033-3033011113320033-2300203303133323-2021301213021212-3320210321011012-2303332023002012-1233200230101322-2020012113333030"></a>

## xcsh_application_profiles — xcsh_application_profiles / 002321230033 / 2

Breadcrumbs:

- xcsh_application_profiles

Manages Application Profiles in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-1203110001310313-0230001113222021-2313320010233020-2313001030302230-2003333321202313-3321210301032113-1210113322031223-0003020201103132"></a>

## Prerequisites — xcsh_application_profiles / 002321230033 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0022222323131201-3333002100231112-0222221220023311-0033310033320000-0232021210130123-2101311033112111-3212013122312310-1223311222122001"></a>

## Minimal configuration — xcsh_application_profiles / 002321230033 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ApplicationProfiles Resource Example
# Manages Application Profiles in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ApplicationProfiles configuration
resource "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}
```

<a id="canonical-0102331000131002-3301111202100320-2003021213223030-2122311230200123-1320100320122212-0023332021303221-3300002123102010-1320112131201121"></a>

## Root configuration — xcsh_application_profiles / 002321230033 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0001212323203201-3100300222111330-3102013212012333-0123013221210112-0012211311233100-2302323202103011-1122030223023323-1011011233213331"></a>

## Next pages — xcsh_application_profiles / 002321230033 / 6

- [Property reference](../guides/resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [Examples](../guides/resources--application_profiles--examples--group-001.md#canonical-0121211302010133-1133122330330220-2313210032111110-0221132112232120-2100033100002131-2020331200001231-2002232202120301-1003301312102122)
- [Import](../guides/resources--application_profiles--lifecycle--group-001.md#canonical-0303202323020332-2131113030101331-2210312303123203-0112030310133301-2020311011101330-0000300002100100-2102213313200223-2031030210012030)
- [Timeouts](../guides/resources--application_profiles--lifecycle--group-001.md#canonical-0103003302231111-1010123021120230-3010130121023123-3010120012130132-2330103312203232-1030030001013010-2010133221131321-2000002112330031)
