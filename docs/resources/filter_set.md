---
page_title: "xcsh_filter_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set landing."
---

# xcsh_filter_set landing

<a id="canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bb6666c72e153d882f8293a5620500cbceee00441513089e78b31e79052a864"></a>

## xcsh_filter_set — xcsh_filter_set / 642c6f50c63e / 2

Breadcrumbs:

- xcsh_filter_set

Manages specification in F5 Distributed Cloud.

<a id="canonical-6c043a08ce7df925a907a50346509defd4ac91eb112663e42ae72dc7980c9f8d"></a>

## Prerequisites — xcsh_filter_set / 642c6f50c63e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ab431ccb3c9cb5337b35d8496ad14da51d4c22b5acab1d6e87fda2d42ea477ea"></a>

## Minimal configuration — xcsh_filter_set / 642c6f50c63e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FilterSet Resource Example
# Manages specification in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FilterSet configuration
resource "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"

  context_key = "example-value"
}
```

<a id="canonical-db0b1ef230a3d3c82fae78f6d453b50ccea49871dd84e7114bc4790c218d1507"></a>

## Root configuration — xcsh_filter_set / 642c6f50c63e / 5

Required root properties: `context_key`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-6798ab120dac98892d0856125aaad9417004907d18863e8e1ed250bdf4300ae1"></a>

## Next pages — xcsh_filter_set / 642c6f50c63e / 6

- [Property reference](../guides/resources--filter_set--reference--group-001.md#canonical-4b8193f66a3093ca8d9043cb60cc840d73ecc30ec51abc6a7ec8b5264dc822b5)
- [Examples](../guides/resources--filter_set--examples--group-001.md#canonical-ea95c2ddcc082e916032a6c060bd5936a855ab6e5ac9587a6bc454afd9d21723)
- [Import](../guides/resources--filter_set--lifecycle--group-001.md#canonical-c41648d28d83546d80aaaee6f2946680c25f2b8c6b916809afd29e8a710238b1)
- [Timeouts](../guides/resources--filter_set--lifecycle--group-001.md#canonical-30b65814688cf4dbcecb8f534bd494f2c4994055f41667b310a1251fc3c3cf98)
