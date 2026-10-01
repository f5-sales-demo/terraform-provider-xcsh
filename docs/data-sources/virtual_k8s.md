---
page_title: "xcsh_virtual_k8s landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s landing."
---

# xcsh_virtual_k8s landing

<a id="canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212002003321131-1201333012210312-1031322211313102-1113103011210302-3002030201300012-0032302103031323-1210022103232200-1030023223320133"></a>

## xcsh_virtual_k8s — xcsh_virtual_k8s / 232012112300 / 2

Breadcrumbs:

- xcsh_virtual_k8s

Manages virtual\_k8s will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-0022200100311230-1133201301323121-3200201301121233-0002000331000031-1121331111022333-3320301020123013-3100220131030033-0112013203111111"></a>

## Prerequisites — xcsh_virtual_k8s / 232012112300 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `workload`.

- workload: Container workloads in this namespace

<a id="canonical-3301011101320211-2132231000212232-2003300231100130-0321322012132021-2021031300203023-1100111112010220-0313202320101231-2123100232033001"></a>

## Minimal configuration — xcsh_virtual_k8s / 232012112300 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualK8S Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualK8S by name
data "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}

output "virtual_k8s_id" {
  value = data.xcsh_virtual_k8s.example.id
}
```

<a id="canonical-3032102321221231-1030032220130313-0022010133031122-0333111133200101-2322123032101022-0020313022022033-0132303331123131-0222020320212313"></a>

## Root configuration — xcsh_virtual_k8s / 232012112300 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2132321021021030-2222332211310122-0322112331032313-0003123212103213-0320310201311022-2011101220023320-0021033101031031-2222210110321101"></a>

## Next pages — xcsh_virtual_k8s / 232012112300 / 6

- [Property reference](../guides/data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- [Examples](../guides/data-sources--virtual_k8s--examples--group-001.md#canonical-3003212211022232-3322110222230213-0231321321013322-1203302321122120-1333320121200133-0312320223001013-0332223203022313-1331301003223120)
