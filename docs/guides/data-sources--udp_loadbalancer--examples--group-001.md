---
page_title: "xcsh_udp_loadbalancer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer examples."
---

# xcsh_udp_loadbalancer examples

<a id="canonical-f833d9606febec82d58db6db4015687e7ba0a4ad447343f297ba85f5f428ae9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e798b9fa0d6aad963e306e1e5966f6b68e3c8df701e5c9dac36a662036ff5d65"></a>

## Examples — Examples / 37e8069132c0 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- Examples

<a id="canonical-1a5530babbba71b27493524550ab7617662aac01767eaf14f4d06e08e3007dcb"></a>

## Complete configurations — Examples / 37e8069132c0 / 3

- [Data source](data-sources--udp_loadbalancer--examples--group-001.md#canonical-15f16df3de3e1638cbbb1cec05fb1ec20412af422a5b877090655d2389c84b17): valid configuration.

<a id="canonical-bdca0135c75a6edb6a82bcdfa2e05e957e6bfb8dbd0b10143168c226929f7316"></a>

## Next pages — Examples / 37e8069132c0 / 4

- [Data source](data-sources--udp_loadbalancer--examples--group-001.md#canonical-15f16df3de3e1638cbbb1cec05fb1ec20412af422a5b877090655d2389c84b17)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-15f16df3de3e1638cbbb1cec05fb1ec20412af422a5b877090655d2389c84b17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7f60ae5fea2fe029540f0df3dabb76fd3ee71fd241e950774725e145f352d1e"></a>

## Data source — Data source / 6cd8ccb11c2f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Examples](data-sources--udp_loadbalancer--examples--group-001.md#canonical-f833d9606febec82d58db6db4015687e7ba0a4ad447343f297ba85f5f428ae9e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_udp_loadbalancer/data-source.tf`; digest `sha256:90436c5e0e42b785b65d7370ba89219e7c36c5236adccb1abeed194590bc9460`.

```terraform
# UDPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UDPLoadBalancer by name
data "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}

output "udp_loadbalancer_id" {
  value = data.xcsh_udp_loadbalancer.example.id
}
```

<a id="canonical-b21e85ab8ae1940db57e7fcb35bc510a121fa3b3263ce5fdc464b8680e70cfe9"></a>

## Next pages — Data source / 6cd8ccb11c2f / 3

- [Examples](data-sources--udp_loadbalancer--examples--group-001.md#canonical-f833d9606febec82d58db6db4015687e7ba0a4ad447343f297ba85f5f428ae9e)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
