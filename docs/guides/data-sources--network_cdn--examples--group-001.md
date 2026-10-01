---
page_title: "xcsh_network_cdn examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_cdn examples."
---

# xcsh_network_cdn examples

<a id="canonical-505237567fe5f04878c86fd50cdc9d89964a065b8c895e7d8eb1f6358bf91dba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e600a06faa43522f1451f4df961ee47c80b7a30f890fa66628567e619dd60d30"></a>

## Examples — Examples / 552c95f82ad6 / 2

Breadcrumbs:

- [xcsh_network_cdn](../data-sources/network_cdn.md#canonical-fe445c78360c578eaf7df047899b5aa7530327fec5e46dfc02c8b8d89d741656)
- Examples

<a id="canonical-4689af4838206346ddaf44896d293f4defd85d7e355c2980e4502a607cfd2388"></a>

## Complete configurations — Examples / 552c95f82ad6 / 3

- [Data source](data-sources--network_cdn--examples--group-001.md#canonical-01a00fbfcab7710280d974aeea4b21b9bcdcbc3c9cff890d848a535ca0916b08): valid configuration.

<a id="canonical-e41b10711d1da77cc374df8347e91745ed0a73123441b455e94d24ba42547cbc"></a>

## Next pages — Examples / 552c95f82ad6 / 4

- [Data source](data-sources--network_cdn--examples--group-001.md#canonical-01a00fbfcab7710280d974aeea4b21b9bcdcbc3c9cff890d848a535ca0916b08)
- [xcsh_network_cdn](../data-sources/network_cdn.md#canonical-fe445c78360c578eaf7df047899b5aa7530327fec5e46dfc02c8b8d89d741656)

<a id="canonical-01a00fbfcab7710280d974aeea4b21b9bcdcbc3c9cff890d848a535ca0916b08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52f034d8657ae2d7f5e5e3fa28ad9f487a08fe326db7b505850741b3e0407791"></a>

## Data source — Data source / abc944c5e570 / 2

Breadcrumbs:

- [xcsh_network_cdn](../data-sources/network_cdn.md#canonical-fe445c78360c578eaf7df047899b5aa7530327fec5e46dfc02c8b8d89d741656)
- [Examples](data-sources--network_cdn--examples--group-001.md#canonical-505237567fe5f04878c86fd50cdc9d89964a065b8c895e7d8eb1f6358bf91dba)
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

<a id="canonical-f3ad45b6b44bdfdbed02e02dc37ea3a69c6e15e5106266936947871a82d81526"></a>

## Next pages — Data source / abc944c5e570 / 3

- [Examples](data-sources--network_cdn--examples--group-001.md#canonical-505237567fe5f04878c86fd50cdc9d89964a065b8c895e7d8eb1f6358bf91dba)
- [xcsh_network_cdn](../data-sources/network_cdn.md#canonical-fe445c78360c578eaf7df047899b5aa7530327fec5e46dfc02c8b8d89d741656)
