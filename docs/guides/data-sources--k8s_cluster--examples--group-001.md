---
page_title: "xcsh_k8s_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster examples."
---

# xcsh_k8s_cluster examples

<a id="canonical-70d43c25095d127645acfd8acd8ab5cca96170f2dc83264bcc3ed7b0b4e0de2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcf8e28532f8252171fe0134f2d386c412d35cfcc6d6ec4bf25a9b86d2f0f163"></a>

## Examples — Examples / 4535c15ab65d / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- Examples

<a id="canonical-aac6c780c3f5a5cfea83f5a8fbdec8745c02dde7aaac7ef1ed9665dd01bf3b38"></a>

## Complete configurations — Examples / 4535c15ab65d / 3

- [Data source](data-sources--k8s_cluster--examples--group-001.md#canonical-d15829148bcba5f4c15a2d189a9e7beaa4fb649e4cc0231ddc5f16d83a3e4a40): valid configuration.

<a id="canonical-ae0bedfb344d5ff76411603d5a7f2db0cd64639df1c67c4c727b24630abcc4e1"></a>

## Next pages — Examples / 4535c15ab65d / 4

- [Data source](data-sources--k8s_cluster--examples--group-001.md#canonical-d15829148bcba5f4c15a2d189a9e7beaa4fb649e4cc0231ddc5f16d83a3e4a40)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)

<a id="canonical-d15829148bcba5f4c15a2d189a9e7beaa4fb649e4cc0231ddc5f16d83a3e4a40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebbda8064b29bd35eae183a01ea829a21927983af7ab18fbbe1cd29ae800cd62"></a>

## Data source — Data source / 853a5ca461f8 / 2

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
- [Examples](data-sources--k8s_cluster--examples--group-001.md#canonical-70d43c25095d127645acfd8acd8ab5cca96170f2dc83264bcc3ed7b0b4e0de2e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster/data-source.tf`; digest `sha256:1ffe54fa3443473ef381429123673c75783a908cc4dae691369d96028c415887`.

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

<a id="canonical-ce343a782060c8d912166fcfcf5f717f7eae15b6fc0ec12c1e3610974487aa89"></a>

## Next pages — Data source / 853a5ca461f8 / 3

- [Examples](data-sources--k8s_cluster--examples--group-001.md#canonical-70d43c25095d127645acfd8acd8ab5cca96170f2dc83264bcc3ed7b0b4e0de2e)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-80d7cda51b5cbe0a9301672f8d5c626a79fc062e8e044ad4a6d04ad0a8e66989)
