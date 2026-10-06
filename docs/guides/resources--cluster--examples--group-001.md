---
page_title: "xcsh_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster examples."
---

# xcsh_cluster examples

<a id="canonical-0222121032033202-3301310203120212-3210221131013113-2100232232000123-2211032310121200-0300120221121213-0032213223032312-0220012231030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- Examples

<a id="canonical-0303022000031311-2102112100222222-3313321013322323-1222121231331233-0223303330323122-3301330031321003-0333201210332301-0203320211100113"></a>

### Complete configurations for `xcsh_cluster`

- [Resource](resources--cluster--examples--group-001.md#canonical-0331332022232123-1322210312322323-3011020222012033-3110310002112202-1232010201201122-0302122102013332-0301020330012110-0312002302322022): valid configuration.

<a id="canonical-0331332022232123-1322210312322323-3011020222012033-3110310002112202-1232010201201122-0302122102013332-0301020330012110-0312002302322022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Examples](resources--cluster--examples--group-001.md#canonical-0222121032033202-3301310203120212-3210221131013113-2100232232000123-2211032310121200-0300120221121213-0032213223032312-0220012231030221)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cluster/resource.tf`; digest `sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e`.

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
