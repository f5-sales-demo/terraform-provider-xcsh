---
page_title: "xcsh_dc_cluster_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group landing."
---

# xcsh_dc_cluster_group landing

<a id="canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d87ecc40f26184b0192bd3595f16459c71040ec44f796443b85f940ee87a652"></a>

## xcsh_dc_cluster_group — xcsh_dc_cluster_group / 1ede2471ac32 / 2

Breadcrumbs:

- xcsh_dc_cluster_group

Manages DC Cluster group in given namespace in F5 Distributed Cloud.

<a id="canonical-2dbafd3d9c6eeaf96dfbffd3f68c928c17d77a7396716b76bc934af525545427"></a>

## Prerequisites — xcsh_dc_cluster_group / 1ede2471ac32 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4b8f13239b552a5203ca8f90488434cd4de4f1c13406dab9442e3bf6f286064b"></a>

## Minimal configuration — xcsh_dc_cluster_group / 1ede2471ac32 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DcClusterGroup Resource Example
# Manages DC Cluster group in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DcClusterGroup configuration
resource "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}
```

<a id="canonical-f36d7ce0b61acecd0a7f64dba31225ff29a89aeb0df0717563a265c5a2bfca99"></a>

## Root configuration — xcsh_dc_cluster_group / 1ede2471ac32 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-758dd2752ab7433657e9d76b9a3798b1f5f84c736ce9941cb906d27e05a892f8"></a>

## Next pages — xcsh_dc_cluster_group / 1ede2471ac32 / 6

- [Property reference](../guides/resources--dc_cluster_group--reference--group-001.md#canonical-c8ed235d98953d15c6c0a0fbcbc6d019c220f7a721c3d418daf1d4995e5095f9)
- [Examples](../guides/resources--dc_cluster_group--examples--group-001.md#canonical-2f4264098e0152fbf2779a6bbae0471ac64aac33e9f2333b107598f53731b774)
- [Import](../guides/resources--dc_cluster_group--lifecycle--group-001.md#canonical-9df0ec08df74e8daf8425f927cfbf91840df16fd91d7ac83009c6cf841d69624)
- [Timeouts](../guides/resources--dc_cluster_group--lifecycle--group-001.md#canonical-cfe18cbfca70988146634c38e4a55623ab5fcd5177a22a483fabae99a239c8dd)
