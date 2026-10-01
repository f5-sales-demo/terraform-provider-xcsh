---
page_title: "xcsh_bgp_asn_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set examples."
---

# xcsh_bgp_asn_set examples

<a id="canonical-75d91db6035cd23ac099067374b4d55f332b44657b056ee864a34c304fa3610b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d097bef6fbc08f0ec96bd3c6fea0fd2d944d9fae550f4097018c57d6bfc0a564"></a>

## Examples — Examples / 2460086b97f0 / 2

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)
- Examples

<a id="canonical-5ae2b86850a589791b1f1adcf71a169a0bf5972c40e01dddd269f802f3aa6ff9"></a>

## Complete configurations — Examples / 2460086b97f0 / 3

- [Resource](resources--bgp_asn_set--examples--group-001.md#canonical-68ced3fa70ae751b35e37d5e00a101e0d55c83bde0433399a219bc8b085d7448): valid configuration.

<a id="canonical-f393aa818ed9bbb0021162fd69fed8014e83caf2d1770b6c7f242234b78975c5"></a>

## Next pages — Examples / 2460086b97f0 / 4

- [Resource](resources--bgp_asn_set--examples--group-001.md#canonical-68ced3fa70ae751b35e37d5e00a101e0d55c83bde0433399a219bc8b085d7448)
- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)

<a id="canonical-68ced3fa70ae751b35e37d5e00a101e0d55c83bde0433399a219bc8b085d7448"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-939f3b768c89cb974bd4698cfcf56d1f9fb9b1ddfd0b40a81da2edc69480775c"></a>

## Resource — Resource / dd66f528b70f / 2

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)
- [Examples](resources--bgp_asn_set--examples--group-001.md#canonical-75d91db6035cd23ac099067374b4d55f332b44657b056ee864a34c304fa3610b)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp_asn_set/resource.tf`; digest `sha256:3c2bb9b5fa3b8555e63a99053456d35fa40e857a94daa94de302421615af065a`.

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

<a id="canonical-025a8173ec86e73148606979e321cffaaebc5c65f5913644d79dbcd0133998ff"></a>

## Next pages — Resource / dd66f528b70f / 3

- [Examples](resources--bgp_asn_set--examples--group-001.md#canonical-75d91db6035cd23ac099067374b4d55f332b44657b056ee864a34c304fa3610b)
- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-c4139b5a1ff4b2663ce284b2126b8739541425b2c4962c040d4f86ceec7cc653)
