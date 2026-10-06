---
page_title: "xcsh_route examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route examples."
---

# xcsh_route examples

<a id="canonical-2103210131020030-3332100111120032-0121012330312101-3320321023201002-0230023100223111-1033211232100203-1122111100333010-3310033010101013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- Examples

<a id="canonical-1303032321122021-2302231301121203-3031000323120031-3122011132310322-1233113033221023-1320111001023021-1000310120031212-2330000012210300"></a>

### Complete configurations for `xcsh_route`

- [Data source](data-sources--route--examples--group-001.md#canonical-2111001110122330-3313200121202320-1032123300123212-3030231013303020-2311021122201333-1201002113331113-2200322112201000-2203220100210013): valid configuration.

<a id="canonical-2111001110122330-3313200121202320-1032123300123212-3030231013303020-2311021122201333-1201002113331113-2200322112201000-2203220100210013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Examples](data-sources--route--examples--group-001.md#canonical-2103210131020030-3332100111120032-0121012330312101-3320321023201002-0230023100223111-1033211232100203-1122111100333010-3310033010101013)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_route/data-source.tf`; digest `sha256:2077e687d5b6480f640ada5b5b9db38081cc538c7b97d40a3dfe626debd65a3f`.

```terraform
# Route Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Route by name
data "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}

output "route_id" {
  value = data.xcsh_route.example.id
}
```
