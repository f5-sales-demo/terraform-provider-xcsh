---
page_title: "xcsh_bigip_http_proxy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy landing."
---

# xcsh_bigip_http_proxy landing

<a id="canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9959ca33b9fdb4a1fd21f3d713a6d84d8a6a5b2a0336e00f17ba97a6356767be"></a>

## xcsh_bigip_http_proxy — xcsh_bigip_http_proxy / e532e45ea080 / 2

Breadcrumbs:

- xcsh_bigip_http_proxy

Manages BIG-IP HTTP Proxy in a given namespace. If one already exists, it will give an error in F5
Distributed Cloud.

<a id="canonical-49c8b4b88eac064978a1bc145d52a2fd368d662178c0961adbe87a80dbc3bed8"></a>

## Prerequisites — xcsh_bigip_http_proxy / e532e45ea080 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8f749a529caff0be26fc08faa7b697a4ba413ecc09fcac002b41bea76ef67ff1"></a>

## Minimal configuration — xcsh_bigip_http_proxy / e532e45ea080 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPHTTPProxy Resource Example
# Manages BIG-IP HTTP Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BigIPHTTPProxy configuration
resource "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}
```

<a id="canonical-cbbbd1651315663f31a712930c551f76fe2bb137c52b6baa969da9c82ebc15dc"></a>

## Root configuration — xcsh_bigip_http_proxy / e532e45ea080 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-aeb6178d4d3c93ba9294fc01b94650f34a0c0d3853a35f52f2a5f0eb0fef0ef4"></a>

## Next pages — xcsh_bigip_http_proxy / e532e45ea080 / 6

- [Property reference](../guides/resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [Examples](../guides/resources--bigip_http_proxy--examples--group-001.md#canonical-b7b12e3406784cbb40e85608e016edc3ea1d6c4c7b7321810294fc48c7273f64)
- [Import](../guides/resources--bigip_http_proxy--lifecycle--group-001.md#canonical-7a43975bf5af7ed7ac0424dfd0b760392a058e64ba8b6a2d5d07a06aa386c346)
- [Timeouts](../guides/resources--bigip_http_proxy--lifecycle--group-001.md#canonical-0e26e27a4d09cd4d0ce2b4d4ce4c39062d5c20f01484ce9c87fd8fb07b89d665)
