---
page_title: "xcsh_bot_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure examples."
---

# xcsh_bot_infrastructure examples

<a id="canonical-0312311210210122-3303012200112332-3333133000021021-0301113321021223-3322130112200312-3222123113000213-3101322023012213-3012120222310020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
- Examples

<a id="canonical-0112322233200021-0003301130003223-0313210110103113-0021113220212211-0002100212120220-0212133121102301-3000001231233321-3203203230321323"></a>

### Complete configurations for `xcsh_bot_infrastructure`

- [Resource](resources--bot_infrastructure--examples--group-001.md#canonical-2222320201313333-3321231120231130-1320121030130311-1021311232121221-2131002020323331-3203232230001003-1023233102213300-0323233122330130): valid configuration.

<a id="canonical-2222320201313333-3321231120231130-1320121030130311-1021311232121221-2131002020323331-3203232230001003-1023233102213300-0323233122330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md#canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033)
- [Examples](resources--bot_infrastructure--examples--group-001.md#canonical-0312311210210122-3303012200112332-3333133000021021-0301113321021223-3322130112200312-3222123113000213-3101322023012213-3012120222310020)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bot_infrastructure/resource.tf`; digest `sha256:8bbb63f8f4c5acb98b0adb6309584a0770310129d290ad3dfbcb6667452aa9e4`.

```terraform
# BotInfrastructure Resource Example
# Manages Bot Infrastructure in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotInfrastructure configuration
resource "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}
```
