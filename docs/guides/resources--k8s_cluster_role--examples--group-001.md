---
page_title: "xcsh_k8s_cluster_role examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role examples."
---

# xcsh_k8s_cluster_role examples

<a id="canonical-ce3a232593bab97db089ecb898a23d4a4451bfae8bbddb21d8687777b9b1d0e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df87cecaab6b8cf27b9c6eecbc8c2d2b17fe058b57a92f1453fd2301421eff3e"></a>

## Examples — Examples / d13a3e39b00c / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- Examples

<a id="canonical-ea2c97d3a866fc1cda19073c6d3bfddaf6a035a97eb9917a8336512fb175615d"></a>

## Complete configurations — Examples / d13a3e39b00c / 3

- [Resource](resources--k8s_cluster_role--examples--group-001.md#canonical-2404536fa733c2cf2c3a755008579f8759e17c856d787ce654adedf6745387d1): valid configuration.

<a id="canonical-06d069c9c533906e41ab6f7a92d4e382d6fdc05fb54eda5c90d56acc258bdd5f"></a>

## Next pages — Examples / d13a3e39b00c / 4

- [Resource](resources--k8s_cluster_role--examples--group-001.md#canonical-2404536fa733c2cf2c3a755008579f8759e17c856d787ce654adedf6745387d1)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)

<a id="canonical-2404536fa733c2cf2c3a755008579f8759e17c856d787ce654adedf6745387d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dcb7df7d310c01bcc1e0ac6bebecd76f5379081b9825be7ebd29a370ef3c430"></a>

## Resource — Resource / 9ccabd03a69e / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
- [Examples](resources--k8s_cluster_role--examples--group-001.md#canonical-ce3a232593bab97db089ecb898a23d4a4451bfae8bbddb21d8687777b9b1d0e0)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster_role/resource.tf`; digest `sha256:f4470002db951c71b74b86c89c38ea07204065a7ca36bf4ba1c80737618e16ac`.

```terraform
# K8SClusterRole Resource Example
# Manages k8s_cluster_role will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SClusterRole configuration
resource "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}
```

<a id="canonical-2fd76bd7429bfa6e02d3528bfc7b962c854edfccd60367789a3c6ba541b53488"></a>

## Next pages — Resource / 9ccabd03a69e / 3

- [Examples](resources--k8s_cluster_role--examples--group-001.md#canonical-ce3a232593bab97db089ecb898a23d4a4451bfae8bbddb21d8687777b9b1d0e0)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b)
