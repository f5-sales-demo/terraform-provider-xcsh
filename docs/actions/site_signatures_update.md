---
page_title: "xcsh_site_signatures_update landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_signatures_update landing."
---

# xcsh_site_signatures_update landing

<a id="canonical-997cfc6e0541d9b7078c9e8c892fe43f791342c6fe0d09cd7f128fc478b34edf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43c8412747e7497efd3e05fc88523738ae7aa25e80ba382c9f68226ecb2c4dbd"></a>

## xcsh_site_signatures_update — xcsh_site_signatures_update / 3c95fdcf0f80 / 2

Breadcrumbs:

- xcsh_site_signatures_update

Resource creation operation.

<a id="canonical-c6b857918e7bf33ba31e98689c85a56fc7f431da0a458123cb71ed1ecb59ddc8"></a>

## Prerequisites — xcsh_site_signatures_update / 3c95fdcf0f80 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a4edfcaf7801534b952533a027501b0d2c299b4cb47f6bf774e23079e1da804f"></a>

## Minimal configuration — xcsh_site_signatures_update / 3c95fdcf0f80 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-731538d755c710a5ce13518214719528ba90ad13bc20caf7d21a1d78729adf96"></a>

## Root configuration — xcsh_site_signatures_update / 3c95fdcf0f80 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-cee9e8a88516677796d8268f1225fc3ff44482e4584b140958b7a17191aa5c76"></a>

## Next pages — xcsh_site_signatures_update / 3c95fdcf0f80 / 6

- [Property reference](../guides/actions--site_signatures_update--reference--group-001.md#canonical-93d629fdfa3cd299d906173b156c98b1b70aa3c0bfa0665fac781e3cef9f1ce2)
- [Examples](../guides/actions--site_signatures_update--examples--group-001.md#canonical-5f9409c365fddfafc467f46ef8caa853731689277c363e22f35bd55e440ba1d9)
- [Lifecycle](../guides/actions--site_signatures_update--lifecycle--group-001.md#canonical-1ca67fef219122a782008fae2f0e92c009b323253ebb95dad619893bf2261ecc)
