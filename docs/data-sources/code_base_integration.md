---
page_title: "xcsh_code_base_integration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration landing."
---

# xcsh_code_base_integration landing

<a id="canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1eb7cbfec959fb8f6e58765d9d02dc41ef7b040706b977f210f0e30d90a30d4"></a>

## xcsh_code_base_integration — xcsh_code_base_integration / 2abd5159de93 / 2

Breadcrumbs:

- xcsh_code_base_integration

Manages integration details in F5 Distributed Cloud.

<a id="canonical-b705c696929efd8293d35be909a2e6444722e13a658f948b01dcc003eb6c8bf5"></a>

## Prerequisites — xcsh_code_base_integration / 2abd5159de93 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b731f8f31f8b7f3fc26829f33503e8c18af34cb856601e00be2174e922744c24"></a>

## Minimal configuration — xcsh_code_base_integration / 2abd5159de93 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CodeBaseIntegration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CodeBaseIntegration by name
data "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}

output "code_base_integration_id" {
  value = data.xcsh_code_base_integration.example.id
}
```

<a id="canonical-2ebdec4fd6760e4c63bfce3aa7d5993cfa677b8d57a74a4a2c9e558c12525690"></a>

## Root configuration — xcsh_code_base_integration / 2abd5159de93 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a2ea337ba99147abc175da850b6357573409f1214d39ac5af11a5b340fbf3060"></a>

## Next pages — xcsh_code_base_integration / 2abd5159de93 / 6

- [Property reference](../guides/data-sources--code_base_integration--reference--group-001.md#canonical-bf7364062ab003f6a5afc722ec7c5ff0dffa6a5c90e531ab18f12c5fdd2db00a)
- [Examples](../guides/data-sources--code_base_integration--examples--group-001.md#canonical-1c960e71d8f0331c3ff5be704f697493d965f155409f7befe21e1e6bbd7bfbb3)
