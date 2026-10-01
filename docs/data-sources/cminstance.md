---
page_title: "xcsh_cminstance landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance landing."
---

# xcsh_cminstance landing

<a id="canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9ba3a789bbc9682ad501337a2656c0140c1ae1c842f7826a5c4134b5c28bb11"></a>

## xcsh_cminstance — xcsh_cminstance / 7762dc8c376b / 2

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-aec393d1be545ed7744c908e4fce8652b7456dfa9074b5ea66616ae45e345279"></a>

## Prerequisites — xcsh_cminstance / 7762dc8c376b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a5e49144a054ee4d3790ab7d20eb36db1649256ccb379c448c13dc08875bf47a"></a>

## Minimal configuration — xcsh_cminstance / 7762dc8c376b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```

<a id="canonical-3934195b0293bbb33e8c9ca2fd9d0589dd8cb4fcd8ee125e68a088e9cc7893f1"></a>

## Root configuration — xcsh_cminstance / 7762dc8c376b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-562fd4079e41bfb3277aa20c686bc512f1b31927c8ba4f757aa12d441c352a43"></a>

## Next pages — xcsh_cminstance / 7762dc8c376b / 6

- [Property reference](../guides/data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [Examples](../guides/data-sources--cminstance--examples--group-001.md#canonical-a7fb1bdcf2f4a1339b3dacfc81a14bb5d9562839eee61c351ac173a8891f74e3)
