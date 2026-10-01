---
page_title: "xcsh_cdn_loadbalancer landing"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer landing."
---

# xcsh_cdn_loadbalancer landing

<a id="canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fb2fad2cafb93b3b836cb07eddcd3f90506e3018204234735174d5a47c196d3"></a>

## xcsh_cdn_loadbalancer — xcsh_cdn_loadbalancer / 2195c2f31d3c / 2

Breadcrumbs:

- xcsh_cdn_loadbalancer

Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching
with load balancing.

<a id="canonical-53158dcb9d9aedb57718a953d3f35a6c56084e778bbf02de66f78e75ae39af30"></a>

## Prerequisites — xcsh_cdn_loadbalancer / 2195c2f31d3c / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cdn_origin_pool`.

- cdn_origin_pool: Origin servers for CDN content

<a id="canonical-d1198c7180fea6690d63ca962edbf3240a84cd55b47ce2950c8f0c371a4ea209"></a>

## Minimal configuration — xcsh_cdn_loadbalancer / 2195c2f31d3c / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNLoadBalancer by name
data "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"
}

output "cdn_loadbalancer_id" {
  value = data.xcsh_cdn_loadbalancer.example.id
}
```

<a id="canonical-93a23a71df5e071e3bb9eb2bce1c8e838fe5f01c34f700b326c32d35aa18a6bf"></a>

## Root configuration — xcsh_cdn_loadbalancer / 2195c2f31d3c / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-cd537d14415be07f3db0be00b0d5c5e5169156e592dc4c7129640eef40c80bf0"></a>

## Next pages — xcsh_cdn_loadbalancer / 2195c2f31d3c / 6

- [Property reference](../guides/data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [Examples](../guides/data-sources--cdn_loadbalancer--examples--group-001.md#canonical-d465eef705caed1a30c11c7311fd6668fc60bad15c23d46b23c44c5afc61bab2)
