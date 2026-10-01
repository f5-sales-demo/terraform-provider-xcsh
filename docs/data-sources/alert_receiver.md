---
page_title: "xcsh_alert_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver landing."
---

# xcsh_alert_receiver landing

<a id="canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50b535c8c082c0ecdde7082427158543599a851f734603205e2f03dc2da57a3c"></a>

## xcsh_alert_receiver — xcsh_alert_receiver / b2c50b4b3385 / 2

Breadcrumbs:

- xcsh_alert_receiver

Manages new Alert Receiver object in F5 Distributed Cloud.

<a id="canonical-82aece4663ed1582f15148f351678a9127a21aed502a2ef982aac8dbc65aad6d"></a>

## Prerequisites — xcsh_alert_receiver / b2c50b4b3385 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f40e9fe1d172e42415db5429204967f22e50142ec7664d6b66062c9c4fda9af3"></a>

## Minimal configuration — xcsh_alert_receiver / b2c50b4b3385 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AlertReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertReceiver by name
data "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}

output "alert_receiver_id" {
  value = data.xcsh_alert_receiver.example.id
}
```

<a id="canonical-c4f95b34e03bf3d5d95d45a2931a4407015c58735bbf0104d4a740d501498dd5"></a>

## Root configuration — xcsh_alert_receiver / b2c50b4b3385 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-9663580acd9e562872b6c334f940c86a662c1d570b9bd46c263dc430618ea862"></a>

## Next pages — xcsh_alert_receiver / b2c50b4b3385 / 6

- [Property reference](../guides/data-sources--alert_receiver--reference--group-001.md#canonical-9e77ba4988d239ee5fcf37742214d6c3d97c59c318f7ff0d06d4ec6a09a48451)
- [Examples](../guides/data-sources--alert_receiver--examples--group-001.md#canonical-9c611f5c98034d913775613ac7408603296d5248701a996dd4d98f6a2bd352b4)
