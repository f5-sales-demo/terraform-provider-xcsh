---
page_title: "xcsh_network_dnslb_health_checks examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_dnslb_health_checks examples."
---

# xcsh_network_dnslb_health_checks examples

<a id="canonical-33d31580d0a9942499cdade7b7ab7db335b8c5e8687bfa6ff986ff99d6735606"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d2f2ce70fb6a709fda0a07302c9bfb7848e89a9e052a5b1c6df6b32f090d588"></a>

## Examples — Examples / 2a20e35b493c / 2

Breadcrumbs:

- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md#canonical-7fe15c0ea60ec26d670552c2fb5cea28d2a66ac5c60327c73de2c1f6647d76b6)
- Examples

<a id="canonical-2ae56ce47bc779acc220fa0222992fa52a80513c41f44bd0a897afb3547acd74"></a>

## Complete configurations — Examples / 2a20e35b493c / 3

- [Data source](data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-395781db2bfad350c0a2309baea4b9bfbde18c383acbd7494a74488655ebcd01): valid configuration.

<a id="canonical-4d13af750ff45887ec26a4030e3d101c8f69b2d36f59231296cab9496e68f985"></a>

## Next pages — Examples / 2a20e35b493c / 4

- [Data source](data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-395781db2bfad350c0a2309baea4b9bfbde18c383acbd7494a74488655ebcd01)
- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md#canonical-7fe15c0ea60ec26d670552c2fb5cea28d2a66ac5c60327c73de2c1f6647d76b6)

<a id="canonical-395781db2bfad350c0a2309baea4b9bfbde18c383acbd7494a74488655ebcd01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fccdf6e9467cde3636501792e5f491d72310993b3eabe2bc79faa5f7ae978b8"></a>

## Data source — Data source / 5c6f3e861a76 / 2

Breadcrumbs:

- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md#canonical-7fe15c0ea60ec26d670552c2fb5cea28d2a66ac5c60327c73de2c1f6647d76b6)
- [Examples](data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-33d31580d0a9942499cdade7b7ab7db335b8c5e8687bfa6ff986ff99d6735606)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_dnslb_health_checks/data-source.tf`; digest `sha256:806558a3420657577ad88d195487a5583b29bf785d14c8b410da5c5fe1315a79`.

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

<a id="canonical-1472500c20c58aa7c0eac8e195b9cc0176e0f2d910988b706ad09f1119fe47b1"></a>

## Next pages — Data source / 5c6f3e861a76 / 3

- [Examples](data-sources--network_dnslb_health_checks--examples--group-001.md#canonical-33d31580d0a9942499cdade7b7ab7db335b8c5e8687bfa6ff986ff99d6735606)
- [xcsh_network_dnslb_health_checks](../data-sources/network_dnslb_health_checks.md#canonical-7fe15c0ea60ec26d670552c2fb5cea28d2a66ac5c60327c73de2c1f6647d76b6)
