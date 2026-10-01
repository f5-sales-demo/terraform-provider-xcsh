---
page_title: "xcsh_nginx_service_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery examples."
---

# xcsh_nginx_service_discovery examples

<a id="canonical-766ba811139fdb5dc6fb3bb8d510b3535c22c310114d846a562ea2bab8392701"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2c4f57180b87d780c12d4855307db9f33f1e32ab058ab776794f7389e5e6ca9"></a>

## Examples — Examples / 072cea164338 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- Examples

<a id="canonical-58237389e9ae2160fac52f0351e4fa35062363884c1438269681f67a38a7b9c6"></a>

## Complete configurations — Examples / 072cea164338 / 3

- [Data source](data-sources--nginx_service_discovery--examples--group-001.md#canonical-34c1bb7f86626e492c39a356a559d08812a66c9738938664a6477b595748a609): valid configuration.

<a id="canonical-2e20df4c0d5584dd55a2730cc8c439422b149a171e33dfc3225a3d36f0d94322"></a>

## Next pages — Examples / 072cea164338 / 4

- [Data source](data-sources--nginx_service_discovery--examples--group-001.md#canonical-34c1bb7f86626e492c39a356a559d08812a66c9738938664a6477b595748a609)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)

<a id="canonical-34c1bb7f86626e492c39a356a559d08812a66c9738938664a6477b595748a609"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e6d921127d62ebe22ec7ee04cf0b2e956b5d066cbf1d99a846ec3fc23539a9f"></a>

## Data source — Data source / a642e5b86906 / 2

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
- [Examples](data-sources--nginx_service_discovery--examples--group-001.md#canonical-766ba811139fdb5dc6fb3bb8d510b3535c22c310114d846a562ea2bab8392701)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_service_discovery/data-source.tf`; digest `sha256:c46b908211689c5888bfaaf70a487544236d409e407922fa907fe86ca0588fa6`.

```terraform
# NginxServiceDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServiceDiscovery by name
data "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}

output "nginx_service_discovery_id" {
  value = data.xcsh_nginx_service_discovery.example.id
}
```

<a id="canonical-03db9cdb3e29c31021ac3f1cd8a71c02d29ca2dd9283ee96c36c1305adc786e0"></a>

## Next pages — Data source / a642e5b86906 / 3

- [Examples](data-sources--nginx_service_discovery--examples--group-001.md#canonical-766ba811139fdb5dc6fb3bb8d510b3535c22c310114d846a562ea2bab8392701)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md#canonical-20a77ba34df172f8899cd77b9d06e83d5056bcba93fb611466b7f24b213ccfb5)
