---
page_title: "xcsh_dns_lb_health_check examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check examples."
---

# xcsh_dns_lb_health_check examples

<a id="canonical-50d223bed08b617ecb0b68269b4c2165ce73e90bcfaabfaac6924c01a15472df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e9942cf9323f74e5ef97c8cdae6257a61a4735d7ed4cfed119488a6bf27736d"></a>

## Examples — Examples / 7f5a2d1ba2fd / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- Examples

<a id="canonical-e3bfdbfb6dc0d0e423ded27250e7a8ccadcfdfd202f03721e7da5e21f2ff10bd"></a>

## Complete configurations — Examples / 7f5a2d1ba2fd / 3

- [Data source](data-sources--dns_lb_health_check--examples--group-001.md#canonical-8392fd6a91b9d22f1f8bab2eab296da0772b40dbd01c269d9883c9513e21fc64): valid configuration.

<a id="canonical-eb4365300c0ef6981bef4ecc9aa20dfc96127e5df7e56027fb26c092751649cf"></a>

## Next pages — Examples / 7f5a2d1ba2fd / 4

- [Data source](data-sources--dns_lb_health_check--examples--group-001.md#canonical-8392fd6a91b9d22f1f8bab2eab296da0772b40dbd01c269d9883c9513e21fc64)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)

<a id="canonical-8392fd6a91b9d22f1f8bab2eab296da0772b40dbd01c269d9883c9513e21fc64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-baa4ab35d26eda2e64dc50cbb9e550634b335bb013a988256adb216f488a2996"></a>

## Data source — Data source / f99f86e52f1c / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
- [Examples](data-sources--dns_lb_health_check--examples--group-001.md#canonical-50d223bed08b617ecb0b68269b4c2165ce73e90bcfaabfaac6924c01a15472df)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_lb_health_check/data-source.tf`; digest `sha256:687cdf9541bb7302cd31e410ec561a4c9e1516b1f24c1acc1a51bcdfad32aecc`.

```terraform
# DNSLBHealthCheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLBHealthCheck by name
data "xcsh_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}

output "dns_lb_health_check_id" {
  value = data.xcsh_dns_lb_health_check.example.id
}
```

<a id="canonical-1393c07dbc7a550977e32b4a20b5eccfebcb6106a29227e5071ba72a0e9bdba4"></a>

## Next pages — Data source / f99f86e52f1c / 3

- [Examples](data-sources--dns_lb_health_check--examples--group-001.md#canonical-50d223bed08b617ecb0b68269b4c2165ce73e90bcfaabfaac6924c01a15472df)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-d068f5afc6b6bc57e00aa2fc61b83add8b359cf271ca48b98dca29937d9e0d69)
