---
page_title: "xcsh_usb_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy landing."
---

# xcsh_usb_policy landing

<a id="canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a8065d3dfe59eac241609f7446816f292968ceaef9646e1a55fb92c249d83a0"></a>

## xcsh_usb_policy — xcsh_usb_policy / a66ab44c373f / 2

Breadcrumbs:

- xcsh_usb_policy

Manages new USB policy object in F5 Distributed Cloud.

<a id="canonical-1a9789de77a0da74aa5e3585ada01b364ba6cfd0170c8848db31c9f0e3ba018a"></a>

## Prerequisites — xcsh_usb_policy / a66ab44c373f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-9d7debecff03639742592b961fdb3ab5a62e7120150885ed90c7277ae1a1041b"></a>

## Minimal configuration — xcsh_usb_policy / a66ab44c373f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UsbPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UsbPolicy by name
data "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}

output "usb_policy_id" {
  value = data.xcsh_usb_policy.example.id
}
```

<a id="canonical-fddc6e965ffc738911d928aaf35c91612db0d30b733bc81a10744e2cea643186"></a>

## Root configuration — xcsh_usb_policy / a66ab44c373f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-363d881d28ef2e71cbe7afa6c2a8c24c536fe7922ba4dc92c8eeafc1c8521495"></a>

## Next pages — xcsh_usb_policy / a66ab44c373f / 6

- [Property reference](../guides/data-sources--usb_policy--reference--group-001.md#canonical-407b1cbd64e749afbe04a859df95b1539bfd8fdf4088bc42bfa8bc99410033c0)
- [Examples](../guides/data-sources--usb_policy--examples--group-001.md#canonical-128b611eeb9ec26ed0bae7757101e8240a1f4e34454d09f8d7e3424f000f5c67)
