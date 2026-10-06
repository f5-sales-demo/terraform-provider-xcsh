---
page_title: "xcsh_api_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery examples."
---

# xcsh_api_discovery examples

<a id="canonical-3301213201311031-3110013301031111-0113212001331001-1333230011010331-3000023322111303-2320221233022331-1103132033021121-3232301013323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- Examples

<a id="canonical-2231223001330232-3222323310223032-2201320312302213-3111113330022200-1201032212130213-0312300031103310-0022032121201121-0230232032021310"></a>

### Complete configurations for `xcsh_api_discovery`

- [Data source](data-sources--api_discovery--examples--group-001.md#canonical-1020010203113302-2202232221321101-2020332001232022-3210013310130031-1120320311223300-3012133011202200-0132102332301033-2320210212012310): valid configuration.

<a id="canonical-1020010203113302-2202232221321101-2020332001232022-3210013310130031-1120320311223300-3012133011202200-0132102332301033-2320210212012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_api_discovery](../data-sources/api_discovery.md#canonical-3123302132003013-1112002130101021-3331313310310030-3002230123333210-1113203220331333-0131231011213031-0211202023001022-1322201233022220)
- [Examples](data-sources--api_discovery--examples--group-001.md#canonical-3301213201311031-3110013301031111-0113212001331001-1333230011010331-3000023322111303-2320221233022331-1103132033021121-3232301013323113)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_discovery/data-source.tf`; digest `sha256:f37cea7bc8746642eef8cfd7a0d2f33975184479df2f41465a9bdda95de57bb2`.

```terraform
# APIDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDiscovery by name
data "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}

output "api_discovery_id" {
  value = data.xcsh_api_discovery.example.id
}
```
