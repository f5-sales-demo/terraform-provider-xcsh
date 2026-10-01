---
page_title: "xcsh_enhanced_firewall_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy landing."
---

# xcsh_enhanced_firewall_policy landing

<a id="canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50867a2373fc744d97fac966ec685a674efbdcf1187075307dd44f01e0293194"></a>

## xcsh_enhanced_firewall_policy — xcsh_enhanced_firewall_policy / a7e2d5e9db78 / 2

Breadcrumbs:

- xcsh_enhanced_firewall_policy

Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy
specification. configuration.

<a id="canonical-42f000cd62269f799664f3ab33b50891faf4ee37efcbae68df2e6be490d60188"></a>

## Prerequisites — xcsh_enhanced_firewall_policy / a7e2d5e9db78 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ef0f4c81fcb6992b4af7cf5d8c417580b885accd991b6783f8d4fd331b05ce18"></a>

## Minimal configuration — xcsh_enhanced_firewall_policy / a7e2d5e9db78 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# EnhancedFirewallPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing EnhancedFirewallPolicy by name
data "xcsh_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "staging"
}

output "enhanced_firewall_policy_id" {
  value = data.xcsh_enhanced_firewall_policy.example.id
}
```

<a id="canonical-971c0d400fb5637b710b47811b82faf6008f3905561bd8041ee97edc94fdaaff"></a>

## Root configuration — xcsh_enhanced_firewall_policy / a7e2d5e9db78 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c759e2ccfea6560c7993cf7187466d5681c1c7ef91af683e034de6e5f6a1e962"></a>

## Next pages — xcsh_enhanced_firewall_policy / a7e2d5e9db78 / 6

- [Property reference](../guides/data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-66f0e294308cd8b4cbed5e1382a33e0142349b0cd3c550d25d932648ead94194)
- [Examples](../guides/data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-70c80eaf95e4cca8d9d72e1682069637496af9b4930875c6410467d72a0aacaa)
