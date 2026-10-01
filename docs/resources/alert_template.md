---
page_title: "xcsh_alert_template landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template landing."
---

# xcsh_alert_template landing

<a id="canonical-97fd669adea834d391f4ccd360fb87c16cdfc5afd855377f3e5665ed96c729d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d09a79cbfb8745db75eb4d3ed46f07c5af7889ad00d0673192671ea74d9fc4a2"></a>

## xcsh_alert_template — xcsh_alert_template / 52f7aefe3f1d / 2

Breadcrumbs:

- xcsh_alert_template

Manages Domain to protect in F5 Distributed Cloud.

<a id="canonical-a4dee82c7466f74f712d2f1a952b6ed5547385fddf3720a624281fca3816f821"></a>

## Prerequisites — xcsh_alert_template / 52f7aefe3f1d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-782a1d58939bdcac46dd07b612ee915280ab635fe47c094c6893fb73a6b5358f"></a>

## Minimal configuration — xcsh_alert_template / 52f7aefe3f1d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-f16def9ddc75f026683b7ff0bb9c9d8341ea36c6f0f2a8f298d28d958adde77c"></a>

## Root configuration — xcsh_alert_template / 52f7aefe3f1d / 5

Required root properties: `alert_message`, `alert_message_details`, `alert_name`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d1b37c8692274e3c2f2a9a658c203c476b5337580e13bb3387ee843056a276ee"></a>

## Next pages — xcsh_alert_template / 52f7aefe3f1d / 6

- [Property reference](../guides/resources--alert_template--reference--group-001.md#canonical-5fdfbe7749f50797d48aa600091b2ae0aae8cf488d15e19a3a5448fc91baf7ef)
- [Examples](../guides/resources--alert_template--examples--group-001.md#canonical-7ec8b75b4bc0ba306d7f1dc1964880f19646ce5e88533bfa6737b7d85e35ccdb)
- [Import](../guides/resources--alert_template--lifecycle--group-001.md#canonical-c1ff8d22a3a82ff3b3ef385e9f7d43a288ce00f11fabc12b2927cd8e614e6483)
- [Timeouts](../guides/resources--alert_template--lifecycle--group-001.md#canonical-1964042c27228e53abf7cec2a6c9f835da18d8385c88f07fc1a0bcc93dd2de30)
