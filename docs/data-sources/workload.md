---
page_title: "xcsh_workload landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload landing."
---

# xcsh_workload landing

<a id="canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78efd7b8d765195d54ada22dbeace47c72c40d00728c4572ea75b6bd204695bf"></a>

## xcsh_workload — xcsh_workload / e13edaef8d35 / 2

Breadcrumbs:

- xcsh_workload

Manages a Workload resource in F5 Distributed Cloud for workload. configuration.

<a id="canonical-8ba98ea690a7f5572068556ed49c3b8f9f8f17e19b3fb8477f45a630e3dab97d"></a>

## Prerequisites — xcsh_workload / e13edaef8d35 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_k8s`.

- virtual_k8s: Namespace for workload deployment

<a id="canonical-c597474ef94018211d0ad9394abd59ab6f4fd3febe3c2cc36061b2899f597233"></a>

## Minimal configuration — xcsh_workload / e13edaef8d35 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Workload Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Workload by name
data "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}

output "workload_id" {
  value = data.xcsh_workload.example.id
}
```

<a id="canonical-69122af91978b3369949100ac4bd554fbcf40c514e5f53c2fdfc65e06f03fac7"></a>

## Root configuration — xcsh_workload / e13edaef8d35 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-fd9f1d2cb095e736a58fb8fa57feac2a8f9af25ac4b26aaaae38f0bc2a885d26"></a>

## Next pages — xcsh_workload / e13edaef8d35 / 6

- [Property reference](../guides/data-sources--workload--reference--group-001.md#canonical-3128aa1366d0810fef144b5c757c17c27d984a4a52c3ca575620ceeae5ab24cd)
- [Examples](../guides/data-sources--workload--examples--group-001.md#canonical-d495c5379762ac477e71189db6d070866bf97284034ec2aa67a79b9d27182312)
