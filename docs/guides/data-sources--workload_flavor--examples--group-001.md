---
page_title: "xcsh_workload_flavor examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor examples."
---

# xcsh_workload_flavor examples

<a id="canonical-a9945d616ff369efe638b8eb3f7f840bfb4552c8f0287e9b4578032eb5f36892"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47f94a6781dcadaf6e4b5729837aa187dffe484c6826b5fe81f0a78a672c1f63"></a>

## Examples — Examples / e2a8a3107f1f / 2

Breadcrumbs:

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-8c8c0ecc2e6ee25fc807e107477b515e49f3e31343c02d3b3fb96b576dfd8acc)
- Examples

<a id="canonical-de2b62d5f2403e24a79037b6305147f70078003aad038c0d554f863650fb7eb9"></a>

## Complete configurations — Examples / e2a8a3107f1f / 3

- [Data source](data-sources--workload_flavor--examples--group-001.md#canonical-e334815de23f34337a9e8bf2f27bd2d6d76caeef5b2a6f4e15aaad3d06d18dc7): valid configuration.

<a id="canonical-767962ca8dd3d371a843190413e86d8990810a9b13e796a77f1c329015d3f790"></a>

## Next pages — Examples / e2a8a3107f1f / 4

- [Data source](data-sources--workload_flavor--examples--group-001.md#canonical-e334815de23f34337a9e8bf2f27bd2d6d76caeef5b2a6f4e15aaad3d06d18dc7)
- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-8c8c0ecc2e6ee25fc807e107477b515e49f3e31343c02d3b3fb96b576dfd8acc)

<a id="canonical-e334815de23f34337a9e8bf2f27bd2d6d76caeef5b2a6f4e15aaad3d06d18dc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cce6c98b483a9d0e0fffcc877b751da42c53365c22a958be36a541e946951c98"></a>

## Data source — Data source / f1be55dc8945 / 2

Breadcrumbs:

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-8c8c0ecc2e6ee25fc807e107477b515e49f3e31343c02d3b3fb96b576dfd8acc)
- [Examples](data-sources--workload_flavor--examples--group-001.md#canonical-a9945d616ff369efe638b8eb3f7f840bfb4552c8f0287e9b4578032eb5f36892)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_workload_flavor/data-source.tf`; digest `sha256:6a1b8a9b6022de8c3cc75f425f01b978df5bf741969fc4bc0998b22eb6f2b173`.

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

<a id="canonical-fe55d76a171a393c99e92cc1c05550b267603bfefe77a32494acbdf2b9dbe035"></a>

## Next pages — Data source / f1be55dc8945 / 3

- [Examples](data-sources--workload_flavor--examples--group-001.md#canonical-a9945d616ff369efe638b8eb3f7f840bfb4552c8f0287e9b4578032eb5f36892)
- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-8c8c0ecc2e6ee25fc807e107477b515e49f3e31343c02d3b3fb96b576dfd8acc)
