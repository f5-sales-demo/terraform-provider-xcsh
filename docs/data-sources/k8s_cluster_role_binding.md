---
page_title: "xcsh_k8s_cluster_role_binding landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding landing."
---

# xcsh_k8s_cluster_role_binding landing

<a id="canonical-2021222210012102-1330130200013102-0220323113122312-2121230113313023-1123103021232230-1310303332323332-2100321301113012-3031120031031013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321223320123322-3031110330223100-0132212001100032-2110200001131230-2101321230220220-0120122302110323-3233111013001010-2203031203113032"></a>

## xcsh_k8s_cluster_role_binding — xcsh_k8s_cluster_role_binding / 000220131021 / 2

Breadcrumbs:

- xcsh_k8s_cluster_role_binding

Manages k8s\_cluster\_role\_binding will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-0020111213102310-1233202131001110-1101222033132100-0000002201020332-0310220310210303-3313022213002302-3202311322130230-3131120200101033"></a>

## Prerequisites — xcsh_k8s_cluster_role_binding / 000220131021 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0130310121130003-3232212203330221-0301202003122231-2010320130213223-0300103133212220-3232103013000031-1132120132102212-0202222321103232"></a>

## Minimal configuration — xcsh_k8s_cluster_role_binding / 000220131021 / 4

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

<a id="canonical-2120220321221200-2112233200302122-1030210211230220-3300100222303202-1133233333320302-1120300021022001-0121031011010010-0020200302301000"></a>

## Root configuration — xcsh_k8s_cluster_role_binding / 000220131021 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3201023333333133-2330021222110332-3203310133033033-1321123302120331-2021020230300222-0321332301333331-0021002111001213-0333132211113122"></a>

## Next pages — xcsh_k8s_cluster_role_binding / 000220131021 / 6

- [Property reference](../guides/data-sources--k8s_cluster_role_binding--reference--group-001.md#canonical-2330203132111311-2030223031030132-1113210123013110-2331032322332103-0220202113111221-2130320110300133-2230301113102200-0202032220222113)
- [Examples](../guides/data-sources--k8s_cluster_role_binding--examples--group-001.md#canonical-0033002303321210-1002121323201310-3212010333102211-0001332311110131-1312330232323121-0000210220233021-1231021331320321-3121121223312012)
