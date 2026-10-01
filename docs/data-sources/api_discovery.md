---
page_title: "xcsh_api_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery landing."
---

# xcsh_api_discovery landing

<a id="canonical-dbc9e0c75609c449fddf4d0cc2b1bfe4578e8f7f1db459cd2588b04a7a86f2a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5801c6442584df6e1079a510ea4409dc8d00b375ce279c0538f6281474a8c8e"></a>

## xcsh_api_discovery — xcsh_api_discovery / 9ec7404ebd4a / 2

Breadcrumbs:

- xcsh_api_discovery

Manages API discovery creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-9b4f8ef3e0db3c498de7c9b195c5d0b412a32f8834ffe00731edcc49fd57f9c2"></a>

## Prerequisites — xcsh_api_discovery / 9ec7404ebd4a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a5ae7d6b0fbd9758c6f8aff068d12ab50d0b81af02a04707c8ef5f69e1e53a58"></a>

## Minimal configuration — xcsh_api_discovery / 9ec7404ebd4a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDiscovery by name
data "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}

output "api_discovery_id" {
  value = data.xcsh_api_discovery.example.id
}
```

<a id="canonical-6f76c9e0cd1465770145c4f867e34fb9861364a9806ec4f4dc198e05a128b1ea"></a>

## Root configuration — xcsh_api_discovery / 9ec7404ebd4a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-054e2641c023519dbc7a3da7d5af25bf36968e80fbe8ca27cf97972bdc9a914e"></a>

## Next pages — xcsh_api_discovery / 9ec7404ebd4a / 6

- [Property reference](../guides/data-sources--api_discovery--reference--group-001.md#canonical-1f0ff0303fc3ca83f14ad6309047cfbe032c19f845366b31dfa90ad8354282f4)
- [Examples](../guides/data-sources--api_discovery--examples--group-001.md#canonical-f19e1d4dd41f135517981f417fb0513dc02fa573b8a6f2bd5378f259eec47ed7)
