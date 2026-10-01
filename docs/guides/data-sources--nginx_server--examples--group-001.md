---
page_title: "xcsh_nginx_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_server examples."
---

# xcsh_nginx_server examples

<a id="canonical-b0fd1de3f10926ffab2217c5eb0d3ad5844ce69cc1a748f0b031faca928b3c5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a7cfa621ef400dc8115241db8b19fd148ebaa5bb3d635105c51b113ae37454b"></a>

## Examples — Examples / 435318c4e1a1 / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- Examples

<a id="canonical-ab137c73e6e773408943bc96ffccd337f09eb8cbc02a2a3ba420d8da70f748b0"></a>

## Complete configurations — Examples / 435318c4e1a1 / 3

- [Data source](data-sources--nginx_server--examples--group-001.md#canonical-eaf31cf02ef293aef0c69714682e5d06f0766f2193f129e07162e45db6342974): valid configuration.

<a id="canonical-eddcee1526c307541d2253cdfbef2ec682b10df156e9c0c46fdbe4105ba0e966"></a>

## Next pages — Examples / 435318c4e1a1 / 4

- [Data source](data-sources--nginx_server--examples--group-001.md#canonical-eaf31cf02ef293aef0c69714682e5d06f0766f2193f129e07162e45db6342974)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)

<a id="canonical-eaf31cf02ef293aef0c69714682e5d06f0766f2193f129e07162e45db6342974"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e6ddd0a978a4713b6bae4398abdfe8b7f6da82646e2ffd6e8704cf7d852b593"></a>

## Data source — Data source / ff9942c0304d / 2

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
- [Examples](data-sources--nginx_server--examples--group-001.md#canonical-b0fd1de3f10926ffab2217c5eb0d3ad5844ce69cc1a748f0b031faca928b3c5d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_server/data-source.tf`; digest `sha256:255e0e3089bba8c6998465d72659b209b4c83e33fec880e96e0bccd476d39e6e`.

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

<a id="canonical-d643c4bc6cf1fa64a2a461c3c4dd5f0c90bacfb0fe4111e014cac5ea5052d487"></a>

## Next pages — Data source / ff9942c0304d / 3

- [Examples](data-sources--nginx_server--examples--group-001.md#canonical-b0fd1de3f10926ffab2217c5eb0d3ad5844ce69cc1a748f0b031faca928b3c5d)
- [xcsh_nginx_server](../data-sources/nginx_server.md#canonical-81ac0adffbfbb989fac7fac7321455b46042e5aa76bf6fb2f8ad3b6e9b4a4792)
