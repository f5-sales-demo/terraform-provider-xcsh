---
page_title: "xcsh_k8s_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster examples."
---

# xcsh_k8s_cluster examples

<a id="canonical-369842889ab7ce2a0ae26413e95efe2ebe911e63b57fd9106d5a694ca13ad942"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ad10df98da3608d356d79bc42cf18184903f784d8429fceb5e1b68ac1b33076"></a>

## Examples — Examples / 4c1a75c6b326 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- Examples

<a id="canonical-5b27df54c63a57b09e9ae9befd21b4c160af364a5642b047b996b1decb1f2ae6"></a>

## Complete configurations — Examples / 4c1a75c6b326 / 3

- [Resource](resources--k8s_cluster--examples--group-001.md#canonical-cfc278ec72834cc5edf6e0ff4c5f3afac3790d16ddadbe44fe828afae4d872ed): valid configuration.

<a id="canonical-65e7019a73c01686a7ece8bbf56fce65e1bce614a839a240c06fbf70c8de7edf"></a>

## Next pages — Examples / 4c1a75c6b326 / 4

- [Resource](resources--k8s_cluster--examples--group-001.md#canonical-cfc278ec72834cc5edf6e0ff4c5f3afac3790d16ddadbe44fe828afae4d872ed)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)

<a id="canonical-cfc278ec72834cc5edf6e0ff4c5f3afac3790d16ddadbe44fe828afae4d872ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05b7c325cad5d58e240d78f86b7725f4d14fd7e97cef24e2c2418baaaf2254c2"></a>

## Resource — Resource / f4983b94444b / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
- [Examples](resources--k8s_cluster--examples--group-001.md#canonical-369842889ab7ce2a0ae26413e95efe2ebe911e63b57fd9106d5a694ca13ad942)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster/resource.tf`; digest `sha256:4b234e2e10f8643667634d9286b365cede58a97ccb384ebff96249fbe01b28a1`.

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

<a id="canonical-891c332de5647a5d4f706fb61c7326779e59b9c182bd8de152a14fc77e9770fd"></a>

## Next pages — Resource / f4983b94444b / 3

- [Examples](resources--k8s_cluster--examples--group-001.md#canonical-369842889ab7ce2a0ae26413e95efe2ebe911e63b57fd9106d5a694ca13ad942)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-2f8408128bac4154d293c68f236e0ceda708f0d32acbd5c5f0e51cfa3164bce6)
