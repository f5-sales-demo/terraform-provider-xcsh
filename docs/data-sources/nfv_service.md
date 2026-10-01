---
page_title: "xcsh_nfv_service landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service landing."
---

# xcsh_nfv_service landing

<a id="canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b334db1446c63ce2100aa7c5e9bb852421adbd5552158780542718dc8edca75"></a>

## xcsh_nfv_service — xcsh_nfv_service / f588a2816730 / 2

Breadcrumbs:

- xcsh_nfv_service

Manages new NFV service with configured parameters in F5 Distributed Cloud.

<a id="canonical-2b0c8250bf0c15057b3d5240dcbf3131a7633fd9faa7d6976d3e505b48aa7b64"></a>

## Prerequisites — xcsh_nfv_service / f588a2816730 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c85d26543d7c8ebf94f54550543ae6895b39cc3e0666f6b2753dfd5d098f17ee"></a>

## Minimal configuration — xcsh_nfv_service / f588a2816730 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NfvService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NfvService by name
data "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}

output "nfv_service_id" {
  value = data.xcsh_nfv_service.example.id
}
```

<a id="canonical-b11f594507eb30f64ba0dd74de6d0df4366a43f45e57d7e6ba75609468f4eff3"></a>

## Root configuration — xcsh_nfv_service / f588a2816730 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8d3af14e42b2f8bb07c20ecc5250298d0dfef1b27fa55aac0e4dce7ae054b823"></a>

## Next pages — xcsh_nfv_service / f588a2816730 / 6

- [Property reference](../guides/data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [Examples](../guides/data-sources--nfv_service--examples--group-001.md#canonical-17846793b391435f50864456d6d6124f1ec9432c80f12fa51c56820307807d96)
