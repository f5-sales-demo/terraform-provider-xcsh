---
page_title: "xcsh_k8s_cluster_role landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role landing."
---

# xcsh_k8s_cluster_role landing

<a id="canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe31cd270a3a1efacb3a25ca49fe1c277a94d543eefe972c6a478f5e146fe38d"></a>

## xcsh_k8s_cluster_role — xcsh_k8s_cluster_role / bec782ffc216 / 2

Breadcrumbs:

- xcsh_k8s_cluster_role

Manages k8s\_cluster\_role will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-ec23eef6487ac4de9c53ce3fac793686d9680d5a8e16d1032dd89e35c90e94ef"></a>

## Prerequisites — xcsh_k8s_cluster_role / bec782ffc216 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-bf14b927d3a6dc45c29c14cd0606b2d601a3947c2f09993165ea7dbdc8a2546c"></a>

## Minimal configuration — xcsh_k8s_cluster_role / bec782ffc216 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SClusterRole Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SClusterRole by name
data "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}

output "k8s_cluster_role_id" {
  value = data.xcsh_k8s_cluster_role.example.id
}
```

<a id="canonical-6764c27be5eec382072ad7cfe72643d1017275dac2b38cb551fc1e9a385eca4e"></a>

## Root configuration — xcsh_k8s_cluster_role / bec782ffc216 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-9cdee0f42aea58e77c95cb43a686cd0fc7db0134962ab5a8e33943e80ec5f4ce"></a>

## Next pages — xcsh_k8s_cluster_role / bec782ffc216 / 6

- [Property reference](../guides/data-sources--k8s_cluster_role--reference--group-001.md#canonical-81d3736ec3cd7378005627ab8c00194ddb3f8ab7f6499d71f90ba016ef6a2050)
- [Examples](../guides/data-sources--k8s_cluster_role--examples--group-001.md#canonical-8882b7df6715e590e1c8e903bdfa86640fea7361d10c232053079c13b6107864)
