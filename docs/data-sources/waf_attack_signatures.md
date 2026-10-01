---
page_title: "xcsh_waf_attack_signatures landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_attack_signatures landing."
---

# xcsh_waf_attack_signatures landing

<a id="canonical-4263664dad62728ed7ee4477dbb626c200a26b4fc673ffeb3c69ebfb8fec8479"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64472d538a18f1b1151a44215be4369a4f05acf824bb298a9cdc35d2b8a60e31"></a>

## xcsh_waf_attack_signatures — xcsh_waf_attack_signatures / d36ba71a92d4 / 2

Breadcrumbs:

- xcsh_waf_attack_signatures

Resource retrieval operation.

<a id="canonical-dd73b150c0d03f3ed6dabaeaa3af7993fd30013459c1a3c951ea4cdeee1a165b"></a>

## Prerequisites — xcsh_waf_attack_signatures / d36ba71a92d4 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6f6101895193f3aa62f6d232eb2d7fd39e04149acaae220dae9766ebcdd55092"></a>

## Minimal configuration — xcsh_waf_attack_signatures / d36ba71a92d4 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFAttackSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_attack_signatures" "example" {
}

output "waf_attack_signatures_result" {
  value = data.xcsh_waf_attack_signatures.example
}
```

<a id="canonical-82f35e35e585b2bd96b946db176ef88291aa697163fdf611d2385a6a007af6a9"></a>

## Root configuration — xcsh_waf_attack_signatures / d36ba71a92d4 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-7e2aebae13d2baab1e312f775e4607f1f06db68c3f63b5507ce836ea39e8e636"></a>

## Next pages — xcsh_waf_attack_signatures / d36ba71a92d4 / 6

- [Property reference](../guides/data-sources--waf_attack_signatures--reference--group-001.md#canonical-d8800c568006a12d4c0110d8027440d15955ab5599affd2d76e682b49dec38be)
- [Examples](../guides/data-sources--waf_attack_signatures--examples--group-001.md#canonical-406df1b7f29a187b4eac690cc0fb9071ea95aff78f5a777330f7d2afc40cd572)
