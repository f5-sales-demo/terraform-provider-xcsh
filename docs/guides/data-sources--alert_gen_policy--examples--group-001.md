---
page_title: "xcsh_alert_gen_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy examples."
---

# xcsh_alert_gen_policy examples

<a id="canonical-3101133110101332-1212313230200233-0212133030031121-3000033232011331-3002320201122210-1121300201310231-3102220301010033-0300333030203210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-1212223311231031-2303130101102331-2121323302132120-3001003212230233-2233103321211211-1300022131031130-1101122301003332-3330023301222220)
- Examples

<a id="canonical-2200021100022021-2010103322013301-2131233000002213-3130211003310231-3121332013333102-2233101331101223-0030121222010101-0112010333020233"></a>

### Complete configurations for `xcsh_alert_gen_policy`

- [Data source](data-sources--alert_gen_policy--examples--group-001.md#canonical-1103023110200103-2022110230331111-2121301210003300-0333200010031033-2013330110211221-2313110133210132-2012222310222220-2032213300022312): valid configuration.

<a id="canonical-1103023110200103-2022110230331111-2121301210003300-0333200010031033-2013330110211221-2313110133210132-2012222310222220-2032213300022312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-1212223311231031-2303130101102331-2121323302132120-3001003212230233-2233103321211211-1300022131031130-1101122301003332-3330023301222220)
- [Examples](data-sources--alert_gen_policy--examples--group-001.md#canonical-3101133110101332-1212313230200233-0212133030031121-3000033232011331-3002320201122210-1121300201310231-3102220301010033-0300333030203210)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_gen_policy/data-source.tf`; digest `sha256:349cbe33e45f53c629e7ce220830862d4e4d7dd9b1fd6b5c5a42c3556d1a1bed`.

```terraform
# AlertGenPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertGenPolicy by name
data "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}

output "alert_gen_policy_id" {
  value = data.xcsh_alert_gen_policy.example.id
}
```
