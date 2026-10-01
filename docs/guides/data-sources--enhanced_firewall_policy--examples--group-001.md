---
page_title: "xcsh_enhanced_firewall_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy examples."
---

# xcsh_enhanced_firewall_policy examples

<a id="canonical-70c80eaf95e4cca8d9d72e1682069637496af9b4930875c6410467d72a0aacaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-738aa600fb188d9e3d88e6582d093a08df4c86570075adf2ee20c8a190fba0a1"></a>

## Examples — Examples / 63f414778073 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- Examples

<a id="canonical-0db2d278f78bd676a322b9c059f2a5b043883e5851294faf9a9e23e1d167c2f8"></a>

## Complete configurations — Examples / 63f414778073 / 3

- [Data source](data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-e6c586c2980aad9a0ca5d12da2d68aef8389b1abf3e92020e3d7b40d9dae64de): valid configuration.

<a id="canonical-6b2bc96fd7906ef24fb885010339db77a55ee442041afa2bc87b7dc852ec08e4"></a>

## Next pages — Examples / 63f414778073 / 4

- [Data source](data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-e6c586c2980aad9a0ca5d12da2d68aef8389b1abf3e92020e3d7b40d9dae64de)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)

<a id="canonical-e6c586c2980aad9a0ca5d12da2d68aef8389b1abf3e92020e3d7b40d9dae64de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e843ccf2f93789264eb0b9f4380d1d5ba811082dbf9943beeeb0d2899b98041f"></a>

## Data source — Data source / eface25d69bc / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
- [Examples](data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-70c80eaf95e4cca8d9d72e1682069637496af9b4930875c6410467d72a0aacaa)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_enhanced_firewall_policy/data-source.tf`; digest `sha256:4fad9180d1e11ef8fde0a09a6c59611ed259b2e0dc9e6684bd5a7a00db211172`.

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

<a id="canonical-9428a4402b700809b868ba80106997f63ff1eb9daefc8ee60a4224ee1d54922d"></a>

## Next pages — Data source / eface25d69bc / 3

- [Examples](data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-70c80eaf95e4cca8d9d72e1682069637496af9b4930875c6410467d72a0aacaa)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-75b50fd72c5f699362531939e227fe2dc3f549a792b531c93efa229b70ad0de0)
