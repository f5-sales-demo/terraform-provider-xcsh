---
page_title: "xcsh_network_interface landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface landing."
---

# xcsh_network_interface landing

<a id="canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8995b4e321af9ab3beb8a46b473103b1a5644fef64b0cc97db3412d6560bd3e4"></a>

## xcsh_network_interface — xcsh_network_interface / b5e997a1e07a / 2

Breadcrumbs:

- xcsh_network_interface

Manages a Network Interface resource in F5 Distributed Cloud for network interface represents
configuration of a network device. it is created by users in system namespace. configuration.

<a id="canonical-42d0ec24bf53f2ee86fa404834ed9167d1ff101a56ef039e377b4186067fb2fe"></a>

## Prerequisites — xcsh_network_interface / b5e997a1e07a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-459317fb786e00a8b2c7e078118f1c2ad08fe650dfb6f05602160ccc43abbe83"></a>

## Minimal configuration — xcsh_network_interface / b5e997a1e07a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkInterface Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkInterface by name
data "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}

output "network_interface_id" {
  value = data.xcsh_network_interface.example.id
}
```

<a id="canonical-b1a1af318a9c5272f8a137c3f0138a03ca27268ffe0e1a115eef05b20193baac"></a>

## Root configuration — xcsh_network_interface / b5e997a1e07a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1262bb557ced0b38752bc671ec10dac724ae9f9769582fde1cac59c181f4b24c"></a>

## Next pages — xcsh_network_interface / b5e997a1e07a / 6

- [Property reference](../guides/data-sources--network_interface--reference--group-001.md#canonical-727acd3bc463226c03ce24381e7e30db3be51700329ff5aea905c4046ff507a4)
- [Examples](../guides/data-sources--network_interface--examples--group-001.md#canonical-9fa7fa1fc3b9dbc53b75c7440e35e8a3e37ba004e065e2deff3541442c1ec5cc)
