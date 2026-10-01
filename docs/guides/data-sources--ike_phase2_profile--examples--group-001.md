---
page_title: "xcsh_ike_phase2_profile examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile examples."
---

# xcsh_ike_phase2_profile examples

<a id="canonical-59588c09ef53762a3a76590a75aaee307169f6d19df7a933e040c76d7162e176"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7816b037380ba32864dee4c02a20038a8b43559ae87162516fcfe569ba1643aa"></a>

## Examples — Examples / 81a80226b5de / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- Examples

<a id="canonical-89dc384095d88e1b0ef7b51f7d440c66be4dd1712b758e00252b6300ad7661d0"></a>

## Complete configurations — Examples / 81a80226b5de / 3

- [Data source](data-sources--ike_phase2_profile--examples--group-001.md#canonical-4b4f665797567c4c0fc68adb801abd62d7632aefe9fe8e207afb6c1a52cf1ecd): valid configuration.

<a id="canonical-be351a4c06b3ec02cdd2fdd5ce82832990683fa1f8fa323afdce8ce152b7ea38"></a>

## Next pages — Examples / 81a80226b5de / 4

- [Data source](data-sources--ike_phase2_profile--examples--group-001.md#canonical-4b4f665797567c4c0fc68adb801abd62d7632aefe9fe8e207afb6c1a52cf1ecd)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)

<a id="canonical-4b4f665797567c4c0fc68adb801abd62d7632aefe9fe8e207afb6c1a52cf1ecd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd2f4a2dc450765d65efac3cda32e851815402fedcbc07d9b6081517d9c05c45"></a>

## Data source — Data source / 492718915f14 / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- [Examples](data-sources--ike_phase2_profile--examples--group-001.md#canonical-59588c09ef53762a3a76590a75aaee307169f6d19df7a933e040c76d7162e176)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike_phase2_profile/data-source.tf`; digest `sha256:a5a7c67365793986b4f75d4d921baeb441a3835387aa16c35286802039150551`.

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

<a id="canonical-81eb94b57b96e386f96017eebe15e791d94e52b7b9792b0916b202f57f448231"></a>

## Next pages — Data source / 492718915f14 / 3

- [Examples](data-sources--ike_phase2_profile--examples--group-001.md#canonical-59588c09ef53762a3a76590a75aaee307169f6d19df7a933e040c76d7162e176)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
