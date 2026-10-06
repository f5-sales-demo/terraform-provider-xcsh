---
page_title: "xcsh_data_group"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group."
---

# xcsh_data_group

<a id="canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_data_group

Reads Data Group information from F5 Distributed Cloud.

<a id="canonical-1012122222133011-2032312012012200-1303321321020220-0331013200101303-1310122302200313-2122322032212232-1001331102331320-0021010002100023"></a>

### Prerequisites for `xcsh_data_group`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0331210033331213-1321211320311101-2031122122211003-0301312321103302-0231130022113332-2003111200112001-3120321200311013-3230120031302021"></a>

### Minimal configuration for `xcsh_data_group`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0130211220133002-3211032332312303-3021223122102321-3311011330313312-1333102001022231-1322032320332232-0000033232300232-1021131021100111"></a>

### Root configuration for `xcsh_data_group`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1033100113311321-3203033001110112-3203321103012213-1032333123113020-3322020220320020-1211311230212330-1202330223310333-2210020110103102"></a>

### Explore this collection for `xcsh_data_group`

- [Property reference](../guides/data-sources--data_group--reference--group-001.md#canonical-2122332000002322-0030233311323023-1021312312333322-3212232000101113-1300221310101123-3101301013233100-3021022110110100-0302132122222001)
- [Examples](../guides/data-sources--data_group--examples--group-001.md#canonical-0300111313221131-3320021111213311-2331132211132120-3220103130330022-1202100002010223-1133300133311230-2333320313310033-1223131223230033)
