---
page_title: "xcsh_infraprotect_mitigation_ips landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_mitigation_ips landing."
---

# xcsh_infraprotect_mitigation_ips landing

<a id="canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fc35487791a6853da2daf7ec4498e8817f33ae62913c648f7f49e5a2bf8cfb8"></a>

## xcsh_infraprotect_mitigation_ips — xcsh_infraprotect_mitigation_ips / 5f4e634d1b39 / 2

Breadcrumbs:

- xcsh_infraprotect_mitigation_ips

Resource retrieval operation.

<a id="canonical-29cc456716225b7d508aed6b77f575385209656cef87ac6ba90362000105407d"></a>

## Prerequisites — xcsh_infraprotect_mitigation_ips / 5f4e634d1b39 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-de83035f598abcb1ea04146fdd641d01c3bf181a94253ec1abfc5a132042260d"></a>

## Minimal configuration — xcsh_infraprotect_mitigation_ips / 5f4e634d1b39 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```

<a id="canonical-df585738585966b17f9e493013cd5465c1b1d31f8d0afa1ef5a555cf57bb5b1f"></a>

## Root configuration — xcsh_infraprotect_mitigation_ips / 5f4e634d1b39 / 5

Required root properties: `mitigation_id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8daa6c511615869656e09a9528851e6ca99d9a3455d0189268852305e88e8d1d"></a>

## Next pages — xcsh_infraprotect_mitigation_ips / 5f4e634d1b39 / 6

- [Property reference](../guides/data-sources--infraprotect_mitigation_ips--reference--group-001.md#canonical-e50d1fb885e742097142394b7c46c24d524be2fe0d2364e28add4376fdca80f4)
- [Examples](../guides/data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-acf837bcc3ba76e94b3d1c3511fae39e4803523592aaec119cad377aab2d3fa0)
