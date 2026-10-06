---
page_title: "xcsh_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster examples."
---

# xcsh_cluster examples

<a id="canonical-0231101023131313-2212022212310233-2202223132101223-3112213331100203-2232030313130130-1330212111330303-3120212202100033-2101033112202301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- Examples

<a id="canonical-1300303201030133-2220213003330110-3220132331321223-1131213311311323-1001010132101210-2130001321001122-0302202030321012-2332030131201020"></a>

### Complete configurations for `xcsh_cluster`

- [Data source](data-sources--cluster--examples--group-001.md#canonical-3113213201302132-3020232133203000-1212202212233003-0033333300300310-1110111323132220-1001202322033223-0120321221121102-2221021311131021): valid configuration.

<a id="canonical-3113213201302132-3020232133203000-1212202212233003-0033333300300310-1110111323132220-1001202322033223-0120321221121102-2221021311131021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Examples](data-sources--cluster--examples--group-001.md#canonical-0231101023131313-2212022212310233-2202223132101223-3112213331100203-2232030313130130-1330212111330303-3120212202100033-2101033112202301)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cluster/data-source.tf`; digest `sha256:657c300c6f1b904145d0c5ed8c458dbdc20cbe652caf3ef6b0a40bc6ebb5e4ee`.

```terraform
# Cluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cluster by name
data "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}

output "cluster_id" {
  value = data.xcsh_cluster.example.id
}
```
