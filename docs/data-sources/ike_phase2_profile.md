---
page_title: "xcsh_ike_phase2_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile landing."
---

# xcsh_ike_phase2_profile landing

<a id="canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222202011120133-2201213003123002-0333231233023023-3021011302002233-3011001221101323-2012212002032302-3331113020011312-2210011022203131"></a>

## xcsh_ike_phase2_profile — xcsh_ike_phase2_profile / 201232122032 / 2

Breadcrumbs:

- xcsh_ike_phase2_profile

Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.
configuration.

<a id="canonical-3302211233130013-1011320313131101-0232313130130230-3033130222332222-0231033323102232-2012223231030133-3132222120033210-2012233030012220"></a>

## Prerequisites — xcsh_ike_phase2_profile / 201232122032 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3020231230123000-1213313220300210-1302321212321122-1012333020230310-3222033320123331-1222111011203212-3132010313303223-0112111120133320"></a>

## Minimal configuration — xcsh_ike_phase2_profile / 201232122032 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase2Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase2Profile by name
data "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"
}

output "ike_phase2_profile_id" {
  value = data.xcsh_ike_phase2_profile.example.id
}
```

<a id="canonical-1301230332333003-3111100113230123-0013321100200330-2012111210030010-2300130213313111-2113112230302313-1310202301102313-3102322102002212"></a>

## Root configuration — xcsh_ike_phase2_profile / 201232122032 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3101300223300320-0020332000022330-2020011010212102-0000133001310021-3303013313320111-1002133123033221-1003321011300302-0201032013233223"></a>

## Next pages — xcsh_ike_phase2_profile / 201232122032 / 6

- [Property reference](../guides/data-sources--ike_phase2_profile--reference--group-001.md#canonical-1221020020223301-3321201330233132-0012122311312233-0211331333321112-1210022031000211-2123121212122103-1000331210101312-1310102310302122)
- [Examples](../guides/data-sources--ike_phase2_profile--examples--group-001.md#canonical-1121112020300021-3233110313120222-0322131211210022-1311222232320300-1301122133123101-2131331322210303-3200100030131231-1301120232011312)
