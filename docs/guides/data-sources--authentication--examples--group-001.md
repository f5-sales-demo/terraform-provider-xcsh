---
page_title: "xcsh_authentication examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication examples."
---

# xcsh_authentication examples

<a id="canonical-f543f75ce3357d8e276af5544257737200167f9301ffcd387f66063716548ac0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8d0a36620973ce32af78f5cefafc42e6b0125ad3a1f557321e743cf6494a2af"></a>

## Examples — Examples / e20fd712aaa7 / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- Examples

<a id="canonical-45824f9b47defa47047a1ebee5c898eb16c2a63a7f597e1896ba72a9fd779374"></a>

## Complete configurations — Examples / e20fd712aaa7 / 3

- [Data source](data-sources--authentication--examples--group-001.md#canonical-1ba3e4f37ad79339ed06bfc03b86c12d3ef3c8282f674633db67ad72f160b88f): valid configuration.

<a id="canonical-6c10ed3c97750ef5724ede4f9ba8f825fdaaab05984a423e717c5dfc4bcfe205"></a>

## Next pages — Examples / e20fd712aaa7 / 4

- [Data source](data-sources--authentication--examples--group-001.md#canonical-1ba3e4f37ad79339ed06bfc03b86c12d3ef3c8282f674633db67ad72f160b88f)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)

<a id="canonical-1ba3e4f37ad79339ed06bfc03b86c12d3ef3c8282f674633db67ad72f160b88f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-723620444f4954190b67706325196bb4f3b5bf623197d0f55e06ef16654e2fed"></a>

## Data source — Data source / 4fd6927d7dcb / 2

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
- [Examples](data-sources--authentication--examples--group-001.md#canonical-f543f75ce3357d8e276af5544257737200167f9301ffcd387f66063716548ac0)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authentication/data-source.tf`; digest `sha256:fd8b227ee8d771e5d805c51da258f79cea8ff5031e710c318c14f6217d6efd78`.

```terraform
# Authentication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Authentication by name
data "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}

output "authentication_id" {
  value = data.xcsh_authentication.example.id
}
```

<a id="canonical-dcba30b20bd2ca8323903bc3b68cd1bdcf5efd5bb53a8f1aeaca557e57affc43"></a>

## Next pages — Data source / 4fd6927d7dcb / 3

- [Examples](data-sources--authentication--examples--group-001.md#canonical-f543f75ce3357d8e276af5544257737200167f9301ffcd387f66063716548ac0)
- [xcsh_authentication](../data-sources/authentication.md#canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413)
