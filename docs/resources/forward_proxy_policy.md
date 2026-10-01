---
page_title: "xcsh_forward_proxy_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy landing."
---

# xcsh_forward_proxy_policy landing

<a id="canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd86b304b62950e96fb0f466486fec9fc860808f868a1d535133f44c4e175928"></a>

## xcsh_forward_proxy_policy — xcsh_forward_proxy_policy / 2abf9f083654 / 2

Breadcrumbs:

- xcsh_forward_proxy_policy

Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy
specification. configuration.

<a id="canonical-dd06d0a7582e5ad592a535cef421b0297bda99c08bec5a35806521db36c19edd"></a>

## Prerequisites — xcsh_forward_proxy_policy / 2abf9f083654 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-78b154cacd5b8a2bada2826ac7d6ba754ecdbc8f628309887812ffb5a4fc5ee2"></a>

## Minimal configuration — xcsh_forward_proxy_policy / 2abf9f083654 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardProxyPolicy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardProxyPolicy configuration
resource "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}
```

<a id="canonical-92efeef13e2fa5d429e3965e1786a2b502e0b90e9079f29867cc28a2929254fe"></a>

## Root configuration — xcsh_forward_proxy_policy / 2abf9f083654 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-dbaf23f81fb14b6352a8d7625dc329906ca9aa064d396bd3d59ec9e04ef46bc4"></a>

## Next pages — xcsh_forward_proxy_policy / 2abf9f083654 / 6

- [Property reference](../guides/resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [Examples](../guides/resources--forward_proxy_policy--examples--group-001.md#canonical-8c48d5959f98bb2710e87756ddab35325fdfa094b84af6db832080dd50ddbb18)
- [Import](../guides/resources--forward_proxy_policy--lifecycle--group-001.md#canonical-882d28c3a9657b28c4c7c4cb1edb52a1f43686d15fc043d2c1152c5973812d92)
- [Timeouts](../guides/resources--forward_proxy_policy--lifecycle--group-001.md#canonical-e2d82b10a41c23d967cf70675230b1101aab7aecfa3fe4b8dc64e0493ec391fc)
