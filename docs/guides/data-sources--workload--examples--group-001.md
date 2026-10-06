---
page_title: "xcsh_workload examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload examples."
---

# xcsh_workload examples

<a id="canonical-3110211130110313-2113120222301013-1332130101202131-2312310013002012-1223332113022010-0003103230022222-1213221321232131-0213012002030102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- Examples

<a id="canonical-3131331331212012-0102233332012011-2323213322201110-2100331000101213-0103232023122030-2001131110311001-1321323220010002-1332230020210032"></a>

### Complete configurations for `xcsh_workload`

- [Data source](data-sources--workload--examples--group-001.md#canonical-0021101122100012-0321032002123223-3311132110313023-3022101022212002-1133333033021220-3331101233200213-3111312331210312-2120022012122033): valid configuration.

<a id="canonical-0021101122100012-0321032002123223-3311132110313023-3022101022212002-1133333033021220-3331101233200213-3111312331210312-2120022012122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Examples](data-sources--workload--examples--group-001.md#canonical-3110211130110313-2113120222301013-1332130101202131-2312310013002012-1223332113022010-0003103230022222-1213221321232131-0213012002030102)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_workload/data-source.tf`; digest `sha256:cc2cda57d20aa47819844a194fd93abc59f8f0880458fe8ebeae6e323411954a`.

```terraform
# Workload Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Workload by name
data "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}

output "workload_id" {
  value = data.xcsh_workload.example.id
}
```
