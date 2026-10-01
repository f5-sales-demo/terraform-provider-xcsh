---
page_title: "xcsh_ike_phase1_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile landing."
---

# xcsh_ike_phase1_profile landing

<a id="canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122103010121000-3321003320031231-1121200030121012-1111011120212021-3202023013133131-3030201100332020-1201111021232220-3211130020301220"></a>

## xcsh_ike_phase1_profile — xcsh_ike_phase1_profile / 032223120233 / 2

Breadcrumbs:

- xcsh_ike_phase1_profile

Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.
configuration.

<a id="canonical-1310133303221200-3131331301310021-3021302203203103-1323320130302030-1221221033212323-1110220203131011-2221312213233020-3113130123223002"></a>

## Prerequisites — xcsh_ike_phase1_profile / 032223120233 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2111023030110011-0333232323311222-3030330232033131-2320023332211230-2301212031302003-0212211311320301-2312230000211313-3033000220313312"></a>

## Minimal configuration — xcsh_ike_phase1_profile / 032223120233 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase1Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase1Profile by name
data "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"
}

output "ike_phase1_profile_id" {
  value = data.xcsh_ike_phase1_profile.example.id
}
```

<a id="canonical-3223102131122023-0002200213122212-2030321123212102-3223132032323031-1331010011233320-3111100322300110-0332321103010111-2322002000003110"></a>

## Root configuration — xcsh_ike_phase1_profile / 032223120233 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3110312023221132-0012203013032330-2100213230231011-2312221123111220-3111313120023011-2302313202223212-2013203221332213-0122003232313133"></a>

## Next pages — xcsh_ike_phase1_profile / 032223120233 / 6

- [Property reference](../guides/data-sources--ike_phase1_profile--reference--group-001.md#canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123)
- [Examples](../guides/data-sources--ike_phase1_profile--examples--group-001.md#canonical-3131103321231102-0203133013231322-3031033002100130-2320212022331101-0123010211103000-3322032303232021-1321012003023000-2103110133313222)
