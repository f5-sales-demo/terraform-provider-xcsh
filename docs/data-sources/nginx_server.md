---
page_title: "xcsh_nginx_server landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_server landing."
---

# xcsh_nginx_server landing

<a id="canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c70c1b82cad720727bf802c2ca6cc1770c6bb8a978b39dc26780f2bb9949b2df"></a>

## xcsh_nginx_server — xcsh_nginx_server / c1e15b22a748 / 2

Breadcrumbs:

- xcsh_nginx_server

Manages a Nginx Server resource in F5 Distributed Cloud for get nginx server block configuration.
configuration. (read-only data source)

<a id="canonical-a850f711cb580dc47601b9095395b1068e50d06dbadec724f8928fa2a7e9ddda"></a>

## Prerequisites — xcsh_nginx_server / c1e15b22a748 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-eb12a11f69069d674f6b01ae3f1cd373818862ed7ae58f4f8500765f79c7c541"></a>

## Minimal configuration — xcsh_nginx_server / c1e15b22a748 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServer by name
data "xcsh_nginx_server" "example" {
  name      = "example-nginx-server"
  namespace = "staging"
}

output "nginx_server_id" {
  value = data.xcsh_nginx_server.example.id
}
```

<a id="canonical-86b25ab8dbd02fa2760e4a689e2a1b1794a7a853b349e0de50d8d0929304d090"></a>

## Root configuration — xcsh_nginx_server / c1e15b22a748 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3b899add0cd3680d0df56eae5665b73569968c0a38d04383d075e0d23f1eb3a7"></a>

## Next pages — xcsh_nginx_server / c1e15b22a748 / 6

- [Property reference](../guides/data-sources--nginx_server--reference--group-001.md#canonical-4beef9df327d559710a9ea2ff0a10f025a18a9e27f5211c9c423685fdb688084)
- [Examples](../guides/data-sources--nginx_server--examples--group-001.md#canonical-b0fd1de3f10926ffab2217c5eb0d3ad5844ce69cc1a748f0b031faca928b3c5d)
