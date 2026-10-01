---
page_title: "xcsh_subnet landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet landing."
---

# xcsh_subnet landing

<a id="canonical-285e02e33caffc9b0ccb490d0d8dc76b095484047258d3416444a6aef03e3a41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc24e78f1cacf8e72b106a99a4c61bc5fc040e54e5da3ddf0195f184816473fb"></a>

## xcsh_subnet — xcsh_subnet / 814970675e8b / 2

Breadcrumbs:

- xcsh_subnet

Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an
interface of a vm/pod. it is created in user or shared namespace. configuration.

<a id="canonical-fe2fef8b0958158dc5435e1d030c8e60c1ab679210b1d73cadd765de8c7bc11f"></a>

## Prerequisites — xcsh_subnet / 814970675e8b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3ccda1eb91d85a9f2d39f2be8e4792746f1ed07a5326d7946b2e23f30ae9274b"></a>

## Minimal configuration — xcsh_subnet / 814970675e8b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}
```

<a id="canonical-5170a80b69cde496738c779c55b041741d4227b43745fd5e21b5a0ba248c9274"></a>

## Root configuration — xcsh_subnet / 814970675e8b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-420b6cbf481f9c4d4944e90b823500f8b7b8518ee4e26088b1be787df151ad85"></a>

## Next pages — xcsh_subnet / 814970675e8b / 6

- [Property reference](../guides/resources--subnet--reference--group-001.md#canonical-0af76c82047a5ef49e34e177ac5b41716ecac250d2abccdb52897da51dccc33d)
- [Examples](../guides/resources--subnet--examples--group-001.md#canonical-673f0e65e6cb760bcb080dd331c2cf6d8f2132a3a7647cef482d76d0c71001a7)
- [Import](../guides/resources--subnet--lifecycle--group-001.md#canonical-9fc872f637a726f17717c837d6f9dfdb4bbd0d6945a72c056ad96c2466fe7288)
- [Timeouts](../guides/resources--subnet--lifecycle--group-001.md#canonical-2fdbf6f30a0ff5bbf312928ab0c0121753e7c749f51d3430e515f6c8fdf759fb)
