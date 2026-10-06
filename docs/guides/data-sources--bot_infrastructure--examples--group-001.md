---
page_title: "xcsh_bot_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure examples."
---

# xcsh_bot_infrastructure examples

<a id="canonical-3320013121211030-2031323200002323-2132020301312100-1033303231202032-1322332132203020-3120200230130301-3023000303232123-0301210121302203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)
- Examples

<a id="canonical-1200203210012013-3111201020122201-3320223123122033-2330323012013103-1201011132021003-0322120101311133-3330302310222313-1201003011221123"></a>

### Complete configurations for `xcsh_bot_infrastructure`

- [Data source](data-sources--bot_infrastructure--examples--group-001.md#canonical-2122303003330312-2112310320230020-0311223203032121-2200300011032333-2223202002211123-1212233100203201-2203011311112220-2303003133112113): valid configuration.

<a id="canonical-2122303003330312-2112310320230020-0311223203032121-2200300011032333-2223202002211123-1212233100203201-2203011311112220-2303003133112113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210)
- [Examples](data-sources--bot_infrastructure--examples--group-001.md#canonical-3320013121211030-2031323200002323-2132020301312100-1033303231202032-1322332132203020-3120200230130301-3023000303232123-0301210121302203)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_infrastructure/data-source.tf`; digest `sha256:f72b753447890fa3999ee0487a20af9414eac1431bcf20d1ef50048bd9025f06`.

```terraform
# BotInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotInfrastructure by name
data "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}

output "bot_infrastructure_id" {
  value = data.xcsh_bot_infrastructure.example.id
}
```
