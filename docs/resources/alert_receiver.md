---
page_title: "xcsh_alert_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver landing."
---

# xcsh_alert_receiver landing

<a id="canonical-0fb886588ec1c70113d1093e2a82a5de928b18b20a2a373435526407a441bff5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0544da6be8c51d4803bf9a6c6454aed8198ea4d002c2406dacd0cd91cef82b15"></a>

## xcsh_alert_receiver — xcsh_alert_receiver / 87d24bc2006c / 2

Breadcrumbs:

- xcsh_alert_receiver

Manages new Alert Receiver object in F5 Distributed Cloud.

<a id="canonical-e13ad0b59f87055f202abe7e73f24c5697f9afc31709d0e3232fed4da03616cf"></a>

## Prerequisites — xcsh_alert_receiver / 87d24bc2006c / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-61977ded5154b8ba7e59307310df269f3f84f8d6ba42f546ec1f419de7c34005"></a>

## Minimal configuration — xcsh_alert_receiver / 87d24bc2006c / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-91ee31432804a677973f470091057b1dcf751728bab7c65ebe5ccbce380be9d3"></a>

## Root configuration — xcsh_alert_receiver / 87d24bc2006c / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-48a6c25b310db37f3f92f207590bf9eeb701a0756288d00bf62214606d57f4ac"></a>

## Next pages — xcsh_alert_receiver / 87d24bc2006c / 6

- [Property reference](../guides/resources--alert_receiver--reference--group-001.md#canonical-3bdd8813233b978c5891e876eb017fa9fba9a927108a24d7e64bc84768f2c444)
- [Examples](../guides/resources--alert_receiver--examples--group-001.md#canonical-ba89dbcb9dac7c377504e401a1b7ed2f328cb8fc18ad6bd9894953be2fc2502a)
- [Import](../guides/resources--alert_receiver--lifecycle--group-001.md#canonical-98b097c7e805a760f02fbb39cfea8d793b0484449340483bdf1576bb0498b4a9)
- [Timeouts](../guides/resources--alert_receiver--lifecycle--group-001.md#canonical-db028321f45805788de9d55afee044e8a465949cccc6d5265cb7103541282c3f)
