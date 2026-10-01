---
page_title: "xcsh_sensitive_data_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy landing."
---

# xcsh_sensitive_data_policy landing

<a id="canonical-097e4ed086586b49bc303a1ea5095eb9ea6db5c59db3d4ba9f0f6b743c2b8d44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b88c2478392d1cc308fffa7db46ad339a3bfadd17ee70c22ede5e78c3b14666"></a>

## xcsh_sensitive_data_policy — xcsh_sensitive_data_policy / e2152190c13e / 2

Breadcrumbs:

- xcsh_sensitive_data_policy

Manages sensitive\_data\_policy creates a new object in the storage backend for metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-0159305efd45360324bd62f9f3409954275af2fe2d38b460becac900e1e41beb"></a>

## Prerequisites — xcsh_sensitive_data_policy / e2152190c13e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-54f16b9504e719c3414317cacd874508afe8a3fa4b20e582f056d4625391c2ea"></a>

## Minimal configuration — xcsh_sensitive_data_policy / e2152190c13e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SensitiveDataPolicy Resource Example
# Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SensitiveDataPolicy configuration
resource "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}
```

<a id="canonical-54b3f1b4ea96156536f485288fc3a5d644231902b802082fe25c558e38316f67"></a>

## Root configuration — xcsh_sensitive_data_policy / e2152190c13e / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-7be144d3558f1d13df7882ecc8444be4c181e6c2a8e872876b2e475063f3e480"></a>

## Next pages — xcsh_sensitive_data_policy / e2152190c13e / 6

- [Property reference](../guides/resources--sensitive_data_policy--reference--group-001.md#canonical-237e6c635e5bbef2763802cde769bc7b230606d9ae2fd8af6641d74700f6ef41)
- [Examples](../guides/resources--sensitive_data_policy--examples--group-001.md#canonical-9a55abc4607becc0302d726239e05592c1088a771f57f7445193af0f804f2f8b)
- [Import](../guides/resources--sensitive_data_policy--lifecycle--group-001.md#canonical-4e0bc2c2901630fb7387023a2c9a4ea5f79de443965a790626d296ec8198f1c2)
- [Timeouts](../guides/resources--sensitive_data_policy--lifecycle--group-001.md#canonical-7c18416fab5cab2cc5ef86360951b23d19f759ac0f14f0005bb75bb971763808)
