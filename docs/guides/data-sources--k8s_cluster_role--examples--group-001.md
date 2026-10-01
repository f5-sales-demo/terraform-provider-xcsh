---
page_title: "xcsh_k8s_cluster_role examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role examples."
---

# xcsh_k8s_cluster_role examples

<a id="canonical-8882b7df6715e590e1c8e903bdfa86640fea7361d10c232053079c13b6107864"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05f3b2e40cb0e8e523abbe3df0e86f83c7df24398cd7aadd1208a90f288b4264"></a>

## Examples — Examples / 4c4fdcc3b8bb / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- Examples

<a id="canonical-53f61b2c9ef533442a10787c7f273b10778f1af1ee35742d9b6df99c79cfecc0"></a>

## Complete configurations — Examples / 4c4fdcc3b8bb / 3

- [Data source](data-sources--k8s_cluster_role--examples--group-001.md#canonical-664064eb0b104e7e0d750f54874fb3eee5e25e2668542df12e13900201b69efd): valid configuration.

<a id="canonical-f4a7b98cb7dee2c5e105cb4422628bf2bacce63320361741201c9c8fb01db571"></a>

## Next pages — Examples / 4c4fdcc3b8bb / 4

- [Data source](data-sources--k8s_cluster_role--examples--group-001.md#canonical-664064eb0b104e7e0d750f54874fb3eee5e25e2668542df12e13900201b69efd)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)

<a id="canonical-664064eb0b104e7e0d750f54874fb3eee5e25e2668542df12e13900201b69efd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3de0bf6eb3b84e10f976df47918ce6e2bbb78dc520b5c5e304938c7200ba1b0"></a>

## Data source — Data source / 1529230ce441 / 2

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
- [Examples](data-sources--k8s_cluster_role--examples--group-001.md#canonical-8882b7df6715e590e1c8e903bdfa86640fea7361d10c232053079c13b6107864)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster_role/data-source.tf`; digest `sha256:3a2a8196bb01ed845c9acfb23f02e2761ef76a0ab6bfc6ad04f8767ce9170112`.

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

<a id="canonical-167474857f69867cea1381311c7dd6a0f8bb78b08dc5db842448a878f4390baa"></a>

## Next pages — Data source / 1529230ce441 / 3

- [Examples](data-sources--k8s_cluster_role--examples--group-001.md#canonical-8882b7df6715e590e1c8e903bdfa86640fea7361d10c232053079c13b6107864)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md#canonical-ae676024de71b8439d6e1da589b80d914e1b787ad3c0876d229fcc1a32cbbeda)
