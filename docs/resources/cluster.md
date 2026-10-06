---
page_title: "xcsh_cluster"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster."
---

# xcsh_cluster

<a id="canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_cluster

Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5
Distributed Cloud.

<a id="canonical-2013112211333221-2233021221230313-0320021003230203-3303331110100033-2232003023013233-1202113231301223-0312301110323121-0021213313332302"></a>

### Prerequisites for `xcsh_cluster`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0001112012313102-0230322301030032-0323202310203223-2100000213300300-2202032311000030-0232120112110020-3122202020311323-1311211023002322"></a>

### Minimal configuration for `xcsh_cluster`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cluster configuration
resource "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}
```

<a id="canonical-1000002000210113-2003201322121231-3133133302030110-1123321201322103-2112233011130233-3303133132320232-3030323331012322-0000211010200212"></a>

### Root configuration for `xcsh_cluster`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0101201110300213-3130030312031103-1132103120010301-1323220302313203-3301231211012031-1120121133303303-0303011323130313-2023031103122300"></a>

### Explore this collection for `xcsh_cluster`

- [Property reference](../guides/resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [Examples](../guides/resources--cluster--examples--group-001.md#canonical-0222121032033202-3301310203120212-3210221131013113-2100232232000123-2211032310121200-0300120221121213-0032213223032312-0220012231030221)
- [Import](../guides/resources--cluster--lifecycle--group-001.md#canonical-3011110330023100-0223201213132312-2223331123111022-2211103133210212-2200320323320003-1013331133211300-1201233201011103-3013233120010320)
- [Timeouts](../guides/resources--cluster--lifecycle--group-001.md#canonical-0021032101101013-2210121022021022-3201022301002222-2301132131010212-3233110020013030-0332301123112101-2021310313310022-3333302023300312)
