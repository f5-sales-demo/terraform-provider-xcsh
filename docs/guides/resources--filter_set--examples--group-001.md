---
page_title: "xcsh_filter_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set examples."
---

# xcsh_filter_set examples

<a id="canonical-3222211130023131-3030002002322101-1200030222123000-1200233111210312-2220111122231232-1122302111201322-1223301011102233-3121310201130203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- Examples

<a id="canonical-3313330031011322-1010303310023202-1022000123112213-2303302301111332-3323133312001301-2132032303313010-1110032113200332-1333213100322001"></a>

### Complete configurations for `xcsh_filter_set`

- [Resource](resources--filter_set--examples--group-001.md#canonical-2313012200213310-2312010001331203-2320031131200101-3010310321330001-2010300010123111-0202311102003120-3021311011232101-0012112122113133): valid configuration.

<a id="canonical-2313012200213310-2312010001331203-2320031131200101-3010310321330001-2010300010123111-0202311102003120-3021311011232101-0012112122113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001)
- [Examples](resources--filter_set--examples--group-001.md#canonical-3222211130023131-3030002002322101-1200030222123000-1200233111210312-2220111122231232-1122302111201322-1223301011102233-3121310201130203)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_filter_set/resource.tf`; digest `sha256:034749c02446cbe5e781e600fc1e24249a070ab78e205de7268da87fdb24c610`.

```terraform
# FilterSet Resource Example
# Manages specification in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FilterSet configuration
resource "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"

  context_key = "example-value"
}
```
