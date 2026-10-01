---
page_title: "xcsh_waf_threats examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threats examples."
---

# xcsh_waf_threats examples

<a id="canonical-ec4d28c7443fcfd9835e6bf5700cc562dffd31d6cc15938fae92732a710ee386"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-673092929c80bb0c5797562f45b3d045daaf91946dc762215bc1d4b623957272"></a>

## Examples — Examples / acc89bf97ecd / 2

Breadcrumbs:

- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)
- Examples

<a id="canonical-55ec2ea361b94d78a7a011782425248ca0a81d507634e63c1273188ea1d6af99"></a>

## Complete configurations — Examples / acc89bf97ecd / 3

- [Data source](data-sources--waf_threats--examples--group-001.md#canonical-a2da63262117bb5e53bcd14ac4ca21bf65f84ca71de75339c09d9128241edd29): valid configuration.

<a id="canonical-2ece42c2ba0e2faee2cab67c4a9315416a6a77394d5ada8f03eae26641eefe24"></a>

## Next pages — Examples / acc89bf97ecd / 4

- [Data source](data-sources--waf_threats--examples--group-001.md#canonical-a2da63262117bb5e53bcd14ac4ca21bf65f84ca71de75339c09d9128241edd29)
- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)

<a id="canonical-a2da63262117bb5e53bcd14ac4ca21bf65f84ca71de75339c09d9128241edd29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-540ae1cd752d5f74b917cb6f3099b39ad61b81c3ed7c45f3a319b73740f415c7"></a>

## Data source — Data source / 871a6fa4f1e3 / 2

Breadcrumbs:

- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)
- [Examples](data-sources--waf_threats--examples--group-001.md#canonical-ec4d28c7443fcfd9835e6bf5700cc562dffd31d6cc15938fae92732a710ee386)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_threats/data-source.tf`; digest `sha256:9ac0cdd709cbc9a9ca710dea9bf40dd5b343227de03ab5599937418e7cc2c055`.

```terraform
# WAFThreats DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threats" "example" {
}

output "waf_threats_result" {
  value = data.xcsh_waf_threats.example
}
```

<a id="canonical-e247b7d4d76fe03d5046e069d791529140938ff322cbd8b9d4d7dfd94f3a2ab3"></a>

## Next pages — Data source / 871a6fa4f1e3 / 3

- [Examples](data-sources--waf_threats--examples--group-001.md#canonical-ec4d28c7443fcfd9835e6bf5700cc562dffd31d6cc15938fae92732a710ee386)
- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)
