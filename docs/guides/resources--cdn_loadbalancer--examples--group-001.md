---
page_title: "xcsh_cdn_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer examples."
---

# xcsh_cdn_loadbalancer examples

<a id="canonical-abfd3355bd1d9030ad4bf5366d5a54f3431a0ea24b3a1f50fa5f3c9767812356"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb9ffc48e71a52206e311088ebf39fd98829a242c674a0fd2fb65248de714b17"></a>

## Examples — Examples / 2b3f340dc6a4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- Examples

<a id="canonical-1dd8409b86128b38f3c0287103cb36d5b45ab1ff87c74511662ea6c62010e4bb"></a>

## Complete configurations — Examples / 2b3f340dc6a4 / 3

- [Resource](resources--cdn_loadbalancer--examples--group-001.md#canonical-afa819c350cbfed9ac2d18e04c25d78ea4f748064105a07bf612876517c5059c): valid configuration.

<a id="canonical-68fd501092ced56806db0e2901f28b0320edcdebc89af5d6a70cf6eb580cf967"></a>

## Next pages — Examples / 2b3f340dc6a4 / 4

- [Resource](resources--cdn_loadbalancer--examples--group-001.md#canonical-afa819c350cbfed9ac2d18e04c25d78ea4f748064105a07bf612876517c5059c)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-afa819c350cbfed9ac2d18e04c25d78ea4f748064105a07bf612876517c5059c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a59f1816c8588c3b0984a0e1d6e6ec80a9c2d4a6f9edc700fe9e656941c41c4"></a>

## Resource — Resource / ea7c02f085b4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Examples](resources--cdn_loadbalancer--examples--group-001.md#canonical-abfd3355bd1d9030ad4bf5366d5a54f3431a0ea24b3a1f50fa5f3c9767812356)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_loadbalancer/resource.tf`; digest `sha256:9dbe7d10d66d8f67145599515660dfb082a890fef657898fc4e2002f5c72f6dd`.

```terraform
# CDNLoadBalancer Resource Example
# Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNLoadBalancer configuration
resource "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

<a id="canonical-a37f27410834e318380ed5daba1013cbf19135038f875b4d6f0198ae738bf3a2"></a>

## Next pages — Resource / ea7c02f085b4 / 3

- [Examples](resources--cdn_loadbalancer--examples--group-001.md#canonical-abfd3355bd1d9030ad4bf5366d5a54f3431a0ea24b3a1f50fa5f3c9767812356)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
