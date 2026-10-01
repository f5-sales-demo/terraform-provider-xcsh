---
page_title: "xcsh_endpoint landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint landing."
---

# xcsh_endpoint landing

<a id="canonical-4cf8c7cfcd5876238831708e5962c7ec0ef4cf3a14b5dd028d0037907895b506"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aaab05f936e36defe888f92ce3c4ffb6c219a80ec5e0563634ccf7d8a5cdc1df"></a>

## xcsh_endpoint — xcsh_endpoint / 8b9099c23cc6 / 2

Breadcrumbs:

- xcsh_endpoint

Manages endpoint will create the object in the storage backend for namespace metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-15f0f76b58352c0d6723ac1aa12fa462d10644af776af13b96ace7dfd32cb725"></a>

## Prerequisites — xcsh_endpoint / 8b9099c23cc6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-8c2c5e7daafdf295dc28b8f7a364f8cfb20286b3d30a87a55c17cbfc922d967d"></a>

## Minimal configuration — xcsh_endpoint / 8b9099c23cc6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```

<a id="canonical-37d8f45193bf2be8833d7bfb22af480c0ce22da15a7e4e8188bc527d09b6d61f"></a>

## Root configuration — xcsh_endpoint / 8b9099c23cc6 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-922e2cebb7e5fa46ca03ec36fffb71b4116c1fccd55fc22b738fe898be8437e1"></a>

## Next pages — xcsh_endpoint / 8b9099c23cc6 / 6

- [Property reference](../guides/resources--endpoint--reference--group-001.md#canonical-a2069cc7bd5814e4f0146ee477f37273fa022196ffceca0b64578ac1c6059bdd)
- [Examples](../guides/resources--endpoint--examples--group-001.md#canonical-59458105bf2576af045aeb7a0d428b8b8147593e44f513c8b62c4f392b72181a)
- [Import](../guides/resources--endpoint--lifecycle--group-001.md#canonical-8550c658667519834e127ac6117774f22d5a657a26eaf5ef2c6337a0a8696a3a)
- [Timeouts](../guides/resources--endpoint--lifecycle--group-001.md#canonical-9d724647f290d77db38ff0889cdf8e28360a5ab517ca1f09692224092ce725f3)
