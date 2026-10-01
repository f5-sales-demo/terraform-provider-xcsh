---
page_title: "xcsh_infraprotect_synchronize_configuration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_synchronize_configuration landing."
---

# xcsh_infraprotect_synchronize_configuration landing

<a id="canonical-c3e1b3aa794f6032e8b56effdea59b5c380dd1c1dc0035df959f94ebabf3a655"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9ec3f03b1c9c312b9886dd1b2d56dc0ebbed8da2c47b04b5e00a2040a7e57b3"></a>

## xcsh_infraprotect_synchronize_configuration — xcsh_infraprotect_synchronize_configuration / b62a90890173 / 2

Breadcrumbs:

- xcsh_infraprotect_synchronize_configuration

Resource creation operation.

<a id="canonical-c042f0398e834039ac4ffa737af6c4b8865fb8b22ad41f7f63790b22a5008454"></a>

## Prerequisites — xcsh_infraprotect_synchronize_configuration / b62a90890173 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-37941bc584b6bbc0f40fc529dc1de4ae2cd80e3a90fd9811da861373b451d2fe"></a>

## Minimal configuration — xcsh_infraprotect_synchronize_configuration / b62a90890173 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectSynchronizeConfiguration Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_infraprotect_synchronize_configuration" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-b885be67b42633511f7e955e11a2e381129ffe1cc5c4cb4d2108ec7033eb1086"></a>

## Root configuration — xcsh_infraprotect_synchronize_configuration / b62a90890173 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2877d9726a8b3afd7828972fe4de83bf04305ce17ceacc059a62b47af6cb0c5b"></a>

## Next pages — xcsh_infraprotect_synchronize_configuration / b62a90890173 / 6

- [Property reference](../guides/actions--infraprotect_synchronize_configuration--reference--group-001.md#canonical-426a080eaa64a02bca284f5524e405a95c9484c4e5bde17ebe9bc6b5f8657438)
- [Examples](../guides/actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-592bbdecfb48f78b5766a613695205d3e5f9dd31e0ac81e921edcb2d1794284c)
- [Lifecycle](../guides/actions--infraprotect_synchronize_configuration--lifecycle--group-001.md#canonical-fed5546b597259316f2e9fae36d81929c34359398abe5239921f138a66d14df9)
