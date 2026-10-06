---
page_title: "xcsh_k8s_cluster_role_binding examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding examples."
---

# xcsh_k8s_cluster_role_binding examples

<a id="canonical-0033002303321210-1002121323201310-3212010333102211-0001332311110131-1312330232323121-0000210220233021-1231021331320321-3121121223312012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)
- Examples

<a id="canonical-3133231002001030-3211011100323022-0223302012012212-2322330111233311-2111131003222103-3131203013010200-2112002232223130-0101203330312100"></a>

### Complete configurations for `xcsh_k8s_cluster_role_binding`

- [Data source](data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-3210032221132111-0233321033302212-3112002113223102-3110320322320313-1231112113212331-1303302100212231-1310222330330331-0310220030303030): valid configuration.

<a id="canonical-3210032221132111-0233321033302212-3112002113223102-3110320322320313-1231112113212331-1303302100212231-1310222330330331-0310220030303030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../data-sources/k8s_cluster_role_binding.md#canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013)
- [Examples](data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-0033002303321210-1002121323201310-3212010333102211-0001332311110131-1312330232323121-0000210220233021-1231021331320321-3121121223312012)
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
