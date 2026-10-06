---
page_title: "xcsh_sensitive_data_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy examples."
---

# xcsh_sensitive_data_policy examples

<a id="canonical-3130112210311310-3220231221203112-2212010312211031-2031330001130100-3203332022210322-0021312303111210-0231223120331200-2020130310211003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-2321213032331000-2132213131031233-2100233222312121-0113013100321030-1011013210020203-2003033231120211-3313210223013012-0130113122220200)
- Examples

<a id="canonical-1111322000001033-2003310121231030-3331001231120122-2310333322012311-0322021323011303-3222312303232023-1110031022310220-0102310011222032"></a>

### Complete configurations for `xcsh_sensitive_data_policy`

- [Data source](data-sources--sensitive_data_policy--examples--group-001.md#canonical-3000113000010210-3110121213122023-0330003221011312-2213310130020331-2301323102013130-0031021032220203-2302310032030012-3023202303013101): valid configuration.

<a id="canonical-3000113000010210-3110121213122023-0330003221011312-2213310130020331-2301323102013130-0031021032220203-2302310032030012-3023202303013101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_sensitive_data_policy](../data-sources/sensitive_data_policy.md#canonical-2321213032331000-2132213131031233-2100233222312121-0113013100321030-1011013210020203-2003033231120211-3313210223013012-0130113122220200)
- [Examples](data-sources--sensitive_data_policy--examples--group-001.md#canonical-3130112210311310-3220231221203112-2212010312211031-2031330001130100-3203332022210322-0021312303111210-0231223120331200-2020130310211003)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_sensitive_data_policy/data-source.tf`; digest `sha256:cfa788d1ad2d015969f47a23be03161a048878a194ca7cc1ede99b43ed21acb8`.

```terraform
# SensitiveDataPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SensitiveDataPolicy by name
data "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}

output "sensitive_data_policy_id" {
  value = data.xcsh_sensitive_data_policy.example.id
}
```
