---
page_title: "xcsh_k8s_cluster_role"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role."
---

# xcsh_k8s_cluster_role

<a id="canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_cluster_role

Manages k8s\_cluster\_role will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-0230021113203032-2233113121330012-1213223021303122-2332331210222020-3011001323211113-3002221113233321-2100210210233012-1232121122103233"></a>

### Prerequisites for `xcsh_k8s_cluster_role`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-1230303032220120-3010033320121032-2221202312332230-0223101222110012-0331113223211321-3202120331031333-1313313332113021-2100220212212310"></a>

### Minimal configuration for `xcsh_k8s_cluster_role`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SClusterRole Resource Example
# Manages k8s_cluster_role will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SClusterRole configuration
resource "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}
```

<a id="canonical-3301330320033230-2232331312112110-2031123000030102-0110232312310023-1010012120122100-0022111330100331-2333012323032031-0030333233302001"></a>

### Root configuration for `xcsh_k8s_cluster_role`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2101333210302003-3212201303000332-3332313003013120-1323122103020012-2030203230321101-1010031031323112-1333222302323321-1030111310311321"></a>

### Explore this collection for `xcsh_k8s_cluster_role`

- [Property reference](../guides/resources--k8s_cluster_role--reference--group-001.md#canonical-2133203023220330-1102213322212330-0132102322203020-3312031132113031-0113221210031112-2310320230002102-0032121131200320-1022122031320230)
- [Examples](../guides/resources--k8s_cluster_role--examples--group-001.md#canonical-3032032202030211-2103232223211331-2300202132302320-2120220203311022-1010110123332232-2023233131230201-3120122013131313-2321230131003200)
- [Import](../guides/resources--k8s_cluster_role--lifecycle--group-001.md#canonical-1311100021231112-1103201000203120-1230302223123300-2033030131200232-0101221011113101-2030130202312213-1323300030210202-2230301011323023)
- [Timeouts](../guides/resources--k8s_cluster_role--lifecycle--group-001.md#canonical-0012112112011322-1230003232011311-0211032231302000-2023023312200230-1011211220010202-3001333102222221-3000103220211032-1132311021100001)
