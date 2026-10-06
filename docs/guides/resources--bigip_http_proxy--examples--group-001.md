---
page_title: "xcsh_bigip_http_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy examples."
---

# xcsh_bigip_http_proxy examples

<a id="canonical-2313230102320310-0012132010302323-1000322011120020-3200011232313003-3222013112301030-1323130302012001-0002211033301020-3013021303331210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- Examples

<a id="canonical-3023222330312121-2111302301332101-2013313003123013-2111021323210012-0021330233203213-0011301230321313-3113330200130311-0121210212311301"></a>

### Complete configurations for `xcsh_bigip_http_proxy`

- [Resource](resources--bigip_http_proxy--examples--group-001.md#canonical-3001200003221011-3032223322001133-0000130213231133-0103130103101113-3133232033000130-3201133320001233-2302011022303010-3302111133330110): valid configuration.

<a id="canonical-3001200003221011-3032223322001133-0000130213231133-0103130103101113-3133232033000130-3201133320001233-2302011022303010-3302111133330110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Examples](resources--bigip_http_proxy--examples--group-001.md#canonical-2313230102320310-0012132010302323-1000322011120020-3200011232313003-3222013112301030-1323130302012001-0002211033301020-3013021303331210)
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
