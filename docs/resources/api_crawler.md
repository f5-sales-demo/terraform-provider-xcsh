---
page_title: "xcsh_api_crawler landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler landing."
---

# xcsh_api_crawler landing

<a id="canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75a5ae2011b77c95236424260c644e0de8d0b72aaddfe3b336940d9c915415ee"></a>

## xcsh_api_crawler — xcsh_api_crawler / f2c1a634a966 / 2

Breadcrumbs:

- xcsh_api_crawler

Manages a API Crawler resource in F5 Distributed Cloud.

<a id="canonical-1699e127b859fc4a8f05e43f185ec55e9c9f20a3d811f74519023a2a468d967c"></a>

## Prerequisites — xcsh_api_crawler / f2c1a634a966 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-9931d34fe7a2207b0b367944a4fdac1320e1a8c9f238940f0ae60be3c4a9bcc7"></a>

## Minimal configuration — xcsh_api_crawler / f2c1a634a966 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APICrawler Resource Example
# Manages a API Crawler resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APICrawler configuration
resource "xcsh_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"
}
```

<a id="canonical-1022eed0ed27c458b0e0a75dbd5fb81724ff377ad4cfb773348df78dd2dc90d3"></a>

## Root configuration — xcsh_api_crawler / f2c1a634a966 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e1390d20d7099af24be37297b469def0121b91e2f6bea71db49a38262371e6e1"></a>

## Next pages — xcsh_api_crawler / f2c1a634a966 / 6

- [Property reference](../guides/resources--api_crawler--reference--group-001.md#canonical-c4a827c8a5e49740b10aec165bceeda2801428c38358cb3a9397450748ba86c5)
- [Examples](../guides/resources--api_crawler--examples--group-001.md#canonical-0da3f92e3e5ef92bbae93921afc86c6f674d415f4e126bbffe8924e640a7f6a5)
- [Import](../guides/resources--api_crawler--lifecycle--group-001.md#canonical-5557aabb1d0a3f857f12aa62a2dad22779ffda8f8decf3dc7e389203304c6105)
- [Timeouts](../guides/resources--api_crawler--lifecycle--group-001.md#canonical-5549eb9e8e5729d1bbe4341c76b5bd80b47a046a7b8e1f313fdd5398dba89dcc)
