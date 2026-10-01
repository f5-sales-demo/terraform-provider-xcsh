---
page_title: "xcsh_workload_flavor landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor landing."
---

# xcsh_workload_flavor landing

<a id="canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64e630fdcd636313c8b359229352e7865ecd979db1dec622591683d84d4995cc"></a>

## xcsh_workload_flavor — xcsh_workload_flavor / c02b64b085b3 / 2

Breadcrumbs:

- xcsh_workload_flavor

Manages workload\_flavor in F5 Distributed Cloud.

<a id="canonical-f52c5037bdcec707365dd70cc5ac3df66dd69ec1db9878190b93b5ea1fd83040"></a>

## Prerequisites — xcsh_workload_flavor / c02b64b085b3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-9a689cb77d139cd3c106675c84701b5ecbb4b8a107e9f951a9932fd21f803ec7"></a>

## Minimal configuration — xcsh_workload_flavor / c02b64b085b3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WorkloadFlavor Resource Example
# Manages workload_flavor in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WorkloadFlavor configuration
resource "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}
```

<a id="canonical-08ec3b46fc4af37e1887fe2fe7c8871d6d74006d585a079299777c6e231f0876"></a>

## Root configuration — xcsh_workload_flavor / c02b64b085b3 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-f8c7bb353b52c6acdbe028ddcdc1abb541422179d43cc2e97d1d14d42d8f27e6"></a>

## Next pages — xcsh_workload_flavor / c02b64b085b3 / 6

- [Property reference](../guides/resources--workload_flavor--reference--group-001.md#canonical-a1b2df5c2fff8d0b8fbb2c32a43b469b29d4036f48d674235772118f3099a1a7)
- [Examples](../guides/resources--workload_flavor--examples--group-001.md#canonical-0986ae5df41648085a9c7c99f600c126da00a829a494b8601789f82abbc5dee4)
- [Import](../guides/resources--workload_flavor--lifecycle--group-001.md#canonical-bb25a902a65154102fe6fba37a0e1a02fb70da7846800fcd33caac5f95a9d72b)
- [Timeouts](../guides/resources--workload_flavor--lifecycle--group-001.md#canonical-c864a07acee5fa8627617f1c184f76a1d78f9103620170accde6da36bfd42bde)
