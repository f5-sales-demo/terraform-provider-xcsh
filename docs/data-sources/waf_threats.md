---
page_title: "xcsh_waf_threats landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threats landing."
---

# xcsh_waf_threats landing

<a id="canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b44c0ecbdf81dae6c221f7d5db00a6da41491347b70d874281ee54812807018"></a>

## xcsh_waf_threats — xcsh_waf_threats / 79b1bf7cbd63 / 2

Breadcrumbs:

- xcsh_waf_threats

Resource creation operation.

<a id="canonical-a3002064d13ab7d75e45273e714cf5abbbb9bcf8465701242de2a8964a73907e"></a>

## Prerequisites — xcsh_waf_threats / 79b1bf7cbd63 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-83450a547f0ff9e5138f645907150f0da1ef7f2f0c4e419c775543449f9befd6"></a>

## Minimal configuration — xcsh_waf_threats / 79b1bf7cbd63 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFThreats DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threats" "example" {
}

output "waf_threats_result" {
  value = data.xcsh_waf_threats.example
}
```

<a id="canonical-e02f3f307dc65f2ec5a03cf14990aa993d162652c7c425d62bc8636d48b4767b"></a>

## Root configuration — xcsh_waf_threats / 79b1bf7cbd63 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2a7e5a2c239bd721ebe289cfa7b14cb805385b3a1ada577536f6e8626f6f7539"></a>

## Next pages — xcsh_waf_threats / 79b1bf7cbd63 / 6

- [Property reference](../guides/data-sources--waf_threats--reference--group-001.md#canonical-694a9838f896e006462dd3aa27eeba62fd0d019c7247bec80b961803c8526d8e)
- [Examples](../guides/data-sources--waf_threats--examples--group-001.md#canonical-ec4d28c7443fcfd9835e6bf5700cc562dffd31d6cc15938fae92732a710ee386)
