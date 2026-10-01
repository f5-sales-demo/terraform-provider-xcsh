---
page_title: "xcsh_container_registry landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry landing."
---

# xcsh_container_registry landing

<a id="canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112032133213213-3311013132211033-2110330032300033-1212331300230211-3330322322000000-0212301201032200-3001321322232133-3200201012213300"></a>

## xcsh_container_registry — xcsh_container_registry / 211223003302 / 2

Breadcrumbs:

- xcsh_container_registry

Manages a Container Registry resource in F5 Distributed Cloud for container image registry
configuration.

<a id="canonical-1233300022003120-2201000320011023-3102100032121201-0021100320231010-0311222023231212-1320003111123133-1123333333200322-3313020320301303"></a>

## Prerequisites — xcsh_container_registry / 211223003302 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-3213302110200112-3010002033302131-2231101221131103-3301301111311331-3121000022312110-0211332132101310-1013210130312021-0021113010133311"></a>

## Minimal configuration — xcsh_container_registry / 211223003302 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ContainerRegistry Resource Example
# Manages a Container Registry resource in F5 Distributed Cloud for container image registry configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ContainerRegistry configuration
resource "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"

  registry  = "example-value"
  user_name = "example-value"
}
```

<a id="canonical-1021322300331110-2300133133002122-1323010103313203-1323310020230333-2012001310302101-2133133331103203-3133102122221120-1300022122133011"></a>

## Root configuration — xcsh_container_registry / 211223003302 / 5

Required root properties: `name`, `namespace`, `registry`, `user_name`. Full root flags and choices appear in the property reference.

<a id="canonical-3111232030000032-0310320202232230-3322111102000211-1212000000231123-0203101311001311-0212133121000033-2133112202012103-0211101132213223"></a>

## Next pages — xcsh_container_registry / 211223003302 / 6

- [Property reference](../guides/resources--container_registry--reference--group-001.md#canonical-1032120112320001-2222011121203202-3330011001232203-0333133013110231-2202321322121232-1121012101311030-0100022220223123-1211111333222203)
- [Examples](../guides/resources--container_registry--examples--group-001.md#canonical-0132120323120123-0223213303233201-1021012301310010-0103313121312020-3001233203211120-1133301233201321-3003031002201103-2201223212203310)
- [Import](../guides/resources--container_registry--lifecycle--group-001.md#canonical-1232313310001120-0013101312033331-2100002211100101-3203130201321101-1203113000232330-0132211032022212-1300223300001310-0320122201310122)
- [Timeouts](../guides/resources--container_registry--lifecycle--group-001.md#canonical-2302210310221301-0332223010211001-2221223302001301-0201312231223111-1203230220121303-2323231031333331-0323030203132132-2302113101320230)
