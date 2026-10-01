---
page_title: "xcsh_log_receiver examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver examples."
---

# xcsh_log_receiver examples

<a id="canonical-11fae6b302311be6d0fb8acdf7f91098ebe095a218fd88c691fab809d0318fbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a18c2981b7ecfb0080c23bc0b21f8d0fc77c844ae6c17fd3ada5a0ac9da35935"></a>

## Examples — Examples / c8f4c6bf0627 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- Examples

<a id="canonical-44af287c18c593751bd042ade911ce240e9346626fda67db91081301d376292d"></a>

## Complete configurations — Examples / c8f4c6bf0627 / 3

- [Resource](resources--log_receiver--examples--group-001.md#canonical-6305a0265eccb406cdedc5ff6def0f8ed194b3bb88648c623deff25d4914bb7a): valid configuration.

<a id="canonical-1ddb18aaaaffce483ca28bc138a8428d497f7f5b99943fc9946eb61437d36051"></a>

## Next pages — Examples / c8f4c6bf0627 / 4

- [Resource](resources--log_receiver--examples--group-001.md#canonical-6305a0265eccb406cdedc5ff6def0f8ed194b3bb88648c623deff25d4914bb7a)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)

<a id="canonical-6305a0265eccb406cdedc5ff6def0f8ed194b3bb88648c623deff25d4914bb7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2087584fd2853116a6e14ec914c31481eb6c1b7eda5bdd5ae205cc11f0bc7de1"></a>

## Resource — Resource / 3657c92632f5 / 2

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
- [Examples](resources--log_receiver--examples--group-001.md#canonical-11fae6b302311be6d0fb8acdf7f91098ebe095a218fd88c691fab809d0318fbc)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_log_receiver/resource.tf`; digest `sha256:5e9b1e9ee7893ce8f260198c000519dd49a791826b61c5ded6feeb808a6b615d`.

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

<a id="canonical-55366c3997650bb496bf6cdd75f52866725ea4b386e917c76d8bd1089ce7ec5c"></a>

## Next pages — Resource / 3657c92632f5 / 3

- [Examples](resources--log_receiver--examples--group-001.md#canonical-11fae6b302311be6d0fb8acdf7f91098ebe095a218fd88c691fab809d0318fbc)
- [xcsh_log_receiver](../resources/log_receiver.md#canonical-375d2736af6e65bedaf1d43b754773d707d996cab5ebd55470571936157cd683)
