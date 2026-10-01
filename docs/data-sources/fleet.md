---
page_title: "xcsh_fleet landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet landing."
---

# xcsh_fleet landing

<a id="canonical-f4c9b93c4771c8961bdf548833b4eb42ad13d7fe37bb59509edc38c238a94f2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dabb3bcf0879995b6c7f8c10f6fa3e7e40949fc2584447ce56915ee6f98fb81"></a>

## xcsh_fleet — xcsh_fleet / 36a03be8ac28 / 2

Breadcrumbs:

- xcsh_fleet

Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

<a id="canonical-f7233a49c3399e4e3bca4185149084dcc2f48693d3adac7ebcde1457e7a617cd"></a>

## Prerequisites — xcsh_fleet / 36a03be8ac28 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-9d65745e5faa0df5b57245209cb2c5921dc262d0d4a311ee83972ac08ce00c3b"></a>

## Minimal configuration — xcsh_fleet / 36a03be8ac28 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Fleet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Fleet by name
data "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"
}

output "fleet_id" {
  value = data.xcsh_fleet.example.id
}
```

<a id="canonical-980d036d100c0b3b7b7849e5889c871ab0af0665208ab00d54c0c7099a3c50cc"></a>

## Root configuration — xcsh_fleet / 36a03be8ac28 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-aa50ae843ff8ed9711922e0d82ddf886c9695ba5aef7b4d57555c6adc1ffbd33"></a>

## Next pages — xcsh_fleet / 36a03be8ac28 / 6

- [Property reference](../guides/data-sources--fleet--reference--group-001.md#canonical-505760c4f93fab63fb2a0a602ddb76a1a640623711e5b0413ada6dfd870e2dc9)
- [Examples](../guides/data-sources--fleet--examples--group-001.md#canonical-6b4f913b2d40c1f26fc8e814243719a8906bcedb349bf38a8429fa1910f02b67)
