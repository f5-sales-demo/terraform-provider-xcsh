---
page_title: "xcsh_k8s_cluster landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster landing."
---

# xcsh_k8s_cluster landing

<a id="canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-844cae08514a1bd6f19e3b65d794b3f899fb6a8fde4d20ec61902e70b3cd3986"></a>

## xcsh_k8s_cluster — xcsh_k8s_cluster / 4764cd801e31 / 2

Breadcrumbs:

- xcsh_k8s_cluster

Manages k8s\_cluster will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-050482e5c081eef278c659f0bef41e9bf054d27a516209ee304b7e63a569d7d0"></a>

## Prerequisites — xcsh_k8s_cluster / 4764cd801e31 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1a271dce85acc7a28af7e7a56fd01aa742c469a2a5c8d99c1ff63751c5d2040f"></a>

## Minimal configuration — xcsh_k8s_cluster / 4764cd801e31 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SCluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SCluster by name
data "xcsh_k8s_cluster" "example" {
  name      = "example-k8s-cluster"
  namespace = "system"
}

output "k8s_cluster_id" {
  value = data.xcsh_k8s_cluster.example.id
}
```

<a id="canonical-a6b86d450e9df44e4296a2a7a07c28faca4b6ca478927c3eb9051015747ff20b"></a>

## Root configuration — xcsh_k8s_cluster / 4764cd801e31 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-e78123e1f2bd685d198a748426d6694c67f201015e74a679fd1ab8ed92bb5f62"></a>

## Next pages — xcsh_k8s_cluster / 4764cd801e31 / 6

- [Property reference](../guides/data-sources--k8s_cluster--reference--group-001.md#canonical-cfb0210391a5f24aede2a50fa5e556e1d6ccaf8fcc953ef127edc76145098620)
- [Examples](../guides/data-sources--k8s_cluster--examples--group-001.md#canonical-70d43c25095d127645acfd8acd8ab5cca96170f2dc83264bcc3ed7b0b4e0de2e)
