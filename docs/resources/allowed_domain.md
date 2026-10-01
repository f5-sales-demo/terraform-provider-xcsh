---
page_title: "xcsh_allowed_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain landing."
---

# xcsh_allowed_domain landing

<a id="canonical-dbe8b7d2442b45128c46f867dab30be75f23c65c250d4451f1beb4e00c8d65f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22c6bf8712f43af22712df15b7d4661a86c1f26a3f9aee00c87b37d047996ed8"></a>

## xcsh_allowed_domain — xcsh_allowed_domain / 7b691519ee46 / 2

Breadcrumbs:

- xcsh_allowed_domain

Manages allowed domain in F5 Distributed Cloud.

<a id="canonical-a9703f84a0302333eb3c5f478b3d4800a08121c8225c160fb19245023aba1c90"></a>

## Prerequisites — xcsh_allowed_domain / 7b691519ee46 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3f9a7a0616fcc5e919c2200021601a5f21fb611bb459da972c3b9eadd4b5fb41"></a>

## Minimal configuration — xcsh_allowed_domain / 7b691519ee46 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```

<a id="canonical-132a14cc3087c699e25e1b1e70a333b886e1ee45e50f3c4715ddbbd206b92609"></a>

## Root configuration — xcsh_allowed_domain / 7b691519ee46 / 5

Required root properties: `allowed_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ae6370b3038253d41d78539c4b6f64b193dc6b87d37a1af665fb6c4927dfa9fc"></a>

## Next pages — xcsh_allowed_domain / 7b691519ee46 / 6

- [Property reference](../guides/resources--allowed_domain--reference--group-001.md#canonical-b4f5d2267a3810dd15ca31a5e12ece95139a49ccbbc3defd7a52e44af9653777)
- [Examples](../guides/resources--allowed_domain--examples--group-001.md#canonical-104c527e8be9f381bf482f9625410f3ecb948b959e49bbecf63ce5facc703ae7)
- [Import](../guides/resources--allowed_domain--lifecycle--group-001.md#canonical-b699ff7a0af626f3a0fe6d2929e083115601af13d99cf8cacac3a8cf997ba778)
- [Timeouts](../guides/resources--allowed_domain--lifecycle--group-001.md#canonical-9b4dbaac9bf9cb1ece1cb74e19965e6b8864986f574f4ef4ade4cbb7ec4ab236)
