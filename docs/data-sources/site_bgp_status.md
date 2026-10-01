---
page_title: "xcsh_site_bgp_status landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_bgp_status landing."
---

# xcsh_site_bgp_status landing

<a id="canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f065403f584840574ab569b8bc51b084f18ac84c8e7ddefb46677b634f6c0bb"></a>

## xcsh_site_bgp_status — xcsh_site_bgp_status / 604ca8b5d5d8 / 2

Breadcrumbs:

- xcsh_site_bgp_status

Polls authoritative F5 XC BGP and route observations until they agree with MAC-bound AWS
expectations.

<a id="canonical-a4cacbe896318f25313b79c2a33a2c39ed3cd47cb605bad78a39b2709d5fc867"></a>

## Prerequisites — xcsh_site_bgp_status / 604ca8b5d5d8 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0ed140d384a8cd8c6f7bcfdd5b5249d26f3e6f15a2bf62f7abaf566f59fdd0ec"></a>

## Minimal configuration — xcsh_site_bgp_status / 604ca8b5d5d8 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1f7c5d9cdf8e3bce195aad7be39f83b19583af752990d15e11238b7fd5087d01"></a>

## Root configuration — xcsh_site_bgp_status / 604ca8b5d5d8 / 5

Required root properties: `expected_exported_routes`, `expected_peers`, `namespace`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-7def36d6523bb2fec5af4e6b93e80cb7ab69d06bf12f5e0e05b267e199f977d5"></a>

## Next pages — xcsh_site_bgp_status / 604ca8b5d5d8 / 6

- [Property reference](../guides/data-sources--site_bgp_status--reference--group-001.md#canonical-b3d7a08e010462cf8b9e36f2952a57bc79b1f43a2d6740a60d4ab1649bb5cb27)
- [Examples](../guides/data-sources--site_bgp_status--examples--group-001.md#canonical-2285a7d280f3f9856f43ceebf42040f58085458e3c003570355cba1aa26f96d6)
