---
page_title: "xcsh_k8s_cluster landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster landing."
---

# xcsh_k8s_cluster landing

<a id="canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23d930555a32057cb56f5a4f45ca885f442834c132f2e0393b8527cab40b1c3c"></a>

## xcsh_k8s_cluster — xcsh_k8s_cluster / 483d5c44c02b / 2

Breadcrumbs:

- xcsh_k8s_cluster

Manages k8s\_cluster will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-416f5f1607f8b8e44327d6c7fb0ca1ce54f3eb42e8f0cb6ce52f30f76474a923"></a>

## Prerequisites — xcsh_k8s_cluster / 483d5c44c02b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a41d1995225891a0ff32a47b8a79b2cd8f11225d4d6c7d1ef8fddefdf9bec42d"></a>

## Minimal configuration — xcsh_k8s_cluster / 483d5c44c02b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SCluster Resource Example
# Manages k8s_cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SCluster configuration
resource "xcsh_k8s_cluster" "example" {
  name      = "example-k8s-cluster"
  namespace = "system"
}
```

<a id="canonical-1cd1bfb0b91934f1caefec389f1dbe1a0cc5053780740e8f9273bec57430b45c"></a>

## Root configuration — xcsh_k8s_cluster / 483d5c44c02b / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-ecb8e23332d8dbe2881a7e589f4b93d247831ab5d1e0fdc5c4940527a7521891"></a>

## Next pages — xcsh_k8s_cluster / 483d5c44c02b / 6

- [Property reference](../guides/resources--k8s_cluster--reference--group-001.md#canonical-90f1e9cc519426a194ab7e4acdb90f791c068205f7aaf5c13bfe9ca7fe982476)
- [Examples](../guides/resources--k8s_cluster--examples--group-001.md#canonical-369842889ab7ce2a0ae26413e95efe2ebe911e63b57fd9106d5a694ca13ad942)
- [Import](../guides/resources--k8s_cluster--lifecycle--group-001.md#canonical-b2d53ccfb7f4a2575f10ed7cb7d637dde88cbd5605e1499816a8e110abf8a198)
- [Timeouts](../guides/resources--k8s_cluster--lifecycle--group-001.md#canonical-f9e06870da9ab5fdd76f79578308fd2f79e3342c8398d21305f34047939b8dbe)
