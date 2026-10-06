---
page_title: "xcsh_bot_allowlist_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_allowlist_policy examples."
---

# xcsh_bot_allowlist_policy examples

<a id="canonical-2033001321211220-3332032010012310-1231001310030222-3330121110002310-2012103030201232-3003112023123132-1202202012130003-2231210030102303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md#canonical-0022111012002123-0020202233333300-0020222322013201-0023201320232112-2223012222333122-3222112110213212-1033221302121302-2123032122301213)
- Examples

<a id="canonical-3303002113101130-1231003013330133-3022003311010232-3212003102123313-0112000221210300-2331123030020223-2212300323033100-1232021201312222"></a>

### Complete configurations for `xcsh_bot_allowlist_policy`

- [Data source](data-sources--bot_allowlist_policy--examples--group-001.md#canonical-0220001130331330-3110113221021010-3131112022300313-2232121223321222-3121221110311100-1223101022213101-2021023130333132-1102320312332230): valid configuration.

<a id="canonical-0220001130331330-3110113221021010-3131112022300313-2232121223321222-3121221110311100-1223101022213101-2021023130333132-1102320312332230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bot_allowlist_policy](../data-sources/bot_allowlist_policy.md#canonical-0022111012002123-0020202233333300-0020222322013201-0023201320232112-2223012222333122-3222112110213212-1033221302121302-2123032122301213)
- [Examples](data-sources--bot_allowlist_policy--examples--group-001.md#canonical-2033001321211220-3332032010012310-1231001310030222-3330121110002310-2012103030201232-3003112023123132-1202202012130003-2231210030102303)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_allowlist_policy/data-source.tf`; digest `sha256:e79cb351210b30d9b4d2bf6ff1c7b6c44610a7fd5a55e86ba638364a3bd45c0b`.

```terraform
# BotAllowlistPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotAllowlistPolicy by name
data "xcsh_bot_allowlist_policy" "example" {
  name      = "example-bot-allowlist-policy"
  namespace = "staging"
}

output "bot_allowlist_policy_id" {
  value = data.xcsh_bot_allowlist_policy.example.id
}
```
