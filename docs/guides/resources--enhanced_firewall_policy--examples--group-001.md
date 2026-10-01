---
page_title: "xcsh_enhanced_firewall_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy examples."
---

# xcsh_enhanced_firewall_policy examples

<a id="canonical-3067c7f78f0773f9525f7cd6dabdd7e7c632b4e17610ccc3f70bf0a902a73356"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c176e209c3fef7a37d2cfacb8527c4b1a2f22e510499ea64088f2d5d1630f09f"></a>

## Examples — Examples / 912336911415 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- Examples

<a id="canonical-2e97166b8bad05f6155cd7224e6450b076bf00b337496f387793c7bbe6fa1c22"></a>

## Complete configurations — Examples / 912336911415 / 3

- [Resource](resources--enhanced_firewall_policy--examples--group-001.md#canonical-54893d72d80f6d6530596540b7515dcfbca22e25ae960f6399e95b72cef97b59): valid configuration.

<a id="canonical-99292e0898f23209d5083e31a8552a62e2639ad36d2e524c534418b0ec643442"></a>

## Next pages — Examples / 912336911415 / 4

- [Resource](resources--enhanced_firewall_policy--examples--group-001.md#canonical-54893d72d80f6d6530596540b7515dcfbca22e25ae960f6399e95b72cef97b59)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)

<a id="canonical-54893d72d80f6d6530596540b7515dcfbca22e25ae960f6399e95b72cef97b59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22dc306091715de4241a3a8923574402d1f3fdd4767d42eb4bf7bcb62e92f1d7"></a>

## Resource — Resource / 8a6cd5ee2326 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
- [Examples](resources--enhanced_firewall_policy--examples--group-001.md#canonical-3067c7f78f0773f9525f7cd6dabdd7e7c632b4e17610ccc3f70bf0a902a73356)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_enhanced_firewall_policy/resource.tf`; digest `sha256:de754ad7eba55318affca16381d010b7282e358ea583d9fad1b36b6208539006`.

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

<a id="canonical-e62af7430ae39696ec6040725c506851324d7824cd75b3f27d30ae486737a7d4"></a>

## Next pages — Resource / 8a6cd5ee2326 / 3

- [Examples](resources--enhanced_firewall_policy--examples--group-001.md#canonical-3067c7f78f0773f9525f7cd6dabdd7e7c632b4e17610ccc3f70bf0a902a73356)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-2c68365799ebb1e43e903dcddc71904838c745479d157dd46ede05f4cbf18058)
