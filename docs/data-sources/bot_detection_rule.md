---
page_title: "xcsh_bot_detection_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_detection_rule landing."
---

# xcsh_bot_detection_rule landing

<a id="canonical-a904c12467d27c1ff259b9084cd492f62e963591ee8010367082ed7499705052"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a7bf070cc83aaccb190b134b20e9b42530551c18834f0fb55436e874043d1ba"></a>

## xcsh_bot_detection_rule — xcsh_bot_detection_rule / c4ddc57be508 / 2

Breadcrumbs:

- xcsh_bot_detection_rule

Manages a Bot Detection Rule resource in F5 Distributed Cloud for get bot detection rule.
configuration. (read-only data source)

<a id="canonical-a2f9d1b62086a94062ec42e8194abf0081efa040e7b085e19f72119ac4a2663d"></a>

## Prerequisites — xcsh_bot_detection_rule / c4ddc57be508 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b895922567357243e9ff3daeda9a52ce5316d141ce8a44dc703576ed0769b27b"></a>

## Minimal configuration — xcsh_bot_detection_rule / c4ddc57be508 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-69d03cfa4281cf700922904cbbafdf9ae92020cda7ba5f3fca5134b31117ed0d"></a>

## Root configuration — xcsh_bot_detection_rule / c4ddc57be508 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-99adbdade33065dabf6932960e7fa0353c2c80b77449c990c77e3fc0cdffcdd7"></a>

## Next pages — xcsh_bot_detection_rule / c4ddc57be508 / 6

- [Property reference](../guides/data-sources--bot_detection_rule--reference--group-001.md#canonical-7003e9dcedcdc81d72637cd8ac57d9509b0c46a0c2dd75e7391e39c0ba1ec277)
- [Examples](../guides/data-sources--bot_detection_rule--examples--group-001.md#canonical-ae5f7faec0025da533bf32857ec6eece19301ba3f0323f8e2964ccac10dfbf02)
