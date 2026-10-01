---
page_title: "xcsh_kubernetes_manifests landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_kubernetes_manifests landing."
---

# xcsh_kubernetes_manifests landing

<a id="canonical-c83216e14c73f7838248127c78a0137201c23f4a90e3676b65e4448f86ede548"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bac1b8e0c33f5a3d21311ce0c069d131ce5b4d1668f7906573562ded23246e34"></a>

## xcsh_kubernetes_manifests — xcsh_kubernetes_manifests / 5c970db1dddd / 2

Breadcrumbs:

- xcsh_kubernetes_manifests

Kubernetes workload configuration.

<a id="canonical-ca25ade0b09e4ec7b683929724e50a18b9b3c022b59b4067dc6f226bf67d808f"></a>

## Prerequisites — xcsh_kubernetes_manifests / 5c970db1dddd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-bf5a4a7f1835c8484bd32563a9466e8a9d5b2ad897f8dea0d5a7513772ab5e11"></a>

## Minimal configuration — xcsh_kubernetes_manifests / 5c970db1dddd / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# KubernetesManifests EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_kubernetes_manifests" "example" {
  site = "example-value"
}
```

<a id="canonical-85b0024713031bcdd0057b0f41d9d02096653f8ddc3a71bff4520a004ea65801"></a>

## Root configuration — xcsh_kubernetes_manifests / 5c970db1dddd / 5

Required root properties: `site`. Full root flags and choices appear in the property reference.

<a id="canonical-bf8403d5274ef035f36aa8827c60231314ff399e890391db25a0252cad5e6c0b"></a>

## Next pages — xcsh_kubernetes_manifests / 5c970db1dddd / 6

- [Property reference](../guides/ephemeral-resources--kubernetes_manifests--reference--group-001.md#canonical-2ac4a07b42a4e33d39b56bd91d9913d9cb0a8431f6e40485e5151778ee5864c8)
- [Examples](../guides/ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-facff2d02f16b29739081fecca5d71af469e84f62929ac9ed28ace112c3855fe)
- [Lifecycle](../guides/ephemeral-resources--kubernetes_manifests--lifecycle--group-001.md#canonical-1db5d4ebfb494161bd7d29024efcd30b5eabba5803149c890860ccc09c5f8476)
