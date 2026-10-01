---
page_title: "xcsh_network_interface landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface landing."
---

# xcsh_network_interface landing

<a id="canonical-f30a0bb8fdbdd02d9556b2c7b7d0a922bfacce9fb54c00286944679f8a83fef4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3baa3a3950571813011e0ef24c2fc8ca113815b27d90739b18afb28a2506b5ee"></a>

## xcsh_network_interface — xcsh_network_interface / 74badff7b6e0 / 2

Breadcrumbs:

- xcsh_network_interface

Manages a Network Interface resource in F5 Distributed Cloud for network interface represents
configuration of a network device. it is created by users in system namespace. configuration.

<a id="canonical-8b64da3651916f9472fcaa70c3bb37643283c6dbd8f2e57972b308227b463c15"></a>

## Prerequisites — xcsh_network_interface / 74badff7b6e0 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-28f80eb5ef468f4cf26cc9aab5acde3cc615911af39697326ec01782f1eae832"></a>

## Minimal configuration — xcsh_network_interface / 74badff7b6e0 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkInterface Resource Example
# Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkInterface configuration
resource "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}
```

<a id="canonical-bcd50ffefc1543997eab8339c980f618ed3443ae332a1a8bb038bcf84c7bf7b8"></a>

## Root configuration — xcsh_network_interface / 74badff7b6e0 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-f54f4ab54279af3147df2a029acf6530cb8d630500c04b41e74e83c853d4c4b5"></a>

## Next pages — xcsh_network_interface / 74badff7b6e0 / 6

- [Property reference](../guides/resources--network_interface--reference--group-001.md#canonical-ba29b31fc0d930ff348a867c346cc9c8cafc4e5deace11f22012b29e80d5bc66)
- [Examples](../guides/resources--network_interface--examples--group-001.md#canonical-deb23a5ecdf3dfc02a1fdea55794fc92d32fd586e58dcc71f144a1a5b39f9e4d)
- [Import](../guides/resources--network_interface--lifecycle--group-001.md#canonical-e2eb5f8b10aed3dca706cfd1b2ab76093bbe301df458f74109d4ca5955ce5c56)
- [Timeouts](../guides/resources--network_interface--lifecycle--group-001.md#canonical-8fc8628676ba55ab49d94c5e34584deccb7647b2f569c80027a924df38666369)
