---
page_title: "xcsh_bot_defense_app_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_defense_app_infrastructure examples."
---

# xcsh_bot_defense_app_infrastructure examples

<a id="canonical-2133001221300330-1201323101022110-0302232333033321-0230133312113321-1010131012132012-0233102022313333-2302101232213122-3113330201323310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- Examples

<a id="canonical-3130101033231200-2021231231002212-2233330013212222-3113310223121223-0231310331301201-1032130122132133-3020313122210132-0321023211200322"></a>

### Complete configurations for `xcsh_bot_defense_app_infrastructure`

- [Resource](resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-3110122032110102-2033130301023300-2113300333023101-0323203101322332-2031233123012213-3130231323023211-1031230112313130-0232233131122120): valid configuration.

<a id="canonical-3110122032110102-2033130301023300-2113300333023101-0323203101322332-2031233123012213-3130231323023211-1031230112313130-0232233131122120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md#canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121)
- [Examples](resources--bot_defense_app_infrastructure--examples--group-001.md#canonical-2133001221300330-1201323101022110-0302232333033321-0230133312113321-1010131012132012-0233102022313333-2302101232213122-3113330201323310)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bot_defense_app_infrastructure/resource.tf`; digest `sha256:3d3368cb414ef6890130a5c64346a138d87252424badf8402f264c1532265f55`.

```terraform
# BotDefenseAppInfrastructure Resource Example
# Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotDefenseAppInfrastructure configuration
resource "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}
```
