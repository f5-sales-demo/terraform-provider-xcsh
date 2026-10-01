---
page_title: "xcsh_protocol_inspection landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection landing."
---

# xcsh_protocol_inspection landing

<a id="canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23dd275ee6e3ea1421ae7006587aacc6c36904cd66c16f89af72dea7c9ef112d"></a>

## xcsh_protocol_inspection — xcsh_protocol_inspection / 182f5e9b2592 / 2

Breadcrumbs:

- xcsh_protocol_inspection

Manages Protocol Inspection Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

<a id="canonical-692aea8a420ddab70d3fe5f0be0ed3be46fa1bddac183afeefe393c304abdd1c"></a>

## Prerequisites — xcsh_protocol_inspection / 182f5e9b2592 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f2981b0f70e7fdf39883a23df884c1a37193d39a1c2b04c9a8e21581fcb580b8"></a>

## Minimal configuration — xcsh_protocol_inspection / 182f5e9b2592 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolInspection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolInspection by name
data "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}

output "protocol_inspection_id" {
  value = data.xcsh_protocol_inspection.example.id
}
```

<a id="canonical-12995ed8d1f4821b29f5dfa1e31f3553bfbaf765648ca54193dca134fae571ab"></a>

## Root configuration — xcsh_protocol_inspection / 182f5e9b2592 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-6b8e419c96cb1e38bc9655f9801907d54eed31d3e705c06099454bfb451bde2d"></a>

## Next pages — xcsh_protocol_inspection / 182f5e9b2592 / 6

- [Property reference](../guides/data-sources--protocol_inspection--reference--group-001.md#canonical-8d7eb163bd367f14551787af5c3a1edacc58ccd9ad3d628927f70e0bdc1d96cb)
- [Examples](../guides/data-sources--protocol_inspection--examples--group-001.md#canonical-aa80b641987380a6c2cb4bc8c9b5799bd83daa72873da274772ef5d74b57bf9b)
