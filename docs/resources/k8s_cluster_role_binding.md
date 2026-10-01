---
page_title: "xcsh_k8s_cluster_role_binding landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding landing."
---

# xcsh_k8s_cluster_role_binding landing

<a id="canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6e5e458173570530390c9daf173a7cdbd6e7f62a4b1eeb9a7d7464772160b0d"></a>

## xcsh_k8s_cluster_role_binding — xcsh_k8s_cluster_role_binding / b1a9d19d0fc8 / 2

Breadcrumbs:

- xcsh_k8s_cluster_role_binding

Manages k8s\_cluster\_role\_binding will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-854c73d5c68d4cd79d39613c3e59197eda755144cff135d03cf20fca795356f5"></a>

## Prerequisites — xcsh_k8s_cluster_role_binding / b1a9d19d0fc8 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1a0729cecc06618c973ffd106bc68db229028addbab56551e4f814ae039055aa"></a>

## Minimal configuration — xcsh_k8s_cluster_role_binding / b1a9d19d0fc8 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SClusterRoleBinding Resource Example
# Manages k8s_cluster_role_binding will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SClusterRoleBinding configuration
resource "xcsh_k8s_cluster_role_binding" "example" {
  name      = "example-k8s-cluster-role-binding"
  namespace = "staging"
}
```

<a id="canonical-71369c1da75a5ee53d4b7347a8e3de272669462f87832abc6273df95db60ff5e"></a>

## Root configuration — xcsh_k8s_cluster_role_binding / b1a9d19d0fc8 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-627c3713faec59c6fd26162638a9bd2a420390d5f8f11581042f33d45406c52b"></a>

## Next pages — xcsh_k8s_cluster_role_binding / b1a9d19d0fc8 / 6

- [Property reference](../guides/resources--k8s_cluster_role_binding--reference--group-001.md#canonical-daa8129e558b368a8f6eac480ce48c930cc68a1d199f4322ed9f9803d63de406)
- [Examples](../guides/resources--k8s_cluster_role_binding--examples--group-001.md#canonical-e2b707def45e554e19cd5fd74f95944fea5a8e14f31191ebd90db8649358c6fb)
- [Import](../guides/resources--k8s_cluster_role_binding--lifecycle--group-001.md#canonical-a0525dc51116c91df4d00e06f2b005bbeb21a648c76000fae70e74fe444797a5)
- [Timeouts](../guides/resources--k8s_cluster_role_binding--lifecycle--group-001.md#canonical-f3b36fbe368fdd8d8f3910de95c8f7ed0b3bfdbf16ce04191809760d19522691)
