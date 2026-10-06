---
page_title: "xcsh_user_identification examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification examples."
---

# xcsh_user_identification examples

<a id="canonical-2332023102210113-1231211101202322-2303231032012131-0320023202113003-3233222222302313-1020333202322011-0000013021301000-1003303030200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- Examples

<a id="canonical-1323133232033311-0010123303011311-0001003013213210-3103211322203130-1002320221002211-2300032033231012-0222002130222333-3112331102122301"></a>

### Complete configurations for `xcsh_user_identification`

- [Data source](data-sources--user_identification--examples--group-001.md#canonical-1113012020110212-3203110103303312-2010033223110313-3013231211121132-3210300300233322-3300021101211023-1120313212023220-0333121030121331): valid configuration.

<a id="canonical-1113012020110212-3203110103303312-2010033223110313-3013231211121132-3210300300233322-3300021101211023-1120313212023220-0333121030121331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321)
- [Examples](data-sources--user_identification--examples--group-001.md#canonical-2332023102210113-1231211101202322-2303231032012131-0320023202113003-3233222222302313-1020333202322011-0000013021301000-1003303030200123)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_user_identification/data-source.tf`; digest `sha256:c30020dfd6aecd49ffb0273704b9b9e1c85f494c0e9e28cd883c479b6b6234a3`.

```terraform
# UserIdentification Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UserIdentification by name
data "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}

output "user_identification_id" {
  value = data.xcsh_user_identification.example.id
}
```
