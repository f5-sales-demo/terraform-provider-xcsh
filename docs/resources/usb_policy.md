---
page_title: "xcsh_usb_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy landing."
---

# xcsh_usb_policy landing

<a id="canonical-f73e60f042f8d4c87b077020c3b0feaed53615253bc90e4ee6084a07e1526dae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d270e13f950621a8eb638039b6dba3c22169f19e4fab6d2cd7d0aad3d13f89db"></a>

## xcsh_usb_policy — xcsh_usb_policy / 6c1ff4a72a03 / 2

Breadcrumbs:

- xcsh_usb_policy

Manages new USB policy object in F5 Distributed Cloud.

<a id="canonical-8385fe51d7a73889916baa9dd1c063b374c713f7e968eaa3567bb48e6acc3fab"></a>

## Prerequisites — xcsh_usb_policy / 6c1ff4a72a03 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-be110c86b73a79ed8674ba53cb2a069ade658205837a00f2156255e77e6626d8"></a>

## Minimal configuration — xcsh_usb_policy / 6c1ff4a72a03 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UsbPolicy Resource Example
# Manages new USB policy object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UsbPolicy configuration
resource "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}
```

<a id="canonical-2829c7559891792c393d3791e1011a476567c10165a9aa3bda0c2d89e6578751"></a>

## Root configuration — xcsh_usb_policy / 6c1ff4a72a03 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1217b69fac5e012aa7abf5b2377e6b3d864374918e4b6a4ac1129aaa9a569f69"></a>

## Next pages — xcsh_usb_policy / 6c1ff4a72a03 / 6

- [Property reference](../guides/resources--usb_policy--reference--group-001.md#canonical-238bd627cdad2a74f3c3fd88afa01e42e80e01fe627622780fa621ade10ec15d)
- [Examples](../guides/resources--usb_policy--examples--group-001.md#canonical-1adaa97e282905281884b4084b173c708d144e90757a4955cc9d306a91eb2732)
- [Import](../guides/resources--usb_policy--lifecycle--group-001.md#canonical-552f3b53c95bc3ee9d1bb215f5dd44fe8e6b01d6d54d7d1303c079550ff07752)
- [Timeouts](../guides/resources--usb_policy--lifecycle--group-001.md#canonical-0505d4f4745402d4ddd666c6c82bb406d48781e0b3749451c663ab0bbee9d29a)
