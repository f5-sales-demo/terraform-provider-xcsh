---
page_title: "xcsh_cluster landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster landing."
---

# xcsh_cluster landing

<a id="canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b04a3afc1025262f95765899c77c9052e1bdc2c88a1adb0ceb412e500e7c77ee"></a>

## xcsh_cluster — xcsh_cluster / 0857e8915c44 / 2

Breadcrumbs:

- xcsh_cluster

Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5
Distributed Cloud.

<a id="canonical-0d7f45b6741dec2dacafd59a2cb639d5020f37254a825d48f606b1af7589c4c1"></a>

## Prerequisites — xcsh_cluster / 0857e8915c44 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-16cb57c3903de6d4e861a83f9422b675bc0119daba6539bea0f4513646f6a74c"></a>

## Minimal configuration — xcsh_cluster / 0857e8915c44 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cluster by name
data "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}

output "cluster_id" {
  value = data.xcsh_cluster.example.id
}
```

<a id="canonical-d0b34590cb108d6aa9049b0e13699394ed0e9a8802ef09405de708fa61e38a12"></a>

## Root configuration — xcsh_cluster / 0857e8915c44 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a5a2c13927c16e7dd3cefba055659a1ccfbee94428364b758c5bd2745bfd5e8d"></a>

## Next pages — xcsh_cluster / 0857e8915c44 / 6

- [Property reference](../guides/data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [Examples](../guides/data-sources--cluster--examples--group-001.md#canonical-2d44b777a62a6d2fa2ade46bd69fd423ae33771c7c995f33d89a240f913d68b1)
