---
page_title: "xcsh_service_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy landing."
---

# xcsh_service_policy landing

<a id="canonical-dcb176e0206b31dfc3d035d78a3f4865f3fc119cbf2c464c9ecd2483a15412b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0286ab122f4ab552b6a631cc1a49d026c36977ec635020ca3d89560a058ee8c"></a>

## xcsh_service_policy — xcsh_service_policy / 9071da22cb0b / 2

Breadcrumbs:

- xcsh_service_policy

Manages service\_policy creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-60ec72df20586696a353c8576a47c250b5cf4f19b744a21190ba1a9d6ff7384a"></a>

## Prerequisites — xcsh_service_policy / 9071da22cb0b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-7750e52f92e39235458a6fe161f462ca00e99e2f61af47616d2cb70b3e68564e"></a>

## Minimal configuration — xcsh_service_policy / 9071da22cb0b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicy Resource Example
# Manages service_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicy configuration
resource "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}
```

<a id="canonical-8dec4f43486aa9f3f77d417273619ee55e9cff6e3ce35e937ef66e7986e14036"></a>

## Root configuration — xcsh_service_policy / 9071da22cb0b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9ab96fb810f8d63593a891bf450ca1bc6b8d4e7cadd9c4033f1cddb5103e61e3"></a>

## Next pages — xcsh_service_policy / 9071da22cb0b / 6

- [Property reference](../guides/resources--service_policy--reference--group-001.md#canonical-ef66b5745b2615658857f1c4dbfacf1be8619f41e200443445dd2d0c768a74fc)
- [Examples](../guides/resources--service_policy--examples--group-001.md#canonical-68e63fb92b908802c10cf1ef9156256ce1670aa780fe9c56ebd52738f5bb6d39)
- [Import](../guides/resources--service_policy--lifecycle--group-001.md#canonical-c1904e1c8e5cab08ef44523cd83dcf899ea621af041a272714410f3c7190ccc4)
- [Timeouts](../guides/resources--service_policy--lifecycle--group-001.md#canonical-48436e810a4b813a3733ab227538fec8086c14819622f245aa86a813f6afb096)
