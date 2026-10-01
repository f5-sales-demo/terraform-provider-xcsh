---
page_title: "xcsh_bot_allowlist_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_allowlist_policy landing."
---

# xcsh_bot_allowlist_policy landing

<a id="canonical-0022111012002123-0020202233333300-0020222322013201-0023201320232112-2223012222333122-3222112110213212-1033221302121302-2123032122301213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221322203301101-0131032131100213-3302320020030100-3021132231303320-2123210322010203-2320311002213132-0110222111010331-1113012000020310"></a>

## xcsh_bot_allowlist_policy — xcsh_bot_allowlist_policy / 030331122122 / 2

Breadcrumbs:

- xcsh_bot_allowlist_policy

Manages a Bot Allowlist Policy resource in F5 Distributed Cloud for get bot allowlist policy.
configuration. (read-only data source)

<a id="canonical-2123010210330131-2213010001100020-3310333102112232-0122110232020010-3100301023033011-2101130222301022-0312012022022232-1121222332021300"></a>

## Prerequisites — xcsh_bot_allowlist_policy / 030331122122 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2321101010212030-0203222130233202-3111010212201002-3302201301132103-0320013230202111-2110210310000003-1321231230122303-2233021032003323"></a>

## Minimal configuration — xcsh_bot_allowlist_policy / 030331122122 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1321233212210010-1110000003101212-3200210012013100-2222220312131322-0211221321331231-3233233212130222-1031322002032331-3323133212200100"></a>

## Root configuration — xcsh_bot_allowlist_policy / 030331122122 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2303013013313132-2313111022031322-1121202220313033-3300133223123100-0130303212130203-2131000120320200-1002322321303311-3010213031222211"></a>

## Next pages — xcsh_bot_allowlist_policy / 030331122122 / 6

- [Property reference](../guides/data-sources--bot_allowlist_policy--reference--group-001.md#canonical-1011231013012133-2121222203110231-0131022322212213-0313203222020020-1033321330133221-3030213123203133-3221132122003221-3323000023122122)
- [Examples](../guides/data-sources--bot_allowlist_policy--examples--group-001.md#canonical-2033001321211220-3332032010012310-1231001310030222-3330121110002310-2012103030201232-3003112023123132-1202202012130003-2231210030102303)
