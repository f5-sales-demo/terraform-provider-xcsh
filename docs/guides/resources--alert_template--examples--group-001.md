---
page_title: "xcsh_alert_template examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template examples."
---

# xcsh_alert_template examples

<a id="canonical-7ec8b75b4bc0ba306d7f1dc1964880f19646ce5e88533bfa6737b7d85e35ccdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99fc0f5bb28fcf3a94a436b2b032c40bb110b54ea32aa9ebd60b8157a7ef2298"></a>

## Examples — Examples / 4760bb5f3619 / 2

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)
- Examples

<a id="canonical-23735219707c016173bf342543e0477812fc355bbdc4d10712533cbe843b26e7"></a>

## Complete configurations — Examples / 4760bb5f3619 / 3

- [Resource](resources--alert_template--examples--group-001.md#canonical-19e1f32e245675f34b21b307d1499a592464c9df1c0356e19dd2f68b207f8c93): valid configuration.

<a id="canonical-95d65ed33b7d5dba73fe91a921204af96f255077ace449a9283b695e65aaf2ca"></a>

## Next pages — Examples / 4760bb5f3619 / 4

- [Resource](resources--alert_template--examples--group-001.md#canonical-19e1f32e245675f34b21b307d1499a592464c9df1c0356e19dd2f68b207f8c93)
- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)

<a id="canonical-19e1f32e245675f34b21b307d1499a592464c9df1c0356e19dd2f68b207f8c93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7f299364db47aa1969fc762b1e3da54a83cd3dfff37f4af3c3464a94fa995be"></a>

## Resource — Resource / 5a16100c3a60 / 2

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)
- [Examples](resources--alert_template--examples--group-001.md#canonical-7ec8b75b4bc0ba306d7f1dc1964880f19646ce5e88533bfa6737b7d85e35ccdb)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_template/resource.tf`; digest `sha256:fca951c776be93b0c4b2e24ed4f4a814383784763f4abf35cbade1403b91638c`.

```terraform
# AlertTemplate Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertTemplate configuration
resource "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"

  alert_message         = "example-value"
  alert_message_details = "example-value"
  alert_name            = "example-value"
}
```

<a id="canonical-5fe54bb9981fa13ef7771b67990c6a323d6665443bd368cf0cfdc6cbeb23fc5f"></a>

## Next pages — Resource / 5a16100c3a60 / 3

- [Examples](resources--alert_template--examples--group-001.md#canonical-7ec8b75b4bc0ba306d7f1dc1964880f19646ce5e88533bfa6737b7d85e35ccdb)
- [xcsh_alert_template](../resources/alert_template.md#canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7)
