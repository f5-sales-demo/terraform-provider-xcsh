---
page_title: "xcsh_k8s_cluster"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster."
---

# xcsh_k8s_cluster

<a id="canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_cluster

Reads Kubernetes cluster information from F5 Distributed Cloud.

<a id="canonical-2010103022320020-1101102201233112-3301213203231211-3113211023033320-2121332312222033-3132103102003230-1201210002321300-2303303103212012"></a>

### Prerequisites for `xcsh_k8s_cluster`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0011001020023211-3000200132323302-1320301211213300-2332331001322123-3300111031021322-1101120200213232-0300102313321203-2211122131133100"></a>

### Minimal configuration for `xcsh_k8s_cluster`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0122021301313032-2011223030132202-2022331332132211-1233310001222213-1002301012212202-2211302031212130-0133331203131101-3011310200100033"></a>

### Root configuration for `xcsh_k8s_cluster`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2212232012311011-0032213133101032-1002211222022213-2200133002203322-3022102312302210-1320210213300332-2321001101000111-1310133333020023"></a>

### Explore this collection for `xcsh_k8s_cluster`

- [Property reference](../guides/data-sources--k8s_cluster--reference--group-001.md#canonical-3033230002010003-2101221133021022-3231320222110033-2211321111123201-3112303022332033-3030211103323301-0213323130131201-1011002120120200)
- [Examples](../guides/data-sources--k8s_cluster--examples--group-001.md#canonical-1300311003300211-0021113101021312-1011223033312022-3031202223113030-2221120113003302-3130200302121023-3030033231132300-2310320031320232)
