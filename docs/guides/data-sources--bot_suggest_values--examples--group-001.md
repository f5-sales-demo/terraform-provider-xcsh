---
page_title: "xcsh_bot_suggest_values examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_suggest_values examples."
---

# xcsh_bot_suggest_values examples

<a id="canonical-3023011213133012-0111020023130020-1102232223001103-1310002010213032-0302100103210013-3130303103311003-0123122011212000-3022210211021131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-3223113331303120-3211131233310020-2103101120000000-3302331313210231-3330320100131032-1303233132310303-1120022100002302-3203202022101303)
- Examples

<a id="canonical-1022102322020002-3220233233223110-0232013200220112-0231233122230332-0331112103003213-0122111322232330-0102132220313123-3231300021021110"></a>

### Complete configurations for `xcsh_bot_suggest_values`

- [Data source](data-sources--bot_suggest_values--examples--group-001.md#canonical-0313312323011010-1031111310122203-3012110312110231-3120311032102232-2001012313110231-1121323110102121-3010110120222122-2232101100310000): valid configuration.

<a id="canonical-0313312323011010-1031111310122203-3012110312110231-3120311032102232-2001012313110231-1121323110102121-3010110120222122-2232101100310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-3223113331303120-3211131233310020-2103101120000000-3302331313210231-3330320100131032-1303233132310303-1120022100002302-3203202022101303)
- [Examples](data-sources--bot_suggest_values--examples--group-001.md#canonical-3023011213133012-0111020023130020-1102232223001103-1310002010213032-0302100103210013-3130303103311003-0123122011212000-3022210211021131)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_suggest_values/data-source.tf`; digest `sha256:c6c1b3639717b7c716bae9d4f7857fddc676935e02dc229f788549aecb0a97ae`.

```terraform
# BotSuggestValues DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_suggest_values" "example" {
  namespace = "example-value"
}

output "bot_suggest_values_result" {
  value = data.xcsh_bot_suggest_values.example
}
```
