---
page_title: "xcsh_udp_loadbalancer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer landing."
---

# xcsh_udp_loadbalancer landing

<a id="canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1cc91f6d1cca3309759e6f662651ae628c78443072cecd8c7974f27b638b528"></a>

## xcsh_udp_loadbalancer — xcsh_udp_loadbalancer / 6102c82f6c9f / 2

Breadcrumbs:

- xcsh_udp_loadbalancer

Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across
origin pools.

<a id="canonical-7ee6a7a1238ad775ca1ccddcd02143f02e0c8f46fc7986bfe3035202f1a13128"></a>

## Prerequisites — xcsh_udp_loadbalancer / 6102c82f6c9f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4ba6533e4c88b91a1b07bcdef25aff36c4f8113ed3b89f26af2ca088b3902148"></a>

## Minimal configuration — xcsh_udp_loadbalancer / 6102c82f6c9f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-a159c981ea0daaebcb5adca2cdbdd66bd80edd79b9e8bcffca34e6a292765b60"></a>

## Root configuration — xcsh_udp_loadbalancer / 6102c82f6c9f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e8dbcc55f621ac598e241d3bf870108bf60c6ea16909412b5fb4f9b11543cd6e"></a>

## Next pages — xcsh_udp_loadbalancer / 6102c82f6c9f / 6

- [Property reference](../guides/data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [Examples](../guides/data-sources--udp_loadbalancer--examples--group-001.md#canonical-f833d9606febec82d58db6db4015687e7ba0a4ad447343f297ba85f5f428ae9e)
