---
page_title: "xcsh_workload_flavor landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor landing."
---

# xcsh_workload_flavor landing

<a id="canonical-8c8c0ecc2e6ee25fc807e107477b515e49f3e31343c02d3b3fb96b576dfd8acc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd8bd8bd5de56ba87c034acf620559460ae2e17da920f1a343cf2833bb62f473"></a>

## xcsh_workload_flavor — xcsh_workload_flavor / ac133c0c04a1 / 2

Breadcrumbs:

- xcsh_workload_flavor

Manages workload\_flavor in F5 Distributed Cloud.

<a id="canonical-9085d48dc3249804ada61863c0dcc97fd57e9c31363400f862a80af9a4e4eb4f"></a>

## Prerequisites — xcsh_workload_flavor / ac133c0c04a1 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c7c8f07e9ae0982eafc99f361fb1fea47236cec2948ff8150e648aab5e6bdd75"></a>

## Minimal configuration — xcsh_workload_flavor / ac133c0c04a1 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WorkloadFlavor Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WorkloadFlavor by name
data "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}

output "workload_flavor_id" {
  value = data.xcsh_workload_flavor.example.id
}
```

<a id="canonical-636a70b2c459f2b67241c7c5bce64293a564d521578d0c305a4717b0e2af20e0"></a>

## Root configuration — xcsh_workload_flavor / ac133c0c04a1 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-19b65502f6aa971cb128df99c90b138bfc46853065e18d2c4b76a3638d637716"></a>

## Next pages — xcsh_workload_flavor / ac133c0c04a1 / 6

- [Property reference](../guides/data-sources--workload_flavor--reference--group-001.md#canonical-9193b6067e4e8e493c8d2e08f2541bfa4d5b090c7074a32eae1dc6e93462dee9)
- [Examples](../guides/data-sources--workload_flavor--examples--group-001.md#canonical-a9945d616ff369efe638b8eb3f7f840bfb4552c8f0287e9b4578032eb5f36892)
