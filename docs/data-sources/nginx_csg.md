---
page_title: "xcsh_nginx_csg landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_csg landing."
---

# xcsh_nginx_csg landing

<a id="canonical-c410c59f0b8d71705523fc80bce3a448bc2e5b40e768ceaa424debd127388095"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd2f8ae0660ae882054100f05958079dfea926ee745631462f9b3f8b9e07f445"></a>

## xcsh_nginx_csg — xcsh_nginx_csg / 69c1aa49dc31 / 2

Breadcrumbs:

- xcsh_nginx_csg

Manages a Nginx Csg resource in F5 Distributed Cloud for get nginx csg configuration. configuration.
(read-only data source)

<a id="canonical-46458e078527e27c2a737f72ede5566c78b5d8f3188155d402dc92ba89bc5762"></a>

## Prerequisites — xcsh_nginx_csg / 69c1aa49dc31 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c60b29251e0d14becd810a0d62200ff4b862d3c70bbbb4bf4457f17cfa38cd62"></a>

## Minimal configuration — xcsh_nginx_csg / 69c1aa49dc31 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxCsg Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxCsg by name
data "xcsh_nginx_csg" "example" {
  name      = "example-nginx-csg"
  namespace = "staging"
}

output "nginx_csg_id" {
  value = data.xcsh_nginx_csg.example.id
}
```

<a id="canonical-d864f8bbe07350ddd91c949c7d66b64255d2480bdfb01f0bfcc99536aea3f1e3"></a>

## Root configuration — xcsh_nginx_csg / 69c1aa49dc31 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0570c9c6154932ae98b29b57076bc041ea8253617798ec369b4aaff81c8b7fcb"></a>

## Next pages — xcsh_nginx_csg / 69c1aa49dc31 / 6

- [Property reference](../guides/data-sources--nginx_csg--reference--group-001.md#canonical-d475c7367cf5faac52b1c3c5588be6a77c688dee95c47c7e686b714ba7ed1ac0)
- [Examples](../guides/data-sources--nginx_csg--examples--group-001.md#canonical-086165534725127b74852bb823de8e81842f7adbfe62b524f29a1fbc2ee653b6)
