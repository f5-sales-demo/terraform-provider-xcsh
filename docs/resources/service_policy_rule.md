---
page_title: "xcsh_service_policy_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule landing."
---

# xcsh_service_policy_rule landing

<a id="canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e1cfa7dacf92e693f39dbe2562be442378cdfda00f1449a2d47daf1d6ce9adb"></a>

## xcsh_service_policy_rule — xcsh_service_policy_rule / 789beaa68e89 / 2

Breadcrumbs:

- xcsh_service_policy_rule

Manages service\_policy\_rule creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-8500285005e6798324bf0b69faada56d1e09b61a3ed5b1e00f4ab51b5752cb85"></a>

## Prerequisites — xcsh_service_policy_rule / 789beaa68e89 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-20040b58912c5e522e65837e1e8acc146fc216623ea11494272e5c6f694211b7"></a>

## Minimal configuration — xcsh_service_policy_rule / 789beaa68e89 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicyRule Resource Example
# Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicyRule configuration
resource "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"

  action = "DENY"
}
```

<a id="canonical-9abd25c0c78d8b23baeb936b1ed8421548e5f8897f2c8d27c73acb6f3762a215"></a>

## Root configuration — xcsh_service_policy_rule / 789beaa68e89 / 5

Required root properties: `action`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-aca2cfe3281d653d84af59bbe7cd7d714072012924e3f390dd29ecc82dc635d4"></a>

## Next pages — xcsh_service_policy_rule / 789beaa68e89 / 6

- [Property reference](../guides/resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [Examples](../guides/resources--service_policy_rule--examples--group-001.md#canonical-0b47f8afb67032c78d6321166b882fd0532f386cde338c2c92d9c8bc15c25118)
- [Import](../guides/resources--service_policy_rule--lifecycle--group-001.md#canonical-e22c3b7aed96d383a4fa826a8df645de490a6ffed0d96920b265b1a68498a69a)
- [Timeouts](../guides/resources--service_policy_rule--lifecycle--group-001.md#canonical-3ede20bdfc6da2ff9cee70c702369e73620e706b61a92d557a7b5a480d354b18)
