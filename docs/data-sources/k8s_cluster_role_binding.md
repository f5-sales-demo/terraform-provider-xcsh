---
page_title: "xcsh_k8s_cluster_role_binding landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding landing."
---

# xcsh_k8s_cluster_role_binding landing

<a id="canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9af86facd53cad01e98140e9480176c91e6ca28186b253bef547044a33635ce"></a>

## xcsh_k8s_cluster_role_binding — xcsh_k8s_cluster_role_binding / 74a339028749 / 2

Breadcrumbs:

- xcsh_k8s_cluster_role_binding

Manages k8s\_cluster\_role\_binding will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-085674b46f89d05451a8f790000a123e34a34933f72a70b2e2d7a72cdd62044f"></a>

## Prerequisites — xcsh_k8s_cluster_role_binding / 74a339028749 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1cd19703ee9a3f29318836ad84e1c9eb304df9a8ee4c700d5e61e4a622ab94ee"></a>

## Minimal configuration — xcsh_k8s_cluster_role_binding / 74a339028749 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SClusterRoleBinding Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SClusterRoleBinding by name
data "xcsh_k8s_cluster_role_binding" "example" {
  name      = "example-k8s-cluster-role-binding"
  namespace = "staging"
}

output "k8s_cluster_role_binding_id" {
  value = data.xcsh_k8s_cluster_role_binding.example.id
}
```

<a id="canonical-98a39a6096be0c9a4c925b28f042ace25fbffe3258c092811934510408832c40"></a>

## Root configuration — xcsh_k8s_cluster_role_binding / 74a339028749 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e12fffdfbc26a53ee3d1f3cf796f263d8922cc2a39fb1ffd090950673f7a55da"></a>

## Next pages — xcsh_k8s_cluster_role_binding / 74a339028749 / 6

- [Property reference](../guides/data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-bc8de5758cacd31e5791b1d4bd3baf93288975699ce14c1facc574a0223a8a97)
- [Examples](../guides/data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-0f0b3e644267b874e613f4a501fb551d76f2eed900928bc96d27de39d966bd86)
