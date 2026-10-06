---
page_title: "xcsh_k8s_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster examples."
---

# xcsh_k8s_cluster examples

<a id="canonical-0312212010022020-2122231330320222-0022320212100103-3221113233320232-2332210101321203-2311133331210100-1231112212211030-2201032231211002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- Examples

<a id="canonical-1022310100313321-2031220312002031-0311123113212330-1002303301200120-1021000333132010-3120100221333032-2311320123122022-3001230303001312"></a>

### Complete configurations for `xcsh_k8s_cluster`

- [Resource](resources--k8s_cluster--examples--group-001.md#canonical-3033300213203230-1302200310303011-3231331232003333-1030113303223322-3003132100310112-3131223123321010-3332200220223322-3210312013023231): valid configuration.

<a id="canonical-3033300213203230-1302200310303011-3231331232003333-1030113303223322-3003132100310112-3131223123321010-3332200220223322-3210312013023231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Examples](resources--k8s_cluster--examples--group-001.md#canonical-0312212010022020-2122231330320222-0022320212100103-3221113233320232-2332210101321203-2311133331210100-1231112212211030-2201032231211002)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster/resource.tf`; digest `sha256:4b234e2e10f8643667634d9286b365cede58a97ccb384ebff96249fbe01b28a1`.

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
