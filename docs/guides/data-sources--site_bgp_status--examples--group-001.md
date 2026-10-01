---
page_title: "xcsh_site_bgp_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_bgp_status examples."
---

# xcsh_site_bgp_status examples

<a id="canonical-2285a7d280f3f9856f43ceebf42040f58085458e3c003570355cba1aa26f96d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cef623411aad7ca104c2bcb8bdb53719e3776ece8b18e775fcc9a35068984e14"></a>

## Examples — Examples / dcd4d51e2a35 / 2

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)
- Examples

<a id="canonical-3c0dcada04da72a15cbf802d49b8633dbcdccfb75d826a9165ef40641020e347"></a>

## Complete configurations — Examples / dcd4d51e2a35 / 3

- [Data source](data-sources--site_bgp_status--examples--group-001.md#canonical-22a60d3e939d3c019519a37c33929e2fa15c178f0c01a137fe8b78b24b4de20e): valid configuration.

<a id="canonical-09fe25a0ef7ae7633f341512f440235427e3cc34613059cb2a430c071d1b2ce5"></a>

## Next pages — Examples / dcd4d51e2a35 / 4

- [Data source](data-sources--site_bgp_status--examples--group-001.md#canonical-22a60d3e939d3c019519a37c33929e2fa15c178f0c01a137fe8b78b24b4de20e)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)

<a id="canonical-22a60d3e939d3c019519a37c33929e2fa15c178f0c01a137fe8b78b24b4de20e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2280467b7917e7e286e236cef9725942da8dece52a99c81999cc16f3bb31ffe"></a>

## Data source — Data source / ecd6471bfe90 / 2

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)
- [Examples](data-sources--site_bgp_status--examples--group-001.md#canonical-2285a7d280f3f9856f43ceebf42040f58085458e3c003570355cba1aa26f96d6)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_bgp_status/data-source.tf`; digest `sha256:78479126aafed9697f5378abaad404d355901539446a65e0ee8620a492c5b9cd`.

```terraform
# Wait for MAC-correlated BGP peers and both BGP and simplified route views to
# converge. Peer addresses and expected routes come from authoritative AWS
# TGW Connect resource attributes in a real configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_site_bgp_status" "site" {
  namespace = "system"
  site      = "example-smsv2-site"

  expected_exported_routes = ["10.40.0.10/32"]

  expected_peers = {
    node_0_slo = {
      node                     = "node-0"
      role                     = "slo"
      mac                      = "02:00:00:00:00:10"
      peer_address             = "169.254.100.1"
      expected_imported_routes = ["10.20.0.0/16"]
    }
    node_0_sli = {
      node                     = "node-0"
      role                     = "sli"
      mac                      = "02:00:00:00:00:11"
      peer_address             = "169.254.101.1"
      expected_imported_routes = ["10.30.0.0/16"]
    }
  }

  timeout_seconds       = 300
  poll_interval_seconds = 10
}

output "bgp_converged" {
  value = data.xcsh_site_bgp_status.site.converged
}

output "bgp_peers" {
  value = data.xcsh_site_bgp_status.site.peers
}
```

<a id="canonical-4a17df58c3b6733f3e103a637d3520a4a8d4a060065c4634ee78648606ce4080"></a>

## Next pages — Data source / ecd6471bfe90 / 3

- [Examples](data-sources--site_bgp_status--examples--group-001.md#canonical-2285a7d280f3f9856f43ceebf42040f58085458e3c003570355cba1aa26f96d6)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)
