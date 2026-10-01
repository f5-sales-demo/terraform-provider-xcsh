---
page_title: "xcsh_virtual_k8s landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s landing."
---

# xcsh_virtual_k8s landing

<a id="canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ebc637ef24080a3349015d515e379d4df7ec40448d5e9135c183661601c347d"></a>

## xcsh_virtual_k8s — xcsh_virtual_k8s / a77b04424b28 / 2

Breadcrumbs:

- xcsh_virtual_k8s

Manages virtual\_k8s will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-ff1ab70487c51d58841c66ed1270af6decc2ad09c0d1d44b07aa764d4a3f6044"></a>

## Prerequisites — xcsh_virtual_k8s / a77b04424b28 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `workload`.

- workload: Container workloads in this namespace

<a id="canonical-988beaa088358fb48bc632a7986506906468944c0109d1fffcafd2c81eee73f1"></a>

## Minimal configuration — xcsh_virtual_k8s / a77b04424b28 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualK8S Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualK8S configuration
resource "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}
```

<a id="canonical-550e9674d8ed84f793dc378f2680f80653e8eb3f3ef0c90325c326d643550745"></a>

## Root configuration — xcsh_virtual_k8s / a77b04424b28 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a7e52307dfe424c8fc4197efd41d4eeaf9cf3891bbfc0c98a7857ee12168cee9"></a>

## Next pages — xcsh_virtual_k8s / a77b04424b28 / 6

- [Property reference](../guides/resources--virtual_k8s--reference--group-001.md#canonical-66258665f48680e99a5481bf758757375f97426daea337d5e8870822b96771c8)
- [Examples](../guides/resources--virtual_k8s--examples--group-001.md#canonical-b9dfe44a63b9ab731f1bf8981d28b8e26f5cfe8a38945fcd155244c144832a31)
- [Import](../guides/resources--virtual_k8s--lifecycle--group-001.md#canonical-24f312347127b4009dc077b6f85e751ea6d1dd8064fd8b2246ce946b359d32af)
- [Timeouts](../guides/resources--virtual_k8s--lifecycle--group-001.md#canonical-ac5c7a1a86f8048c811e6b23406804f98a5424bde7a6aa36f3e889fd12946fef)
