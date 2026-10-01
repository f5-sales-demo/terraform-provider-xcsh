---
page_title: "xcsh_bigip_http_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy examples."
---

# xcsh_bigip_http_proxy examples

<a id="canonical-b7b12e3406784cbb40e85608e016edc3ea1d6c4c7b7321810294fc48c7273f64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbabcd9995cb1f9187dc36c79527b90609f2f8e705c6ce77d7f2073519926d71"></a>

## Examples — Examples / 69d7ec6d9bcd / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- Examples

<a id="canonical-5d7c7d0e201aabfac5222807528845dff6e79bad7e9ed67f70d332ea433225e6"></a>

## Complete configurations — Examples / 69d7ec6d9bcd / 3

- [Resource](resources--bigip_http_proxy--examples--group-001.md#canonical-c1803a45ceafa05f00727b5f13713457dfb8f01ce17f806fb214acc4f255ff14): valid configuration.

<a id="canonical-3822901665066599fbb844fe93b957d94c912c8dd1c0e33fc1ccf154ac900c66"></a>

## Next pages — Examples / 69d7ec6d9bcd / 4

- [Resource](resources--bigip_http_proxy--examples--group-001.md#canonical-c1803a45ceafa05f00727b5f13713457dfb8f01ce17f806fb214acc4f255ff14)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-c1803a45ceafa05f00727b5f13713457dfb8f01ce17f806fb214acc4f255ff14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-060ca3b2f451f3d85bcd825f3a6367de532167f8304c0ed7b416ea249501a014"></a>

## Resource — Resource / d257d4a3ac77 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Examples](resources--bigip_http_proxy--examples--group-001.md#canonical-b7b12e3406784cbb40e85608e016edc3ea1d6c4c7b7321810294fc48c7273f64)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bigip_http_proxy/resource.tf`; digest `sha256:52717a495ae2d7db46b088a16cabbff85cbe56ae60a8545793e0ce939a5f736c`.

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

<a id="canonical-edd7eadc53912c3b361244de7ecfbbbaf275d0fa663c3795ee58a8b2f1ffa5d9"></a>

## Next pages — Resource / d257d4a3ac77 / 3

- [Examples](resources--bigip_http_proxy--examples--group-001.md#canonical-b7b12e3406784cbb40e85608e016edc3ea1d6c4c7b7321810294fc48c7273f64)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
