---
page_title: "xcsh_network_data_intelligence examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_data_intelligence examples."
---

# xcsh_network_data_intelligence examples

<a id="canonical-0323132310222120-3020213120021003-2220113000222110-0102223210033320-0112313202331031-3230221031033212-2011132110311232-2203020213322023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-3123013012321330-1002111020200012-2102200013011121-0212213012311312-3102011113113022-1211131123200320-0301213033123012-1213000322121013)
- Examples

<a id="canonical-3231302111120323-1202000222030210-0233110332303331-2121033131100211-1113002113132123-0022223032101303-3131002301121302-3122231003111221"></a>

### Complete configurations for `xcsh_network_data_intelligence`

- [Data source](data-sources--network_data_intelligence--examples--group-001.md#canonical-2310133133233022-2132112222002311-0131322010333011-3133102122000021-1223111121232020-3223130120320103-2312031313322200-1322132211030230): valid configuration.

<a id="canonical-2310133133233022-2132112222002311-0131322010333011-3133102122000021-1223111121232020-3223130120320103-2312031313322200-1322132211030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-3123013012321330-1002111020200012-2102200013011121-0212213012311312-3102011113113022-1211131123200320-0301213033123012-1213000322121013)
- [Examples](data-sources--network_data_intelligence--examples--group-001.md#canonical-0323132310222120-3020213120021003-2220113000222110-0102223210033320-0112313202331031-3230221031033212-2011132110311232-2203020213322023)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_data_intelligence/data-source.tf`; digest `sha256:529d85fb608129abcf053afa1a4c7516767c8567633cbc06efe8d845d135d233`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_data_intelligence" "us" {
  regions = ["us"]
}

output "data_intelligence_https_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_data_intelligence.us.cidr_blocks
  }
}
```
