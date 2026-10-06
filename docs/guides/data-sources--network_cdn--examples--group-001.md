---
page_title: "xcsh_network_cdn examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_cdn examples."
---

# xcsh_network_cdn examples

<a id="canonical-1100110203131112-1333321133001020-1320302012333111-0030313021312021-2112102200121123-2030202111321331-2032230133120311-2023332101312322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_cdn](../data-sources/network_cdn.md#canonical-3332101011301320-0312003011132032-2233133133001013-2021212311222213-1103000302133332-3011321012313330-0002302023203120-2131131001121112)
- Examples

<a id="canonical-3212000022001233-2222100311020233-0110110133103133-2112013232101330-2000231322030033-2021003322121212-0220111213321201-2131311200310300"></a>

### Complete configurations for `xcsh_network_cdn`

- [Data source](data-sources--network_cdn--examples--group-001.md#canonical-0001220000332333-3022231313010002-2000312113102232-3222102302012321-2330313023300330-2130333320210031-2010202211031130-2200210112230020): valid configuration.

<a id="canonical-0001220000332333-3022231313010002-2000312113102232-3222102302012321-2330313023300330-2130333320210031-2010202211031130-2200210112230020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_cdn](../data-sources/network_cdn.md#canonical-3332101011301320-0312003011132032-2233133133001013-2021212311222213-1103000302133332-3011321012313330-0002302023203120-2131131001121112)
- [Examples](data-sources--network_cdn--examples--group-001.md#canonical-1100110203131112-1333321133001020-1320302012333111-0030313021312021-2112102200121123-2030202111321331-2032230133120311-2023332101312322)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_cdn/data-source.tf`; digest `sha256:e239b2d01c82b76466c9b21dfd9551a3649d04c98c8f874243856b6d5396e241`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_cdn" "origin_ingress" {}

output "cdn_https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_cdn.origin_ingress.cidr_blocks
  }
}
```
