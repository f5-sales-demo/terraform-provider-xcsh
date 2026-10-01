---
page_title: "xcsh_bot_suggest_values landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_suggest_values landing."
---

# xcsh_bot_suggest_values landing

<a id="canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5b3b92488807553d63a7a43dda88a3150fe82777c4f5525b7718a78cca57e9c"></a>

## xcsh_bot_suggest_values — xcsh_bot_suggest_values / 75008e866ab6 / 2

Breadcrumbs:

- xcsh_bot_suggest_values

Resource creation operation.

<a id="canonical-1c2481c176ca8162a9ddfcb4413b22605fe02fb5b298d54fb8022db866cdcfc0"></a>

## Prerequisites — xcsh_bot_suggest_values / 75008e866ab6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b6e98ea0e9ea35453144b3584c7d5923b44f484e02bf135a5bcd78fa861174e3"></a>

## Minimal configuration — xcsh_bot_suggest_values / 75008e866ab6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-420f660c54361d983ddf7fb10127caa90ed5423a471617f3f08c05657ccfcde4"></a>

## Root configuration — xcsh_bot_suggest_values / 75008e866ab6 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-035145d3a88564bcdbc4a91f93fe47c400efa17d44c3062663f41395fa5b0a37"></a>

## Next pages — xcsh_bot_suggest_values / 75008e866ab6 / 6

- [Property reference](../guides/data-sources--bot_suggest_values--reference--group-001.md#canonical-1fd46f22c3a31585d917103e1b4692467b3651cd75d9fe25d8cf70497503b572)
- [Examples](../guides/data-sources--bot_suggest_values--examples--group-001.md#canonical-cb1677c61520b70852bab053740849ce32413907dccd3d431b685980ca92525d)
