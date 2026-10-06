---
page_title: "xcsh_k8s_cluster_role"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role."
---

# xcsh_k8s_cluster_role

<a id="canonical-2232121312000210-3132130123201003-2131123201312211-2021232000312101-1032012313201322-3103300020131231-0202213330300122-0302302323323122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_cluster_role

Reads Kubernetes cluster role information from F5 Distributed Cloud.

<a id="canonical-3332030130310213-0022032201323322-3023032202113022-1021333201300213-1322211031111003-3232333221130230-1222101320331132-0110123332032031"></a>

### Prerequisites for `xcsh_k8s_cluster_role`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-3230020332323312-1020132230103132-2130110330320333-2230132103122012-3121122000311122-2032011231010003-0231312021320311-3021003221103233"></a>

### Minimal configuration for `xcsh_k8s_cluster_role`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SClusterRole Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SClusterRole by name
data "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}

output "k8s_cluster_role_id" {
  value = data.xcsh_k8s_cluster_role.example.id
}
```

<a id="canonical-2333011023210213-3103221231301011-3002213001103031-0012001223023112-0001220321101330-0233002121210301-1211322213312331-3020220211101230"></a>

### Root configuration for `xcsh_k8s_cluster_role`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1213121030021323-3211323230032002-0013022231133033-3213021210033101-0001130213113122-3002230320302311-1101333001322122-0320113230221032"></a>

### Explore this collection for `xcsh_k8s_cluster_role`

- [Property reference](../guides/data-sources--k8s_cluster_role--reference--group-001.md#canonical-2001310313031232-3003303113031320-0000111202132223-2030000001211031-3123033320222313-3312102121311301-3321002322000112-3233122202001100)
- [Examples](../guides/data-sources--k8s_cluster_role--examples--group-001.md#canonical-2020200223133133-1213011132112100-3201302032210003-2331332220121210-0033322213031201-3101003002030200-1103001321300103-2312010013201210)
