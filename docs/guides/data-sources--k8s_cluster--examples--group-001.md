---
page_title: "xcsh_k8s_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster examples."
---

# xcsh_k8s_cluster examples

<a id="canonical-1300311003300211-0021113101021312-1011223033312022-3031202223113030-2221120113003302-3130200302121023-3030033231132300-2310320031320232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- Examples

<a id="canonical-3130332032022011-0302332002110201-1301333200010310-3302310320123010-0102310311303330-3012311232301023-3302112221232012-3102330033011203"></a>

### Complete configurations for `xcsh_k8s_cluster`

- [Data source](data-sources--k8s_cluster--examples--group-001.md#canonical-3101112002210110-2023302322113310-3001112202310120-2122213213233222-2210332312102132-1030300002030131-3130113301123120-0322033210221000): valid configuration.

<a id="canonical-3101112002210110-2023302322113310-3001112202310120-2122213213233222-2210332312102132-1030300002030131-3130113301123120-0322033210221000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md#canonical-2000311330312211-0123113023320022-2103000112130233-2031113012021222-1321333000120232-2032001010223110-2212310010223100-2220321212212021)
- [Examples](data-sources--k8s_cluster--examples--group-001.md#canonical-1300311003300211-0021113101021312-1011223033312022-3031202223113030-2221120113003302-3130200302121023-3030033231132300-2310320031320232)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster/data-source.tf`; digest `sha256:1ffe54fa3443473ef381429123673c75783a908cc4dae691369d96028c415887`.

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
