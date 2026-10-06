---
page_title: "xcsh_k8s_cluster"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster."
---

# xcsh_k8s_cluster

<a id="canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_cluster

Manages k8s\_cluster will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-0203312103001111-1122030200111330-2311123311221033-1011302220201133-1010022003103001-0302330232000321-0323201102133022-2310002301300330"></a>

### Prerequisites for `xcsh_k8s_cluster`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1001123311330112-0013332023203210-1003021331123013-3323003022013032-1110330332231002-3220330030231230-3211023303003313-1210131022210203"></a>

### Minimal configuration for `xcsh_k8s_cluster`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SCluster Resource Example
# Manages k8s_cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SCluster configuration
resource "xcsh_k8s_cluster" "example" {
  name      = "example-k8s-cluster"
  namespace = "system"
}
```

<a id="canonical-2210013101212111-0202112021012200-3333030222101323-2022132123023031-2033010102021131-1031123013310132-3320333131323331-3321233230100231"></a>

### Root configuration for `xcsh_k8s_cluster`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0130310123332300-2321012103103301-3022323332300320-2133013123320122-0030301100110313-2000131000322033-2102130323323011-1310030023101130"></a>

### Explore this collection for `xcsh_k8s_cluster`

- [Property reference](../guides/resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [Examples](../guides/resources--k8s_cluster--examples--group-001.md#canonical-0312212010022020-2122231330320222-0022320212100103-3221113233320232-2332210101321203-2311133331210100-1231112212211030-2201032231211002)
- [Import](../guides/resources--k8s_cluster--lifecycle--group-001.md#canonical-2302311103303033-2313331022021113-1133010032311330-2313311203133131-3220203023311112-0011320110212120-0112222032010100-2223332022012120)
- [Timeouts](../guides/resources--k8s_cluster--lifecycle--group-001.md#canonical-3321320012201300-3122212223113331-3113123313211113-2003002033310233-1321320303100230-2003212031020103-0011330310001013-2103212320312332)
