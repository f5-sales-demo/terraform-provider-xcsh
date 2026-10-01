---
page_title: "xcsh_irule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule examples."
---

# xcsh_irule examples

<a id="canonical-2de7ae51f53e24a111116c100bfd19a60484b4b51f3ad177bb6724e12299aa5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc4fb1a55bc3cd732c7942c9e916583b98d3f6933847e9bbc1fc1d83c824ec7e"></a>

## Examples — Examples / cb26730a4a25 / 2

Breadcrumbs:

- [xcsh_irule](../data-sources/irule.md#canonical-2d46463305002f54ff3a1ae4bb70fcb0aa618cbcb1816cb90bb6c2a6ed549962)
- Examples

<a id="canonical-9003f02535725f24f4ed7f320a9e3ebf22c30d165cc1f4a5db2e16946ad43e54"></a>

## Complete configurations — Examples / cb26730a4a25 / 3

- [Data source](data-sources--irule--examples--group-001.md#canonical-88d3de9b6dde73932871165e252a5c862129babf8983331478eb0f23b348294d): valid configuration.

<a id="canonical-76f1f0f52ca15a808ae7d9fe43830f76b2ebe0bb7ac339e684119ea79bf394ed"></a>

## Next pages — Examples / cb26730a4a25 / 4

- [Data source](data-sources--irule--examples--group-001.md#canonical-88d3de9b6dde73932871165e252a5c862129babf8983331478eb0f23b348294d)
- [xcsh_irule](../data-sources/irule.md#canonical-2d46463305002f54ff3a1ae4bb70fcb0aa618cbcb1816cb90bb6c2a6ed549962)

<a id="canonical-88d3de9b6dde73932871165e252a5c862129babf8983331478eb0f23b348294d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e75288d7d794aa5423925c6878853748cf0a2b3c99b2487f76f00da2f4b1e91a"></a>

## Data source — Data source / 9add6b8ece66 / 2

Breadcrumbs:

- [xcsh_irule](../data-sources/irule.md#canonical-2d46463305002f54ff3a1ae4bb70fcb0aa618cbcb1816cb90bb6c2a6ed549962)
- [Examples](data-sources--irule--examples--group-001.md#canonical-2de7ae51f53e24a111116c100bfd19a60484b4b51f3ad177bb6724e12299aa5d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_irule/data-source.tf`; digest `sha256:567750397ff4df80dac106ded5aa33f7fcb698c11cea8ada22ebaa5fc6727f22`.

```terraform
# Irule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Irule by name
data "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"
}

output "irule_id" {
  value = data.xcsh_irule.example.id
}
```

<a id="canonical-f3370bbe36187e61d704f1b8aa592c5893ad789ebd8c20de18cd1c30e7076314"></a>

## Next pages — Data source / 9add6b8ece66 / 3

- [Examples](data-sources--irule--examples--group-001.md#canonical-2de7ae51f53e24a111116c100bfd19a60484b4b51f3ad177bb6724e12299aa5d)
- [xcsh_irule](../data-sources/irule.md#canonical-2d46463305002f54ff3a1ae4bb70fcb0aa618cbcb1816cb90bb6c2a6ed549962)
