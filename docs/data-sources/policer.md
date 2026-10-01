---
page_title: "xcsh_policer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer landing."
---

# xcsh_policer landing

<a id="canonical-fc38c1c976cffbbb7d1001c95e37ff2eb2f3d9753461e9f18f25729eae4994c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2e7c3264d55fd6ada954229331343471a5d16df0bab7dee87b5c1a8c18807fc"></a>

## xcsh_policer — xcsh_policer / b31432c0cbcb / 2

Breadcrumbs:

- xcsh_policer

Manages new policer with traffic rate limits in F5 Distributed Cloud.

<a id="canonical-a21409c8834325da6284a112c708259936ab513cfc1f7913e0aba945add14487"></a>

## Prerequisites — xcsh_policer / b31432c0cbcb / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f67588a88e70365cfdcef78bca9c6d7b4a9e49dfa869a5c43cae5e4fd8bd97c3"></a>

## Minimal configuration — xcsh_policer / b31432c0cbcb / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Policer by name
data "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"
}

output "policer_id" {
  value = data.xcsh_policer.example.id
}
```

<a id="canonical-c32acf54e84680df97f2a0145d0b6fe82e0509dc06665e1715e79b17265cd578"></a>

## Root configuration — xcsh_policer / b31432c0cbcb / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b1ea50810840f79dd8799d82b1a6c87d09b16721c690fb5a54a00fedafaccca1"></a>

## Next pages — xcsh_policer / b31432c0cbcb / 6

- [Property reference](../guides/data-sources--policer--reference--group-001.md#canonical-f87bd9bfc4cb7f2e29ae21826a8d294799efb0c555b9ca771f74fa96c2470bc7)
- [Examples](../guides/data-sources--policer--examples--group-001.md#canonical-1ef5a3f8e94ec063f1bd6e0f9cd0fd48a7ac2bcac02a6bd4e066dc01077cf67d)
