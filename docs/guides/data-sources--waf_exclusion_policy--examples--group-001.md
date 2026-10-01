---
page_title: "xcsh_waf_exclusion_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy examples."
---

# xcsh_waf_exclusion_policy examples

<a id="canonical-c81efcf5a52a4fc6a6d88f4e932312fe4a94dd7084d183754fbb93bf646c0abf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae79e91f4ff905538ba0d905f4241a9a40f4c2a1d63bd026d980f49ea548da05"></a>

## Examples — Examples / aa2617a32f59 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- Examples

<a id="canonical-9105b7a7231bbda767f4242cc25d6ab797857ca46dedbad2aae5b2a2983ef551"></a>

## Complete configurations — Examples / aa2617a32f59 / 3

- [Data source](data-sources--waf_exclusion_policy--examples--group-001.md#canonical-cd5dc41bdee58f1122ee72f10dd689aa0d902a45ab491174d61219d4dd2e92ba): valid configuration.

<a id="canonical-1037337a73b704e8d80b4cae3cde63d04323e7e79f48b7248e1d1baa03b11c67"></a>

## Next pages — Examples / aa2617a32f59 / 4

- [Data source](data-sources--waf_exclusion_policy--examples--group-001.md#canonical-cd5dc41bdee58f1122ee72f10dd689aa0d902a45ab491174d61219d4dd2e92ba)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)

<a id="canonical-cd5dc41bdee58f1122ee72f10dd689aa0d902a45ab491174d61219d4dd2e92ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-084faf6168b047c9df8cc71b9c65c07e95236ed4d2c707eb98293f96a3ab48da"></a>

## Data source — Data source / cf534d8b5ddd / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
- [Examples](data-sources--waf_exclusion_policy--examples--group-001.md#canonical-c81efcf5a52a4fc6a6d88f4e932312fe4a94dd7084d183754fbb93bf646c0abf)
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

<a id="canonical-cc81d5106e0f0c9a1f8c77348bcaaeeb8a33dd71853eedcbaf8a9af2421febbb"></a>

## Next pages — Data source / cf534d8b5ddd / 3

- [Examples](data-sources--waf_exclusion_policy--examples--group-001.md#canonical-c81efcf5a52a4fc6a6d88f4e932312fe4a94dd7084d183754fbb93bf646c0abf)
- [xcsh_waf_exclusion_policy](../data-sources/waf_exclusion_policy.md#canonical-00d01dd131d2721a36f10a4902fffb30f51636069e4609e56a6fd98293061663)
