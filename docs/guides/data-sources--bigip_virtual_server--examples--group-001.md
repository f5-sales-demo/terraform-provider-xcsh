---
page_title: "xcsh_bigip_virtual_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_virtual_server examples."
---

# xcsh_bigip_virtual_server examples

<a id="canonical-803e305ddcc14d26305861439a5a6225a5a00bb0c19fe90bd14bec7fbdf53bd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4ce562b7c526a5a6d42e4855fcde109d0c48898b002a7e68d66f7947d688a46"></a>

## Examples — Examples / 136bf6117449 / 2

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md#canonical-5de83ca3dfb9cf73d57be7ed86f617e30ad06fc103f2f9653fda1240ecf86f88)
- Examples

<a id="canonical-4c32a967fad5d24d84de8978c5e0879f2184e8c917346ff4f7e86586bb969b20"></a>

## Complete configurations — Examples / 136bf6117449 / 3

- [Data source](data-sources--bigip_virtual_server--examples--group-001.md#canonical-b144fc5a7d98d4119da47a8875c459ea4595fdcb0a684835ea552f061a9d954a): valid configuration.

<a id="canonical-b55a262a5ce7f8be974fb494b44e8598ab73a59178f05a8362f481f3752fa9a0"></a>

## Next pages — Examples / 136bf6117449 / 4

- [Data source](data-sources--bigip_virtual_server--examples--group-001.md#canonical-b144fc5a7d98d4119da47a8875c459ea4595fdcb0a684835ea552f061a9d954a)
- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md#canonical-5de83ca3dfb9cf73d57be7ed86f617e30ad06fc103f2f9653fda1240ecf86f88)

<a id="canonical-b144fc5a7d98d4119da47a8875c459ea4595fdcb0a684835ea552f061a9d954a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c6e2b58929d014ccd17481a5a247189f729aacc0ee8fbc019cb9b68c905f6e8"></a>

## Data source — Data source / 92f90a6c7a97 / 2

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md#canonical-5de83ca3dfb9cf73d57be7ed86f617e30ad06fc103f2f9653fda1240ecf86f88)
- [Examples](data-sources--bigip_virtual_server--examples--group-001.md#canonical-803e305ddcc14d26305861439a5a6225a5a00bb0c19fe90bd14bec7fbdf53bd9)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bigip_virtual_server/data-source.tf`; digest `sha256:4b98a932b2c1a0b5e2b9c608b25b6b3f74aea2df606bdee9675da50e166ce11d`.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```

<a id="canonical-59c32cba5bfc5e0193bdb244f35a5869fb26522b067419049e5e030e716d466d"></a>

## Next pages — Data source / 92f90a6c7a97 / 3

- [Examples](data-sources--bigip_virtual_server--examples--group-001.md#canonical-803e305ddcc14d26305861439a5a6225a5a00bb0c19fe90bd14bec7fbdf53bd9)
- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md#canonical-5de83ca3dfb9cf73d57be7ed86f617e30ad06fc103f2f9653fda1240ecf86f88)
