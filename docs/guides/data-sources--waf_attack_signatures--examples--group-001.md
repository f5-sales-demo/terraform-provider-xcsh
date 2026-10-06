---
page_title: "xcsh_waf_attack_signatures examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_attack_signatures examples."
---

# xcsh_waf_attack_signatures examples

<a id="canonical-1000123133012313-3302212201201323-1032223012210030-3000332321001301-3222211122333313-2033112213131303-0300331331022233-3010003031111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md#canonical-1002120312121031-2231120213022032-3113323210101313-3123231202123002-0000220212231033-3012130333333223-0330122132233323-2033323020101321)
- Examples

<a id="canonical-3330333222130030-1320313011102110-3213103001123213-0030330312300032-2331231023113203-3222222000300202-3233110100103321-1010133233331121"></a>

### Complete configurations for `xcsh_waf_attack_signatures`

- [Data source](data-sources--waf_attack_signatures--examples--group-001.md#canonical-2311322111003122-1330012212233023-2223001232321231-3021213221010010-3032301301220122-0110233000232021-1322330311323001-3111302312322213): valid configuration.

<a id="canonical-2311322111003122-1330012212233023-2223001232321231-3021213221010010-3032301301220122-0110233000232021-1322330311323001-3111302312322213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md#canonical-1002120312121031-2231120213022032-3113323210101313-3123231202123002-0000220212231033-3012130333333223-0330122132233323-2033323020101321)
- [Examples](data-sources--waf_attack_signatures--examples--group-001.md#canonical-1000123133012313-3302212201201323-1032223012210030-3000332321001301-3222211122333313-2033112213131303-0300331331022233-3010003031111302)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_attack_signatures/data-source.tf`; digest `sha256:5857e356bd3dea3a26e0c6c17fd28eb9757681e0e4bfa5915066871a8e38e467`.

```terraform
# WAFAttackSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_attack_signatures" "example" {
}

output "waf_attack_signatures_result" {
  value = data.xcsh_waf_attack_signatures.example
}
```
