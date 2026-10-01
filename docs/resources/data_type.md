---
page_title: "xcsh_data_type landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type landing."
---

# xcsh_data_type landing

<a id="canonical-1cff15081e577df3cc4acee98c37a779281dfe15bfd78338a144f83da6db7bdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c828ab2119454762d59faa050a520eeb43b4f3a4afb35ceb55988a6ea362b99c"></a>

## xcsh_data_type — xcsh_data_type / 71f312f3765b / 2

Breadcrumbs:

- xcsh_data_type

Manages data\_type creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-b66aa8c4f5464967cb4f64fedbfc8c04b9ef20a3d491374bd88a624ec6c844e5"></a>

## Prerequisites — xcsh_data_type / 71f312f3765b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-26f36bfdb65983a81ad5ccc49bc1000f420c2bb764763cda0437fe028b6dbf8e"></a>

## Minimal configuration — xcsh_data_type / 71f312f3765b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```

<a id="canonical-a1c887252fa0bfdd863b68a938fa9661768de0e946566ff5913bcf5140562811"></a>

## Root configuration — xcsh_data_type / 71f312f3765b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-290bbefabf77a76f0d761b29a99db10f096edb4648c5645cfd0faf0701e84563"></a>

## Next pages — xcsh_data_type / 71f312f3765b / 6

- [Property reference](../guides/resources--data_type--reference--group-001.md#canonical-0b4bcef4cf9c0ef8b8a8c71ba751ea78d2753b38ce148712e866e6dc49645b88)
- [Examples](../guides/resources--data_type--examples--group-001.md#canonical-f6fcb5b6928c9a403046c371b1e2041d0b724dd72cc231fccb0851e9f5346050)
- [Import](../guides/resources--data_type--lifecycle--group-001.md#canonical-9a35c59b2f222d3733cc3f5fed11513984d6ffddf3e84f238fcea63a275f41d3)
- [Timeouts](../guides/resources--data_type--lifecycle--group-001.md#canonical-3cbbb8c92b860749df4036ae205b8b2f50e705acdd543f1dd527b92118162937)
