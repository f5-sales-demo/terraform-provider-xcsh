---
page_title: "xcsh_tunnel landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel landing."
---

# xcsh_tunnel landing

<a id="canonical-4a29ad0424b0245b80e17e625816f6cf1350032e2eedd0aa4c64c636f219e081"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0fbe18aac8a085826370aca5cb0a38bc2aed695fd94a40f11da1ad3a504015c"></a>

## xcsh_tunnel — xcsh_tunnel / bde3443da9de / 2

Breadcrumbs:

- xcsh_tunnel

Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

<a id="canonical-319c820fed440dc01b9a427c32b38ecf0fa2b4feab27d4ba56c8e1cb80095449"></a>

## Prerequisites — xcsh_tunnel / bde3443da9de / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-659629fce2d8d997cd9553d92b1dab474bf578d15e363b731ae28411ca93c89f"></a>

## Minimal configuration — xcsh_tunnel / bde3443da9de / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Tunnel by name
data "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}

output "tunnel_id" {
  value = data.xcsh_tunnel.example.id
}
```

<a id="canonical-a4229c025ee5c371478f20412a67ca8ab9a03536b083c36b075227580bc21ed6"></a>

## Root configuration — xcsh_tunnel / bde3443da9de / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-77e04517b44c812d49cd2158ccf383a0a5d114ce36fdd2c3f3aa87018ec8e6ea"></a>

## Next pages — xcsh_tunnel / bde3443da9de / 6

- [Property reference](../guides/data-sources--tunnel--reference--group-001.md#canonical-34a9b4f3823b8e24e5b8fd12ce61ffd5e7ed68ad007c35cd4cfd37ceb9d5aefa)
- [Examples](../guides/data-sources--tunnel--examples--group-001.md#canonical-f0f2d675aa3ad4df19b03048afbed9c24424795f3a5748460812f25e5fbb78f4)
