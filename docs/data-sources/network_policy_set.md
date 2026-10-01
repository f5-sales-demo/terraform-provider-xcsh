---
page_title: "xcsh_network_policy_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_set landing."
---

# xcsh_network_policy_set landing

<a id="canonical-ca548052a89d7f647ebf8e8303af31a758be406e22a36f86b6a13ff023dd5e7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-455fb5e1d63b2f1ad9207c59a41dad8a7af9703fca19e643f923e9845d45ef1f"></a>

## xcsh_network_policy_set — xcsh_network_policy_set / bd04e482b5d0 / 2

Breadcrumbs:

- xcsh_network_policy_set

Manages a Network Policy Set resource in F5 Distributed Cloud for get network policy set in a given
namespace. configuration. (read-only data source)

<a id="canonical-82321fdfa6f29e88cd3e6afd43443ba30805a8f9711d5185cc42a5774a0b3261"></a>

## Prerequisites — xcsh_network_policy_set / bd04e482b5d0 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-17b0defec152c3fa9e8519cfedfc15c36b3878a58cab5d1510afe4f3efd46dc6"></a>

## Minimal configuration — xcsh_network_policy_set / bd04e482b5d0 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicySet by name
data "xcsh_network_policy_set" "example" {
  name      = "example-network-policy-set"
  namespace = "staging"
}

output "network_policy_set_id" {
  value = data.xcsh_network_policy_set.example.id
}
```

<a id="canonical-13463d2180c4ebe426a1c4fade06659bc4d7de995d7eb28f324b1134af37c3dd"></a>

## Root configuration — xcsh_network_policy_set / bd04e482b5d0 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ce7641c1066ee0b987d6ae31c5ba30bd6f65280a2b16fe5c49bc7439ee77f9ae"></a>

## Next pages — xcsh_network_policy_set / bd04e482b5d0 / 6

- [Property reference](../guides/data-sources--network_policy_set--reference--group-001.md#canonical-e2fecfde17f2b3ba8ddee21eece6bfcd1b43c5e09d0ec7715a8cabd0437a8cca)
- [Examples](../guides/data-sources--network_policy_set--examples--group-001.md#canonical-6e5d96c0f0c95a4764658184ce943e95d603c429be3d17a52e3aa762848638d6)
