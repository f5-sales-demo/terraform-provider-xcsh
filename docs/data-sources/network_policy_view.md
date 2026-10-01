---
page_title: "xcsh_network_policy_view landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view landing."
---

# xcsh_network_policy_view landing

<a id="canonical-5172c506ca44043c027038b651531cb7225374c63ff3675a36a173a5f3eaaa18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b2c8ff7040b431a2d9997f6b4749ac6eca3bb80fdd727355bb5b7421fb5082e"></a>

## xcsh_network_policy_view — xcsh_network_policy_view / f40655bec1ff / 2

Breadcrumbs:

- xcsh_network_policy_view

Manages a Network Policy View resource in F5 Distributed Cloud for network policy view
specification. configuration.

<a id="canonical-5fa05ebd7950b1ff20655fe107bb5e7a3e14caa5c559594f68c7fe246a458831"></a>

## Prerequisites — xcsh_network_policy_view / f40655bec1ff / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-602c47f79f3a28cdec6d0479dd127f2ee3dea64a890e245fbc32adb201be784e"></a>

## Minimal configuration — xcsh_network_policy_view / f40655bec1ff / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyView Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyView by name
data "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.xcsh_network_policy_view.example.id
}
```

<a id="canonical-a1382c2be68c0ccbff6f08f9713d2a4eec5cf97e3c08a78d831fd2214c624e7a"></a>

## Root configuration — xcsh_network_policy_view / f40655bec1ff / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-38ff5aeb6fcaee1979690d6592dd629fb5a9f1eb71fb12514398a855b9ff4d81"></a>

## Next pages — xcsh_network_policy_view / f40655bec1ff / 6

- [Property reference](../guides/data-sources--network_policy_view--reference--group-001.md#canonical-4aac946c6a413f00401908db4f809d0dba762087cdfb8bf7ceb4b1fcb00b6799)
- [Examples](../guides/data-sources--network_policy_view--examples--group-001.md#canonical-3c9c66dfd1fa47b761e332e76bd2fd4801b1dd120cf96d2162cd28460590654d)
