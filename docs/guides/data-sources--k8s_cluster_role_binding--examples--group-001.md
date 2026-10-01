---
page_title: "xcsh_k8s_cluster_role_binding examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding examples."
---

# xcsh_k8s_cluster_role_binding examples

<a id="canonical-0f0b3e644267b874e613f4a501fb551d76f2eed900928bc96d27de39d966bd86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfb4204ce5150eca2bc861a6baf15bf595743a93dd8c7120960aeadc118fcd90"></a>

## Examples — Examples / 9fef401ff86c / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
- Examples

<a id="canonical-abcf320455a05aa5b8ec7c36c84a2125da9cb1529c1ca998ed95c4cc6cb31cf9"></a>

## Complete configurations — Examples / 9fef401ff86c / 3

- [Data source](data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-e43a97952fe4fca6d6097ad2d4e3ae376d5979bd73c909ad74abcf3d34a0cccc): valid configuration.

<a id="canonical-067ebebe0480a13ee4f9df73ee023eb802b98dff02b3aa8284a187087af6c493"></a>

## Next pages — Examples / 9fef401ff86c / 4

- [Data source](data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-e43a97952fe4fca6d6097ad2d4e3ae376d5979bd73c909ad74abcf3d34a0cccc)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)

<a id="canonical-e43a97952fe4fca6d6097ad2d4e3ae376d5979bd73c909ad74abcf3d34a0cccc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63e8757abadffb133137bbbd610d35967a87e930b3024c9468fc530d4d28922d"></a>

## Data source — Data source / 9486a457dec4 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
- [Examples](data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-0f0b3e644267b874e613f4a501fb551d76f2eed900928bc96d27de39d966bd86)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster_role_binding/data-source.tf`; digest `sha256:9cfd5eabfe904c328529b82ac99f7f41ca30b8745f9d6eda0a7716bc22bfe86e`.

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

<a id="canonical-f44d27946358371ec3205c978ed7284f75c70b70b75f06efc9bd29abdcb65858"></a>

## Next pages — Data source / 9486a457dec4 / 3

- [Examples](data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-0f0b3e644267b874e613f4a501fb551d76f2eed900928bc96d27de39d966bd86)
- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-89aa41927c7201d228ed76b699b17dcb5b4c9bac74cfeefe90e715c6cd60d347)
