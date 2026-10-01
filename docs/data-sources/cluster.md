---
page_title: "xcsh_cluster landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster landing."
---

# xcsh_cluster landing

<a id="canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300102203223330-0100021102120233-2111131211202121-3013133021001102-3201233130023020-2022012231230030-3223100102321100-0032133013133232"></a>

## xcsh_cluster — xcsh_cluster / 210111301010 / 2

Breadcrumbs:

- xcsh_cluster

Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5
Distributed Cloud.

<a id="canonical-0031133310112312-1310013132300231-2230223331112122-0230231203213111-0002003303130211-1022200211311020-3312001223012233-1311202130103001"></a>

## Prerequisites — xcsh_cluster / 210111301010 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0112302311133003-2100033132123110-3220120122200333-2110020223121311-2330000101213122-2322121103212332-2200331011010312-1012331222131030"></a>

## Minimal configuration — xcsh_cluster / 210111301010 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cluster by name
data "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}

output "cluster_id" {
  value = data.xcsh_cluster.example.id
}
```

<a id="canonical-3100230310112100-3023010020311222-2221001021230032-0103122121032110-3231003221222020-0002323300211000-1131321300203322-1201320320220102"></a>

## Root configuration — xcsh_cluster / 210111301010 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2211220230010321-0213300112321331-3103303233232200-1111121121220130-3033233232211010-0220031210231311-2030112331021310-1123333111322031"></a>

## Next pages — xcsh_cluster / 210111301010 / 6

- [Property reference](../guides/data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [Examples](../guides/data-sources--cluster--examples--group-001.md#canonical-0231101023131313-2212022212310233-2202223132101223-3112213331100203-2232030313130130-1330212111330303-3120212202100033-2101033112202301)
