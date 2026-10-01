---
page_title: "xcsh_log_receiver landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver landing."
---

# xcsh_log_receiver landing

<a id="canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f33e48ec85ca337766dae3fd29e5fbea1c32a225f1630c9c6f6751f20bab5c05"></a>

## xcsh_log_receiver — xcsh_log_receiver / 2da5c0060516 / 2

Breadcrumbs:

- xcsh_log_receiver

Manages new Log Receiver object in F5 Distributed Cloud.

<a id="canonical-3e878d86d8d52259f3b24c853b76342e8c7f0511b197adc9d23cb0bfde60489b"></a>

## Prerequisites — xcsh_log_receiver / 2da5c0060516 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-6159c6cf0f12a2d2927f95996ae2b5ace19a98bd1dbf9df5a3165b9bb19d192c"></a>

## Minimal configuration — xcsh_log_receiver / 2da5c0060516 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# LogReceiver Resource Example
# Manages new Log Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic LogReceiver configuration
resource "xcsh_log_receiver" "example" {
  name      = "example-log-receiver"
  namespace = "staging"
}
```

<a id="canonical-8a726020c4af28cc18d629e8302ee2c76666ed26abad4d250c78c1d3c4f4bec4"></a>

## Root configuration — xcsh_log_receiver / 2da5c0060516 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-7c37fd866db2e601226b5a66784f556ede8dff28d9502ea5d568a13c7eb2c927"></a>

## Next pages — xcsh_log_receiver / 2da5c0060516 / 6

- [Property reference](../guides/resources--log_receiver--reference--group-001.md#canonical-ef02783c49c50c98ba0aaaf5e5b83465c290bcf3f2f090952d79c6692f4536e1)
- [Examples](../guides/resources--log_receiver--examples--group-001.md#canonical-11fae6b302311be6d0fb8acdf7f91098ebe095a218fd88c691fab809d0318fbc)
- [Import](../guides/resources--log_receiver--lifecycle--group-001.md#canonical-dbad79429a01eefa2233734bbd07902e181ce8756cbf60c21abfb6a2e9c324a1)
- [Timeouts](../guides/resources--log_receiver--lifecycle--group-001.md#canonical-ecbe32c890b9175f7e4a5ff8480174e602ef9290459539cc5ec8f130ad404719)
