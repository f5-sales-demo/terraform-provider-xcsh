---
page_title: "xcsh_bgp_asn_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set landing."
---

# xcsh_bgp_asn_set landing

<a id="canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37de172b57d669155bf985917a266b9500fee8758b8316437221a21988e2b52c"></a>

## xcsh_bgp_asn_set — xcsh_bgp_asn_set / eea48d16e2e3 / 2

Breadcrumbs:

- xcsh_bgp_asn_set

Manages bgp\_asn\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-a6aa486904c4a9b03f36ca272db45a335851d17f6080509478d9b1cc743658d7"></a>

## Prerequisites — xcsh_bgp_asn_set / eea48d16e2e3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-351345a73637e1a3e06ace74650ccd99c67c623c4a23b8a064d825e52c0a1391"></a>

## Minimal configuration — xcsh_bgp_asn_set / eea48d16e2e3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPAsnSet Resource Example
# Manages bgp_asn_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPAsnSet configuration
resource "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"

  as_numbers = [1]
}
```

<a id="canonical-ceea4c0fe9fe97fb182c47bb5c15cf9abc2b1deb206e926ddb130f5ae98f3885"></a>

## Root configuration — xcsh_bgp_asn_set / eea48d16e2e3 / 5

Required root properties: `as_numbers`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e7f9b9d1797f8b027da8b92c48edbfd69ff2fe2cd2d2622e0e9966e6c8250d91"></a>

## Next pages — xcsh_bgp_asn_set / eea48d16e2e3 / 6

- [Property reference](../guides/resources--bgp_asn_set--reference--group-001.md#canonical-a9058fc90b7a6c3e30352fc706291f4deeabb26d142eebfe18a505b212bc5a60)
- [Examples](../guides/resources--bgp_asn_set--examples--group-001.md#canonical-75d91db6035cd23ac099067374b4d55f332b44657b056ee864a34c304fa3610b)
- [Import](../guides/resources--bgp_asn_set--lifecycle--group-001.md#canonical-44460d5c4a3bcbca5146850b4f6803edea768512008ea9af0f5e8cbedc4bc590)
- [Timeouts](../guides/resources--bgp_asn_set--lifecycle--group-001.md#canonical-ad45eeb50d8f2f45538ac8298838e4fa028c47fe3afe5bac0abcd436bf5726cf)
