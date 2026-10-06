---
page_title: "xcsh_irule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule examples."
---

# xcsh_irule examples

<a id="canonical-0231321322321101-3311033202102201-0101010112300100-0023333101212212-0010201023102311-0133032231011313-2323121302103201-0202212122221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_irule](../data-sources/irule.md#canonical-0231101210120303-0011000002331110-3333032201223210-2323130033302300-2222120120302330-2301200112302321-0023231230022212-3231111021211202)
- Examples

<a id="canonical-3130103323012211-1123300330311303-0230132110023021-3221011211200323-2120310333122103-0320101332212323-3001333001312003-3020021032301332"></a>

### Complete configurations for `xcsh_irule`

- [Data source](data-sources--irule--examples--group-001.md#canonical-2020310331322123-1231313213032103-0220130101121132-0211022211302012-0201022123222333-2021200303030110-1320322300330203-2303102002211031): valid configuration.

<a id="canonical-2020310331322123-1231313213032103-0220130101121132-0211022211302012-0201022123222333-2021200303030110-1320322300330203-2303102002211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_irule](../data-sources/irule.md#canonical-0231101210120303-0011000002331110-3333032201223210-2323130033302300-2222120120302330-2301200112302321-0023231230022212-3231111021211202)
- [Examples](data-sources--irule--examples--group-001.md#canonical-0231321322321101-3311033202102201-0101010112300100-0023333101212212-0010201023102311-0133032231011313-2323121302103201-0202212122221131)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_irule/data-source.tf`; digest `sha256:567750397ff4df80dac106ded5aa33f7fcb698c11cea8ada22ebaa5fc6727f22`.

```terraform
# Irule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Irule by name
data "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"
}

output "irule_id" {
  value = data.xcsh_irule.example.id
}
```
