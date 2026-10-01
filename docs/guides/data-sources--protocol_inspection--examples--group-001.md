---
page_title: "xcsh_protocol_inspection examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection examples."
---

# xcsh_protocol_inspection examples

<a id="canonical-aa80b641987380a6c2cb4bc8c9b5799bd83daa72873da274772ef5d74b57bf9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d03fe53a127b8e8ce74942ed83a91cab538ae9b1307e69ea4f7286b9b7f681ed"></a>

## Examples — Examples / 245d1464a2e3 / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- Examples

<a id="canonical-e6aeeb95a9d53a28382bb168f56ad1bce34b78da54a48d38347a8fb9086785bc"></a>

## Complete configurations — Examples / 245d1464a2e3 / 3

- [Data source](data-sources--protocol_inspection--examples--group-001.md#canonical-b93d8648d1c9ff119d9105876be898cce59c788bd69ed9069e5cee2d6319c4b9): valid configuration.

<a id="canonical-f62949ecea8eaa64042c6c189ba42c76d49379fbe6064ace01cb8648cd1617fe"></a>

## Next pages — Examples / 245d1464a2e3 / 4

- [Data source](data-sources--protocol_inspection--examples--group-001.md#canonical-b93d8648d1c9ff119d9105876be898cce59c788bd69ed9069e5cee2d6319c4b9)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)

<a id="canonical-b93d8648d1c9ff119d9105876be898cce59c788bd69ed9069e5cee2d6319c4b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d94466069f5d4a6ab1a3f0e8fbd9e918046096a44e11ed8fc623c9beac4a23b6"></a>

## Data source — Data source / fc2d9b678e9b / 2

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
- [Examples](data-sources--protocol_inspection--examples--group-001.md#canonical-aa80b641987380a6c2cb4bc8c9b5799bd83daa72873da274772ef5d74b57bf9b)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protocol_inspection/data-source.tf`; digest `sha256:8e357f2e82e2e660e9b2335aec0a0bf6fb5efafdd6746da23a0a982fc6492237`.

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

<a id="canonical-ba0b282fd174be0f8ba82ce2bec05ddcc730a2b71ef462c41cd360a19b72f0b0"></a>

## Next pages — Data source / fc2d9b678e9b / 3

- [Examples](data-sources--protocol_inspection--examples--group-001.md#canonical-aa80b641987380a6c2cb4bc8c9b5799bd83daa72873da274772ef5d74b57bf9b)
- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-77c725c112dfe8c6a26fd1bebe3fb78dbea0cb86c9b03fb1a6550cb324fde634)
