---
page_title: "xcsh_usb_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy examples."
---

# xcsh_usb_policy examples

<a id="canonical-128b611eeb9ec26ed0bae7757101e8240a1f4e34454d09f8d7e3424f000f5c67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8114a8ad2b710be982084ebb05dcac9daa1f28d8d44d11796d731a3845296b2"></a>

## Examples — Examples / a91dcb398b20 / 2

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)
- Examples

<a id="canonical-6369bc2de5c88fb492e774322b08087b5266e4b396984730788c0c03febfcadb"></a>

## Complete configurations — Examples / a91dcb398b20 / 3

- [Data source](data-sources--usb_policy--examples--group-001.md#canonical-2fd4441ad38a3c92ab599604bda7d453c2589abdcb4efe13ce06152f6f7ec1b3): valid configuration.

<a id="canonical-86b0276a3909823b825514f02cb9c270c46259648eb6b19865902af67aaebe7c"></a>

## Next pages — Examples / a91dcb398b20 / 4

- [Data source](data-sources--usb_policy--examples--group-001.md#canonical-2fd4441ad38a3c92ab599604bda7d453c2589abdcb4efe13ce06152f6f7ec1b3)
- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)

<a id="canonical-2fd4441ad38a3c92ab599604bda7d453c2589abdcb4efe13ce06152f6f7ec1b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32a7f5fbf31c0d45b2f147b9971d11eaaafd538b1f551abc8f60635d105e54e5"></a>

## Data source — Data source / bf08dd6068fd / 2

Breadcrumbs:

- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)
- [Examples](data-sources--usb_policy--examples--group-001.md#canonical-128b611eeb9ec26ed0bae7757101e8240a1f4e34454d09f8d7e3424f000f5c67)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_usb_policy/data-source.tf`; digest `sha256:40aac30ea686e00cce92cecb2aee01cb69e14aef947ae725b686bf0b6e9af7b0`.

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

<a id="canonical-7f26a47d2cd0593c141a18776c3b532dd79833fc1daf35ab7084e860cec57799"></a>

## Next pages — Data source / bf08dd6068fd / 3

- [Examples](data-sources--usb_policy--examples--group-001.md#canonical-128b611eeb9ec26ed0bae7757101e8240a1f4e34454d09f8d7e3424f000f5c67)
- [xcsh_usb_policy](../data-sources/usb_policy.md#canonical-dfe63a993fbbc222b5a61bd9657afcb60469e4538d376446c02e9e1722a38540)
