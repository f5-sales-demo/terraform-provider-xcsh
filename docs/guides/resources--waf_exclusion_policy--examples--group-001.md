---
page_title: "xcsh_waf_exclusion_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy examples."
---

# xcsh_waf_exclusion_policy examples

<a id="canonical-146b94aad58f86e5ea0e61f3bfb3ddcd8362cbd7dbb53ea2486b5251c5f1dec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9736741b870549760f55fc262d872089327e2f6628b9960b6deebc94f365e754"></a>

## Examples — Examples / 5bc89fa30dd6 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- Examples

<a id="canonical-32ea849e1316477ab5c5334f8e2fe84002b5ca351636bea132c05c49a5faa325"></a>

## Complete configurations — Examples / 5bc89fa30dd6 / 3

- [Resource](resources--waf_exclusion_policy--examples--group-001.md#canonical-aceaf6ac51db81fce31500ee57ba32fb46888d657acf7856e7dfcd75ba13c867): valid configuration.

<a id="canonical-874a7300dd6c09f2b9e98d23bad3ff5825c90cae5b0087549dd2fae055869531"></a>

## Next pages — Examples / 5bc89fa30dd6 / 4

- [Resource](resources--waf_exclusion_policy--examples--group-001.md#canonical-aceaf6ac51db81fce31500ee57ba32fb46888d657acf7856e7dfcd75ba13c867)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)

<a id="canonical-aceaf6ac51db81fce31500ee57ba32fb46888d657acf7856e7dfcd75ba13c867"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77e68fa255dc715e2ed1fe0073a437da261ce1f0081f33583246669f4ed436c3"></a>

## Resource — Resource / 47e163546743 / 2

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
- [Examples](resources--waf_exclusion_policy--examples--group-001.md#canonical-146b94aad58f86e5ea0e61f3bfb3ddcd8362cbd7dbb53ea2486b5251c5f1dec0)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_waf_exclusion_policy/resource.tf`; digest `sha256:873fdc1e55067d3a5c09e53e53d5c6c7bcd1d7d8f83e917fef809f43dd1b0325`.

```terraform
# WAFExclusionPolicy Resource Example
# Manages WAF exclusion policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WAFExclusionPolicy configuration
resource "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}
```

<a id="canonical-4fbd8bd2f8586f244c6590d7276e8ccc31e7f4fcceb1e8a4a2e749bc3f0aa5fa"></a>

## Next pages — Resource / 47e163546743 / 3

- [Examples](resources--waf_exclusion_policy--examples--group-001.md#canonical-146b94aad58f86e5ea0e61f3bfb3ddcd8362cbd7dbb53ea2486b5251c5f1dec0)
- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89)
