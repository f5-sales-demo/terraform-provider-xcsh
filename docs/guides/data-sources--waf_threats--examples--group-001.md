---
page_title: "xcsh_waf_threats examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threats examples."
---

# xcsh_waf_threats examples

<a id="canonical-3230103102203013-1010033330333121-2003113212233311-1300003030111202-3133333103013112-3030011121032033-2232210213030222-1301003232032012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-2231102302300310-3000321032301121-1323113333312231-2132011213333010-3012303030303021-1322212103212101-0130001033031000-2100323021012200)
- Examples

<a id="canonical-1213030021022102-2130200023230030-1113211311120233-1011230331001011-3122223321012110-1231301312020201-1123300131102312-0203211113021302"></a>

### Complete configurations for `xcsh_waf_threats`

- [Data source](data-sources--waf_threats--examples--group-001.md#canonical-2202312212030212-0201011323231132-1103233031011022-3010302202012333-1211332010302213-0131321311030321-3000213121010220-0210013231310221): valid configuration.

<a id="canonical-2202312212030212-0201011323231132-1103233031011022-3010302202012333-1211332010302213-0131321311030321-3000213121010220-0210013231310221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-2231102302300310-3000321032301121-1323113333312231-2132011213333010-3012303030303021-1322212103212101-0130001033031000-2100323021012200)
- [Examples](data-sources--waf_threats--examples--group-001.md#canonical-3230103102203013-1010033330333121-2003113212233311-1300003030111202-3133333103013112-3030011121032033-2232210213030222-1301003232032012)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_threats/data-source.tf`; digest `sha256:9ac0cdd709cbc9a9ca710dea9bf40dd5b343227de03ab5599937418e7cc2c055`.

```terraform
# WAFThreats DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threats" "example" {
}

output "waf_threats_result" {
  value = data.xcsh_waf_threats.example
}
```
