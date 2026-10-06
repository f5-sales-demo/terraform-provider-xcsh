---
page_title: "xcsh_dns_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy examples."
---

# xcsh_dns_proxy examples

<a id="canonical-2133222100312232-1221112223202121-2002020000120023-2022033232120030-3110313330333121-1201030313001033-2320210131022111-0212023110003211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- Examples

<a id="canonical-0120131330311023-0210033322120232-0321023110022001-2223121221233333-1213323023111013-3302011123223213-1211222010120202-1222033001100021"></a>

### Complete configurations for `xcsh_dns_proxy`

- [Resource](resources--dns_proxy--examples--group-001.md#canonical-1023023213302301-1013332103111020-1223023020331033-1222112203110332-1121320322130100-1232001012320001-2002332201010123-1323321100230330): valid configuration.

<a id="canonical-1023023213302301-1013332103111020-1223023020331033-1222112203110332-1121320322130100-1232001012320001-2002332201010123-1323321100230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Examples](resources--dns_proxy--examples--group-001.md#canonical-2133222100312232-1221112223202121-2002020000120023-2022033232120030-3110313330333121-1201030313001033-2320210131022111-0212023110003211)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_proxy/resource.tf`; digest `sha256:943728fe8d11ffeb7bc560a6a0e3b2ff777ab7fb03003f4fd0ef80073c303500`.

```terraform
# DNSProxy Resource Example
# Manages DNS Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSProxy configuration
resource "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}
```
