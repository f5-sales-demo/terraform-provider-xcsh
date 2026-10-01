---
page_title: "xcsh_enhanced_firewall_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy landing."
---

# xcsh_enhanced_firewall_policy landing

<a id="canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-390f179e075232526b52a6cca321bea421e24a027fd43cd49849b29bae304cb0"></a>

## xcsh_enhanced_firewall_policy — xcsh_enhanced_firewall_policy / 0448a8d2e807 / 2

Breadcrumbs:

- xcsh_enhanced_firewall_policy

Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy
specification. configuration.

<a id="canonical-4672b7b06734ac9359dd2a21640f35911dd600b39b364136f0bf03f9bdf5391f"></a>

## Prerequisites — xcsh_enhanced_firewall_policy / 0448a8d2e807 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-93cc7bd3b2b6f91bb40f3b98cae2f67bd91a10b26f49e7d6ebbdfca7d3770a68"></a>

## Minimal configuration — xcsh_enhanced_firewall_policy / 0448a8d2e807 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# EnhancedFirewallPolicy Resource Example
# Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic EnhancedFirewallPolicy configuration
resource "xcsh_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "staging"
}
```

<a id="canonical-7a9f017703f7410f86f58b8adc7ace0f73e544d11c9a2e90f2f585c7c96a01c8"></a>

## Root configuration — xcsh_enhanced_firewall_policy / 0448a8d2e807 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-53e35e19b40c92e2b11699fbe637c05d2f8c259261157c69ea750f9765f6735f"></a>

## Next pages — xcsh_enhanced_firewall_policy / 0448a8d2e807 / 6

- [Property reference](../guides/resources--enhanced_firewall_policy--reference--group-001.md#canonical-d897218d98eda688244d962679d4a02f9097e7cf0efaf995e3fdbba6e8ab503d)
- [Examples](../guides/resources--enhanced_firewall_policy--examples--group-001.md#canonical-3067c7f78f0773f9525f7cd6dabdd7e7c632b4e17610ccc3f70bf0a902a73356)
- [Import](../guides/resources--enhanced_firewall_policy--lifecycle--group-001.md#canonical-a9fe2fabd663a232ff7fc13d1134fee3e7ed4b6b6ec65eae2bbda9847ad5bce0)
- [Timeouts](../guides/resources--enhanced_firewall_policy--lifecycle--group-001.md#canonical-8674ae53e944f898260824d54ad8c538ef3d84072b15c08b9ecbca0107dd44e8)
