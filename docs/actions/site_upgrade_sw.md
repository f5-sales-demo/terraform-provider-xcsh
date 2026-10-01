---
page_title: "xcsh_site_upgrade_sw landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_sw landing."
---

# xcsh_site_upgrade_sw landing

<a id="canonical-18736db82b100397af6d4421dc51a7a0bb99b97fd8de86af0b094a5638396565"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b983ee5b3fe394d7d1e33b68198fe76b14f4ec4a470eca9dd901a7e4896635f0"></a>

## xcsh_site_upgrade_sw — xcsh_site_upgrade_sw / f5e6663c8f8f / 2

Breadcrumbs:

- xcsh_site_upgrade_sw

Request an in-place site software upgrade.

<a id="canonical-a43823969e9d7690158f934b0e7f780c15f24d60fd639506add52b1cae0897d0"></a>

## Prerequisites — xcsh_site_upgrade_sw / f5e6663c8f8f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8f24f6b5c6ea20a4ee20aac9d604e6cfabc571c787486e2119012878249a0607"></a>

## Minimal configuration — xcsh_site_upgrade_sw / f5e6663c8f8f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-bbddb4e1f9a367627adc2bfb0355ed0143ced9c4340ef99036aa09bb7013aa6e"></a>

## Root configuration — xcsh_site_upgrade_sw / f5e6663c8f8f / 5

Required root properties: `site`, `software_version`. Full root flags and choices appear in the property reference.

<a id="canonical-9bb397876254c61db76e13522b1b9f6c36d747defa4c1cd3ce680ccc6d460548"></a>

## Next pages — xcsh_site_upgrade_sw / f5e6663c8f8f / 6

- [Property reference](../guides/actions--site_upgrade_sw--reference--group-001.md#canonical-0c382e518064cf3a2721e69913fd349c991e7047f5c7a6025d6c71eb8269340e)
- [Examples](../guides/actions--site_upgrade_sw--examples--group-001.md#canonical-5f5bbc00fe0ddc9db04112bacf135860a397d4686eb33770fd9bb25aba8c9394)
- [Lifecycle](../guides/actions--site_upgrade_sw--lifecycle--group-001.md#canonical-23a990217858f0772fd3794724ea67680f0ad99f8f1529d4d949476ca1a8332f)
