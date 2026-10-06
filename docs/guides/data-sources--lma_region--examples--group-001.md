---
page_title: "xcsh_lma_region examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_lma_region examples."
---

# xcsh_lma_region examples

<a id="canonical-3201110113300321-1111333210230110-1310113001212123-0211303112100221-0303113020300311-3210230100202001-3211313112222002-3001132233133330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_lma_region](../data-sources/lma_region.md#canonical-2212011320333232-1123311200100001-1321311330220000-3033022231230311-3012011130132021-2120003020013220-1120200123320031-0302331121003101)
- Examples

<a id="canonical-0021220020031101-1022323102311332-3300121021330121-3121013233123013-2133222033223031-3121112233130121-3201221012220131-1132011231322302"></a>

### Complete configurations for `xcsh_lma_region`

- [Data source](data-sources--lma_region--examples--group-001.md#canonical-1231011002111003-3313012131132321-3230212301000031-2201000110013133-2223103203300302-2002132103232003-1020001030201221-0113310210300231): valid configuration.

<a id="canonical-1231011002111003-3313012131132321-3230212301000031-2201000110013133-2223103203300302-2002132103232003-1020001030201221-0113310210300231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_lma_region](../data-sources/lma_region.md#canonical-2212011320333232-1123311200100001-1321311330220000-3033022231230311-3012011130132021-2120003020013220-1120200123320031-0302331121003101)
- [Examples](data-sources--lma_region--examples--group-001.md#canonical-3201110113300321-1111333210230110-1310113001212123-0211303112100221-0303113020300311-3210230100202001-3211313112222002-3001132233133330)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_lma_region/data-source.tf`; digest `sha256:6e44cbfbf0065cf8526c29152a263e7eb100be46cd3be2092eca1d815333bf89`.

```terraform
# LmaRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LmaRegion by name
data "xcsh_lma_region" "example" {
  name      = "example-lma-region"
  namespace = "staging"
}

output "lma_region_id" {
  value = data.xcsh_lma_region.example.id
}
```
