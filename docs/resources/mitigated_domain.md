---
page_title: "xcsh_mitigated_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain landing."
---

# xcsh_mitigated_domain landing

<a id="canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0dbf28903df0dbe5712ed8549ca78aec0439964d2416ca95ec8cbaceb11dc72"></a>

## xcsh_mitigated_domain — xcsh_mitigated_domain / e681c16403da / 2

Breadcrumbs:

- xcsh_mitigated_domain

Manages Mitigated Domain in F5 Distributed Cloud.

<a id="canonical-bbc4ef5882e4188d67aaf368e5b090047f3828505bf1cae6fd52ca730fcfa2e4"></a>

## Prerequisites — xcsh_mitigated_domain / e681c16403da / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a1c528fb9ee17ac2078bfc190aa2c1d32c90a577f3fce80b158efea77dd822d2"></a>

## Minimal configuration — xcsh_mitigated_domain / e681c16403da / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MitigatedDomain Resource Example
# Manages Mitigated Domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MitigatedDomain configuration
resource "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"

  mitigated_domain = "example-value"
}
```

<a id="canonical-da24b7c2d5e53f59f088d5d6a48ee9aef3ddb0fb489025a6d94e5feb4a568f5f"></a>

## Root configuration — xcsh_mitigated_domain / e681c16403da / 5

Required root properties: `mitigated_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d37074f1e4eb7b23936a74c0123d763938dc14cc81da6c8a48b34e8782b6371f"></a>

## Next pages — xcsh_mitigated_domain / e681c16403da / 6

- [Property reference](../guides/resources--mitigated_domain--reference--group-001.md#canonical-d8a7d36f91ae444348c9b288663ecdfe54e263784cd66b472170a42f06e73224)
- [Examples](../guides/resources--mitigated_domain--examples--group-001.md#canonical-f8cc4ca693157d4c768911d71d3f9378ab6faa7e3fb24d78c4c693a8ced45d35)
- [Import](../guides/resources--mitigated_domain--lifecycle--group-001.md#canonical-699327b09fa7c3f4db717e6b5d5a97fb3259f26a7ff14a9552f6feaf36f4ebc2)
- [Timeouts](../guides/resources--mitigated_domain--lifecycle--group-001.md#canonical-388bcb3e87fdea76fb69f6996f6852d7eab39d6a4bb3d66dc90560c76e1b26e8)
