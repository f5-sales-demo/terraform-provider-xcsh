---
page_title: "xcsh_alert_gen_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy landing."
---

# xcsh_alert_gen_policy landing

<a id="canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33eee06cb2465b8a713a121e1e34d9538ba6972f45b25cca76699beeae4c0162"></a>

## xcsh_alert_gen_policy — xcsh_alert_gen_policy / d391a589b44e / 2

Breadcrumbs:

- xcsh_alert_gen_policy

Manages Alert Generation Policy in F5 Distributed Cloud.

<a id="canonical-88546ba7fa438eb82285efe50b4e397dbb61ca0f174d1de7810e6ea095d2b297"></a>

## Prerequisites — xcsh_alert_gen_policy / d391a589b44e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-79e267167c67bef59b30ebb7e5204d4205623c92ca7662517445ebaaf9122789"></a>

## Minimal configuration — xcsh_alert_gen_policy / d391a589b44e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertGenPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertGenPolicy by name
data "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}

output "alert_gen_policy_id" {
  value = data.xcsh_alert_gen_policy.example.id
}
```

<a id="canonical-73593913add3d06454a62774fa0b9f25d5b53f4048bdaca1504162e85a2cf580"></a>

## Root configuration — xcsh_alert_gen_policy / d391a589b44e / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-041fa9779e23edf89a4272d577effc022f49b4dfddb2d565817c7d3796008aba"></a>

## Next pages — xcsh_alert_gen_policy / d391a589b44e / 6

- [Property reference](../guides/data-sources--alert_gen_policy--reference--group-001.md#canonical-2e8fb0c00c43506e5e88f5573a46c7fc3c5d2904f33ee950e03580f94f30d516)
- [Examples](../guides/data-sources--alert_gen_policy--examples--group-001.md#canonical-d17d447e66dec82f267cc359c03ee17dc2e216a459c21d2dd2a3110f30fcc8e4)
