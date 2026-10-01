---
page_title: "xcsh_nfv_service landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service landing."
---

# xcsh_nfv_service landing

<a id="canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe49eae72535716275c1f1882a674b01e69ff6b7c4902035e0f5f229de2a029a"></a>

## xcsh_nfv_service — xcsh_nfv_service / 12ad0da006bf / 2

Breadcrumbs:

- xcsh_nfv_service

Manages new NFV service with configured parameters in F5 Distributed Cloud.

<a id="canonical-531a541f8f5d60503621372a3672d5da1b0c5aa4f71c422627b106f87d366800"></a>

## Prerequisites — xcsh_nfv_service / 12ad0da006bf / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b8c0768eb90085c33f36c466cdc523739443da43f8bd74b6a4d491d165f50564"></a>

## Minimal configuration — xcsh_nfv_service / 12ad0da006bf / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NfvService Resource Example
# Manages new NFV service with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NfvService configuration
resource "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}
```

<a id="canonical-eef9d5e555eab05979dcff33a69697c59af3a7f6ba1ed678f2a1903100b478f1"></a>

## Root configuration — xcsh_nfv_service / 12ad0da006bf / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-093a3bf4158773a1430a80bfb223b356a67c3f928293edcf3f0bc3d4aea4920b"></a>

## Next pages — xcsh_nfv_service / 12ad0da006bf / 6

- [Property reference](../guides/resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [Examples](../guides/resources--nfv_service--examples--group-001.md#canonical-fb601c4a70a8553becbfe72e87df89c17d140f83c5c1549cbb177631c1d2b373)
- [Import](../guides/resources--nfv_service--lifecycle--group-001.md#canonical-bb22cebc03e8b53a7325c6bd63fe2064586fe8fe7e34de5160f6b6d2b99a006b)
- [Timeouts](../guides/resources--nfv_service--lifecycle--group-001.md#canonical-8aa9ec314fd24507127f25156aaa112a435faeea12c6b76a74d905a0a3319dc2)
