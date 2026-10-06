---
page_title: "xcsh_waf_exclusion_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy examples."
---

# xcsh_waf_exclusion_policy examples

<a id="canonical-3020013233303311-2211022210333012-2212312020331032-2103020301023332-1022211031311300-2010310120031311-1033232321032333-1210123000222333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- Examples

<a id="canonical-2232132132210133-1033332100111103-2023220031210011-3310021001222122-1000331030022201-3112032331000212-3121200033102132-2211102031220011"></a>

### Complete configurations for `xcsh_waf_exclusion_policy`

- [Data source](data-sources--waf_exclusion_policy--examples--group-001.md#canonical-3031113130100123-3132321120330101-0202323213023301-0031311220212222-0031210002221011-2223102101011310-3112010201213110-3131023221022322): valid configuration.

<a id="canonical-3031113130100123-3132321120330101-0202323213023301-0031311220212222-0031210002221011-2223102101011310-3112010201213110-3131023221022322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-0000310001313101-0301310213020122-0312330100221021-0002333333230300-3311011203120012-2132101200213211-1222123331212002-2103001201121203)
- [Examples](data-sources--waf_exclusion_policy--examples--group-001.md#canonical-3020013233303311-2211022210333012-2212312020331032-2103020301023332-1022211031311300-2010310120031311-1033232321032333-1210123000222333)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_exclusion_policy/data-source.tf`; digest `sha256:e7c1c87b69af9a5626274b0968e5e1ac1658731a6d77764810ad0bd8f60e1d5e`.

```terraform
# WAFExclusionPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WAFExclusionPolicy by name
data "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}

output "waf_exclusion_policy_id" {
  value = data.xcsh_waf_exclusion_policy.example.id
}
```
