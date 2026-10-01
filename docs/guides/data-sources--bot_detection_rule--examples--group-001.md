---
page_title: "xcsh_bot_detection_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_detection_rule examples."
---

# xcsh_bot_detection_rule examples

<a id="canonical-ae5f7faec0025da533bf32857ec6eece19301ba3f0323f8e2964ccac10dfbf02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2dc4fd821a3a6ae0e93a2a0191f63c059e28815694bac2c0e5033f8410d9727d"></a>

## Examples — Examples / 6ab95ca4abbd / 2

Breadcrumbs:

- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md#canonical-a904c12467d27c1ff259b9084cd492f62e963591ee8010367082ed7499705052)
- Examples

<a id="canonical-9a7850b6ab67a8b9bede0ce653cc90f7ac7dd819271c0e42a7be79719e3f6a2c"></a>

## Complete configurations — Examples / 6ab95ca4abbd / 3

- [Data source](data-sources--bot_detection_rule--examples--group-001.md#canonical-55ff1228830a402bbeea16b76dcd18a1621c81fd58df9d6818187f2cec2bd0c8): valid configuration.

<a id="canonical-9e320ded43488cffb7e85552777b3cb1c4540ea21cffd7c32ce27e4ef0834970"></a>

## Next pages — Examples / 6ab95ca4abbd / 4

- [Data source](data-sources--bot_detection_rule--examples--group-001.md#canonical-55ff1228830a402bbeea16b76dcd18a1621c81fd58df9d6818187f2cec2bd0c8)
- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md#canonical-a904c12467d27c1ff259b9084cd492f62e963591ee8010367082ed7499705052)

<a id="canonical-55ff1228830a402bbeea16b76dcd18a1621c81fd58df9d6818187f2cec2bd0c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93b31578bf4c2bba58e0e1ba81e9fd96c894e1b3566cab547a0270c4505c32f6"></a>

## Data source — Data source / 5ea712c17e98 / 2

Breadcrumbs:

- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md#canonical-a904c12467d27c1ff259b9084cd492f62e963591ee8010367082ed7499705052)
- [Examples](data-sources--bot_detection_rule--examples--group-001.md#canonical-ae5f7faec0025da533bf32857ec6eece19301ba3f0323f8e2964ccac10dfbf02)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_detection_rule/data-source.tf`; digest `sha256:fb0cc1608165dd7e57a5b8e5f57c447871327b816ef3b596ef84d5ce798774d2`.

```terraform
# BotDetectionRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDetectionRule by name
data "xcsh_bot_detection_rule" "example" {
  name      = "example-bot-detection-rule"
  namespace = "staging"
}

output "bot_detection_rule_id" {
  value = data.xcsh_bot_detection_rule.example.id
}
```

<a id="canonical-0516dc69df44e589f2c0567b21f9c166935db6cadda272bd86aa05205a2c17c7"></a>

## Next pages — Data source / 5ea712c17e98 / 3

- [Examples](data-sources--bot_detection_rule--examples--group-001.md#canonical-ae5f7faec0025da533bf32857ec6eece19301ba3f0323f8e2964ccac10dfbf02)
- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md#canonical-a904c12467d27c1ff259b9084cd492f62e963591ee8010367082ed7499705052)
