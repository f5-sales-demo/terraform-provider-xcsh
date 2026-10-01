---
page_title: "xcsh_dns_zone examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone examples."
---

# xcsh_dns_zone examples

<a id="canonical-e233a06d738d768cdeb5ad8c949ffcb224dccf43bc25ebcdbe6a486857884d9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6290a184814c3127dd56ebb8fb6895f19bd5ca2e25ffc3f1b6802dafdd661395"></a>

## Examples — Examples / 8272d8364b4b / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- Examples

<a id="canonical-b8b0091b03c6a565b260710fabf7b4e038704958e9d07bf31c522cfc7eb76a66"></a>

## Complete configurations — Examples / 8272d8364b4b / 3

- [Resource](resources--dns_zone--examples--group-001.md#canonical-cfce09799748bfb0cf8c1a2d4b7215f05adceb49646eee583635652535fcf7d0): valid configuration.

<a id="canonical-a9f0ca8521d083e8d02c8025fec6caa42fe0b62177f4e913f9d3459d69f9d8a0"></a>

## Next pages — Examples / 8272d8364b4b / 4

- [Resource](resources--dns_zone--examples--group-001.md#canonical-cfce09799748bfb0cf8c1a2d4b7215f05adceb49646eee583635652535fcf7d0)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-cfce09799748bfb0cf8c1a2d4b7215f05adceb49646eee583635652535fcf7d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de7ebc43b8897a2f8f24df4992611c6346bd8117381ab41137f952b38f96066b"></a>

## Resource — Resource / ba31bedcc0f2 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Examples](resources--dns_zone--examples--group-001.md#canonical-e233a06d738d768cdeb5ad8c949ffcb224dccf43bc25ebcdbe6a486857884d9d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_zone/resource.tf`; digest `sha256:a65c60c4a2a466885fcf90c9e747d8f8ba0052be1693d774ae212f433fd43631`.

```terraform
# DNSZone Resource Example
# Manages DNS Zone in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSZone configuration
resource "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"

  primary {
    allow_http_lb_managed_records = true
  }
}
```

<a id="canonical-0efa1caf8cb7abfdd8e41c84ce0027bc0d4cdd2ce1272429af24da84b4a3eec1"></a>

## Next pages — Resource / ba31bedcc0f2 / 3

- [Examples](resources--dns_zone--examples--group-001.md#canonical-e233a06d738d768cdeb5ad8c949ffcb224dccf43bc25ebcdbe6a486857884d9d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
