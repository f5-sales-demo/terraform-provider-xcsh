---
page_title: "xcsh_k8s_cluster_role_binding landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role_binding landing."
---

# xcsh_k8s_cluster_role_binding landing

<a id="canonical-1203120300333222-3203223200312120-0002210222112313-1113120331003323-3102000331030303-1133200312133122-1031101033120033-3233030130121012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212321132101120-0113031113001103-0003210030213122-3301130322133031-2331123213331202-2210230132322321-2213311310121013-1302011200230031"></a>

## xcsh_k8s_cluster_role_binding — xcsh_k8s_cluster_role_binding / 213100333020 / 2

Breadcrumbs:

- xcsh_k8s_cluster_role_binding

Manages k8s\_cluster\_role\_binding will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-2011103013033111-3012203110303113-2131032112010330-0332112101211332-3122131111011010-3033330103113100-0330330200333022-1321110311123311"></a>

## Prerequisites — xcsh_k8s_cluster_role_binding / 213100333020 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0122001302213032-3030001212012030-2113033333310100-1223301220312302-0221000220223131-2322231112111101-3210332001102232-0003210011112222"></a>

## Minimal configuration — xcsh_k8s_cluster_role_binding / 213100333020 / 4

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

<a id="canonical-1301031221300131-2213112211323211-0331102313031013-2220320331320213-0212122110120233-2013200302222330-1202130331332111-3123120033331132"></a>

## Root configuration — xcsh_k8s_cluster_role_binding / 213100333020 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1202133003130103-3322323011213012-3331021201120212-0320222123310222-1002000321003111-3320330101112001-0010023303033110-1110001230110223"></a>

## Next pages — xcsh_k8s_cluster_role_binding / 213100333020 / 6

- [Property reference](../guides/resources--k8s_cluster_role_binding--reference--group-001.md#canonical-3122222001022132-1111202303122022-2033123222301020-0030321020302103-0030301220220131-0121213310030202-3231213321200003-3112033132100012)
- [Examples](../guides/resources--k8s_cluster_role_binding--examples--group-001.md#canonical-3202231300133132-3310113211111032-0121303111333113-1033211121101033-3222112220320110-3303010121013223-3121003123201210-2103112030123323)
- [Import](../guides/resources--k8s_cluster_role_binding--lifecycle--group-001.md#canonical-2200110211313011-0101011230210131-3310310000320012-3302230000112323-3223020122121020-3013120000003322-3213003213103332-1010101321132211)
- [Timeouts](../guides/resources--k8s_cluster_role_binding--lifecycle--group-001.md#canonical-3303230312332332-0312203331312031-2033032101003132-2111302033133231-0023032333312333-0112303200100121-0120002113120031-0121110202122101)
