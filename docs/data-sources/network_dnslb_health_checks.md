---
page_title: "xcsh_network_dnslb_health_checks landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_dnslb_health_checks landing."
---

# xcsh_network_dnslb_health_checks landing

<a id="canonical-7fe15c0ea60ec26d670552c2fb5cea28d2a66ac5c60327c73de2c1f6647d76b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14d2f58ed24e6c0aee7be32d8520bc8e4321beba02047cfcc04148d3448a488b"></a>

## xcsh_network_dnslb_health_checks — xcsh_network_dnslb_health_checks / 2cee541f2a3e / 2

Breadcrumbs:

- xcsh_network_dnslb_health_checks

DNS Load Balancer health-check probe IPv4 addresses. Values are bundled from the pinned OpenAPI
release; this data source performs no network request. Ports and traffic direction are not encoded
in the manifest.

<a id="canonical-f88c03ed03590aa679208830a022452bfaee8b22deece22deb9ab2796daff35f"></a>

## Prerequisites — xcsh_network_dnslb_health_checks / 2cee541f2a3e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-169f607771d9bc82818a7f66675a3a362298402437f2d7fba106ac4970248f88"></a>

## Minimal configuration — xcsh_network_dnslb_health_checks / 2cee541f2a3e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

data "xcsh_network_dnslb_health_checks" "https_probe" {}

# Match this explicit ingress port to the monitored endpoint.
output "https_health_check_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_dnslb_health_checks.https_probe.cidr_blocks
  }
}
```

<a id="canonical-a47d7000372f23b0897aff38a6040591cd13817815439434d42ee5ee5941677e"></a>

## Root configuration — xcsh_network_dnslb_health_checks / 2cee541f2a3e / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-1cde1655c7d71ca67554aaabd38bab5b21e176ae1d41bc313f7e3393c3dacf77"></a>

## Next pages — xcsh_network_dnslb_health_checks / 2cee541f2a3e / 6

- [Property reference](../guides/data-sources--network_dnslb_health_checks--reference--group-001.md#canonical-eb3db1582ccee632fd8bc87ae872f5e5735d27c136a8bf35db85d8a0ae5e35ba)
- [Examples](../guides/data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-33d31580d0a9942499cdade7b7ab7db335b8c5e8687bfa6ff986ff99d6735606)
