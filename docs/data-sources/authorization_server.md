---
page_title: "xcsh_authorization_server landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server landing."
---

# xcsh_authorization_server landing

<a id="canonical-c9a064e11c392b1729b64c4abb5a0c14e6107920f082e94b97489dd17471f27a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0d7ce63d300d1c79888b7b9411474d9644b2a44a9915d977f3bec22e08d865d"></a>

## xcsh_authorization_server — xcsh_authorization_server / 7dc8d97fce0a / 2

Breadcrumbs:

- xcsh_authorization_server

Manages authorization\_server creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-75001cc491ff292d77aa9309306cb1cfdcb78ae5ef25bf381bb4d57b5d9b78ce"></a>

## Prerequisites — xcsh_authorization_server / 7dc8d97fce0a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f32e460b17c1c528379d5c9a3f183ccb3b11adfa40271b951f3886e463148db5"></a>

## Minimal configuration — xcsh_authorization_server / 7dc8d97fce0a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AuthorizationServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AuthorizationServer by name
data "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"
}

output "authorization_server_id" {
  value = data.xcsh_authorization_server.example.id
}
```

<a id="canonical-76ca0a8337fa5d25aed849a66df18fa516afea45c9f3ba9f3bbd07cc11a14b74"></a>

## Root configuration — xcsh_authorization_server / 7dc8d97fce0a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a7233e04d28a7ce822c0bb8fcad8d70cba2345e6b1dce5d5a77f322992a42815"></a>

## Next pages — xcsh_authorization_server / 7dc8d97fce0a / 6

- [Property reference](../guides/data-sources--authorization_server--reference--group-001.md#canonical-4fe633a5d8eadcbda1f3928dac2979247f71fce9af1db4ad727fc19ccb660f8d)
- [Examples](../guides/data-sources--authorization_server--examples--group-001.md#canonical-8fc47a5c0c28ccc2cd7f1fabecc67d23713c219f4b177fc3dca6b5ebb9377006)
