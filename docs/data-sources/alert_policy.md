---
page_title: "xcsh_alert_policy landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy landing."
---

# xcsh_alert_policy landing

<a id="canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ed1007803438fbc39777cfe3a0186122d662a1b3f8c5aad38592ddedde82a6f"></a>

## xcsh_alert_policy — xcsh_alert_policy / 4b9aba9f7399 / 2

Breadcrumbs:

- xcsh_alert_policy

Manages new Alert Policy Object in F5 Distributed Cloud.

<a id="canonical-8b721da1199f2b7e41f7472645f5ca947f1409204620cc37b1c9483097874b3b"></a>

## Prerequisites — xcsh_alert_policy / 4b9aba9f7399 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-2186061ac47d62cd370406dc63176e39f85f6ace9993c887da052d6fde1654eb"></a>

## Minimal configuration — xcsh_alert_policy / 4b9aba9f7399 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertPolicy by name
data "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}

output "alert_policy_id" {
  value = data.xcsh_alert_policy.example.id
}
```

<a id="canonical-4fa54f0af788346f583a291c4c1cb9cf6126587667436b387931c16276389d5d"></a>

## Root configuration — xcsh_alert_policy / 4b9aba9f7399 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3c61419a774db0354c79c870eaa3ec682270a60e852dfd57c24fbad85ee80630"></a>

## Next pages — xcsh_alert_policy / 4b9aba9f7399 / 6

- [Property reference](../guides/data-sources--alert_policy--reference--group-001.md#canonical-96a61606ae7f7cdfa2189a69f67c959ba0d7cb9eb20ab147015fece06408b99b)
- [Examples](../guides/data-sources--alert_policy--examples--group-001.md#canonical-3d81f623ab4dcd5d770dbd48b8bef571212d883a505f979b3fabf28a252d1469)
