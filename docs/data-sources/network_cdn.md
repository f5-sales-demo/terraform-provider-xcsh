---
page_title: "xcsh_network_cdn landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_cdn landing."
---

# xcsh_network_cdn landing

<a id="canonical-fe445c78360c578eaf7df047899b5aa7530327fec5e46dfc02c8b8d89d741656"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75c4ceb8aa32eba612527509a1af73a3ec0472fb1145e780434720a274f72c2c"></a>

## xcsh_network_cdn — xcsh_network_cdn / 8176fb3d2ad1 / 2

Breadcrumbs:

- xcsh_network_cdn

CDN IPv4 networks for origin or network-firewall ingress allowlists. Values are bundled from the
pinned OpenAPI release; this data source performs no network request. Ports and traffic direction
are not encoded in the manifest.

<a id="canonical-41c7ab25f3b9edcc3dc13bde77ddb7d87492c0e0f0c38d40e3d13cbdacdac7a2"></a>

## Prerequisites — xcsh_network_cdn / 8176fb3d2ad1 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-469d3645573dd20af48e763d3eca07bc77db398f8e736dd1cef835865e201805"></a>

## Minimal configuration — xcsh_network_cdn / 8176fb3d2ad1 / 4

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

data "xcsh_network_cdn" "origin_ingress" {}

output "cdn_https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_cdn.origin_ingress.cidr_blocks
  }
}
```

<a id="canonical-7b3ee34744ea46599f9315508d8f3f75acdc6303c9d489251580c50b3880494d"></a>

## Root configuration — xcsh_network_cdn / 8176fb3d2ad1 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-087ecbe9f08ed4f4937587c2792a0560cd7c71adc0f8a3ff7ce5b4f61ff60c82"></a>

## Next pages — xcsh_network_cdn / 8176fb3d2ad1 / 6

- [Property reference](../guides/data-sources--network_cdn--reference--group-001.md#canonical-cbda8b2c4920edab916746235744ca1fdca4a914e282936ed48fa87e4d80bb47)
- [Examples](../guides/data-sources--network_cdn--examples--group-001.md#canonical-505237567fe5f04878c86fd50cdc9d89964a065b8c895e7d8eb1f6358bf91dba)
