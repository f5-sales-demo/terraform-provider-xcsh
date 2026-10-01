---
page_title: "xcsh_registration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration examples."
---

# xcsh_registration examples

<a id="canonical-308dd2706c143d052513a67a5ad0066da2434e408835e3cde8ea716d8d0bf9f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-442e7658eb17eae08c3f9479f1f4581bc59515078614b421377cdfc531cd8cfa"></a>

## Examples — Examples / 71b3ce573d93 / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- Examples

<a id="canonical-6e0d33b64e4d8d9c88ba469e2ccc0190db174097fb0992956b683d72212d3baa"></a>

## Complete configurations — Examples / 71b3ce573d93 / 3

- [Data source](data-sources--registration--examples--group-001.md#canonical-c2763bee48f6e9fb85757bf6f3d7bf72df7870a5af1eb9963498914d517ccbca): valid configuration.

<a id="canonical-a19ac0f8c17a3328a991d6b0aeef74675da091589d2c8eead8e0dcecc3921d65"></a>

## Next pages — Examples / 71b3ce573d93 / 4

- [Data source](data-sources--registration--examples--group-001.md#canonical-c2763bee48f6e9fb85757bf6f3d7bf72df7870a5af1eb9963498914d517ccbca)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)

<a id="canonical-c2763bee48f6e9fb85757bf6f3d7bf72df7870a5af1eb9963498914d517ccbca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff329abfe9890efe35d716198487d8013267f64691e9b7031beabb48b436e50f"></a>

## Data source — Data source / 11f6789d3a3c / 2

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
- [Examples](data-sources--registration--examples--group-001.md#canonical-308dd2706c143d052513a67a5ad0066da2434e408835e3cde8ea716d8d0bf9f7)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_registration/data-source.tf`; digest `sha256:887504fe79e23b9f0b1ae62f2ad6a04c459d805004c72b40be81620b6a16261f`.

```terraform
# Registration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Registration by name
data "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"
}

output "registration_id" {
  value = data.xcsh_registration.example.id
}
```

<a id="canonical-d24c3ea2559bc6461422832beffd6df3ffa76b25cea8c5b8745057fd4cef1649"></a>

## Next pages — Data source / 11f6789d3a3c / 3

- [Examples](data-sources--registration--examples--group-001.md#canonical-308dd2706c143d052513a67a5ad0066da2434e408835e3cde8ea716d8d0bf9f7)
- [xcsh_registration](../data-sources/registration.md#canonical-2746a399e29ec20e70d5fe8bd2126104322cfa8facd6b7656951f85370c6ccf1)
