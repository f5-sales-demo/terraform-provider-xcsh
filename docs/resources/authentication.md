---
page_title: "xcsh_authentication landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication landing."
---

# xcsh_authentication landing

<a id="canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bec52916b952ed6a50ed265ba605975b02ff168746f93cc1cec1edd24c45bf75"></a>

## xcsh_authentication — xcsh_authentication / 477fbdbbdc2b / 2

Breadcrumbs:

- xcsh_authentication

Manages a Authentication resource in F5 Distributed Cloud.

<a id="canonical-c5f38ded6ed0160d26b779f57c6f13a22159870dac610a9293663b6612421d74"></a>

## Prerequisites — xcsh_authentication / 477fbdbbdc2b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-01b84e98929fb86d58576d672d800e08592b20c76292020bdacf7fdcbda897a8"></a>

## Minimal configuration — xcsh_authentication / 477fbdbbdc2b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```

<a id="canonical-1810c4b391fc6d20620452dc730b0a0e281e62cf548feeb0448f5b347441e2b5"></a>

## Root configuration — xcsh_authentication / 477fbdbbdc2b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ae4d4e30b57a3b97966bb10a01e6e637c5f922bd85faa857822c81eea2a850ea"></a>

## Next pages — xcsh_authentication / 477fbdbbdc2b / 6

- [Property reference](../guides/resources--authentication--reference--group-001.md#canonical-c4df54315d421dd70758e3db0bb63b6e69f2aee4a7ec3a11548b735bedc072c3)
- [Examples](../guides/resources--authentication--examples--group-001.md#canonical-aa08cf8ca5220ee6b47053b20724624aa580a5cf67d8719f0af153aac6c46876)
- [Import](../guides/resources--authentication--lifecycle--group-001.md#canonical-ee10d559d28ec3c590dd2a64cb8554308b348b00a8a676fe65559da367b2f02e)
- [Timeouts](../guides/resources--authentication--lifecycle--group-001.md#canonical-472462d769ee8d837b0a170a03e566803cfdc428faa4e4c0f7e5410abfa4ba71)
