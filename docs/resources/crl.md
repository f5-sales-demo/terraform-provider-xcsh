---
page_title: "xcsh_crl landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl landing."
---

# xcsh_crl landing

<a id="canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8adc524340cd43915ebf905ace004e93446ea4cf59f0daaba3bc34ab59c6da95"></a>

## xcsh_crl — xcsh_crl / 66fa7f8f8f57 / 2

Breadcrumbs:

- xcsh_crl

Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.

<a id="canonical-718f21ca98187aee5e9c013b7a2424ced8ad0b9669b77790b87fc7eb640c11e9"></a>

## Prerequisites — xcsh_crl / 66fa7f8f8f57 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3a2e92437c3ffcd0771f4c0ec29a65c4032fdc04c8b48ea8295e843b8e7bf2ab"></a>

## Minimal configuration — xcsh_crl / 66fa7f8f8f57 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```

<a id="canonical-297dcb3e80db8e05663db7727ac2c065717123d56507d7b3f5191502fd2eeee8"></a>

## Root configuration — xcsh_crl / 66fa7f8f8f57 / 5

Required root properties: `name`, `namespace`, `refresh_interval`, `server_address`, `server_port`, `timeout`. Full root flags and choices appear in the property reference.

<a id="canonical-8cf35974edc429df77ed8bdbfac7dd23c9450f677e25bce887eb6ac8a232daea"></a>

## Next pages — xcsh_crl / 66fa7f8f8f57 / 6

- [Property reference](../guides/resources--crl--reference--group-001.md#canonical-0a1bcf7add556b43d14c3820e55e0a6110324a87da60186d84fa20c107de9d30)
- [Examples](../guides/resources--crl--examples--group-001.md#canonical-f4b8e81c172d26ed2416d6b64fa100e009639d05b196f13650b51e15b65a999e)
- [Import](../guides/resources--crl--lifecycle--group-001.md#canonical-f502346e8d8fab0b4cb599b242976f9f11c6047005af1de64ede7e105392498c)
- [Timeouts](../guides/resources--crl--lifecycle--group-001.md#canonical-c626d01c0a67a27c85bb7d977d174d1b7407011f0150d66aae8856bd345a1889)
