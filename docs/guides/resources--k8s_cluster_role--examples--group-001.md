---
page_title: "xcsh_k8s_cluster_role examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role examples."
---

# xcsh_k8s_cluster_role examples

<a id="canonical-3032032202030211-2103232223211331-2300202132302320-2120220203311022-1010110123332232-2023233131230201-3120122013131313-2321230131003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- Examples

<a id="canonical-3133201330323022-2223122320303302-1323213012323230-2330203002310223-0113333200112023-1113222102330110-1103333102030001-1002013233330332"></a>

### Complete configurations for `xcsh_k8s_cluster_role`

- [Resource](resources--k8s_cluster_role--examples--group-001.md#canonical-0210001011031233-2213030330023033-0230032213111100-0020111321332013-1121320113302011-1231132013303212-1110223132313312-1310110320133101): valid configuration.

<a id="canonical-0210001011031233-2213030330023033-0230032213111100-0020111321332013-1121320113302011-1231132013303212-1110223132313312-1310110320133101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md#canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223)
- [Examples](resources--k8s_cluster_role--examples--group-001.md#canonical-3032032202030211-2103232223211331-2300202132302320-2120220203311022-1010110123332232-2023233131230201-3120122013131313-2321230131003200)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster_role/resource.tf`; digest `sha256:f4470002db951c71b74b86c89c38ea07204065a7ca36bf4ba1c80737618e16ac`.

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
