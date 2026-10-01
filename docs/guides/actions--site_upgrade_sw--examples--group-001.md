---
page_title: "xcsh_site_upgrade_sw examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_sw examples."
---

# xcsh_site_upgrade_sw examples

<a id="canonical-5f5bbc00fe0ddc9db04112bacf135860a397d4686eb33770fd9bb25aba8c9394"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d1314f63f2929b594fc7c9a853c2ccc12dd4ea43588e9b24180fd7f9d686d11"></a>

## Examples — Examples / fc725786edd8 / 2

Breadcrumbs:

- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md#canonical-18736db82b100397af6d4421dc51a7a0bb99b97fd8de86af0b094a5638396565)
- Examples

<a id="canonical-14136bcd49de969b401459427771df52136f9934edb827083c860ed95f837658"></a>

## Complete configurations — Examples / fc725786edd8 / 3

- [Action](actions--site_upgrade_sw--examples--group-001.md#canonical-7bed33b991c392ee8ec911b2b748d75c75ffed17c7327fd12b9d319407c80eae): valid configuration.

<a id="canonical-59e4d6b035e2a89c5ad6e33b08243ddabcfcd2a637f7ea8b123172427bcc81a7"></a>

## Next pages — Examples / fc725786edd8 / 4

- [Action](actions--site_upgrade_sw--examples--group-001.md#canonical-7bed33b991c392ee8ec911b2b748d75c75ffed17c7327fd12b9d319407c80eae)
- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md#canonical-18736db82b100397af6d4421dc51a7a0bb99b97fd8de86af0b094a5638396565)

<a id="canonical-7bed33b991c392ee8ec911b2b748d75c75ffed17c7327fd12b9d319407c80eae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c59e8324cd639357c7cf492762422eb0f2c14cce362ed62e91e9576728f0399"></a>

## Action — Action / 6dfa82a7e50a / 2

Breadcrumbs:

- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md#canonical-18736db82b100397af6d4421dc51a7a0bb99b97fd8de86af0b094a5638396565)
- [Examples](actions--site_upgrade_sw--examples--group-001.md#canonical-5f5bbc00fe0ddc9db04112bacf135860a397d4686eb33770fd9bb25aba8c9394)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_upgrade_sw/action.tf`; digest `sha256:f543d404a49e0f6ff11128bf33aab9d859e0913de2e4fcdcef829bf992c11087`.

```terraform
# SiteUpgradeSw Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# The API accepts the upgrade request immediately; convergence is asynchronous.
# This action does not reconcile a site's pinned software_settings.
action "xcsh_site_upgrade_sw" "example" {
  config {
    site             = "example-value"
    software_version = "example-value"
  }
}
```

<a id="canonical-1791e7bcdf85ee3e40beff2b40e6c6eb8125340334b42890f2c2bb9ea7063387"></a>

## Next pages — Action / 6dfa82a7e50a / 3

- [Examples](actions--site_upgrade_sw--examples--group-001.md#canonical-5f5bbc00fe0ddc9db04112bacf135860a397d4686eb33770fd9bb25aba8c9394)
- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md#canonical-18736db82b100397af6d4421dc51a7a0bb99b97fd8de86af0b094a5638396565)
