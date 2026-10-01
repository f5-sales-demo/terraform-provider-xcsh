---
page_title: "xcsh_protected_application landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application landing."
---

# xcsh_protected_application landing

<a id="canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7e731f8665814d3ad9949d2f1e86df0c934e3bd6ff25bc0c19383678977149b"></a>

## xcsh_protected_application — xcsh_protected_application / 25c5efcfa3bb / 2

Breadcrumbs:

- xcsh_protected_application

Manages applications protected by Bot Defense in F5 Distributed Cloud.

<a id="canonical-ce51e72635eda8da20a236f196259c51c634f13f488a2ae263aa00afd4554326"></a>

## Prerequisites — xcsh_protected_application / 25c5efcfa3bb / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6130f3e1ccf87c8fd79308344c1fe06bfd53aa02b2aab7ec217803638b58054e"></a>

## Minimal configuration — xcsh_protected_application / 25c5efcfa3bb / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedApplication Resource Example
# Manages applications protected by Bot Defense in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedApplication configuration
resource "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}
```

<a id="canonical-c6918298787511d633389394617bf299cbdecec5a256c828459fb4f166637797"></a>

## Root configuration — xcsh_protected_application / 25c5efcfa3bb / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-f340f1fb319d84a9d07bbdb1904ee2f991e06f1cc6c3d897dec571c545a7f08e"></a>

## Next pages — xcsh_protected_application / 25c5efcfa3bb / 6

- [Property reference](../guides/resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [Examples](../guides/resources--protected_application--examples--group-001.md#canonical-1451a94d52ce926e002b158e3403caa1e4a2c208b632bc30ad80f7db36f5ab96)
- [Import](../guides/resources--protected_application--lifecycle--group-001.md#canonical-16ca0f43bcb76b4e62a2f1418b1390f136abb11880a1006b2b2507e9bd0a9916)
- [Timeouts](../guides/resources--protected_application--lifecycle--group-001.md#canonical-a09ee54b2d0374dc71e13d4ca7538454b00a50fadc1f1ee8db369abdbef361b5)
