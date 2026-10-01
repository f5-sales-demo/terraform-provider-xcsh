---
page_title: "xcsh_nginx_instance landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_instance landing."
---

# xcsh_nginx_instance landing

<a id="canonical-b9ef2669c0bf7b4c379f2af4596962e7273bf8b58edf51272ccb89570543335a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-641e48bc55474f616d5fc464a5623c39e6284a0d74d553d87abb898ec89ec8d9"></a>

## xcsh_nginx_instance — xcsh_nginx_instance / 6c7b068e5653 / 2

Breadcrumbs:

- xcsh_nginx_instance

Manages a Nginx Instance resource in F5 Distributed Cloud for get nginx instance configuration.
configuration. (read-only data source)

<a id="canonical-a95f1e114a72547575448b73b42f0b8f71d460392f6491a0cad8c3314546f78f"></a>

## Prerequisites — xcsh_nginx_instance / 6c7b068e5653 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-6de6b905f9ae457db41730c0d2777f3c862a6799906db32cc38adfc3d71da6c7"></a>

## Minimal configuration — xcsh_nginx_instance / 6c7b068e5653 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxInstance by name
data "xcsh_nginx_instance" "example" {
  name      = "example-nginx-instance"
  namespace = "staging"
}

output "nginx_instance_id" {
  value = data.xcsh_nginx_instance.example.id
}
```

<a id="canonical-a7f73b3de557b7a14918f5a59cec3fa4891c2576f40ce04a45e9a13764c33548"></a>

## Root configuration — xcsh_nginx_instance / 6c7b068e5653 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-bb29f3a290f0dd460a24d345fdcc8ce8f60f18b1d3b1701c1c1e3c11afea3b5a"></a>

## Next pages — xcsh_nginx_instance / 6c7b068e5653 / 6

- [Property reference](../guides/data-sources--nginx_instance--reference--group-001.md#canonical-dd5056fa9ff48164c06415e85ac04b9f289523c7992d8a1f6a56f04b83798c21)
- [Examples](../guides/data-sources--nginx_instance--examples--group-001.md#canonical-f1cfd45c9eb52b9043c8dddc78a39f8e6d4c57431fafe737175b7666f359758d)
