---
page_title: "xcsh_cluster landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster landing."
---

# xcsh_cluster landing

<a id="canonical-20bb6d7dccbcd66a532c13c39eef64afca905ef17bcf433f8dc42dffbeda9d03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-875a5fe9af269b3738243b23f3f5440fae0cb1ef625edc6b36c54ed9099f7fb2"></a>

## xcsh_cluster — xcsh_cluster / 4125d25387d1 / 2

Breadcrumbs:

- xcsh_cluster

Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5
Distributed Cloud.

<a id="canonical-01586dd22ceb130e3b8b48eb90027c30a23b500c2e616508da888d7b7594b0ba"></a>

## Prerequisites — xcsh_cluster / 4125d25387d1 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-400809178387a66ddf7f23145be61e9396bc572ff37dee2eccefd1ba00944826"></a>

## Minimal configuration — xcsh_cluster / 4125d25387d1 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cluster configuration
resource "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}
```

<a id="canonical-11854c27dc3363535e4d81317ba32de3f1b6518d5865fcf33317b7378b3536b0"></a>

## Root configuration — xcsh_cluster / 4125d25387d1 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b80fc1e9568d4c5bf5468a6936893a3e66388e27fbde2ca2c1b66f84be8a5c8c"></a>

## Next pages — xcsh_cluster / 4125d25387d1 / 6

- [Property reference](../guides/resources--cluster--reference--group-001.md#canonical-d175f294f6c224794e7bae24d22e1403155d4299cbdbb0d190c275f1d70e2cf9)
- [Examples](../guides/resources--cluster--examples--group-001.md#canonical-2a64e3e2f1d23626e4a5d1d790bae01ba53b4660306296670e9eb3b6281ad329)
- [Import](../guides/resources--cluster--lifecycle--group-001.md#canonical-c553c2d02b8677b6abf5b54aa54df926a0e3be0347f5f97061be1153c7bd8138)
- [Timeouts](../guides/resources--cluster--lifecycle--group-001.md#canonical-09391447a464a24ae12b10aab179d126ef5081cc3ec5b59189d37d0affc8bc36)
