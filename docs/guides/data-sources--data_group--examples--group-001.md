---
page_title: "xcsh_data_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group examples."
---

# xcsh_data_group examples

<a id="canonical-0300111313221131-3320021111213311-2331132211132120-3220103130330022-1202100002010223-1133300133311230-2333320313310033-1223131223230033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002)
- Examples

<a id="canonical-1021211213210230-2212201002302121-2113021101022321-1322311032330321-0320120101322033-0300203202113232-1112221011212000-1000220312101310"></a>

### Complete configurations for `xcsh_data_group`

- [Data source](data-sources--data_group--examples--group-001.md#canonical-1120210212232021-2132232322312100-2212111023221221-3033222233313212-2020321132302010-0022221122012023-3033003011122303-3030101222330002): valid configuration.

<a id="canonical-1120210212232021-2132232322312100-2212111023221221-3033222233313212-2020321132302010-0022221122012023-3033003011122303-3030101222330002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002)
- [Examples](data-sources--data_group--examples--group-001.md#canonical-0300111313221131-3320021111213311-2331132211132120-3220103130330022-1202100002010223-1133300133311230-2333320313310033-1223131223230033)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_data_group/data-source.tf`; digest `sha256:b9743d4eb532720e8a79b190fba83b8f94d4cbca361dec928076b2b257a940ad`.

```terraform
# DataGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataGroup by name
data "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}

output "data_group_id" {
  value = data.xcsh_data_group.example.id
}
```
