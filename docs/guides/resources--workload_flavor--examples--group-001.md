---
page_title: "xcsh_workload_flavor examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor examples."
---

# xcsh_workload_flavor examples

<a id="canonical-0986ae5df41648085a9c7c99f600c126da00a829a494b8601789f82abbc5dee4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29a0eaca5389bd21f2643a1ad612dfb054c8cebc4edf8f90ab35e2f297af34ca"></a>

## Examples — Examples / 9a349f582659 / 2

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)
- Examples

<a id="canonical-c4eacb878a185fe07d907f95f2c1bc5c637095101b6c1ae3e19353ee3167b65f"></a>

## Complete configurations — Examples / 9a349f582659 / 3

- [Resource](resources--workload_flavor--examples--group-001.md#canonical-dac22c0b8d19155b3c07e69cfd3d8fcbe24c93f3cd62a03866519d752bcfd392): valid configuration.

<a id="canonical-8316ffc4b8670c4e6b5fad43eb843001dfd18471939950aff6bed3f0650f755d"></a>

## Next pages — Examples / 9a349f582659 / 4

- [Resource](resources--workload_flavor--examples--group-001.md#canonical-dac22c0b8d19155b3c07e69cfd3d8fcbe24c93f3cd62a03866519d752bcfd392)
- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)

<a id="canonical-dac22c0b8d19155b3c07e69cfd3d8fcbe24c93f3cd62a03866519d752bcfd392"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3bd6917c824f5cd73aca9b5808624a7fc84ec70ce1697a68bd1285d30755547"></a>

## Resource — Resource / 21dd0a0459fb / 2

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)
- [Examples](resources--workload_flavor--examples--group-001.md#canonical-0986ae5df41648085a9c7c99f600c126da00a829a494b8601789f82abbc5dee4)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_workload_flavor/resource.tf`; digest `sha256:3a521cc20a27b0aff22bd7904a8747719375d23df70f4cf67625e0e522097422`.

```terraform
# WorkloadFlavor Resource Example
# Manages workload_flavor in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WorkloadFlavor configuration
resource "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}
```

<a id="canonical-6c940158bee896481583a63e82787e7f4a0f9c910a72a46957b5f3a91781d86e"></a>

## Next pages — Resource / 21dd0a0459fb / 3

- [Examples](resources--workload_flavor--examples--group-001.md#canonical-0986ae5df41648085a9c7c99f600c126da00a829a494b8601789f82abbc5dee4)
- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)
