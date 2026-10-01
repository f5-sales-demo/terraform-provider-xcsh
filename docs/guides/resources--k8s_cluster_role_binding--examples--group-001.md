---
page_title: "xcsh_k8s_cluster_role_binding examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding examples."
---

# xcsh_k8s_cluster_role_binding examples

<a id="canonical-e2b707def45e554e19cd5fd74f95944fea5a8e14f31191ebd90db8649358c6fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-810b8f8be55b7504c47d37d6175b80ebc8f06c04afded96a7e4e9af0ee98d212"></a>

## Examples — Examples / e9c42b0ac079 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
- Examples

<a id="canonical-4b49615404afc5d144a19cb9212447157c454fabf5699a2275629833a76800d9"></a>

## Complete configurations — Examples / e9c42b0ac079 / 3

- [Resource](resources--k8s_cluster_role_binding--examples--group-001.md#canonical-485240ba0e9a204c21515680a59e962e7a2a0cad5ab94a4536c2d21bfe979e4b): valid configuration.

<a id="canonical-17a72ffd4ea1f09915d6947774bff43b007c52eb1fdfa98133df24f30cecf3f9"></a>

## Next pages — Examples / e9c42b0ac079 / 4

- [Resource](resources--k8s_cluster_role_binding--examples--group-001.md#canonical-485240ba0e9a204c21515680a59e962e7a2a0cad5ab94a4536c2d21bfe979e4b)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)

<a id="canonical-485240ba0e9a204c21515680a59e962e7a2a0cad5ab94a4536c2d21bfe979e4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baebf76d7ef1db1d85c700eb96b8490d62136c543f4805256e7e14643bf94da7"></a>

## Resource — Resource / 70c69b62ab9c / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
- [Examples](resources--k8s_cluster_role_binding--examples--group-001.md#canonical-e2b707def45e554e19cd5fd74f95944fea5a8e14f31191ebd90db8649358c6fb)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster_role_binding/resource.tf`; digest `sha256:5eb9f86d55919610546b8f56916f90f46bab49e1d41c08733c6886ac7ecd18db`.

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

<a id="canonical-5705ba5c2dab7f1d0d643399e21578bb7e9ee0167c2b2dfb66932992c337d3d7"></a>

## Next pages — Resource / 70c69b62ab9c / 3

- [Examples](resources--k8s_cluster_role_binding--examples--group-001.md#canonical-e2b707def45e554e19cd5fd74f95944fea5a8e14f31191ebd90db8649358c6fb)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md#canonical-63630feae3ae0d980292a5b75763d0fbd203d3335f8367da4d44f60fef31c646)
