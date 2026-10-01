---
page_title: "xcsh_code_base_integration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration examples."
---

# xcsh_code_base_integration examples

<a id="canonical-1c960e71d8f0331c3ff5be704f697493d965f155409f7befe21e1e6bbd7bfbb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90c21f9f71e4ba8250a63273a78bc2e4be3fbc4025010b874934807f14b99fa7"></a>

## Examples — Examples / e974c8dedba3 / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- Examples

<a id="canonical-1921371a6baf4a796862d609aed8dcfbe31b54ba1c18f41a6f34a5d2b5c80946"></a>

## Complete configurations — Examples / e974c8dedba3 / 3

- [Data source](data-sources--code_base_integration--examples--group-001.md#canonical-b708c2df0f189f0c7fe8a5b89a9cacaebf52644c93550b5332d840fc93f86266): valid configuration.

<a id="canonical-e5c15396210ba06a1e147c38467c7fba1ddc59025ab28375792970d6ca5092e3"></a>

## Next pages — Examples / e974c8dedba3 / 4

- [Data source](data-sources--code_base_integration--examples--group-001.md#canonical-b708c2df0f189f0c7fe8a5b89a9cacaebf52644c93550b5332d840fc93f86266)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)

<a id="canonical-b708c2df0f189f0c7fe8a5b89a9cacaebf52644c93550b5332d840fc93f86266"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5966e89b3e0982f33b38488d561eab781bf333584b6d28e4b5c90bc33ff0b78a"></a>

## Data source — Data source / 2adad45a48bb / 2

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
- [Examples](data-sources--code_base_integration--examples--group-001.md#canonical-1c960e71d8f0331c3ff5be704f697493d965f155409f7befe21e1e6bbd7bfbb3)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_code_base_integration/data-source.tf`; digest `sha256:170c39b82f55f5ee2a4de8f66b829b7a3cbefe8bbffc6b09199912a373cb3b88`.

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

<a id="canonical-beccd0d6123c66220e2882f8c0826a8a97b25174665181a932a62bb06bdbddfd"></a>

## Next pages — Data source / 2adad45a48bb / 3

- [Examples](data-sources--code_base_integration--examples--group-001.md#canonical-1c960e71d8f0331c3ff5be704f697493d965f155409f7befe21e1e6bbd7bfbb3)
- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-dfd634386a632b8766c1f79a7fe219dee977271c92ddfb192ce6925365cdbb55)
