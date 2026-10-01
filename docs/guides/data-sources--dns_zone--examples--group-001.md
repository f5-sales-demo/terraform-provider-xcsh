---
page_title: "xcsh_dns_zone examples"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone examples."
---

# xcsh_dns_zone examples

<a id="canonical-17f8d062238fb35a54de166b256b564642f2aadc0f23fe04b8a8b988b597ea70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a0be054bd4d734998f30cad9dfb15f38a14c307fce053fa70554105d996ca03"></a>

## Examples — Examples / 83863e70514b / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- Examples

<a id="canonical-4b24145583da88ed46789dc7f13833ec356bc255992aebe2a163c9b941fab086"></a>

## Complete configurations — Examples / 83863e70514b / 3

- [Data source](data-sources--dns_zone--examples--group-001.md#canonical-4ab54cf195cc566e58d80b7cf30831c2b3d1cc29777173818889bcdb1aa7e05c): valid configuration.

<a id="canonical-4cf4bc2ca31cc7828c3b4896333064a656b0a16df09e745b399912440b4b52a0"></a>

## Next pages — Examples / 83863e70514b / 4

- [Data source](data-sources--dns_zone--examples--group-001.md#canonical-4ab54cf195cc566e58d80b7cf30831c2b3d1cc29777173818889bcdb1aa7e05c)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-4ab54cf195cc566e58d80b7cf30831c2b3d1cc29777173818889bcdb1aa7e05c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa91e77a34aa0bfe89ec8a43bac7170572546d6cfb50a097e0ff2a0bb42817b4"></a>

## Data source — Data source / f7ab74d4a78f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Examples](data-sources--dns_zone--examples--group-001.md#canonical-17f8d062238fb35a54de166b256b564642f2aadc0f23fe04b8a8b988b597ea70)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_zone/data-source.tf`; digest `sha256:45c5f9a8c56f2ed91ab404e89f23a1d58d7ad35539ecaaccb097a0e0066d41b9`.

```terraform
# DNSZone Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSZone by name
data "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"
}

# Fail closed when this stack depends on an externally owned zone.
resource "terraform_data" "require_managed_records" {
  lifecycle {
    precondition {
      condition = try(
        data.xcsh_dns_zone.example.primary.allow_http_lb_managed_records,
        false
      )
      error_message = "The selected DNS zone must enable HTTP LB managed records."
    }
  }
}

output "dns_zone_id" {
  value = data.xcsh_dns_zone.example.id
}
```

<a id="canonical-9c524642f8135ff3552b8081d9388bfda42bcf831d04f1cf3e908d1a044b4612"></a>

## Next pages — Data source / f7ab74d4a78f / 3

- [Examples](data-sources--dns_zone--examples--group-001.md#canonical-17f8d062238fb35a54de166b256b564642f2aadc0f23fe04b8a8b988b597ea70)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
