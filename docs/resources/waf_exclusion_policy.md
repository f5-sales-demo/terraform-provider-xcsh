---
page_title: "xcsh_waf_exclusion_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy landing."
---

# xcsh_waf_exclusion_policy landing

<a id="canonical-fc317230bf1605ae9396126bfb46e87a3f2d201f0b9627104a70a3a0678f8f89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e0cc09fb25e84e138ff1da37c385a8690bd1959dabf8bd72ecf6e1422ae827e"></a>

## xcsh_waf_exclusion_policy — xcsh_waf_exclusion_policy / 793b3b6c05bc / 2

Breadcrumbs:

- xcsh_waf_exclusion_policy

Manages WAF exclusion policy in F5 Distributed Cloud.

<a id="canonical-fdc8d7051ad5e91fca6082ffae3450d85f06a42c2e2e62cb7d916c2a50c6831f"></a>

## Prerequisites — xcsh_waf_exclusion_policy / 793b3b6c05bc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-746ee5e7ffde9e1bcef2ed74a3e8369cba95aff2ec995eefc026dc7c91d256d5"></a>

## Minimal configuration — xcsh_waf_exclusion_policy / 793b3b6c05bc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-edc5f4013d7a25b88e3397f456d9775d817d5a3e7593f15f19d20477b1680dcd"></a>

## Root configuration — xcsh_waf_exclusion_policy / 793b3b6c05bc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-86ad8b5a0625494e07bf3d9149471ceafc4b1757fefc9fe18ee90dac6a0522ac"></a>

## Next pages — xcsh_waf_exclusion_policy / 793b3b6c05bc / 6

- [Property reference](../guides/resources--waf_exclusion_policy--reference--group-001.md#canonical-e11a828832c3e7a264c0045f1ee095803322cdb88fc9c8ad72980de81b5370f5)
- [Examples](../guides/resources--waf_exclusion_policy--examples--group-001.md#canonical-146b94aad58f86e5ea0e61f3bfb3ddcd8362cbd7dbb53ea2486b5251c5f1dec0)
- [Import](../guides/resources--waf_exclusion_policy--lifecycle--group-001.md#canonical-c62af02aed4a213863d69863f3d3edbb995bf329358720ac0d007a5fdeca4bba)
- [Timeouts](../guides/resources--waf_exclusion_policy--lifecycle--group-001.md#canonical-8879781bdcd401ae30f7dd317eb84188c7c3d62be47a334c59569da3844457d4)
