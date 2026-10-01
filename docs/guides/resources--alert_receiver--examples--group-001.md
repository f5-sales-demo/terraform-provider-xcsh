---
page_title: "xcsh_alert_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver examples."
---

# xcsh_alert_receiver examples

<a id="canonical-ba89dbcb9dac7c377504e401a1b7ed2f328cb8fc18ad6bd9894953be2fc2502a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c939a992c96e02425820e433e52537d954b1a32ec5dc6174cd21d1a7e2afdfea"></a>

## Examples — Examples / f577cf845f3e / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- Examples

<a id="canonical-dbfbc3420879ebfc74aaadad8a2b777cbc9be19e46619910a3fd71d41c68e56f"></a>

## Complete configurations — Examples / f577cf845f3e / 3

- [Resource](resources--alert_receiver--examples--group-001.md#canonical-32976783b213b2111dcf85c7416be6434688263d6014823bda7d400da93afb66): valid configuration.

<a id="canonical-06ffe08b80a9ecde631d6ba84a6f3969b211f329d1487cc4d4336885cb234037"></a>

## Next pages — Examples / f577cf845f3e / 4

- [Resource](resources--alert_receiver--examples--group-001.md#canonical-32976783b213b2111dcf85c7416be6434688263d6014823bda7d400da93afb66)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)

<a id="canonical-32976783b213b2111dcf85c7416be6434688263d6014823bda7d400da93afb66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc2b4d640943c03f01967771b7cd44564bdaebf7684c3122a47ed91ba53d38e5"></a>

## Resource — Resource / 3404a124c8c8 / 2

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
- [Examples](resources--alert_receiver--examples--group-001.md#canonical-ba89dbcb9dac7c377504e401a1b7ed2f328cb8fc18ad6bd9894953be2fc2502a)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_receiver/resource.tf`; digest `sha256:7520b9716f7e91da22e687315c0f06dd552d7dfa1c32467bdbbdf9c32f5c2efb`.

```terraform
# AlertReceiver Resource Example
# Manages new Alert Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertReceiver configuration
resource "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}
```

<a id="canonical-9240c850ed31b2728ef0cdce4e09c2193ffed4d4f90906941830758d30a49bb2"></a>

## Next pages — Resource / 3404a124c8c8 / 3

- [Examples](resources--alert_receiver--examples--group-001.md#canonical-ba89dbcb9dac7c377504e401a1b7ed2f328cb8fc18ad6bd9894953be2fc2502a)
- [xcsh_alert_receiver](../resources/alert_receiver.md#canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5)
