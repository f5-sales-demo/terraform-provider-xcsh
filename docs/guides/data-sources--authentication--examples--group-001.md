---
page_title: "xcsh_authentication examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication examples."
---

# xcsh_authentication examples

<a id="canonical-3311100333131130-3203031113312032-0213122233111110-1002111313031302-0000011213332103-0001333330310320-1333121200120313-0112111020223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- Examples

<a id="canonical-2220310022031212-0200211303303203-0222331320331130-3233223330100232-1223000102112231-0322013311111303-0201321310033033-1210211022022233"></a>

### Complete configurations for `xcsh_authentication`

- [Data source](data-sources--authentication--examples--group-001.md#canonical-0123220332103303-1322311321030321-3231001223333000-0323201230010231-0332330330200220-0233121310120303-3123121322311302-3301120023202033): valid configuration.

<a id="canonical-0123220332103303-1322311321030321-3231001223333000-0323201230010231-0332330330200220-0233121310120303-3123121322311302-3301120023202033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103)
- [Examples](data-sources--authentication--examples--group-001.md#canonical-3311100333131130-3203031113312032-0213122233111110-1002111313031302-0000011213332103-0001333330310320-1333121200120313-0112111020223000)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authentication/data-source.tf`; digest `sha256:fd8b227ee8d771e5d805c51da258f79cea8ff5031e710c318c14f6217d6efd78`.

```terraform
# Authentication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Authentication by name
data "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}

output "authentication_id" {
  value = data.xcsh_authentication.example.id
}
```
