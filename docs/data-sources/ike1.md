---
page_title: "xcsh_ike1 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 landing."
---

# xcsh_ike1 landing

<a id="canonical-9615a09559c55d9ac9b674d5ddca753d31c2214560e2a1d2f9188e7f403124f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a8f087b37bac80fc6d5e7d4007ee3a3c294d49f6bbcdbccffab0a2925956a30"></a>

## xcsh_ike1 — xcsh_ike1 / 089ff9747364 / 2

Breadcrumbs:

- xcsh_ike1

Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.

<a id="canonical-3d40feb6439bcaef3d639eff03e1bb4d8d2c704ddea12285d4a829a261ae4ac3"></a>

## Prerequisites — xcsh_ike1 / 089ff9747364 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6713b716178771a58d1b3ad32e5809120f35adf142e4efa4bb76bc3905db062c"></a>

## Minimal configuration — xcsh_ike1 / 089ff9747364 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike1 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike1 by name
data "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}

output "ike1_id" {
  value = data.xcsh_ike1.example.id
}
```

<a id="canonical-6d713358742960ad0aefdace8874a8564b1530e9fb939cc59b0389154752df32"></a>

## Root configuration — xcsh_ike1 / 089ff9747364 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9d8afa5589bcc9b762464191b088bd98865e8a3baa0800839bbe0748e4004c59"></a>

## Next pages — xcsh_ike1 / 089ff9747364 / 6

- [Property reference](../guides/data-sources--ike1--reference--group-001.md#canonical-976497673af1903f2e4f1926d15b090ed29d1eb43535659975e8253a37136f07)
- [Examples](../guides/data-sources--ike1--examples--group-001.md#canonical-defb7512b14d447bc7b57fa79136b54efef5377b69c4d47d608382ae823c9fe4)
